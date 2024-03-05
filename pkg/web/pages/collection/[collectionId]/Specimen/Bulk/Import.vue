<script lang="ts" setup>
const bulk = useBulkStore()

bulk.CollectionId = useNuxtApp().$collectionId.value
</script>

<template>
  <div>
    <ImportCsvFile
      @id-header-selection="(val) => bulk.LoadCsv({
        headers: val.headers,
        rows: val.rows,
        specimenIdHeader: val.specimenIdHeader,
      })"
    />
    <ImportTableMapping />
  </div>
</template>

<style scoped></style>
