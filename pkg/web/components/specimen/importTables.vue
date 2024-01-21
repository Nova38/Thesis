<script lang="ts" setup>
import type {
  QTableProps,
  QTreeNode,
} from 'nuxt-quasar-ui/dist/runtime/adapter'
import type { ParseResult } from 'papaparse'

import Papa from 'papaparse'
import { crush, keys } from 'radash'

import { FlattedSpecimenKeys } from '~/utils/flatten'

type status = 'error' | 'importing' | 'new' | 'pre-existing' | 'success'
function clearKey(key: string) {
  console.log(key)
  console.log(HeaderMapping[key])
  HeaderMapping[key] = ''
}
const nodes: Ref<QTreeNode[]> = ref([
  {
    children: [
      { label: 'sex' },
      { label: 'age' },
      { label: 'weight' },
      { label: 'weightUnits' },
      { label: 'preparations' },
      { label: 'condition' },
      { label: 'molt' },
      { label: 'notes' },
    ],
    label: 'secondary',
  },
])

interface IRowMeta {
  collectionId?: string
  // index: string;
  status: status
  statusMessage?: string
  uuid?: string
}

interface IColMeta {
  header: string
  mapping?: string
  type: string
}

const SexOptions = ref([
  { enumValue: 1, label: 'SEX_UNKNOWN', mappings: [] },
  { enumValue: 2, label: 'SEX_ATYPICAL', mappings: [] },
  { enumValue: 3, label: 'SEX_MALE', mappings: [] },
  { enumValue: 4, label: 'SEX_FEMALE', mappings: [] },
])

const AgeOptions = ref([
  { enumValue: 1, label: 'AGE_UNKNOWN', mappings: [] },
  { enumValue: 2, label: 'AGE_NEST', mappings: [] },
  { enumValue: 3, label: 'AGE_EMBRYO_EGG', mappings: [] },
  { enumValue: 4, label: 'AGE_CHICK_SUBADULT', mappings: [] },
  { enumValue: 5, label: 'AGE_ADULT', mappings: [] },
  { enumValue: 6, label: 'AGE_CONTINGENT', mappings: [] },
])

const numberFelids = [
  'geography.latitude',
  'geography.longitude',
  'secondary.wight',
  'secondary.sex',
  'secondary.age',
]

const columns = [
  {
    key: 'index',
    label: 'index',
  },
  {
    key: 'primary.catalogNumber',
    label: 'catalog Number',
  },
  {
    key: 'title',
    label: 'Job position',
  },
  {
    key: 'email',
    label: 'Email',
  },
  {
    key: 'role',
  },
]

const ColMeta = ref<Record<string, IColMeta>>({})
// const ImportHeaders = ref<string[]>([]);
const sortedImportHeaders = computed(() => {
  return Object.keys(ColMeta.value).sort()
})
const HeaderMapping = ref<Record<string, string>>({})

const keysForImport = crush(
  JSON.parse(MakeEmptySpecimen().toJsonString({ emitDefaultValues: true })),
)
console.log(keysForImport)

const SpecimenKeys: Array<string> = keys(keysForImport)

const FilteredSpecimenKeys = SpecimenKeys.filter(
  k => !['collection_id', 'id', 'last_modified_by'].includes(k),
)

const rawHeaders: Ref<QTableProps['columns']> = ref()
const rawList: Ref<Record<string, number | string>[]> = ref([])
const RowMeta = ref<IRowMeta[]>([])

const sexStrings = ref<string[]>([])
const ageStrings = ref<string[]>([])

const RowsSelected = ref<Record<string, string>[]>([])

const SpecimenFlatList = useArrayMap(RowsSelected, (item) => {
  const s = FlattenEmptySpecimen()

  for (const key of Object.keys(item)) {
    console.log(key)
    console.log(ColMeta.value[key])

    if (ColMeta.value[key]) {
      console.log(ColMeta.value[key])
      s[key] = item[key]
    }
  }
  return s
})

const file = ref(null)

function makeHeaders() {
  const flat: any = []

  for (const key of FlattedSpecimenKeys()) {
    flat.push({
      key,
      label: key,
    })
  }
  return flat
}

const MappingHeaders = makeHeaders()

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

      rawHeaders.value = results.meta.fields?.map((item) => {
        return {
          field: item,
          label: item,
          name: item,
        }
      })
      rawHeaders.value?.unshift({
        field: 'status',
        label: 'status',
        name: 'status',
      })

      rawList.value = results.data.map(
        (value: Record<string, number | string>, index: number) => {
          // rowToUUID.value.set(index, randomUUID());

          RowMeta.value[index] = {
            status: 'new',
            statusMessage: '',
          }

          if (Object.prototype.hasOwnProperty.call(value, ''))
            delete value['']

          // value["secondary.sex"] = 0;
          // value["secondary.age"] = 0;
          value.index = index
          value.status = 'new'

          return reactifyObject(value)
        },
      )

      results.meta.fields?.forEach((header) => {
        const type = numberFelids.includes(header) ? 'number' : 'string'

        if (FilteredSpecimenKeys.includes(header)) {
          ColMeta.value[header] = {
            header,
            mapping: header,
            type,
          }
        }
        ColMeta.value[header] = {
          header,
          mapping: '',
          type,
        }
      })
    },
    // worker: true,
    header: true,
  })
})

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
  <div>
    <UCard>
      <UCard>
        <p class="font-bold">
          Import Specimens
        </p>
      </UCard>
      <UCard>
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

        <UCard>
          <q-table
            v-model:selected="RowsSelected"
            :columns="rawHeaders"
            :rows="rawList"
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
        </UCard>
        {{ SexOptions }}
        <div
          v-for="(item, index) in SexOptions"
          :key="index"
        >
          {{ item.label.substring(4) }}
          : {{ item.mappings }}
        </div>
        <UContainer class="flex flex-row items-center gap-2">
          <UCard>
            <template #header>
              <p>Secondary Sex Mappings</p>
            </template>
            <div class="flex flex-row items-center gap-2">
              <div
                v-for="item in SexOptions"
                :key="item.label"
                class="flex flex-row items-center gap-1"
              >
                <div>{{ item }} :</div>
                <div>
                  <USelectMenu
                    v-model="SexOptions[item.enumValue].mappings"
                    :multiple="true"
                    :options="sexStrings"
                    class="w-full lg:w-30"
                    placeholder="Select strings to map"
                    searchable
                    searchable-placeholder="Select strings to map"
                  >
                    <template #label>
                      <div class="">
                        <span
                          v-if="sexStrings[item.enumValue]"
                          class="truncate"
                        >{{
                          SexOptions[item.enumValue].mappings.join(", ")
                        }}</span>
                        <span v-else>Select Sex Strings</span>
                      </div>
                    </template>
                  </USelectMenu>
                </div>
              </div>
            </div>
          </UCard>
          <UCard>
            <template #header>
              <p>Secondary Age Mappings</p>
            </template>
            <div class="flex flex-row items-center gap-2">
              <div
                v-for="item in AgeOptions"
                :key="item.value.label"
                class="flex flex-row items-center gap-1"
              >
                <div>{{ item.value.label }} :</div>
                <div>
                  <USelectMenu
                    v-model="AgeOptions[item.value.enumValue].value.mappings"
                    :multiple="true"
                    :options="ageStrings"
                    class="w-full lg:w-30"
                    placeholder="Select strings to map"
                    searchable
                    searchable-placeholder="Select strings to map"
                  >
                    <template #label>
                      <div class="">
                        <span
                          v-if="ageStrings[item.value]"
                          class="truncate"
                        >{{
                          ageMapping[item.value].join(", ")
                        }}</span>
                        <span v-else>Select Age Strings</span>
                      </div>
                    </template>
                  </USelectMenu>
                </div>
              </div>
            </div>
          </UCard>
          <div class="flex flex-row items-center gap-2" />
        </UContainer>
      </UCard>
      <UCard>
        <q-table
          :columns="MappingHeaders"
          :rows="SpecimenFlatList"
          dense
        >
          <template #header-cell="props">
            <q-th :props="props">
              <div class="">
                {{ props.col.label }}
              </div>
              <q-select
                v-model="HeaderMapping[props.col.label]"
                :options="sortedImportHeaders"
                dense
                label="key"
                label-color="teal-10"
                stack-label
              >
                <template
                  v-if="HeaderMapping[props.col.label]"
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
        <q-table
          :columns="MappingHeaders"
          :rows="SpecimenFlatList"
          dense
        >
          <template #header-cell="props">
            <q-th :props="props">
              <div class="">
                {{ props.col.label }}
              </div>

              <q-select
                v-model="HeaderMapping[props.col.label]"
                :options="sortedImportHeaders"
                dense
                label="key"
                label-color="teal-10"
                stack-label
              >
                <template
                  v-if="HeaderMapping[props.col.label]"
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
            </q-th>
          </template>
        </q-table>
      </UCard>

      <pre> {{ FlattenObject(MakeEmptySpecimen()) }}</pre>
      <pre> {{ MakeEmptySpecimen() }}</pre>
    </UCard>
  </div>
</template>

<style></style>
