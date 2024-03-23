<script setup lang="ts">
import type { ParseResult } from 'papaparse'

import Papa from 'papaparse'
import { crush, get, keys, set } from 'radash'
// import { ccbio } from 'saacs'
const csv: Ref<Record<string, string>[]> = ref([])

callOnce(() => {
  useCollectionsStore().LoadRows()
})

// Table 1
const headers = ref<string[]>([])
const rawRows = ref<Record<string, string>[]>([])
const RowMeta = ref<UpdateRowMeta[]>([])
const RowsSelected = ref<Record<string, string>[]>([])
const makeNew = ref<boolean>(true)

// This is a mapping between the specimen keys and the column headers
// <column header>: <specimen key>
const specimenMapping = ref<Record<string, string>>({})

/*
  ┌─────────────────────────────────────────────────────────────────────────┐
  │ Process The RowMeta on specimenMapping Update                           │
  └─────────────────────────────────────────────────────────────────────────┘
 */
watch([specimenMapping.value, RowsSelected], () => {
  console.log('SpecimenMapping Updated', specimenMapping.value)

  const catalogNumKey = specimenMapping.value['primary.catalogNumber']
  RowsSelected.value?.forEach((data) => {
    const catalogNum = data[catalogNumKey]
    const index = Number.parseInt(data.index)
    console.log({ catalogNum, index })

    RowMeta.value[index].exist = 'new'

    RowMeta.value[index].status = 'new'
    RowMeta.value[index].statusMessage = "Specimen hasn't been uploaded"

    if (catalogNum) {
      RowMeta.value[index].catalogNumber = catalogNum
      const specimen =
        useCollectionsStore().GetSpecimenFromCatalogNumber(catalogNum)
      console.log({ specimen })

      if (specimen) {
        RowMeta.value[index].currentSpecimen = specimen
        RowMeta.value[index].exist = 'pre-existing'
      }
    }
  })
})

function clearKey(key: string) {
  specimenMapping.value[key] = ''
}

// Table 2
const possessedData = useArrayMap(RowsSelected, (data) => {
  const newSpecimen = MakeEmptySpecimen()

  for (const key in specimenMapping.value) {
    if (Object.prototype.hasOwnProperty.call(specimenMapping.value, key)) {
      if (key === 'primary.catalogNumber') {
        // continue;
      }
      set(newSpecimen, key, data[specimenMapping.value[key]])
    }
  }
  return { index, specimen: newSpecimen }
})

const selectedCurrent = useArrayMap(RowsSelected, (data) => {
  return RowMeta.value[data.index].currentSpecimen
})

const sortedImportHeaders = computed(() => {
  if (!headers.value) return

  return headers.value.sort()
})

const z = JSON.parse(
  MakeEmptySpecimen().toJsonString({ emitDefaultValues: true }),
)

// Set up the mapping objects

const notAllowedKeys = ['specimenId', 'lastModifiedBy', 'collectionId']

const FilteredSpecimenKeys = keys(crush(z)).filter(
  (k) => !notAllowedKeys.includes(k),
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
  if (!file) return

  Papa.parse(file, {
    complete: (results: ParseResult<Record<string, string>>) => {
      headers.value = results.meta.fields || []
      headers.value.unshift('exists')
      headers.value.unshift('status')

      // If the data's header row that contains the value of a specimen key map it to the specimen key to start
      for (const key of FilteredSpecimenKeys) {
        if (headers.value.includes(key)) {
          specimenMapping.value[key] = key

          if (key === 'primary.catalogNumber')
            specimenMapping.value[key] = 'primary.catalogNumber'
        }
      }

      rawRows.value = results.data.map((value: any, index: any) => {
        RowMeta.value[index] = {
          exist: 'new',
          status: 'new',
          statusMessage: '',
          uuid: '',
        }

        if (Object.prototype.hasOwnProperty.call(value, '')) delete value['']

        const catNum: string | undefined = value['primary.catalogNumber']
        if (catNum !== undefined) {
          // console.log(CatNumToUUID(catNum));
          value.specimenId = CatNumToUUID(catNum)
          RowMeta.value[index].uuid = CatNumToUUID(catNum)
        }

        if (
          catNum &&
          useCollectionsStore().SpecimenCatalogNumbers?.includes(catNum)
        ) {
          // console.log("Skipping row as it has already been uploaded", row);
          console.log('pre-existing', catNum)

          RowMeta.value[index].exist = 'pre-existing'
          RowMeta.value[index].uuid = CatNumToUUID(catNum)
          return reactifyObject(
            Object.assign(
              {
                collectionId: useRoute().params.collectionId.toString(),
                exist: 'pre-existing',
                specimenId: CatNumToUUID(catNum),
                status: 'new',
              },
              value,
            ),
          )
        } else {
          RowMeta.value[index].exist = 'new'
          RowMeta.value[index].status = 'new'

          RowMeta.value[index].statusMessage =
            "Specimen Number with Catalog Number doesn't exist"
          return reactifyObject(
            Object.assign(
              {
                collectionId: useRoute().params.collectionId.toString(),
                exist: 'new',
                status: 'new',
              },
              value,
            ),
          )
        }
      })
      console.log('Imported Data')
      console.table(RowMeta.value)
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
// const dataMap = ref(new Map<number, ccbio.Specimen>());

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

async function run() {
  // console.log(rawData.value);
  console.log({ possessedData })

  // const uploadingArray: Promise[] = [];

  // TODO: Make sure they haven't already been uploaded by checking the catalog number
  possessedData.value.forEach(async (value) => {
    // const index = RowMeta.value.findIndex(
    //   (e) => e.catalogNumber === value.primary?.catalogNumber,
    // );
    const index = value.index

    console.log({ index, value })

    if (!value?.specimen.primary?.catalogNumber) {
      console.log('Skipping row as it has no catalog number', value)
      RowMeta.value[index].status = 'error'
      RowMeta.value[index].statusMessage = 'Row has no catalog number'
      return
    }
    value.specimen.specimenId = CatNumToUUID(
      value.specimen?.primary?.catalogNumber,
    )
    value.specimen.collectionId = useRoute().params.collectionId.toString()
    console.log({ value })

    try {
      console.log(RowMeta.value[index])

      if (RowMeta.value[index].status === 'success') {
        console.debug('Skipping row as it has already been uploaded', value)
        RowMeta.value[index].differences = {}
        RowMeta.value[index].statusMessage = 'Row has already been uploaded'
      } else if (RowMeta.value[index].exist === 'new' && makeNew.value) {
        RowMeta.value[index].status = 'loading'
        const specimen = PrepareRow(value, index)
        const meta = await SpecimenUpdate(specimen, new ccbio.Specimen())

        RowMeta.value[index].status = meta.status
        RowMeta.value[index].differences = meta.differences
        RowMeta.value[index].statusMessage = meta.statusMessage
        console.log(RowMeta)
      } else if (RowMeta.value[index].exist === 'pre-existing') {
        const cur = RowMeta.value[index].currentSpecimen
        value.specimenId =
          cur?.specimenId || CatNumToUUID(value.specimen.primary.catalogNumber)

        if (!isDefined(cur)) {
          throw new Error(
            'Row Claims to be Preexisting but has no current specimen',
          )
        }
        RowMeta.value[index].status = 'loading'
        const specimen = PrepareRow(value.specimen, index)
        const meta = await SpecimenUpdate(specimen, new ccbio.Specimen(cur))

        RowMeta.value[index].status = meta.status
        RowMeta.value[index].differences = meta.differences
        RowMeta.value[index].statusMessage = meta.statusMessage
      }
    } catch (err: any) {
      console.log(err)
      RowMeta.value[index].status = 'error'
      RowMeta.value[index].statusMessage = err?.message
    }
  })

  // console.log(possessedData.value);
}

function PrepareRow(specimen: PlainSpecimen, index: number) {
  // specimen.specimenId = CatNumToUUID(value["primary.catalogNumber"]);
  // RowMeta.value[index].uuid = specimen.specimenId;
  // specimen.collectionId = useRoute().params.collectionId.toString();
  console.log({ specimen })

  numberFelids.forEach((e: any) => {
    const val: string = get(specimen, e)
    console.log(val)
    if (isNaN(Number.parseFloat(val))) {
      console.log('isNan')
      set(specimen, e, 0)
    } else {
      set(specimen, e, Number.parseFloat(val))
    }
  })

  return specimen
}

function existToChipColor(exists: ExistStatus) {
  switch (exists) {
    case 'new':
      return 'purple'
    case 'pre-existing':
      return 'pink'
  }
}
function statusToChipColor(status: ProcessingStatus) {
  switch (status) {
    case 'new':
      return 'blue'
    case 'loading':
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
        <p class="font-bold">Bulk Update Specimens</p>
      </QCardSection>

      <ImportCsvFile :headers="headers" :csv="csv" />

      <ImportUpdateTableRaw />

      <QCardSection class="flex flex-row items-center gap-2 justify-center">
        <QTable :hide-bottom="true" :rows="SexOptions" dense />
        <QTable
          :hide-bottom="true"
          :pagination="{ rowsPerPage: 0 }"
          :rows="AgeOptions"
          dense
        />
      </QCardSection>

      <UCard>
        <div class="text-xl font-semibold border-blue-400 border-solid mb-2">
          Current Selected Specimen Values
        </div>
        <QTable :columns="MappingHeaders" :rows="selectedCurrent" dense />
      </UCard>

      <ImportUpdateTablePreview
        :raw-headers="headers"
        :raw-rows="rawRows"
        :specimen-mapping="specimenMapping"
      />

      <!-- :clear-key="clearKey"
        :possessed-data="possessedData"
        :sorted-import-headers="sortedImportHeaders"
        :specimen-mapping="specimenMapping"
      /> -->
      <q-card-section>
        <q-btn
          class="full-width"
          color="secondary"
          label="Upload Selected"
          @click="run()"
        />
      </q-card-section>
    </q-card>
  </q-page>
</template>

<style scoped></style>
