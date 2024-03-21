<script lang="ts" setup>
const bulk = useBulkStore()

bulk.CollectionId = useNuxtApp().$collectionId.value
</script>

<template>
  <div>
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

    <PCard>
      <template v-if="bulk.RawRows.length" #content>
        <UButton
          block
          class="mb-4"
          :disabled="!bulk.RawRows.length"
          @click="bulk.Upload"
        >
          Upload
        </UButton>
        <ImportTableMapping />
      </template>
    </PCard>
    <PCard>
      <template v-if="bulk.MappedSpecimen.length" #content>
        <SpecimenTable :specimen-list="bulk.MappedSpecimen" />
      </template>
    </PCard>
  </div>
</template>

<style scoped></style>
