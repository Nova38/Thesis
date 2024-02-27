<script lang="ts" setup>
const bulk = useBulkUpdate()
</script>

<template>
  <div class="">
    <ImportCsvFile
      @id-header-selection="(val) => // TODO: Switch to emitting so we can reuse this component
        bulk.LoadUpdates({
          headers: val.headers,
          rows: val.rows,
          specimenIdHeader: val.specimenIdHeader,
        })"
    />

    <UCard class="my-4">
      <template #header>
        <h3>
          Raw Rows from CSV:
        </h3>
      </template>
      <PDataTable
        :value="bulk.Rows"
        data-key="id"
        paginator
        show-gridlines
        striped-rows
        size="small"
        :rows="10"
        :rows-per-page-options="[5, 10, 20, 50]"
      >
        <PColumn
          :field="bulk.ImportColumns.id.field"
          header="Specimen ID"
        />
        <PColumn
          v-for="col of bulk.ImportColumns.raw"
          :key="col.name"
          :field="col.field"
          :header="col.name"
        >
          <UDivider />
        </PColumn>
      </PDataTable>
    </UCard>
    <UCard
      class="my-4"
    >
      <template #header>
        <h3>
          Processed Rows
        </h3>
      </template>
      <!-- <PDataTable
        v-if="bulk.ProcessingCSV"
        :value="bulk.RawRows"
        data-key="id"
        paginator
        show-gridlines
        striped-rows
        size="small"
        :rows="10"
        :rows-per-page-options="[5, 10, 20, 50]"
      >
        <PColumn
          v-for="col of bulk.ImportColumns"
          :key="col.field"
          :field="col.field"
        >
          <template #header="{ column }">
            <div class="flex flex-col">
              <div>{{ col.name }}</div>
              <div>
                <USelectMenu />
              </div>
            </div>
          </template>
          <UDivider />
        </PColumn>
      </PDataTable> -->
    </UCard>
  </div>
</template>

<style scoped>

</style>
