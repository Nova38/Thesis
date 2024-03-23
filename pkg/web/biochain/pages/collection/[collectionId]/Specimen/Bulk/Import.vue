<script lang="ts" setup>
const bulk = useBulkStore()

bulk.CollectionId = useNuxtApp().$collectionId.value

onBeforeRouteLeave((to, from, next) => {
  console.log('onBeforeRouteLeave', to, from)
  useBulkStore().$reset()
  next()
})
</script>

<template>
  <div>
    <PCard>
      <template #content>
        <UButton
          block
          color="red"
          :disabled="!bulk.RawRows.length"
          @click="bulk.$reset"
        >
          Reset
        </UButton>
      </template>
    </PCard>

    <ImportCsvFile
      @id-header-selection="
        (val) =>
          bulk.LoadCsv({
            headers: val.headers,
            rows: val.rows,
            specimenIdHeader: val.specimenIdHeader,
          })
      "
    />

    <PCard v-if="bulk.RawRows.length">
      <template #content>
        <ImportTableMapping />
        <UButton
          block
          class="mb-4"
          :disabled="!bulk.RawRows.length"
          @click="bulk.Upload"
        >
          Upload
        </UButton>
      </template>
    </PCard>
    <PCard v-if="bulk.MappedSpecimen.length">
      <template #content>
        <SpecimenTable :specimen-list="bulk.MappedSpecimen" />
      </template>
    </PCard>
  </div>
</template>

<style scoped></style>
