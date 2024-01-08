<template>
  <div>
    <!-- <Testtable /> -->
    <!-- <SpecimenTable /> -->

    <QCard>
      <q-inner-loading
        :showing="loading"
        label="Please wait..."
        label-class="text-teal"
        label-style="font-size: 1.1em"
      />
      <div>
        <QTable
          :rows="rows"
          :columns="colDef"
          bordered
          class="max-h-max"
          :row-key="(row) => row.specimenId"
          flat
          dense
          title="Specimens"
          separator="horizontal"
          :filter="filter"
          :rows-per-page-options="[25, 30, 50, 100, 200, 500, 1000, 0]"
          :visible-columns="visibleColumns"
        >
          <template #top-right>
            <QBtn
              v-model="filter"
              label="Reload"
              icon="refresh"
              round
              flat
              color="primary"
              class="q-mr-sm"
              @click="reload"
            />
            <q-select
              v-model="visibleColumns"
              class="px-2"
              multiple
              outlined
              dense
              options-dense
              display-value="Visible Columns"
              emit-value
              map-options
              :options="colDef"
              option-value="name"
              options-cover
              style="min-width: 150px"
            />
            <q-input
              v-model="filter"
              borderless
              dense
              debounce="300"
              placeholder="Search"
            >
              <template #append>
                <q-icon name="search" />
              </template>
            </q-input>
          </template>
          <!-- <template #header-cell="props">
            <q-th :props="props">
              <q-icon name="lock_open" size="1.5em" />
              {{ props.col.label }}
            </q-th>
          </template> -->
          <template #body-cell-View="props">
            <q-td :props="props">
              <!-- <QBtn @click="navigate(props.row.specimenId)" /> -->
              <div class="">
                <NuxtLink
                  :to="`/collection/${useRoute().params?.collectionId.toString()}/Specimen/View-${
                    props.row.specimenId
                  }`"
                  :no-prefetch="true"
                >
                  <Icon name="carbon:launch" />
                  Open
                </NuxtLink>
              </div>
            </q-td>
          </template>
        </QTable>
      </div>
    </QCard>
  </div>
</template>

<script lang="ts" setup>
import type { PlainMessage } from "@bufbuild/protobuf";
import { ccbio } from "saacs-es";
import { defu } from "defu";

// const links = useBreadcrumbLinks();
const route = useRoute();

const bookmark = ref("");
const keyedRows = useState<Record<string, PlainMessage<ccbio.Specimen>>>(
  `collectionId${route.params.collectionId}`,
  () => ({}),
);

const rows = ref<PlainMessage<ccbio.Specimen>>();

// const { data, pending, error, refresh } = useFetch(
//   `/api/cc/specimens/list?collectionId=${route.params.collectionId}`,
// );

// function navigate(specimenId: string) {
//   return navigateTo({
//     path: `/collection/${useRoute().params?.collectionId.toString()}/Specimen/View-${specimenId}`,
//   });
// }

const filter = ref("");

const oldcolDefs = ref([
  {
    label: "Full Record",
    name: "View",
    field: "specimenId",
    align: "right",
    required: true,
    headerClasses: "w-4",
  },
  {
    label: "Specimen ID",
    name: "SpecimenID",
    field: "specimenId",
    sortable: true,
  },
  {
    label: "Catalog Number",
    name: "Primary_CatalogNumber",
    align: "left",
    field: (row: PlainMessage<ccbio.Specimen>) =>
      row.primary?.catalogNumber ?? "",
    sortable: true,
  },
  {
    label: "Genus",
    name: "Taxon_Genus",
    align: "left",

    field: (row: PlainMessage<ccbio.Specimen>) => row.taxon?.genus ?? "",
    sortable: true,
  },
  {
    label: "Species",
    name: "Taxon_Species",
    align: "left",

    field: (row: PlainMessage<ccbio.Specimen>) => row.taxon?.species ?? "",
    sortable: true,
  },
  {
    label: "Country",
    align: "left",

    name: "Georeference_Country",
    field: (row: PlainMessage<ccbio.Specimen>) =>
      row.georeference?.country ?? "",
    sortable: true,
  },
  {
    label: "State/Province",
    align: "left",

    name: "Georeference_StateProvince",
    field: (row: PlainMessage<ccbio.Specimen>) =>
      row.georeference?.stateProvince ?? "",
    sortable: true,
  },
  {
    label: "Locality",
    align: "left",

    name: "Georeference_Locality",
    field: (row: PlainMessage<ccbio.Specimen>) =>
      row.georeference?.locality ?? "",
  },
  {
    label: "Field Date",
    name: "Primary_FieldDate",
    align: "left",

    field: (row: PlainMessage<ccbio.Specimen>) =>
      row.primary?.fieldDate?.verbatim ?? "",
    required: true,
  },
]);

const colDef = ref([
  {
    label: "Full Record",
    name: "View",
    field: "specimenId",
    align: "right",
    required: true,
    headerClasses: "w-4",
  },
  {
    label: "Specimen ID",
    name: "SpecimenID",
    field: "specimenId",
    sortable: true,
  },
  {
    name: "Catalog Number",
    label: "Catalog Number",

    field: (row: PlainMessage<ccbio.Specimen>) =>
      row.primary?.catalogNumber ?? "",
    sortable: true,
    align: "left",
  },
  {
    name: "Accession Number",
    label: "Accession Number",
    field: (row: PlainMessage<ccbio.Specimen>) =>
      row.primary?.accessionNumber ?? "",
    sortable: true,
    align: "left",
  },
  {
    name: "Field Number",
    label: "Field Number",
    field: (row: PlainMessage<ccbio.Specimen>) =>
      row.primary?.fieldNumber ?? "",
    sortable: true,
    align: "left",
  },
  {
    name: "tissueNumber",
    label: "Tissue Number",
    field: (row: PlainMessage<ccbio.Specimen>) =>
      row.primary?.tissueNumber ?? "",
    sortable: true,
    align: "left",
  },
  {
    name: "cataloger",
    label: "Cataloger",
    field: (row: PlainMessage<ccbio.Specimen>) => row.primary?.cataloger ?? "",
    sortable: true,
    align: "left",
  },
  {
    name: "collector",
    Label: "Collector",
    field: (row: PlainMessage<ccbio.Specimen>) => row.primary?.collector ?? "",
    sortable: true,
    align: "left",
  },
  {
    name: "determiner",
    label: "Determiner",
    field: (row: PlainMessage<ccbio.Specimen>) => row.primary?.determiner ?? "",
    sortable: true,
    align: "left",
  },
  {
    name: "Field Date",
    label: "Field Date",
    field: (row: PlainMessage<ccbio.Specimen>) =>
      row.primary?.fieldDate?.verbatim ?? "",
    sortable: true,
    align: "left",
  },

  {
    name: "Determined Date",
    field: (row: PlainMessage<ccbio.Specimen>) =>
      row.primary?.determinedDate?.verbatim ?? "",
    sortable: true,
    align: "left",
  },

  {
    name: "determinedReason",
    Label: "Determined Reason",
    field: (row: PlainMessage<ccbio.Specimen>) =>
      row.primary?.determinedReason ?? "",
    sortable: true,
    align: "left",
  },
  {
    name: "originalDate",
    label: "Original Date",
    field: (row: PlainMessage<ccbio.Specimen>) =>
      row.primary?.originalDate?.verbatim ?? "",
    sortable: true,
    align: "left",
  },

  {
    name: "sex",
    Label: "Sex",
    field: (row: PlainMessage<ccbio.Specimen>) => row.secondary?.sex ?? "",
    sortable: true,
    align: "left",
  },
  {
    name: "age",
    label: "Age",
    field: (row: PlainMessage<ccbio.Specimen>) => row.secondary?.age ?? "",
    sortable: true,
    align: "left",
  },
  {
    name: "weight",
    label: "weight",
    field: (row: PlainMessage<ccbio.Specimen>) => row.secondary?.weight ?? "",
    sortable: true,
    align: "left",
  },
  {
    name: "weightUnits",
    label: "Weight Units",
    field: (row: PlainMessage<ccbio.Specimen>) =>
      row.secondary?.weightUnits ?? "",
    sortable: true,
    align: "left",
  },
  {
    name: "condition",
    label: "condition",
    field: (row: PlainMessage<ccbio.Specimen>) =>
      row.secondary?.condition ?? "",
    sortable: true,
    align: "left",
  },
  {
    name: "molt",
    label: "Molt",
    field: (row: PlainMessage<ccbio.Specimen>) => row.secondary?.molt ?? "",
    sortable: true,
    align: "left",
  },
  {
    name: "secondaryNotes",
    label: "Secondary Notes",
    field: (row: PlainMessage<ccbio.Specimen>) => row.secondary?.notes ?? "",
    sortable: true,
    align: "left",
  },

  {
    name: "kingdom",
    label: "Kingdom",
    field: (row: PlainMessage<ccbio.Specimen>) => row.taxon?.kingdom ?? "",
    sortable: true,
    align: "left",
  },
  {
    name: "phylum",
    label: "Phylum",
    field: (row: PlainMessage<ccbio.Specimen>) => row.taxon?.phylum ?? "",
    sortable: true,
    align: "left",
  },
  {
    name: "class",
    label: "Class",
    field: (row: PlainMessage<ccbio.Specimen>) => row.taxon?.class ?? "",
    sortable: true,
    align: "left",
  },
  {
    name: "order",
    label: "Order",
    field: (row: PlainMessage<ccbio.Specimen>) => row.taxon?.order ?? "",
    sortable: true,
    align: "left",
  },
  {
    name: "family",
    label: "Family",
    field: (row: PlainMessage<ccbio.Specimen>) => row.taxon?.family ?? "",
    sortable: true,
    align: "left",
  },
  {
    name: "genus",
    label: "Genus",
    field: (row: PlainMessage<ccbio.Specimen>) => row.taxon?.genus ?? "",
    sortable: true,
    align: "left",
  },
  {
    name: "species",
    label: "Species",
    field: (row: PlainMessage<ccbio.Specimen>) => row.taxon?.species ?? "",
    sortable: true,
    align: "left",
  },
  {
    name: "subspecies",
    label: "Subspecies",
    field: (row: PlainMessage<ccbio.Specimen>) => row.taxon?.subspecies ?? "",
    sortable: true,
    align: "left",
  },

  {
    name: "country",
    label: "Country",
    field: (row: PlainMessage<ccbio.Specimen>) =>
      row.georeference?.country ?? "",
    sortable: true,
    align: "left",
  },
  {
    name: "stateProvince",
    label: "State Province",
    field: (row: PlainMessage<ccbio.Specimen>) =>
      row.georeference?.stateProvince ?? "",
    sortable: true,
    align: "left",
  },
  {
    name: "county",
    label: "County",
    field: (row: PlainMessage<ccbio.Specimen>) =>
      row.georeference?.county ?? "",
    sortable: true,
    align: "left",
  },
  {
    name: "locality",
    label: "Locality",
    field: (row: PlainMessage<ccbio.Specimen>) =>
      row.georeference?.locality ?? "",
    sortable: true,
    align: "left",
  },
  {
    name: "latitude",
    label: "Latitude",
    field: (row: PlainMessage<ccbio.Specimen>) =>
      row.georeference?.latitude ?? "",
    sortable: true,
    align: "left",
  },
  {
    name: "longitude",
    label: "Longitude",
    field: (row: PlainMessage<ccbio.Specimen>) =>
      row.georeference?.longitude ?? "",
    sortable: true,
    align: "left",
  },
  {
    name: "habitat",
    label: "Habitat",
    field: (row: PlainMessage<ccbio.Specimen>) =>
      row.georeference?.habitat ?? "",
    sortable: true,
    align: "left",
  },
  {
    name: "continent",
    label: "Continent",
    field: (row: PlainMessage<ccbio.Specimen>) =>
      row.georeference?.continent ?? "",
    sortable: true,
    align: "left",
  },
  {
    name: "locationRemarks",
    label: "Location Remarks",
    field: (row: PlainMessage<ccbio.Specimen>) =>
      row.georeference?.locationRemarks ?? "",
    sortable: true,
    align: "left",
  },
  {
    name: "coordinateUncertaintyInMeters",
    label: "Coordinate Uncertainty In Meters",
    field: (row: PlainMessage<ccbio.Specimen>) =>
      row.georeference?.coordinateUncertaintyInMeters ?? "",
    sortable: true,
    align: "left",
  },
  {
    name: "georeferenceBy",
    label: "Georeference By",
    field: (row: PlainMessage<ccbio.Specimen>) =>
      row.georeference?.georeferenceBy ?? "",
    sortable: true,
    align: "left",
  },
  {
    name: "GeoreferenceDate",
    label: "Georeference Date",
    field: (row: PlainMessage<ccbio.Specimen>) =>
      row.georeference?.georeferenceDate?.verbatim ?? "",
    sortable: true,
    align: "left",
  },
  {
    name: "georeferenceProtocol",
    label: "Georeference Protocol",
    field: (row: PlainMessage<ccbio.Specimen>) =>
      row.georeference?.georeferenceProtocol ?? "",
    sortable: true,
    align: "left",
  },
  {
    name: "geodeticDatum",
    label: "Geodetic Datum",
    field: (row: PlainMessage<ccbio.Specimen>) =>
      row.georeference?.geodeticDatum ?? "",
    sortable: true,
    align: "left",
  },
  {
    name: "footprintWkt",
    label: "Footprint Wkt",
    field: (row: PlainMessage<ccbio.Specimen>) =>
      row.georeference?.footprintWkt ?? "",
    sortable: true,
    align: "left",
  },
  {
    name: "notes",
    label: "Notes",
    field: (row: PlainMessage<ccbio.Specimen>) => row.georeference?.notes ?? "",
    sortable: true,
    align: "left",
  },
]);

const visibleColumns = ref([
  "View",
  "SpecimenID",
  "Catalog Number",
  "genus",
  "species",
  "Country",
  "stateProvince",
  "locality",
  "Field Date",
]);

const endBookmark = "'\u0000biochain.v1.Specimen\u0000ku_orn\u0000'";

const loading = ref(true);

async function loadBatch() {
  loading.value = true;
  let lastBookmark = bookmark.value;
  while (bookmark.value !== endBookmark) {
    console.log("loadBatch", bookmark.value, lastBookmark);
    lastBookmark = bookmark.value;
    const v = await loadRows();
    if (v.error.value) {
      console.log(v.error);
      break;
    }
    if (lastBookmark === bookmark.value) {
      console.log("no change", bookmark.value, lastBookmark);
      break;
    }
  }
  loading.value = false;
}

callOnce(loadBatch);

async function reload() {
  bookmark.value = "";
  keyedRows.value = {};
  rows.value = [];
  await loadBatch();
}

async function loadRows() {
  console.log("loadRows", route.params.collectionId, bookmark.value);
  return await useCustomFetch<PlainMessage<ccbio.Specimen>[]>(
    `/api/cc/specimens/list`,
    {
      key: `collectionId${route.params.collectionId}-bookmark${bookmark.value}`,
      query: {
        collectionId: route.params.collectionId,
        bookmark: bookmark.value,
      },
      onResponse: async ({ response }) => {
        console.log("history", response._data);
        bookmark.value = response._data?.bookmark ?? "";
        console.log(bookmark.value);
        keyedRows.value = defu(keyedRows.value, response._data?.specimenMap);
        rows.value = Object.values(keyedRows.value);
      },
    },
  );
}

// const { data, pending, error, refresh, execute } = await useCustomFetch<
//   PlainMessage<ccbio.Specimen>[]
// >(`/api/cc/specimens/list`, {
//   query: {
//     collectionId: route.params.collectionId,
//     bookmark: bookmark.value,
//   },
//   onResponse: async ({ response }) => {
//     console.log("history", response._data);
//     bookmark.value = response._data?.bookmark ?? "";
//     console.log(bookmark.value);
//     keyedRows.value = defu(keyedRows.value, response._data?.specimenMap);
//     console.log(keyedRows.value);
//     rows.value = Object.values(keyedRows.value);
//     console.log(rows.value);
//   },
// });
</script>

<style>
/* td:last-child {
  text-align: right;
  color: red;
  background: #eee;
  position: sticky;
  z-index: 1;
  right: 0;
} */
/* th:last-child {
  background: #eee;
  position: sticky;
  z-index: 1;
  right: 0;
} */
</style>
