<script lang="ts" setup>
const route = useRoute()
const CollectionId = computed(() => {
  console.log(route)
  return route.params.collectionId.toString()
})
</script>

<template>
  <div></div>
</template>

<style></style>
