<script lang="ts" setup>
import Papa, { type ParseResult } from 'papaparse'
import type { FormKitFileValue } from '@formkit/inputs'
import type { FormKitNode } from '@formkit/core'

export interface ImportRowsData {
  headers: string[]
  specimenIdHeader: string
  rows: Record<string, string>[]
}

const emit = defineEmits<{
  parse: [payload: ParseResult<Record<string, string>>]
  idHeaderSelection: [data: ImportRowsData]
}>()

// const file = ref<File | null>(null)
const form = ref()
const headers = ref<string[]>([])

const rows = ref<Record<string, string>[]>([])

const SpecimenIdField = ref<string>('')
const unique = ref<string[]>([])

function onUpload(value?: FormKitFileValue, node?: FormKitNode) {
  console.log('onUpload', value, node)

  if (value === undefined)
    return
  const file = value[0]?.file
  if (file === undefined)
    return

  Papa.parse(file, {
    header: true,
    worker: true,
    skipEmptyLines: 'greedy',
    complete: (results: ParseResult<Record<string, string>>) => {
      console.log('All done:', results.data)
      // emit('csvParse', results)
      // emit('parse', results)

      headers.value = results.meta.fields ?? []
      rows.value = results.data

      unique.value = Object.keys(UniqueFields(results.data, headers.value))

      if (unique.value.length === 1)
        SpecimenIdField.value = unique.value[0]
    },
  })
}

function handleForm(data: any, node: FormKitNode) {
  console.log('handleForm', data)
  console.log('node', node)

  emit('idHeaderSelection', {
    headers: headers.value,
    rows: rows.value,
    specimenIdHeader: SpecimenIdField.value,
  })
}
</script>

<template>
  <PCard>
    <template #title>
      <h3>
        Raw Rows from CSV:
      </h3>
    </template>
    <template #content>
      <div class="justify-center">
        <FormKit ref="form" name="csvForm" type="form" submit-label="Load Specimens" @submit="handleForm">
          <div class="flex flex-row gap-4">
            <FormKit
              name="csvFile"
              type="file"
              label="Upload CSV"
              :multiple="true"

              accept=".csv"
              @input="onUpload"
            />
            <FormKit
              v-model="SpecimenIdField"
              name="SpecimenIdField"
              type="select"
              label="Column that contains the Collection ID"
              validation="required"
              :options="unique"
            />
          </div>
        </FormKit>
      </div>
    </template>
  </PCard>
</template>

<!-- eslint-disable no-restricted-syntax -->
<style>

</style>
