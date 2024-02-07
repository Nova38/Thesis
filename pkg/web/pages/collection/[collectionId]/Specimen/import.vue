<script setup lang="ts">
import type { ParseResult } from 'papaparse'

import Papa from 'papaparse'
import { crush, get, keys, set } from 'radash'
import { randomUUID } from 'uncrypto'
import { ccbio } from '~/lib'

const store = useCollectionsStore()
const catalogNumbers = store.SpecimenCatalogNumbers
// Reactive variables
const file = ref<File | null>(null)

type status = 'error' | 'importing' | 'new' | 'pre-existing' | 'success'

interface RowMeta {
  status: status
  statusMessage?: string
  // index: string;
  uuid: string
}

const SexOptions = [
  { label: 'SEX_UNKNOWN', value: 1 },
  { label: 'SEX_ATYPICAL', value: 2 },
  { label: 'SEX_MALE', value: 3 },
  { label: 'SEX_FEMALE', value: 4 },
]

const AgeOptions = [
  { label: 'AGE_UNKNOWN', value: 1 },
  { label: 'AGE_NEST', value: 2 },
  { label: 'AGE_EMBRYO_EGG', value: 3 },
  { label: 'AGE_CHICK_SUBADULT', value: 4 },
  { label: 'AGE_ADULT', value: 5 },
  { label: 'AGE_CONTINGENT', value: 6 },
]

// Table 1
const headers = ref<string[]>([])
const rawData = ref<Record<string, string>[]>()
const RowMeta = ref<RowMeta[]>([])
const RowsSelected = ref<Record<string, string>[]>([])

// const rowToUUID = ref<Map<number, string>>(new Map());

// This is a mapping between the specimen keys and the column headers
// <column header>: <specimen key>
const specimenMapping = ref<Record<string, string>>({})

watch(specimenMapping.value, () => {
  console.log('specimenMapping')

  RowsSelected.value?.forEach((data) => {
    if (rawData.value === undefined)
      return

    console.group('import')
    RowsSelected

    const importingSpecimen = MakeEmptySpecimen()

    for (const key in specimenMapping.value) {
      console.log(key)
      if (Object.prototype.hasOwnProperty.call(specimenMapping.value, key)) {
        const element = specimenMapping.value[key]

        const value = data[element]
        // console.log({ key, value, element });
        // Convert the dates
        if (key === 'primary.catalogNumber') {
          console.log(value)
          console.log(store.SpecimenCatalogNumbers)
          // see if the field is in the store.SpecimenCatalogNumbers list
          if (store.SpecimenCatalogNumbers?.includes(value)) {
            // if it is, set the status to pre-existing
            RowMeta.value[Number.parseInt(data.index)].status = 'pre-existing'
          }
          else {
            RowMeta.value[Number.parseInt(data.index)].status = 'new'
          }
        }
        if (key === 'secondary.sex') {
          // if it doesn't exist on the list of sex strings add it
          // if (!sexStrings.value.includes(value)) {
          //   sexStrings.value.push(value);
          // }
          try {
            if (!isNaN(Number.parseInt(value)))
              set(importingSpecimen, key, 0)
            if (isNaN(Number.parseInt(value)))
              set(importingSpecimen, key, Number.parseInt(value))
          }
          catch (error) {
            console.log(error)
            set(importingSpecimen, key, 0)
          }
        }
        if (key === 'secondary.age') {
          // if it doesn't exist on the list of sex strings add it
          // if (!sexStrings.value.includes(value)) {
          //   sexStrings.value.push(value);
          // }
          try {
            if (!isNaN(Number.parseInt(value)))
              set(importingSpecimen, key, 0)
            if (isNaN(Number.parseInt(value)))
              set(importingSpecimen, key, Number.parseInt(value))
            // set(importingSpecimen, key, parseInt(value));
          }
          catch (error) {
            console.log(error)
            set(importingSpecimen, key, 0)
          }
        }
        set(importingSpecimen, key, value)
      }
    }
    if (data.index) {
      importingSpecimen.specimenId
        = RowMeta.value[Number.parseInt(data.index)].uuid

      dataMap.value.set(Number.parseInt(data.index), importingSpecimen)
    }

    console.log(importingSpecimen)
    console.groupEnd()
  })
})

function clearKey(key: string) {
  console.log(key)
  specimenMapping.value[key] = ''
}

const numberFelids = [
  'geography.latitude',
  'geography.longitude',
  'secondary.wight',
  'secondary.sex',
  'secondary.age',
]

// Table 2
const possessedData = computed(() => {
  if (!RowsSelected.value || !specimenMapping)
    return

  const possessedData: ccbio.Specimen[] = []

  RowsSelected.value?.forEach((data) => {
    if (rawData.value === undefined)
      return

    console.group('import')
    RowsSelected

    const importingSpecimen = MakeEmptySpecimen()

    for (const key in specimenMapping.value) {
      if (Object.prototype.hasOwnProperty.call(specimenMapping.value, key)) {
        const element = specimenMapping.value[key]

        const value = data[element]
        // Convert the dates
        // console.log({ key, value, element });
        set(importingSpecimen, key, value)

        numberFelids.forEach((e) => {
          const val: string = get(value, e)
          if (isNaN(Number.parseFloat(val))) {
            console.log('isNaN')
            set(importingSpecimen, e, 0)
          }
          else {
            console.log(Number.parseFloat(val))
            set(importingSpecimen, e, Number.parseFloat(val))
          }
        })
      }
    }

    possessedData.push(importingSpecimen)
  })
  console.groupEnd()
  return possessedData
})

const sortedImportHeaders = computed(() => {
  if (!headers.value)
    return

  return headers.value.sort()
})

const s = new ccbio.Specimen()
const z = JSON.parse(
  MakeEmptySpecimen().toJsonString({ emitDefaultValues: true }),
)

// Set up the mapping objects
const keysForImport = crush(z)
console.log(keysForImport)
console.log(s)
console.log(z)

const notAllowedKeys = ['id', 'last_modified_by', 'collection_id']
const SpecimenKeys: Array<string> = keys(keysForImport)

const FilteredSpecimenKeys = SpecimenKeys.filter(
  k => !notAllowedKeys.includes(k),
)

callOnce(() => {
  useCollectionsStore().LoadRows()
})
/*
  ┌─────────────────────────────────────────────────────────────────────────┐
  │ Load the file on change and parse it with papaparse                     │
  └─────────────────────────────────────────────────────────────────────────┘
 */

// const numberFeilds= ["primary."]

watch(file, (file) => {
  if (!file)
    return

  Papa.parse(file, {
    complete: (results: ParseResult<Record<string, string>>) => {
      console.log(results)

      rawData.value = results.data.map((value: any, index: any) => {
        // rowToUUID.value.set(index, randomUUID());

        RowMeta.value[index] = {
          status: 'new',
          uuid: randomUUID(),
        }

        // if an empty string is found for a key delete it
        // for (const key in value) {
        //   if (Object.prototype.hasOwnProperty.call(value, key)) {
        //     const element = value[key];
        //     if (element === "") {
        //       delete value[key];
        //     }
        //   }
        // }

        if (Object.prototype.hasOwnProperty.call(value, ''))
          delete value['']

        // value["secondary.sex"] = 0;
        // value["secondary.age"] = 0;

        return reactifyObject(
          Object.assign(
            {
              index: `${index}`,
              status: 'new',
            },
            value,
          ),
        )
      })
      headers.value = results.meta.fields || []
      headers.value.unshift('status')

      // If the data's header row that contains the value of a specimen key map it to the specimen key to start
      for (const key of FilteredSpecimenKeys) {
        if (headers.value.includes(key))
          specimenMapping.value[key] = key
      }
    },
    // worker: true,
    header: true,
  })
})

/*
  ┌─────────────────────────────────────────────────────────────────────────┐
  │    Map the headers to the correct specimen fields                       │
  └─────────────────────────────────────────────────────────────────────────┘
 */
const dataMap = ref(new Map<number, ccbio.Specimen>())

function makeHeaders() {
  const flat: any = []

  for (const key of FilteredSpecimenKeys) {
    flat.push({
      field: (row: any) => get(row, key),
      label: key,
    })
  }
  return flat
}

const MappingHeaders = makeHeaders()

function run() {
  console.group('Importing')
  // console.log(rawData.value);

  // const uploadingArray: Promise[] = [];

  // TODO: Make sure they haven't already been uploaded by checking the catalog number

  dataMap.value.forEach((value, key) => {
    console.log(value)
    if (!rawData?.value) {
      console.log('no rawData')
      return
    }
    const row = rawData.value.at(key)
    if (!row) {
      console.log('no row')
      return
    }
    console.log(useCollectionsStore().SpecimenCatalogNumbers)
    console.log(get(value, 'primary.catalogNumber'))
    console.log(
      useCollectionsStore().SpecimenCatalogNumbers?.includes(
        get(value, 'primary.catalogNumber'),
      ),
    )
    if (
      useCollectionsStore().SpecimenCatalogNumbers?.includes(
        get(value, 'primary.catalogNumber'),
      )
    ) {
      console.log('Skipping row as it has already been uploaded', row)
      row.status = 'pre-existing'
      RowMeta.value[key].statusMessage = 'Catalog Number already exists'
      RowMeta.value[key].status = 'pre-existing'
      return
    }

    if (row.status === 'pre-existing') {
      console.debug('Skipping row as it has already been uploaded', row)
      return
    }

    if (row.status === 'success') {
      console.debug('Skipping row as it has already been uploaded', row)
      return
    }
    try {
      row.status = 'importing'

      value.collectionId = useRoute().params.collectionId.toString()
      console.log(value)

      console.log(numberFelids)
      numberFelids.forEach((e) => {
        const val: string = get(value, e)
        console.log(val)
        if (isNaN(Number.parseFloat(val))) {
          console.log('isNan')
          set(value, e, 0)
        }
        else {
          set(value, e, Number.parseFloat(val))
        }
      })

      console.log(value)

      useCreateSpecimen(new ccbio.Specimen(Specimen.parse(value)))
        .then((res) => {
          console.log(res)
          row.status = 'success'
        })
        .catch((err) => {
          console.log(err)
          row.status = 'error'
          RowMeta.value[key].statusMessage = err?.message
        })
    }
    catch (err: any) {
      console.log(err)
      row.status = 'error'
      RowMeta.value[key].statusMessage = err?.message
    }
  })

  // console.log(possessedData.value);
}

function statusToChipColor(status: status) {
  switch (status) {
    case 'new':
      return 'blue'
    case 'pre-existing':
      return 'pink'
    case 'importing':
      return 'orange'
    case 'success':
      return 'green'
    case 'error':
      return 'red'
    default:
      return 'blue'
  }
}
</script>

<template>
  <q-page class="full-height">
    <q-card class="q-pa-md q-ma-md">
      <QCardSection>
        <p class="font-bold">
          Import Specimens
        </p>
      </QCardSection>
      <q-card-section>
        <h2>Select CSV file to import from</h2>
        <q-file
          v-model="file"
          accept=".csv"
          outlined
        >
          <template #prepend>
            <q-icon name="attach_file" />
          </template>
        </q-file>
      </q-card-section>
      <QCardSection class="flex flex-row items-center gap-2 justify-center">
        <QTable
          :hide-bottom="true"
          :rows="SexOptions"
          dense
        />
        <QTable
          :hide-bottom="true"
          :pagination="{ rowsPerPage: 0 }"
          :rows="AgeOptions"
          dense
        />
      </QCardSection>
      <q-card-section>
        <q-table
          v-model:selected="RowsSelected"
          :rows="rawData"
          dense
          row-key="index"
          selection="multiple"
        >
          <template #body-cell-status="props">
            <q-td :props="props">
              <UPopover
                :popper="{ adaptive: true }"
                mode="hover"
              >
                <UBadge
                  :color="statusToChipColor(props.row[props.col.field])"
                  :label="props.row[props.col.field]"
                />
                <q-circular-progress
                  v-if="props.row[props.col.field] == 'loading'"
                  color="warn"
                  indeterminate
                  rounded
                  size="15px"
                />
                <template
                  v-if="RowMeta[props.row.index].statusMessage != ''"
                  #panel
                >
                  <div class="p-4">
                    <pre wrap>
                    {{ RowMeta[props.row.index].statusMessage }}</pre>
                  </div>
                </template>
              </UPopover>

            <!-- <q-chip
                outline
                :color="statusToChipColor(props.row[props.col.field])"
              >
                <span class="q-pr-sm">
                  {{ props.row[props.col.field] }}
                </span>
              </q-chip> -->
            </q-td>
          </template>
        </q-table>
      </q-card-section>

      <!-- <SecondaryMapper
        :sex-strings="sexStrings"
        :sex-mapping="sexMapping"
        :age-strings="ageStrings"
        :age-mapping="ageMapping"
      /> -->
      <!-- {{ (ageMapping, sexMapping) }} -->
      <q-card-section>
        <div v-if="possessedData">
          <q-table
            :columns="MappingHeaders"
            :rows="possessedData"
            dense
          >
            <template #header-cell="props">
              <q-th :props="props">
                <div class="">
                  {{ props.col.label }}
                </div>
                <q-select
                  v-model="specimenMapping[props.col.label]"
                  :options="sortedImportHeaders"
                  dense
                  label="key"
                  label-color="teal-10"
                  stack-label
                >
                  <template
                    v-if="specimenMapping[props.col.label]"
                    #append
                  >
                    <q-icon
                      class="cursor-pointer"
                      color="red"
                      dense
                      name="cancel"
                      size=".75em"
                      @click.stop.prevent="clearKey(props.col.label)"
                    />
                  </template>
                </q-select>
              <!-- <div v-if="props.col.label">hi</div> -->
              </q-th>
            </template>
          </q-table>
        <!-- <q-btn
      color="white"
      text-color="black"
      label=""
      @click="store.upload()"
    /> -->
        </div>
      </q-card-section>
      <q-card-section>
        <q-btn
          class="full-width"
          color="secondary"
          label="Upload Selected"
          @click="run"
        />
      </q-card-section>
    </q-card>
  </q-page>
</template>
