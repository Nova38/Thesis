<script lang="ts" setup>
const bulk = useBulkUpdate()

bulk.CollectionId = useNuxtApp().$collectionId.value

async function myFetch() {
  const x = await $fetch('/api/cc/specimens/fullList?collectionId=k')

  console.log(x)
}
</script>

<template>
  <div>
    <ImportCsvFile
      @id-header-selection="
        (val) =>
          bulk.LoadUpdates({
            headers: val.headers,
            rows: val.rows,
            specimenIdHeader: val.specimenIdHeader,
          })
      "
    />
    <PCard>
      <template #title>
        <h3>Raw Rows from CSV:</h3>
      </template>
      <template #content>
        <PDataTable
          :value="bulk.RawRows"
          data-key="id"
          paginator
          show-gridlines
          striped-rows
          size="small"
          :rows="10"
          :rows-per-page-options="[5, 10, 20, 50]"
        >
          <!-- <PColumn
            :field="bulk.ImportColumns.id.field"
            header="Specimen ID"
          /> -->
          <PColumn
            v-for="col of bulk.ImportColumns.raw"
            :key="col.name"
            :field="col.field"
            :header="col.name"
          >
            <UDivider />
          </PColumn>
        </PDataTable>
      </template>
    </PCard>
    <PCard>
      <template #title>
        <h3>Processed Rows To Import:</h3>
      </template>
      <template #content>
        <PDataTable
          v-if="bulk.MappedRows"
          :value="bulk.MappedRows"
          data-key="specimenId"
          paginator
          show-gridlines
          striped-rows
          size="small"
          :rows="10"
          :rows-per-page-options="[5, 10, 20, 50]"
        >
          <PColumn
            v-for="col of FlatColDefs"
            :key="col.field"
            :field="col.field"
          >
            <template #header="">
              <div class="text-nowrap">
                <FormKit
                  type="select"
                  :label="col.field()"
                  :name="`map-${col.name}`"
                  :options="bulk.ImportedHeaders"
                  label-class="text-nowrap "
                  @input="(val) => bulk.SetMapping(val, col)"
                />
              </div>
            </template>

            <UDivider />
            <template #body="{ data }">
              {{ data[col.field()] }}
            </template>
          </PColumn>
        </PDataTable>
      </template>
    </PCard>

    <PCard>
      <template #title>
        <h3>Processed Rows To Import:</h3>
      </template>
      <template #content>
        <PDataTable
          v-if="bulk.MappedRows"
          :value="bulk.MappedRows"
          data-key="specimenId"
          paginator
          show-gridlines
          striped-rows
          size="small"
          :rows="10"
          :rows-per-page-options="[5, 10, 20, 50]"
        >
          <PColumn
            v-for="col of SpecimenColDefs"
            :key="col.field"
            :field="col.field"
          >
            <template #header="">
              <div class="text-nowrap">
                <FormKit
                  type="select"
                  :label="col.field"
                  :name="`map-${col.name}`"
                  :options="bulk.ImportedHeaders"
                  label-class="text-nowrap "
                  @input="(val) => bulk.SetMapping(val, col)"
                />
              </div>
            </template>

            <UDivider />
            <!-- <template #body="{ data }">
              {{ data[col.field] }}
            </template> -->
          </PColumn>
        </PDataTable>
        <pre wrapped>
        {{ bulk.MappedRows }}

        </pre>
      </template>
    </PCard>
  </div>
</template>

<style scoped></style>
