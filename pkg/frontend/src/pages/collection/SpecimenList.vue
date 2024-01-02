<template>
  <q-page padding class="column justify-evenly">
    <div style="flex: 1 1 auto" class="q-pa-md full-height column">
      <q-card class="q-pa-md">
        <ag-grid-vue
          id="main_item"
          class="ag-theme-material"
          style="height: 600px"
          :rowData="data.items"
          :columnDefs="columnDefs"
          :columnTypes="columnTypes"
          :defaultColDef="defaultColDef"
          :frameworkComponents="{ tableButton: TableButton }"
        >
        </ag-grid-vue>
        <!-- <q-table
          :columns="colDefs"
          :rows="data.items"
          :pagination="{ rowsPerPage: 10 }"
        >
        </q-table> -->
      </q-card>
    </div>
  </q-page>
</template>

<script setup lang="ts">
import { watch, ref, Ref, reactive } from 'vue';

import { Dialog, useQuasar } from 'quasar';

import { ccApi } from 'boot/axios';
import { schema } from 'src/lib/ccbio';

import { AgGridVue } from '@ag-grid-community/vue3';
import TableButton from 'src/components/table/TableButton.vue';
import { QTableProps } from 'quasar';

const props = defineProps<{
  collectionId: string;
}>();

const data = ref(new schema.Specimen_List());

watch(
  props,
  async () => {
    console.log('ActiveCollectionId', props.collectionId);

    const payload = new schema.GetSpecimenByCollectionRequest({
      id: {
        collectionId: props.collectionId,
      },
    });

    console.log(payload);
    data.value = await ccApi.specimen.GetSpecimenByCollection(payload);
    console.log(data.value);
  },
  { immediate: true }
);

// function popup(event: any) {
//   // console.log('clicked on', specimen);

//   // refreshSpecimen(specimen.guid);

//   // console.log(raw_data[row.catalogNumber]);

//   $q.dialog({
//     component: Dialog,
//     // props forwarded to your custom component
//     componentProps: {
//       // guid: specimen.guid,
//       // ...more..props...
//     },
//   })
//     .onOk(() => {
//       console.log(row);
//     })
//     .onDismiss(() => {
//       console.log('Called on OK or Cancel');
//     });
// }

const defaultPageSize = 25;

// const colDefs: QTableProps['columns'] = [
//   {
//     name: 'Specimen_ID',
//     label: 'Specimen ID',
//     field: 'id',
//     // cellRenderer: TableButton,
//     sortable: true,
//   },
//   {
//     name: 'Catalog_Number',
//     label: 'Catalog Number',
//     field: (row) => row.primary.catalog_number,
//     sortable: true,
//   },
//   {
//     name: 'Genus',
//     label: 'Genus',
//     field: 'taxon.genus',
//     sortable: true,
//   },
//   {
//     name: 'Species',
//     label: 'Species',
//     field: 'taxon.species',
//   },
//   {
//     name: 'Country',
//     label: 'Country',
//     field: 'georeference.country',
//   },
//   {
//     name: 'State',
//     label: 'State',
//     field: 'georeference.state_province',
//   },
//   {
//     name: 'Locality',
//     label: 'Locality',
//     field: 'georeference.locality',
//   },
//   {
//     name: 'Collection Date',
//     label: 'Collection Date',
//     field: 'primary.field_date',

//     type: 'dateColumn',
//     filter: 'agDateColumnFilter',
//     columnGroupShow: 'open',
//   },
// ];

const columnDefs = reactive([
  {
    headerName: 'Specimen ID',
    field: 'id',
    cellRenderer: TableButton,
  },
  {
    headerName: 'Catalog Number',
    field: 'primary.catalog_number',
  },
  {
    headerName: 'Genus',
    field: 'taxon.genus',
  },
  {
    headerName: 'Species',
    field: 'taxon.species',
  },
  {
    headerName: 'Country',
    field: 'georeference.country',
  },
  {
    headerName: 'State',
    field: 'georeference.state_province',
  },
  {
    headerName: 'Locality',
    field: 'georeference.locality',
  },
  {
    headerName: 'Collection Date',
    field: 'primary.field_date',

    type: 'dateColumn',
    filter: 'agDateColumnFilter',
    columnGroupShow: 'open',
  },
]);

const defaultColDef = {
  sortable: true,
  filter: true,
  // flex: 1,
  resizable: true,
};
const columnTypes = {
  nonEditableColumn: { editable: false },
  dateColumn: {
    filter: 'agDateColumnFilter',
    valueFormatter: (params: Ref<any>) =>
      new Date(params.value).toLocaleDateString(),
  },
};
// const gridOptions = {
//   // PROPERTIES
//   // Objects like myRowData and myColDefs would be created in your application
//   // rowData: props.data,
//   columnDefs: columnDefs,
//   defaultColDef: defaultColDef,
//   pagination: true,
//   rowSelection: 'single',
//   animateRows: true, // have rows animate to new positions when sorted

//   // EVENTS
//   // Add event handlers
//   onRowClicked: (event) => {
//     popup(event.data);
//   },
//   onColumnResized: (event) => console.log('A column was resized'),
//   onGridReady: (event) => {
//     console.log('Grid is ready');
//     event.api.sizeColumnsToFit();
//     console.log(event);

//     const allColumnIds = [];

//     event.columnApi.getColumns().forEach((column) => {
//       allColumnIds.push(column.getId());
//     });
//     event.columnApi.autoSizeColumns(allColumnIds);
//   },

//   // CALLBACKS
//   getRowHeight: (params) => 25,
// };
</script>

<style>
/* #main_item {
  flex-grow: 1;
  height: 500px;
} */
</style>
