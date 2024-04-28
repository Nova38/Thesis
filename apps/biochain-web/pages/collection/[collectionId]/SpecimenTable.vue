<script lang="ts" setup>
const store = useCollectionsStore()

callOnce(() => {
  store.Reload()
})

onMounted(() => {
  store.AbortLoading = false
})

onBeforeRouteLeave(() => {
  store.AbortLoading = true
})
</script>

<template>
  <div>
    <PCard>
      <template #title>
        <div>
          <h3>Specimens in Collection</h3>
          <div class="py-2">
            <template v-if="store.Loading">
              <UProgress animation="carousel" />
            </template>
            <template v-else-if="store.FullyLoaded">
              <UProgress
                class="green"
                :value="1"
                :max="1"
              />
            </template>
            <template v-else>
              <UProgress
                :value="50"
                :max="100"
              />
            </template>
          </div>

          <div class="flex flex-grow flex-row gap-2">
            <UButton
              color="amber"
              label="Reload"
              size="sm"
              class="flex-grow"
              @click="store.Reload"
            />
            <UButton
              color="red"
              label="Pause Loading"
              size="sm"
              class="flex-grow"
              @click="
                () => {
                  store.AbortLoading = true
                }
              "
            />
            <UButton
              color="green"
              label="Resume Loading"
              size="sm"
              class="flex-grow"
              @click="store.ResumeLoading"
            />
          </div>
        </div>
      </template>

      <template #content>
        <SpecimenTable :specimen-list="store.SpecimenList" />
      </template>
    </PCard>
  </div>
</template>

<style>
.ag-theme-quartz {
  --ag-grid-size: 5px;
  --ag-list-item-height: 20px;
}
</style>
