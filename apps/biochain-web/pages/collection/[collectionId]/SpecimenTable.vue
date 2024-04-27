<script lang="ts" setup>
const store = useCollectionsStore()

callOnce(() => {
  store.Reload()
})
</script>

<template>
  <div>
    <PCard>
      <template #title>
        <div>
          <h3>Specimens in Collection</h3>
          <UButton
            color="primary"
            label="Reload"
            size="sm"
            @click="store.Reload"
          />
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
