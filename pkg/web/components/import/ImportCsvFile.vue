<!-- eslint-disable no-restricted-syntax -->
<script lang="ts" setup>
import Papa, { type ParseResult } from 'papaparse'
import type { FormKitFileValue } from '@formkit/inputs'
import type { FormKitNode } from '@formkit/core'

const emit = defineEmits({
  csvParse(payload: ParseResult<Record<string, string>>) {
    return payload
  },
})

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
      emit('csvParse', results)

      headers.value = results.meta.fields ?? []
      rows.value = results.data

      unique.value = Object.keys(UniqueFields(results.data, headers.value))
    },
  })
}

const bulk = useBulkUpdate()
function handleForm(data: any, node: FormKitNode) {
  console.log('handleForm', data)
  console.log('node', node)

  // TODO: Switch to emitting so we can reuse this component
  bulk.LoadUpdates({
    headers: headers.value,
    rows: rows.value,
    specimenIdHeader: SpecimenIdField.value,
  })
}
// function submitForm() {
//   const node = form.value.node
//   console.log('node', node)
//   node.submit()
// }
</script>

<template>
  <PCard>
    <template #content>
      <!-- <PFileUpload
          accept=".csv"
          mode="basic"
          :auto="true"
          custom-upload
          @uploader="onUpload"
        /> -->
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
              label="Select the column that contains the Specimen ID "
              validation="required"
              :options="unique"
            />
          <!-- <div class="mt-3 self-center">
            <FormKit
              class=""
              type="button" @click="submitForm"
            >
              Submit request
            </FormKit>
          </div> -->
          </div>
        </FormKit>
      </div>

      <!-- <USelectMenu
            class="min-w-52"
            :model-value="SpecimenIdField"
            :options="unique"
          /> -->
    </template>
  </PCard>
</template>

<style></style>
