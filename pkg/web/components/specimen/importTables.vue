<template>
  <div>
    <UCard>
      <UCard>
        <p class="font-bold">Import Specimens</p>
      </UCard>
      <UCard>
        <h2>Select CSV file to import from</h2>
        <q-file v-model="file" outlined accept=".csv">
          <template #prepend>
            <q-icon name="attach_file" />
          </template>
        </q-file>

        <UCard>
          <q-table
            v-model:selected="RowsSelected"
            dense
            :rows="rawList"
            :columns="rawHeaders"
            row-key="index"
            selection="multiple"
          >
            <template #body-cell-status="props">
              <q-td :props="props">
                <UPopover mode="hover" :popper="{ adaptive: true }">
                  <UBadge
                    :label="props.row[props.col.field]"
                    :color="statusToChipColor(props.row[props.col.field])"
                  />
                  <q-circular-progress
                    v-if="props.row[props.col.field] == 'loading'"
                    indeterminate
                    rounded
                    size="15px"
                    color="warn"
                  />
                  <template
                    v-if="RowMeta[props.row.index].statusMessage != ''"
                    #panel
                  >
                    <div class="p-4">
                      <pre wrap>
                    {{ RowMeta[props.row.index].statusMessage }}</pre
                      >
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
        <div v-for="(item, index) in SexOptions" :key="index">
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
                    searchable
                    searchable-placeholder="Select strings to map"
                    class="w-full lg:w-30"
                    :multiple="true"
                    placeholder="Select strings to map"
                    :options="sexStrings"
                  >
                    <template #label>
                      <div class="">
                        <span
                          v-if="sexStrings[item.enumValue]"
                          class="truncate"
                          >{{
                            SexOptions[item.enumValue].mappings.join(", ")
                          }}</span
                        >
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
                    searchable
                    searchable-placeholder="Select strings to map"
                    class="w-full lg:w-30"
                    :multiple="true"
                    placeholder="Select strings to map"
                    :options="ageStrings"
                  >
                    <template #label>
                      <div class="">
                        <span v-if="ageStrings[item.value]" class="truncate">{{
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
          <div class="flex flex-row items-center gap-2"></div>
        </UContainer>
      </UCard>
      <UCard>
        <q-table dense :rows="SpecimenFlatList" :columns="MappingHeaders">
          <template #header-cell="props">
            <q-th :props="props">
              <div class="">
                {{ props.col.label }}
              </div>
              <q-select
                v-model="HeaderMapping[props.col.label]"
                label-color="teal-10"
                label="key"
                stack-label
                dense
                :options="sortedImportHeaders"
              >
                <template v-if="HeaderMapping[props.col.label]" #append>
                  <q-icon
                    name="cancel"
                    dense
                    size=".75em"
                    color="red"
                    class="cursor-pointer"
                    @click.stop.prevent="clearKey(props.col.label)"
                  />
                </template>
              </q-select>
              <!-- <div v-if="props.col.label">hi</div> -->
            </q-th>
          </template>
        </q-table>
        <q-table dense :rows="SpecimenFlatList" :columns="MappingHeaders">
          <template #header-cell="props">
            <q-th :props="props">
              <div class="">
                {{ props.col.label }}
              </div>

              <q-select
                v-model="HeaderMapping[props.col.label]"
                label-color="teal-10"
                label="key"
                stack-label
                dense
                :options="sortedImportHeaders"
              >
                <template v-if="HeaderMapping[props.col.label]" #append>
                  <q-icon
                    name="cancel"
                    dense
                    size=".75em"
                    color="red"
                    class="cursor-pointer"
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

<script lang="ts" setup>
import type {
  QTableProps,
  QTreeNode,
} from "nuxt-quasar-ui/dist/runtime/adapter";
import type { ParseResult } from "papaparse";
import Papa from "papaparse";
import { crush, get, set, keys } from "radash";
import { ccbio } from "saacs-es";
import { FlattedSpecimenKeys } from "~/utils/flatten";
type status = "new" | "importing" | "success" | "error" | "pre-existing";
function clearKey(key: string) {
  console.log(key);
  console.log(HeaderMapping[key]);
  HeaderMapping[key] = "";
}
const nodes: Ref<QTreeNode[]> = ref([
  {
    label: "secondary",
    children: [
      { label: "sex" },
      { label: "age" },
      { label: "weight" },
      { label: "weightUnits" },
      { label: "preparations" },
      { label: "condition" },
      { label: "molt" },
      { label: "notes" },
    ],
  },
]);

interface IRowMeta {
  // index: string;
  status: status;
  collectionId?: string;
  uuid?: string;
  statusMessage?: string;
}

interface IColMeta {
  header: string;
  mapping?: string;
  type: string;
}

const SexOptions = ref([
  { label: "SEX_UNKNOWN", enumValue: 1, mappings: [] },
  { label: "SEX_ATYPICAL", enumValue: 2, mappings: [] },
  { label: "SEX_MALE", enumValue: 3, mappings: [] },
  { label: "SEX_FEMALE", enumValue: 4, mappings: [] },
]);

const AgeOptions = ref([
  { label: "AGE_UNKNOWN", enumValue: 1, mappings: [] },
  { label: "AGE_NEST", enumValue: 2, mappings: [] },
  { label: "AGE_EMBRYO_EGG", enumValue: 3, mappings: [] },
  { label: "AGE_CHICK_SUBADULT", enumValue: 4, mappings: [] },
  { label: "AGE_ADULT", enumValue: 5, mappings: [] },
  { label: "AGE_CONTINGENT", enumValue: 6, mappings: [] },
]);

const numberFelids = [
  "geography.latitude",
  "geography.longitude",
  "secondary.wight",
  "secondary.sex",
  "secondary.age",
];

const columns = [
  {
    key: "index",
    label: "index",
  },
  {
    key: "primary.catalogNumber",
    label: "catalog Number",
  },
  {
    key: "title",
    label: "Job position",
  },
  {
    key: "email",
    label: "Email",
  },
  {
    key: "role",
  },
];

const ColMeta = ref<Record<string, IColMeta>>({});
// const ImportHeaders = ref<string[]>([]);
const sortedImportHeaders = computed(() => {
  return Object.keys(ColMeta.value).sort();
});
const HeaderMapping = ref<Record<string, string>>({});

const keysForImport = crush(
  JSON.parse(MakeEmptySpecimen().toJsonString({ emitDefaultValues: true })),
);
console.log(keysForImport);

const SpecimenKeys: Array<string> = keys(keysForImport);

const FilteredSpecimenKeys = SpecimenKeys.filter(
  (k) => !["id", "last_modified_by", "collection_id"].includes(k),
);

const rawHeaders: Ref<QTableProps["columns"]> = ref();
const rawList: Ref<Record<string, string | number>[]> = ref([]);
const RowMeta = ref<IRowMeta[]>([]);

const sexStrings = ref<string[]>([]);
const ageStrings = ref<string[]>([]);

const RowsSelected = ref<Record<string, string>[]>([]);

const SpecimenFlatList = useArrayMap(RowsSelected, (item) => {
  const s = FlattenEmptySpecimen();

  for (const key of Object.keys(item)) {
    console.log(key);
    console.log(ColMeta.value[key]);

    if (ColMeta.value[key]) {
      console.log(ColMeta.value[key]);
      s[key] = item[key];
    }
  }
  return s;
});

const file = ref(null);

const makeHeaders = () => {
  const flat: any = [];

  for (const key of FlattedSpecimenKeys()) {
    flat.push({
      label: key,
      key: key,
    });
  }
  return flat;
};

const MappingHeaders = makeHeaders();

/*
  ┌─────────────────────────────────────────────────────────────────────────┐
  │ Load the file on change and parse it with papaparse                     │
  └─────────────────────────────────────────────────────────────────────────┘
 */

// const numberFeilds= ["primary."]

watch(file, (file) => {
  if (!file) {
    return;
  }

  Papa.parse(file, {
    // worker: true,
    header: true,
    complete: (results: ParseResult<Record<string, string>>) => {
      console.log(results);

      rawHeaders.value = results.meta.fields?.map((item) => {
        return {
          label: item,
          name: item,
          field: item,
        };
      });
      rawHeaders.value?.unshift({
        label: "status",
        name: "status",
        field: "status",
      });

      rawList.value = results.data.map(
        (value: Record<string, string | number>, index: number) => {
          // rowToUUID.value.set(index, randomUUID());

          RowMeta.value[index] = {
            status: "new",
            statusMessage: "",
          };

          if (Object.prototype.hasOwnProperty.call(value, "")) {
            delete value[""];
          }
          // value["secondary.sex"] = 0;
          // value["secondary.age"] = 0;
          value["index"] = index;
          value["status"] = "new";

          return reactifyObject(value);
        },
      );

      results.meta.fields?.forEach((header) => {
        const type = numberFelids.includes(header) ? "number" : "string";

        if (FilteredSpecimenKeys.includes(header)) {
          ColMeta.value[header] = {
            header,
            mapping: header,
            type: type,
          };
        }
        ColMeta.value[header] = {
          header,
          mapping: "",
          type: type,
        };
      });
    },
  });
});

function statusToChipColor(status: status) {
  switch (status) {
    case "new":
      return "blue";
    case "pre-existing":
      return "pink";
    case "importing":
      return "orange";
    case "success":
      return "green";
    case "error":
      return "red";
    default:
      return "blue";
  }
}
</script>

<style></style>
