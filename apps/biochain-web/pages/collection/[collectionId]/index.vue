<script lang="ts" setup>
const route = useRoute()
const CollectionId = computed(() => {
  console.log(route)
  return route.params.collectionId.toString()
})
</script>

<template>
  <div>
    <!-- <pre wrap>{{ config.app.baseURL }}</pre> -->
    <!-- {{ crush(MakeEmptySpecimen()) }} -->
    {{ CollectionId }}

  </div>
</template>

<style></style>
