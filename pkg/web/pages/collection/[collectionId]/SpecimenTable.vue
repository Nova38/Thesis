<script lang="ts" setup>
import type { PlainMessage } from '@bufbuild/protobuf'
import type { QTableProps } from 'nuxt-quasar-ui/dist/runtime/adapter'
import type { ccbio } from 'saacs-es'

import TableButton from '@/components/collection/TableButton.vue'
import { ClientSideRowModelModule } from '@ag-grid-community/client-side-row-model'
import '@ag-grid-community/styles/ag-grid.css'
import '@ag-grid-community/styles/ag-theme-quartz.css'
import { AgGridVue } from '@ag-grid-community/vue3'

const store = useCollectionsStore()

callOnce(() => {
  store.Reload()
})

// const { data, pending, error, refresh } = useFetch(
//   `/api/cc/specimens/list?collectionId=${route.params.collectionId}`,
// );

// function navigate(specimenId: string) {
//   return navigateTo({
//     path: `/collection/${useRoute().params?.collectionId.toString()}/Specimen/View-${specimenId}`,
//   });
// }

const filter = ref('')

const colDef = ref<QTableProps['columns']>([
  {
    align: 'right',
    field: 'specimenId',
    headerClasses: 'w-4',
    label: 'Full Record',
    name: 'View',
    required: true,
  },
  {
    field: 'specimenId',
    label: 'Specimen ID',
    name: 'SpecimenID',
    sortable: true,
  },
  {
    align: 'left',
    field: (row: PlainMessage<ccbio.Specimen>) =>
      row.primary?.catalogNumber ?? '',

    label: 'Catalog Number',
    name: 'Catalog Number',
    sortable: true,
  },
  {
    align: 'left',
    field: (row: PlainMessage<ccbio.Specimen>) =>
      row.primary?.accessionNumber ?? '',
    label: 'Accession Number',
    name: 'Accession Number',
    sortable: true,
  },
  {
    align: 'left',
    field: (row: PlainMessage<ccbio.Specimen>) =>
      row.primary?.fieldNumber ?? '',
    label: 'Field Number',
    name: 'Field Number',
    sortable: true,
  },
  {
    align: 'left',
    field: (row: PlainMessage<ccbio.Specimen>) =>
      row.primary?.tissueNumber ?? '',
    label: 'Tissue Number',
    name: 'tissueNumber',
    sortable: true,
  },
  {
    align: 'left',
    field: (row: PlainMessage<ccbio.Specimen>) => row.primary?.cataloger ?? '',
    label: 'Cataloger',
    name: 'cataloger',
    sortable: true,
  },
  {
    align: 'left',
    field: (row: PlainMessage<ccbio.Specimen>) => row.primary?.collector ?? '',
    label: 'Collector',
    name: 'collector',
    sortable: true,
  },
  {
    align: 'left',
    field: (row: PlainMessage<ccbio.Specimen>) => row.primary?.determiner ?? '',
    label: 'Determiner',
    name: 'determiner',
    sortable: true,
  },
  {
    align: 'left',
    field: (row: PlainMessage<ccbio.Specimen>) =>
      row.primary?.fieldDate?.verbatim ?? '',
    label: 'Field Date',
    name: 'Field Date',
    sortable: true,
  },

  {
    align: 'left',
    field: (row: PlainMessage<ccbio.Specimen>) =>
      row.primary?.determinedDate?.verbatim ?? '',
    label: 'Determined Date',
    name: 'Determined Date',
    sortable: true,
  },

  {
    align: 'left',
    field: (row: PlainMessage<ccbio.Specimen>) =>
      row.primary?.determinedReason ?? '',
    label: 'Determined Reason',
    name: 'determinedReason',
    sortable: true,
  },
  {
    align: 'left',
    field: (row: PlainMessage<ccbio.Specimen>) =>
      row.primary?.originalDate?.verbatim ?? '',
    label: 'Original Date',
    name: 'originalDate',
    sortable: true,
  },

  {
    align: 'left',
    field: (row: PlainMessage<ccbio.Specimen>) => row.secondary?.sex ?? '',
    label: 'Sex',
    name: 'sex',
    sortable: true,
  },
  {
    align: 'left',
    field: (row: PlainMessage<ccbio.Specimen>) => row.secondary?.age ?? '',
    label: 'Age',
    name: 'age',
    sortable: true,
  },
  {
    align: 'left',
    field: (row: PlainMessage<ccbio.Specimen>) => row.secondary?.weight ?? '',
    label: 'weight',
    name: 'weight',
    sortable: true,
  },
  {
    align: 'left',
    field: (row: PlainMessage<ccbio.Specimen>) =>
      row.secondary?.weightUnits ?? '',
    label: 'Weight Units',
    name: 'weightUnits',
    sortable: true,
  },
  {
    align: 'left',
    field: (row: PlainMessage<ccbio.Specimen>) =>
      row.secondary?.condition ?? '',
    label: 'condition',
    name: 'condition',
    sortable: true,
  },
  {
    align: 'left',
    field: (row: PlainMessage<ccbio.Specimen>) => row.secondary?.molt ?? '',
    label: 'Molt',
    name: 'molt',
    sortable: true,
  },
  {
    align: 'left',
    field: (row: PlainMessage<ccbio.Specimen>) => row.secondary?.notes ?? '',
    label: 'Secondary Notes',
    name: 'secondaryNotes',
    sortable: true,
  },

  {
    align: 'left',
    field: (row: PlainMessage<ccbio.Specimen>) => row.taxon?.kingdom ?? '',
    label: 'Kingdom',
    name: 'kingdom',
    sortable: true,
  },
  {
    align: 'left',
    field: (row: PlainMessage<ccbio.Specimen>) => row.taxon?.phylum ?? '',
    label: 'Phylum',
    name: 'phylum',
    sortable: true,
  },
  {
    align: 'left',
    field: (row: PlainMessage<ccbio.Specimen>) => row.taxon?.class ?? '',
    label: 'Class',
    name: 'class',
    sortable: true,
  },
  {
    align: 'left',
    field: (row: PlainMessage<ccbio.Specimen>) => row.taxon?.order ?? '',
    label: 'Order',
    name: 'order',
    sortable: true,
  },
  {
    align: 'left',
    field: (row: PlainMessage<ccbio.Specimen>) => row.taxon?.family ?? '',
    label: 'Family',
    name: 'family',
    sortable: true,
  },
  {
    align: 'left',
    field: (row: PlainMessage<ccbio.Specimen>) => row.taxon?.genus ?? '',
    label: 'Genus',
    name: 'genus',
    sortable: true,
  },
  {
    align: 'left',
    field: (row: PlainMessage<ccbio.Specimen>) => row.taxon?.species ?? '',
    label: 'Species',
    name: 'species',
    sortable: true,
  },
  {
    align: 'left',
    field: (row: PlainMessage<ccbio.Specimen>) => row.taxon?.subspecies ?? '',
    label: 'Subspecies',
    name: 'subspecies',
    sortable: true,
  },

  {
    align: 'left',
    field: (row: PlainMessage<ccbio.Specimen>) =>
      row.georeference?.country ?? '',
    label: 'Country',
    name: 'country',
    sortable: true,
  },
  {
    align: 'left',
    field: (row: PlainMessage<ccbio.Specimen>) =>
      row.georeference?.stateProvince ?? '',
    label: 'State Province',
    name: 'stateProvince',
    sortable: true,
  },
  {
    align: 'left',
    field: (row: PlainMessage<ccbio.Specimen>) =>
      row.georeference?.county ?? '',
    label: 'County',
    name: 'county',
    sortable: true,
  },
  {
    align: 'left',
    field: (row: PlainMessage<ccbio.Specimen>) =>
      row.georeference?.locality ?? '',
    label: 'Locality',
    name: 'locality',
    sortable: true,
  },
  {
    align: 'left',
    field: (row: PlainMessage<ccbio.Specimen>) =>
      row.georeference?.latitude ?? '',
    label: 'Latitude',
    name: 'latitude',
    sortable: true,
  },
  {
    align: 'left',
    field: (row: PlainMessage<ccbio.Specimen>) =>
      row.georeference?.longitude ?? '',
    label: 'Longitude',
    name: 'longitude',
    sortable: true,
  },
  {
    align: 'left',
    field: (row: PlainMessage<ccbio.Specimen>) =>
      row.georeference?.habitat ?? '',
    label: 'Habitat',
    name: 'habitat',
    sortable: true,
  },
  {
    align: 'left',
    field: (row: PlainMessage<ccbio.Specimen>) =>
      row.georeference?.continent ?? '',
    label: 'Continent',
    name: 'continent',
    sortable: true,
  },
  {
    align: 'left',
    field: (row: PlainMessage<ccbio.Specimen>) =>
      row.georeference?.locationRemarks ?? '',
    label: 'Location Remarks',
    name: 'locationRemarks',
    sortable: true,
  },
  {
    align: 'left',
    field: (row: PlainMessage<ccbio.Specimen>) =>
      row.georeference?.coordinateUncertaintyInMeters ?? '',
    label: 'Coordinate Uncertainty In Meters',
    name: 'coordinateUncertaintyInMeters',
    sortable: true,
  },
  {
    align: 'left',
    field: (row: PlainMessage<ccbio.Specimen>) =>
      row.georeference?.georeferenceBy ?? '',
    label: 'Georeference By',
    name: 'georeferenceBy',
    sortable: true,
  },
  {
    align: 'left',
    field: (row: PlainMessage<ccbio.Specimen>) =>
      row.georeference?.georeferenceDate?.verbatim ?? '',
    label: 'Georeference Date',
    name: 'GeoreferenceDate',
    sortable: true,
  },
  {
    align: 'left',
    field: (row: PlainMessage<ccbio.Specimen>) =>
      row.georeference?.georeferenceProtocol ?? '',
    label: 'Georeference Protocol',
    name: 'georeferenceProtocol',
    sortable: true,
  },
  {
    align: 'left',
    field: (row: PlainMessage<ccbio.Specimen>) =>
      row.georeference?.geodeticDatum ?? '',
    label: 'Geodetic Datum',
    name: 'geodeticDatum',
    sortable: true,
  },
  {
    align: 'left',
    field: (row: PlainMessage<ccbio.Specimen>) =>
      row.georeference?.footprintWkt ?? '',
    label: 'Footprint Wkt',
    name: 'footprintWkt',
    sortable: true,
  },
  {
    align: 'left',
    field: (row: PlainMessage<ccbio.Specimen>) => row.georeference?.notes ?? '',
    label: 'Notes',
    name: 'notes',
    sortable: true,
  },
])

const visibleColumns = ref([
  'View',
  'SpecimenID',
  'Catalog Number',
  'genus',
  'species',
  'Country',
  'stateProvince',
  'locality',
  'Field Date',
])

const modules = [ClientSideRowModelModule]

const agColDef = [
  {
    cellRenderer: TableButton,
    field: 'specimenId',
    filter: true,
    headerName: 'Specimen ID',
    name: 'SpecimenID',
    pin: 'left',
    sortable: true,
  },
  {
    children: [
      {
        align: 'left',
        field: 'primary.catalogNumber',
        filter: true,
        headerName: 'Catalog Number',
        sortable: true,
      },
      {
        align: 'left',
        field: 'primary.accessionNumber',
        headerName: 'Accession Number',
        name: 'Accession Number',
        sortable: true,
      },
      {
        align: 'left',
        field: 'primary?.fieldNumber',
        headerName: 'Field Number',
        name: 'Field Number',
        sortable: true,
      },
      {
        align: 'left',
        field: 'primary?.tissueNumber',
        headerName: 'Tissue Number',
        name: 'tissueNumber',
        sortable: true,
      },
      {
        align: 'left',
        columnGroupShow: 'open',
        field: 'primary?.cataloger',

        headerName: 'Cataloger',
        name: 'cataloger',
        sortable: true,
      },
      {
        align: 'left',
        columnGroupShow: 'open',
        field: 'primary?.collector',
        headerName: 'Collector',
        name: 'collector',
        sortable: true,
      },
      {
        align: 'left',
        columnGroupShow: 'open',
        field: 'primary?.determiner',
        headerName: 'Determiner',
        name: 'determiner',
        sortable: true,
      },
      {
        align: 'left',
        columnGroupShow: 'open',
        field: 'primary?.fieldDate?.verbatim',
        headerName: 'Field Date',
        name: 'Field Date',
        sortable: true,
      },

      {
        align: 'left',
        columnGroupShow: 'open',
        field: 'primary?.determinedDate?.verbatim',
        headerName: 'Determined Date',
        name: 'Determined Date',
        sortable: true,
      },

      {
        align: 'left',
        columnGroupShow: 'open',
        field: 'primary?.determinedReason',
        headerName: 'Determined Reason',
        name: 'determinedReason',
        sortable: true,
      },
      {
        align: 'left',
        columnGroupShow: 'open',
        field: 'primary?.originalDate?.verbatim',
        headerName: 'Original Date',
        name: 'originalDate',
        sortable: true,
      },
    ],
    headerName: 'Primary',
  },
  {
    children: [
      {
        align: 'left',
        field: 'taxon?.kingdom',
        headerName: 'Kingdom',
        name: 'kingdom',
        sortable: true,
      },
      {
        align: 'left',
        field: 'taxon?.phylum',
        headerName: 'Phylum',
        name: 'phylum',
        sortable: true,
      },
      {
        align: 'left',
        field: 'taxon?.class',
        headerName: 'Class',
        name: 'class',
        sortable: true,
      },
      {
        align: 'left',
        field: 'taxon?.order',
        headerName: 'Order',
        name: 'order',
        sortable: true,
      },
      {
        align: 'left',
        field: 'taxon?.family',
        headerName: 'Family',
        name: 'family',
        sortable: true,
      },
      {
        align: 'left',
        field: 'taxon?.genus',
        headerName: 'Genus',
        name: 'genus',
        sortable: true,
      },
      {
        align: 'left',
        field: 'taxon?.species',
        headerName: 'Species',
        name: 'species',
        sortable: true,
      },
      {
        align: 'left',
        field: 'taxon?.subspecies',
        headerName: 'Subspecies',
        name: 'subspecies',
        sortable: true,
      },
    ],
    headerName: 'Taxon',
  },
  {
    children: [
      {
        align: 'left',
        field: 'secondary?.sex',
        headerName: 'Sex',
        name: 'sex',
        sortable: true,
      },
      {
        align: 'left',
        field: 'secondary?.age',
        headerName: 'Age',
        name: 'age',
        sortable: true,
      },
      {
        align: 'left',
        field: 'secondary?.weight',
        headerName: 'weight',
        name: 'weight',
        sortable: true,
      },
      {
        align: 'left',
        field: 'secondary?.weightUnits',
        headerName: 'Weight Units',
        name: 'weightUnits',
        sortable: true,
      },
      {
        align: 'left',
        field: 'secondary?.condition',
        headerName: 'condition',
        name: 'condition',
        sortable: true,
      },
      {
        align: 'left',
        field: 'secondary?.molt',
        headerName: 'Molt',
        name: 'molt',
        sortable: true,
      },
      {
        align: 'left',
        field: 'secondary?.notes',
        headerName: 'Secondary Notes',
        name: 'secondaryNotes',
        sortable: true,
      },
    ],
    headerName: 'Secondary',
  },
  {
    children: [
      {
        align: 'left',
        field: 'georeference?.country',
        headerName: 'Country',
        name: 'country',
        sortable: true,
      },
      {
        align: 'left',
        field: 'georeference?.stateProvince',
        headerName: 'State Province',
        name: 'stateProvince',
        sortable: true,
      },
      {
        align: 'left',
        field: 'georeference?.county',
        headerName: 'County',
        name: 'county',
        sortable: true,
      },
      {
        align: 'left',
        field: 'georeference?.locality',
        headerName: 'Locality',
        name: 'locality',
        sortable: true,
      },
      {
        align: 'left',
        field: 'georeference?.latitude',
        headerName: 'Latitude',
        name: 'latitude',
        sortable: true,
      },
      {
        align: 'left',
        field: 'georeference?.longitude',
        headerName: 'Longitude',
        name: 'longitude',
        sortable: true,
      },
      {
        align: 'left',
        field: 'georeference?.habitat',
        headerName: 'Habitat',
        name: 'habitat',
        sortable: true,
      },
      {
        align: 'left',
        field: 'georeference?.continent',
        headerName: 'Continent',
        name: 'continent',
        sortable: true,
      },
      {
        align: 'left',
        field: 'georeference?.locationRemarks',
        headerName: 'Location Remarks',
        name: 'locationRemarks',
        sortable: true,
      },
      {
        align: 'left',
        field: 'georeference?.coordinateUncertaintyInMeters',
        headerName: 'Coordinate Uncertainty In Meters',
        name: 'coordinateUncertaintyInMeters',
        sortable: true,
      },
      {
        align: 'left',
        field: 'georeference?.georeferenceBy',
        headerName: 'Georeference By',
        name: 'georeferenceBy',
        sortable: true,
      },
      {
        align: 'left',
        field: 'georeference?.georeferenceDate?.verbatim',
        headerName: 'Georeference Date',
        name: 'GeoreferenceDate',
        sortable: true,
      },
      {
        align: 'left',
        field: 'georeference?.georeferenceProtocol',
        headerName: 'Georeference Protocol',
        name: 'georeferenceProtocol',
        sortable: true,
      },
      {
        align: 'left',
        field: 'georeference?.geodeticDatum',
        headerName: 'Geodetic Datum',
        name: 'geodeticDatum',
        sortable: true,
      },
      {
        align: 'left',
        field: 'georeference?.footprintWkt',
        headerName: 'Footprint Wkt',
        name: 'footprintWkt',
        sortable: true,
      },
      {
        align: 'left',
        field: 'georeference?.notes',
        headerName: 'Notes',
        name: 'notes',
        sortable: true,
      },
    ],
    headerName: 'Georeference',
  },
]
</script>

<template>
  <div>
    <!-- <Testtable /> -->
    <!-- <SpecimenTable /> -->

    <QCard>
      <AgGridVue
        :auto-size-strategy="{ type: 'fitCellContents' }"
        :column-defs="agColDef"
        :framework-components="{ tableButton: TableButton }"
        :modules="modules"
        :row-data="store.SpecimenList"
        class="ag-theme-quartz"
        style="height: 500px"
      />
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
          :columns="colDef"
          :filter="filter"
          :row-key="(row: PlainMessage<ccbio.Specimen>) => row.specimenId"
          :rows="store.SpecimenList"
          :rows-per-page-options="[25, 30, 50, 100, 200, 500, 1000, 0]"
          :visible-columns="visibleColumns"
          bordered
          class="max-h-max"
          dense
          flat
          separator="horizontal"
          title="Specimens"
        >
          <template #top-right>
            <QBtn
              v-model="filter"
              class="q-mr-sm"
              color="primary"
              flat
              icon="refresh"
              label="Reload"
              round
              @click="store.Reload"
            />
            <q-select
              v-model="visibleColumns"
              :options="colDef"
              class="px-2"
              dense
              display-value="Visible Columns"
              emit-value
              map-options
              multiple
              option-value="name"
              options-cover
              options-dense
              outlined
              style="min-width: 150px"
            />
            <q-input
              v-model="filter"
              borderless
              debounce="300"
              dense
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
                  :no-prefetch="true"
                  :to="`/collection/${useRoute().params?.collectionId.toString()}/Specimen/View-${
                    props.row.specimenId
                  }`"
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

<style>
.ag-theme-quartz {
  --ag-grid-size: 5px;
  --ag-list-item-height: 20px;
}
</style>
