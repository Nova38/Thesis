<script lang="ts" setup>
import Papa, { type ParseResult } from 'papaparse'
import { QCardSection } from 'quasar'

const file = ref<File | null>(null)

export interface CsvMeta {
  catIdField: string
  headers: string[]
  rowsByCatId: Record<string, Record<string, string>>
}

const rows = defineModel<Record<string, string>[]>('csv', {
  default: () => [],
  required: true,
})

const headers = defineModel<string[]>('headers', {
  default: () => [],
  required: true,
})

watch(file, (newFile) => {
  if (!file) return

  if (newFile) {
    rows.value = []
    headers.value = []

    Papa.parse(newFile, {
      complete: (results: ParseResult<Record<string, string>>) => {
        headers.value = results.meta.fields || []

        for (const row of results.data) rows.value.push(row)
      },

      header: true,
    })
  }
})
</script>

<template>
  <QCardSection>
    <h2>Select CSV file to import from</h2>
    <!-- <UInput type="file" accept=".csv" :ui="{}">
      <template #leading></template>
      hi
    </UInput> -->

    <QFile v-model="file" accept=".csv" outlined>
      <template #prepend>
        <q-icon name="attach_file" />
      </template>
    </QFile>
  </QCardSection>
</template>

<style></style>
