<template>
  <div>
    <!-- <pre wrap>{{ config.app.baseURL }}</pre> -->
    <!-- {{ crush(MakeEmptySpecimen()) }} -->
    {{ CollectionId }}
    <div>
      <NuxtLink to="/collection/k"> K </NuxtLink>
      <NuxtLink to="/collection/kd"> KD </NuxtLink>
      <NuxtLink to="/collection/ku_orn"> ku-orn </NuxtLink>
    </div>
  </div>
</template>

<script lang="ts" setup>
const route = useRoute();
const CollectionId = computed(() => {
  console.log(route);
  return route.params.collectionId.toString();
});
</script>

<style></style>
