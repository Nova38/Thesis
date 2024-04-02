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
    />
    <PCard v-if="bulk.RawRows.length">
      <template #content>
        <div class="flex flex-col">
          <div class="object-center">
            <PSelectButton
              v-model="bulk.Mode"
              :options="['import', 'update', 'hybrid']"
              aria-labelledby="basic"
            />
          </div>
          <PMeterGroup :value="bulk.ImportStatus" />
          <PMeterGroup :value="bulk.UploadStatus" />
        </div>
      </template>
    </PCard>

    <PCard v-if="bulk.RawRows.length">
      <template #content>
        <div class="space-y-2">
          <div>
            <ImportTableMapping />
          </div>

          <div>
            <UButton
              block
              class="mb-4"
              :disabled="!bulk.RawRows.length"
              @click="bulk.Upload"
            >
              Upload
            </UButton>
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
  </div>
</template>

<style scoped></style>
