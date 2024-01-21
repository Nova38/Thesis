<script setup lang="ts">
defineProps<{
  rawHeaders: string[]
}>()

defineModel<{ mapping: SpecimenMapping[] }>('specimenMapping', {
  default: () => ([{
    newKey: '',
    oldKey: '',
  }]),
})
</script>

<template>
  <UCard>
    <div class="text-xl font-semibold border-blue-400 border-solid mb-2">
      Imported Specimen Values
    </div>
    <div v-if="possessedData">
      <q-table
        :columns="MappingHeaders"
        :rows="possessedData"
        dense
      >
        <template #header-cell="props">
          <q-th :props="props">
            <div class="">
              {{ props.col.label }}
            </div>
            <q-select
              v-model="specimenMapping[props.col.label]"
              :options="sortedImportHeaders"
              dense
              label="key"
              label-color="teal-10"
              stack-label
            >
              <template
                v-if="specimenMapping[props.col.label]"
                #append
              >
                <q-icon
                  class="cursor-pointer"
                  color="red"
                  dense
                  name="cancel"
                  size=".75em"
                  @click.stop.prevent="clearKey(props.col.label)"
                />
              </template>
            </q-select>
          </q-th>
        </template>
      </q-table>
    </div>
  </UCard>
</template>
