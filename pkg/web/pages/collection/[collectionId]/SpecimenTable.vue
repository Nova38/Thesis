<template>
  <div>
    <!-- <Testtable /> -->
    <!-- <SpecimenTable /> -->

    <QCard>
      <ag-grid-vue
        style="height: 500px"
        class="ag-theme-quartz"
        :modules="modules"
        :column-defs="agColDef"
        :row-data="store.SpecimenList"
        :auto-size-strategy="{ type: 'fitCellContents' }"
        :framework-components="{ tableButton: TableButton }"
      ></ag-grid-vue>
    </QCard>

    <QCard>
      <q-inner-loading
        :showing="store.Loading"
        label="Please wait..."
        label-class="text-teal"
        label-style="font-size: 1.1em"
      />
      <div>
        <QTable
          :rows="store.SpecimenList"
          :columns="colDef"
          bordered
          class="max-h-max"
          :row-key="(row: PlainMessage<ccbio.Specimen>) => row.specimenId"
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
              @click="store.Reload"
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
import type { QTableProps } from "nuxt-quasar-ui/dist/runtime/adapter";
import { ccbio } from "saacs-es";

import TableButton from "@/components/collection/TableButton.vue";

import { AgGridVue } from "@ag-grid-community/vue3";
import { ClientSideRowModelModule } from "@ag-grid-community/client-side-row-model";
import "@ag-grid-community/styles/ag-grid.css";
import "@ag-grid-community/styles/ag-theme-quartz.css";
import type { c } from "saacs-es/dist/shared/saacs-es.1e75680d";

const store = useCollectionsStore();

callOnce(() => {
  store.Reload();
});

// const { data, pending, error, refresh } = useFetch(
//   `/api/cc/specimens/list?collectionId=${route.params.collectionId}`,
// );

// function navigate(specimenId: string) {
//   return navigateTo({
//     path: `/collection/${useRoute().params?.collectionId.toString()}/Specimen/View-${specimenId}`,
//   });
// }

const filter = ref("");

const colDef = ref<QTableProps["columns"]>([
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
    label: "Collector",
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
    label: "Determined Date",
    field: (row: PlainMessage<ccbio.Specimen>) =>
      row.primary?.determinedDate?.verbatim ?? "",
    sortable: true,
    align: "left",
  },

  {
    name: "determinedReason",
    label: "Determined Reason",
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
    label: "Sex",
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

const modules = [ClientSideRowModelModule];

const agColDef = [
  {
    headerName: "Specimen ID",
    name: "SpecimenID",
    field: "specimenId",
    sortable: true,
    pin: "left",
    filter: true,
    cellRenderer: TableButton,
  },
  {
    headerName: "Primary",
    children: [
      {
        headerName: "Catalog Number",
        filter: true,
        field: "primary.catalogNumber",
        sortable: true,
        align: "left",
      },
      {
        name: "Accession Number",
        headerName: "Accession Number",
        field: "primary.accessionNumber",
        sortable: true,
        align: "left",
      },
      {
        name: "Field Number",
        headerName: "Field Number",
        field: "primary?.fieldNumber",
        sortable: true,
        align: "left",
      },
      {
        name: "tissueNumber",
        headerName: "Tissue Number",
        field: "primary?.tissueNumber",
        sortable: true,
        align: "left",
      },
      {
        name: "cataloger",
        headerName: "Cataloger",
        field: "primary?.cataloger",

        sortable: true,
        align: "left",
        columnGroupShow: "open",
      },
      {
        name: "collector",
        headerName: "Collector",
        field: "primary?.collector",
        sortable: true,
        align: "left",
        columnGroupShow: "open",
      },
      {
        name: "determiner",
        headerName: "Determiner",
        field: "primary?.determiner",
        sortable: true,
        align: "left",
        columnGroupShow: "open",
      },
      {
        name: "Field Date",
        headerName: "Field Date",
        field: "primary?.fieldDate?.verbatim",
        sortable: true,
        align: "left",
        columnGroupShow: "open",
      },

      {
        name: "Determined Date",
        headerName: "Determined Date",
        field: "primary?.determinedDate?.verbatim",
        sortable: true,
        align: "left",
        columnGroupShow: "open",
      },

      {
        name: "determinedReason",
        headerName: "Determined Reason",
        field: "primary?.determinedReason",
        sortable: true,
        align: "left",
        columnGroupShow: "open",
      },
      {
        name: "originalDate",
        headerName: "Original Date",
        field: "primary?.originalDate?.verbatim",
        sortable: true,
        align: "left",
        columnGroupShow: "open",
      },
    ],
  },
  {
    headerName: "Taxon",
    children: [
      {
        name: "kingdom",
        headerName: "Kingdom",
        field: "taxon?.kingdom",
        sortable: true,
        align: "left",
      },
      {
        name: "phylum",
        headerName: "Phylum",
        field: "taxon?.phylum",
        sortable: true,
        align: "left",
      },
      {
        name: "class",
        headerName: "Class",
        field: "taxon?.class",
        sortable: true,
        align: "left",
      },
      {
        name: "order",
        headerName: "Order",
        field: "taxon?.order",
        sortable: true,
        align: "left",
      },
      {
        name: "family",
        headerName: "Family",
        field: "taxon?.family",
        sortable: true,
        align: "left",
      },
      {
        name: "genus",
        headerName: "Genus",
        field: "taxon?.genus",
        sortable: true,
        align: "left",
      },
      {
        name: "species",
        headerName: "Species",
        field: "taxon?.species",
        sortable: true,
        align: "left",
      },
      {
        name: "subspecies",
        headerName: "Subspecies",
        field: "taxon?.subspecies",
        sortable: true,
        align: "left",
      },
    ],
  },
  {
    headerName: "Secondary",
    children: [
      {
        name: "sex",
        headerName: "Sex",
        field: "secondary?.sex",
        sortable: true,
        align: "left",
      },
      {
        name: "age",
        headerName: "Age",
        field: "secondary?.age",
        sortable: true,
        align: "left",
      },
      {
        name: "weight",
        headerName: "weight",
        field: "secondary?.weight",
        sortable: true,
        align: "left",
      },
      {
        name: "weightUnits",
        headerName: "Weight Units",
        field: "secondary?.weightUnits",
        sortable: true,
        align: "left",
      },
      {
        name: "condition",
        headerName: "condition",
        field: "secondary?.condition",
        sortable: true,
        align: "left",
      },
      {
        name: "molt",
        headerName: "Molt",
        field: "secondary?.molt",
        sortable: true,
        align: "left",
      },
      {
        name: "secondaryNotes",
        headerName: "Secondary Notes",
        field: "secondary?.notes",
        sortable: true,
        align: "left",
      },
    ],
  },
  {
    headerName: "Georeference",
    children: [
      {
        name: "country",
        headerName: "Country",
        field: "georeference?.country",
        sortable: true,
        align: "left",
      },
      {
        name: "stateProvince",
        headerName: "State Province",
        field: "georeference?.stateProvince",
        sortable: true,
        align: "left",
      },
      {
        name: "county",
        headerName: "County",
        field: "georeference?.county",
        sortable: true,
        align: "left",
      },
      {
        name: "locality",
        headerName: "Locality",
        field: "georeference?.locality",
        sortable: true,
        align: "left",
      },
      {
        name: "latitude",
        headerName: "Latitude",
        field: "georeference?.latitude",
        sortable: true,
        align: "left",
      },
      {
        name: "longitude",
        headerName: "Longitude",
        field: "georeference?.longitude",
        sortable: true,
        align: "left",
      },
      {
        name: "habitat",
        headerName: "Habitat",
        field: "georeference?.habitat",
        sortable: true,
        align: "left",
      },
      {
        name: "continent",
        headerName: "Continent",
        field: "georeference?.continent",
        sortable: true,
        align: "left",
      },
      {
        name: "locationRemarks",
        headerName: "Location Remarks",
        field: "georeference?.locationRemarks",
        sortable: true,
        align: "left",
      },
      {
        name: "coordinateUncertaintyInMeters",
        headerName: "Coordinate Uncertainty In Meters",
        field: "georeference?.coordinateUncertaintyInMeters",
        sortable: true,
        align: "left",
      },
      {
        name: "georeferenceBy",
        headerName: "Georeference By",
        field: "georeference?.georeferenceBy",
        sortable: true,
        align: "left",
      },
      {
        name: "GeoreferenceDate",
        headerName: "Georeference Date",
        field: "georeference?.georeferenceDate?.verbatim",
        sortable: true,
        align: "left",
      },
      {
        name: "georeferenceProtocol",
        headerName: "Georeference Protocol",
        field: "georeference?.georeferenceProtocol",
        sortable: true,
        align: "left",
      },
      {
        name: "geodeticDatum",
        headerName: "Geodetic Datum",
        field: "georeference?.geodeticDatum",
        sortable: true,
        align: "left",
      },
      {
        name: "footprintWkt",
        headerName: "Footprint Wkt",
        field: "georeference?.footprintWkt",
        sortable: true,
        align: "left",
      },
      {
        name: "notes",
        headerName: "Notes",
        field: "georeference?.notes",
        sortable: true,
        align: "left",
      },
    ],
  },
];
</script>

<style>
.ag-theme-quartz {
  --ag-grid-size: 5px;
  --ag-list-item-height: 20px;
}
</style>
