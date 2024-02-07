<!-- eslint-disable no-restricted-syntax -->
<script lang="ts" setup>
import { QCardSection } from 'quasar'

const file = ref<File | null>(null)

watch(file, (newFile) => {
  if (newFile) {
    useBulkUpdate().LoadFromFile(newFile)
  }
})
</script>

<template>
  <h2>Select CSV file to import from</h2>
<!-- <UInput type="file" accept=".csv" :ui="{}">
      <template #leading></template>
      hi
    </UInput> -->
<QCardSection>
  <QFile v-model="file" accept=".csv" outlined>
    <template #prepend>
      <q-icon name="attach_file" />
    </template>
  </QFile>
</QCardSection>
</template>

<style></style>
``
