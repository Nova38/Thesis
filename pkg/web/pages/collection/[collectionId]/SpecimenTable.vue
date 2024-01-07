<template>
  <div>
    <!-- <Testtable /> -->
    <!-- <SpecimenTable /> -->
    <QCard>
      <div v-if="data && !pending && !error">
        <QTable
          :rows="data"
          :columns="colDefs"
          bordered
          flat
          dense
          title="Specimens"
          separator="horizontal"
          :filter="filter"
        >
          <template #top-right>
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

// const links = useBreadcrumbLinks();
const route = useRoute();
// const { data, pending, error, refresh } = useFetch(
//   `/api/cc/specimens/list?collectionId=${route.params.collectionId}`,
// );

// function navigate(specimenId: string) {
//   return navigateTo({
//     path: `/collection/${useRoute().params?.collectionId.toString()}/Specimen/View-${specimenId}`,
//   });
// }

const filter = ref("");

const colDefs = ref([
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
    field: (row: any) => row.primary?.catalogNumber ?? "",
    sortable: true,
  },
  {
    label: "Genus",
    name: "Taxon_Genus",
    align: "left",

    field: (row: any) => row.taxon?.genus ?? "",
    sortable: true,
  },
  {
    label: "Species",
    name: "Taxon_Species",
    align: "left",

    field: (row: any) => row.taxon?.species ?? "",
    sortable: true,
  },
  {
    label: "Country",
    align: "left",

    name: "Georeference_Country",
    field: (row: any) => row.georeference?.country ?? "",
    sortable: true,
  },
  {
    label: "State/Province",
    align: "left",

    name: "Georeference_StateProvince",
    field: (row: any) => row.georeference?.stateProvince ?? "",
    sortable: true,
  },
  {
    label: "Locality",
    align: "left",

    name: "Georeference_Locality",
    field: (row: any) => row.georeference?.locality ?? "",
  },
  {
    label: "Field Date",
    name: "Primary_FieldDate",
    align: "left",

    field: (row: any) => row.primary?.fieldDate?.verbatim ?? "",
    required: true,
  },
]);

const { data, pending, error } = await useCustomFetch<
  PlainMessage<ccbio.Specimen>[]
>(`/api/cc/specimens/list`, {
  query: {
    collectionId: route.params.collectionId,
  },
});
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
