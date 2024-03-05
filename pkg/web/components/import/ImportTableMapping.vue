<script lang="ts" setup>
const bulk = useBulkStore()
</script>

<template>
  <div>
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
      <PColumn header="Meta" field="id">
        <template #body="row">
          <pre>{{ row.data?.id }}</pre>
          <pre>{{ bulk.RawRowsMeta.get(row.data?.id) }}</pre>
        </template>
      </PColumn>

      <PColumn
        v-for="col of bulk.RawColDefs"
        :key="col.name"
        :field="col.field"
      >
        <template #header="">
          <div class="text-nowrap">
            <FormKit
              type="select"
              :label="col.name"
              :name="`map-${col.name}`"
              :options="['', ...FlattedSpecimenKeys]"
              label-class="text-nowrap "
              @input="(val) => bulk.SetMapping(val, col) "
            />
          </div>
        </template>
      </pcolumn>
    </PDataTable>
  </div>
</template>

<style scoped>

</style>
              <!-- @input="(val) => bulk.SetMapping(val, col)" -->
<!--  -->
