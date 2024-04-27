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
    <ImportCsvFile
      @id-header-selection="
        (val) =>
          bulk.LoadCsv({
            headers: val.headers,
            rows: val.rows,
            specimenIdHeader: val.specimenIdHeader,
          })
      "
      @reset="bulk.$reset"
    />

    <ImportMenu />

    <PCard v-if="bulk.RawRows.length">
      <template #content>
        <div class="space-y-2">
          <div>
            <PBlockUI :blocked="bulk.Loading">
              <ImportTableMapping />
            </PBlockUI>
          </div>
        </div>
      </template>
    </PCard>

    <PCard v-if="bulk.MappedSpecimen.length">
      <template #content>
        <SpecimenTable :specimen-list="bulk.MappedSpecimen" />
      </template>
    </PCard>

    <PCard>
      <template #content> </template>
    </PCard>
  </div>
</template>

<style scoped></style>
