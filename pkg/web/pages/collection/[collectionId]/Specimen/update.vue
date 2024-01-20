<template>
  <q-page class="full-height">
    <q-card class="q-pa-md q-ma-md">
      <QCardSection>
        <p class="font-bold">Bulk Update Specimens</p>
      </QCardSection>

      <CsvFile :csv="csv" />

      <QCardSection class="flex flex-row items-center gap-2 justify-center">
        <QTable dense :hide-bottom="true" :rows="SexOptions" />
        <QTable
          dense
          :hide-bottom="true"
          :pagination="{ rowsPerPage: 0 }"
          :rows="AgeOptions"
        />
      </QCardSection>
      <q-card-section>
        <q-table
          v-model:selected="RowsSelected"
          dense
          :rows="rawData"
          row-key="index"
          selection="multiple"
        >
          <template #body-cell-exist="props">
            <QTd :props="props">
              <UBadge
                :label="RowMeta[props.rowIndex].exist"
                :color="existToChipColor(RowMeta[props.rowIndex].exist)"
              />
            </QTd>
          </template>
          <template #body-cell-status="props">
            <q-td :props="props">
              <UPopover mode="hover" :popper="{ adaptive: true }">
                <UBadge
                  :label="RowMeta[props.rowIndex].status"
                  :color="statusToChipColor(RowMeta[props.rowIndex].status)"
                />
                <q-circular-progress
                  v-if="RowMeta[props.rowIndex].status == 'loading'"
                  indeterminate
                  rounded
                  size="15px"
                  color="warn"
                />
                <template
                  v-if="RowMeta[props.rowIndex].statusMessage != ''"
                  #panel
                >
                  <div class="p-4">
                    <pre wrap>
                    {{ RowMeta[props.rowIndex].statusMessage }}</pre
                    >
                  </div>
                </template>
              </UPopover>
            </q-td>
          </template>
        </q-table>
      </q-card-section>

      <UCard>
        <!-- <template #header> </template> -->
        <div class="text-xl font-semibold border-blue-400 border-solid mb-2">
          Current Selected Specimen Values
        </div>
        <QTable dense :rows="selectedCurrent" :columns="MappingHeaders">
        </QTable>
      </UCard>

      <UCard>
        <div class="text-xl font-semibold border-blue-400 border-solid mb-2">
          Imported Specimen Values
        </div>
        <div v-if="possessedData">
          <q-table dense :rows="possessedData" :columns="MappingHeaders">
            <template #header-cell="props">
              <q-th :props="props">
                <div class="">
                  {{ props.col.label }}
                </div>
                <q-select
                  v-model="specimenMapping[props.col.label]"
                  label-color="teal-10"
                  label="key"
                  stack-label
                  dense
                  :options="sortedImportHeaders"
                >
                  <template v-if="specimenMapping[props.col.label]" #append>
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
        </div>
      </UCard>
      <q-card-section>
        <q-btn
          color="secondary"
          label="Upload Selected"
          class="full-width"
          @click="run()"
        />
      </q-card-section>
    </q-card>
  </q-page>
</template>

<script setup lang="ts">
import type { ParseResult } from "papaparse";
import Papa from "papaparse";
import { crush, get, set, keys } from "radash";
import { ccbio } from "saacs-es";
import {
  createRegistry,
  Timestamp,
  type PlainMessage,
} from "@bufbuild/protobuf";

type ExistStatus = "new" | "pre-existing";
type status = "new" | "loading" | "success" | "error";

const csv: Ref<Record<string, string>[]> = ref([]);

export interface RowMeta {
  uuid: string;
  catalogNumber?: string;
  status: status;
  exist: ExistStatus;
  statusMessage?: string;
  currentSpecimen?: ccbio.Specimen;
  differences?: Record<string, any>;
}
const SexOptions = [
  { label: "SEX_UNKNOWN", value: 1 },
  { label: "SEX_ATYPICAL", value: 2 },
  { label: "SEX_MALE", value: 3 },
  { label: "SEX_FEMALE", value: 4 },
];

const AgeOptions = [
  { label: "AGE_UNKNOWN", value: 1 },
  { label: "AGE_NEST", value: 2 },
  { label: "AGE_EMBRYO_EGG", value: 3 },
  { label: "AGE_CHICK_SUBADULT", value: 4 },
  { label: "AGE_ADULT", value: 5 },
  { label: "AGE_CONTINGENT", value: 6 },
];
const numberFelids = [
  "geography.latitude",
  "geography.longitude",
  "secondary.wight",
  "secondary.sex",
  "secondary.age",
];

callOnce(() => {
  useCollectionsStore().LoadRows();
});

// Table 1
const headers = ref<string[]>([]);
const rawData = ref<Record<string, string>[]>();
const RowMeta = ref<RowMeta[]>([]);
const RowsSelected = ref<Record<string, string>[]>([]);
const makeNew = ref<boolean>(true);

// This is a mapping between the specimen keys and the column headers
// <column header>: <specimen key>
const specimenMapping = ref<Record<string, string>>({});

/*
  ┌─────────────────────────────────────────────────────────────────────────┐
  │ Process The RowMeta on specimenMapping Update                           │
  └─────────────────────────────────────────────────────────────────────────┘
 */
watch([specimenMapping.value, RowsSelected], () => {
  console.log("SpecimenMapping Updated", specimenMapping.value);

  const catalogNumKey = specimenMapping.value["primary.catalogNumber"];
  RowsSelected.value?.forEach((data) => {
    const catalogNum = data[catalogNumKey];
    const index = parseInt(data.index);
    console.log({ catalogNum, index });

    RowMeta.value[index].exist = "new";

    RowMeta.value[index].status = "new";
    RowMeta.value[index].statusMessage = "Specimen hasn't been uploaded";

    if (catalogNum) {
      RowMeta.value[index].catalogNumber = catalogNum;
      const specimen =
        useCollectionsStore().GetSpecimenFromCatalogNumber(catalogNum);
      console.log({ specimen });

      if (specimen) {
        RowMeta.value[index].currentSpecimen = specimen;
        RowMeta.value[index].exist = "pre-existing";
      }
    }
  });
});

function clearKey(key: string) {
  specimenMapping.value[key] = "";
}

// Table 2
const possessedData = useArrayMap(RowsSelected, (data) => {
  const newSpecimen = MakeEmptySpecimen();

  for (const key in specimenMapping.value) {
    if (Object.prototype.hasOwnProperty.call(specimenMapping.value, key)) {
      if (key === "primary.catalogNumber") {
        // continue;
      }
      set(newSpecimen, key, data[specimenMapping.value[key]]);
    }
  }
  return { index, specimen: newSpecimen };
});

const selectedCurrent = useArrayMap(RowsSelected, (data) => {
  return RowMeta.value[data.index].currentSpecimen;
});

const sortedImportHeaders = computed(() => {
  if (!headers.value) {
    return;
  }
  return headers.value.sort();
});

const z = JSON.parse(
  MakeEmptySpecimen().toJsonString({ emitDefaultValues: true }),
);

// Set up the mapping objects

const notAllowedKeys = ["specimenId", "lastModifiedBy", "collectionId"];

const FilteredSpecimenKeys = keys(crush(z)).filter(
  (k) => !notAllowedKeys.includes(k),
);

callOnce(() => {
  useCollectionsStore().LoadRows();
});
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
      headers.value = results.meta.fields || [];
      headers.value.unshift("exists");
      headers.value.unshift("status");

      // If the data's header row that contains the value of a specimen key map it to the specimen key to start
      for (const key of FilteredSpecimenKeys) {
        if (headers.value.includes(key)) {
          specimenMapping.value[key] = key;

          if (key === "primary.catalogNumber") {
            specimenMapping.value[key] = "primary.catalogNumber";
          }
        }
      }

      rawData.value = results.data.map((value: any, index: any) => {
        RowMeta.value[index] = {
          uuid: "",
          status: "new",
          exist: "new",
          statusMessage: "",
        };

        if (Object.prototype.hasOwnProperty.call(value, "")) {
          delete value[""];
        }

        const catNum: string | undefined = value["primary.catalogNumber"];
        if (catNum !== undefined) {
          // console.log(CatNumToUUID(catNum));
          value["specimenId"] = CatNumToUUID(catNum);
          RowMeta.value[index].uuid = CatNumToUUID(catNum);
        }

        if (
          catNum &&
          useCollectionsStore().SpecimenCatalogNumbers?.includes(catNum)
        ) {
          // console.log("Skipping row as it has already been uploaded", row);
          console.log("pre-existing", catNum);

          RowMeta.value[index].exist = "pre-existing";
          RowMeta.value[index].uuid = CatNumToUUID(catNum);
          return reactifyObject(
            Object.assign(
              {
                specimenId: CatNumToUUID(catNum),
                collectionId: useRoute().params.collectionId.toString(),
                status: "new",
                exist: "pre-existing",
              },
              value,
            ),
          );
        } else {
          RowMeta.value[index].exist = "new";
          RowMeta.value[index].status = "new";

          RowMeta.value[index].statusMessage =
            "Specimen Number with Catalog Number doesn't exist";
          return reactifyObject(
            Object.assign(
              {
                collectionId: useRoute().params.collectionId.toString(),
                status: "new",
                exist: "new",
              },
              value,
            ),
          );
        }
      });
      console.log("Imported Data");
      console.table(RowMeta.value);
    },
  });
});

/*
  ┌─────────────────────────────────────────────────────────────────────────┐
  │    Map the headers to the correct specimen fields                       │
  └─────────────────────────────────────────────────────────────────────────┘
 */
// const dataMap = ref(new Map<number, ccbio.Specimen>());

const makeHeaders = () => {
  const flat: any = [];

  for (const key of FilteredSpecimenKeys) {
    flat.push({
      label: key,
      field: (row: any) => get(row, key),
    });
  }
  return flat;
};

const MappingHeaders = makeHeaders();

async function run() {
  // console.log(rawData.value);
  console.log({ possessedData });

  // const uploadingArray: Promise[] = [];

  // TODO: Make sure they haven't already been uploaded by checking the catalog number
  possessedData.value.forEach(async (value) => {
    // const index = RowMeta.value.findIndex(
    //   (e) => e.catalogNumber === value.primary?.catalogNumber,
    // );
    const index = value.index;

    console.log({ index, value });

    if (!value?.specimen.primary?.catalogNumber) {
      console.log("Skipping row as it has no catalog number", value);
      RowMeta.value[index].status = "error";
      RowMeta.value[index].statusMessage = "Row has no catalog number";
      return;
    }
    value.specimen.specimenId = CatNumToUUID(
      value.specimen?.primary?.catalogNumber,
    );
    value.specimen.collectionId = useRoute().params.collectionId.toString();
    console.log({ value });

    try {
      console.log(RowMeta.value[index]);

      if (RowMeta.value[index].status === "success") {
        console.debug("Skipping row as it has already been uploaded", value);
        RowMeta.value[index].differences = {};
        RowMeta.value[index].statusMessage = "Row has already been uploaded";
        return;
      } else if (RowMeta.value[index].exist === "new" && makeNew.value) {
        RowMeta.value[index].status = "loading";
        const specimen = PrepareRow(value, index);
        const meta = await SpecimenUpdate(specimen, new ccbio.Specimen());

        RowMeta.value[index].status = meta.status;
        RowMeta.value[index].differences = meta.differences;
        RowMeta.value[index].statusMessage = meta.statusMessage;
        console.log(RowMeta);
        return;
      } else if (RowMeta.value[index].exist === "pre-existing") {
        const cur = RowMeta.value[index].currentSpecimen;
        value.specimenId =
          cur?.specimenId || CatNumToUUID(value.specimen.primary.catalogNumber);

        if (!isDefined(cur)) {
          throw new Error(
            "Row Claims to be Preexisting but has no current specimen",
          );
        }
        RowMeta.value[index].status = "loading";
        const specimen = PrepareRow(value.specimen, index);
        const meta = await SpecimenUpdate(specimen, new ccbio.Specimen(cur));

        RowMeta.value[index].status = meta.status;
        RowMeta.value[index].differences = meta.differences;
        RowMeta.value[index].statusMessage = meta.statusMessage;
        return;
      }
    } catch (err: any) {
      console.log(err);
      RowMeta.value[index].status = "error";
      RowMeta.value[index].statusMessage = err?.message;
    }
  });

  // console.log(possessedData.value);
}

function PrepareRow(specimen: PlainSpecimen, index: number) {
  // specimen.specimenId = CatNumToUUID(value["primary.catalogNumber"]);
  // RowMeta.value[index].uuid = specimen.specimenId;
  // specimen.collectionId = useRoute().params.collectionId.toString();
  console.log({ specimen });

  numberFelids.forEach((e) => {
    const val: string = get(specimen, e);
    console.log(val);
    if (isNaN(parseFloat(val))) {
      console.log("isNan");
      set(specimen, e, 0);
    } else {
      set(specimen, e, parseFloat(val));
    }
  });

  return specimen;
}

async function SpecimenUpdate(
  specimen: PlainMessage<ccbio.Specimen>,
  curSpecimen: PlainMessage<ccbio.Specimen>,
): Promise<RowMeta> {
  console.log({ specimen, curSpecimen });

  const { differences, mask } = diffCrush(
    toValue(specimen),
    toValue(curSpecimen),
    ["lastModified"],
  );

  const req = new ccbio.SpecimenUpdate({
    specimen: new ccbio.Specimen(Specimen.parse(specimen)),
    mask,
  });

  try {
    const res = await useCustomFetch(`/api/cc/specimens/update`, {
      method: "POST",
      body: req.toJsonString({
        typeRegistry: createRegistry(ccbio.Specimen, Timestamp),
      }),
    });
    console.log(res);
    return {
      uuid: specimen.specimenId,
      status: "success",
      exist: "new",
      statusMessage: "Uploaded Successfully",
      differences,
    };
  } catch (err: any) {
    return {
      uuid: specimen.specimenId,
      status: "error",
      exist: "new",
      statusMessage: err.message,
      differences,
    };
  }
}
function existToChipColor(exists: ExistStatus) {
  switch (exists) {
    case "new":
      return "purple";
    case "pre-existing":
      return "pink";
  }
}
function statusToChipColor(status: status) {
  switch (status) {
    case "new":
      return "blue";
    case "loading":
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

<style scoped></style>
