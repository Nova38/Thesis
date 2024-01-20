<script lang="ts" setup>
import Papa, { type ParseResult } from 'papaparse'

const file = ref<File | null>(null)

watch(file, (newFile) => {
  if (!file)
    return

  if (newFile) {
    csv.value = []
    headers.value = []

    Papa.parse(newFile, {
      header: true,

      complete: (results: ParseResult<Record<string, string>>) => {
        headers.value = results.meta.fields || []

        for (const row of results.data)
          csv.value.push(row)
      },
    })
  }
})

const csv = defineModel<Record<string, string>[]>('csv', {
  default: () => [],
  required: true,
})
const headers = defineModel<string[]>('headers', {
  default: () => [],
  required: true,
})
</script>

<template>
  <div>
    <q-card-section>
      <h2>Select CSV file to import from</h2>
      <q-file
        v-model="file"
        outlined
        accept=".csv"
      >
        <template #prepend>
          <q-icon name="attach_file" />
        </template>
      </q-file>
    </q-card-section>
  </div>
</template>

<style></style>
