<script lang="ts" setup>
import type { ImportRowsData } from '~/components/import/ImportCsvFile.vue'

function LoadCSV(val: ImportRowsData) {
  console.log(val)
}
</script>

<template>
  <div>
    <ImportCsvFile
      @id-header-selection="LoadCSV"
    />
  </div>
</template>

<style scoped></style>
