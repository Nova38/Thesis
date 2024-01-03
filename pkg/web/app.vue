<template>
  <div>
    <NuxtLayout>
      <template #content>
        <NuxtPage :page-key="(route) => route.fullPath" />
      </template>
    </NuxtLayout>
  </div>
</template>

<script lang="ts" setup>
const route = useRoute();
</script>
