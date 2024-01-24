<script setup lang="ts">
import type { QTableProps } from 'quasar'

// Props
const props = defineProps<{
  rawHeaders: string[]
  rawRows: Record<string, string>[]
}>()

// Models
const specimenMapping = defineModel<SpecimenMapping>('specimenMapping', {
  required: true,
})

// Computed
const processedData = useArrayMap(props.rawRows, (row) => {
  return TransformRecordToFlatSpecimen(row, specimenMapping.value)
})

const MappingHeaders: Ref<QTableProps['columns']> = computed(() => {
  return specimenMapping.value.map((map) => {
    return {
      field: map.newKey,
      label: map.newKey,
      name: map.newKey,
    }
  })
})

// Methods
</script>

<template>
  <UCard>
    <div class="text-xl font-semibold border-blue-400 border-solid mb-2">
      Imported Specimen Values
    </div>
    <div v-if="processedData">
      <q-table
        :columns="MappingHeaders"
        :rows="processedData"
        dense
      >
        <template #header-cell="{ col }">
          <q-th>
            <div class="">
              {{ col.label }}
            </div>
            <q-select
              v-model="specimenMapping[col.label].oldKey"
              :options="rawHeaders"
              dense
              label="key"
              label-color="teal-10"
              stack-label
            >
              <template
                v-if="specimenMapping[col.label]"
                #append
              >
                <q-icon
                  class="cursor-pointer"
                  color="red"
                  dense
                  name="cancel"
                  size=".75em"
                  @click.stop.prevent="ClearMapping(specimenMapping, col.label)"
                />
              </template>
            </q-select>
          </q-th>
        </template>
      </q-table>
    </div>
  </UCard>
</template>
