<script lang="ts" setup>
import 'primevue/resources/themes/aura-light-green/theme.css'
</script>

<template>
  <div>
    <NuxtLayout>
      <template #content>
        <NuxtLoadingIndicator />
        <NuxtPage :page-key="(route) => route.fullPath" />
      </template>
    </NuxtLayout>
  </div>
</template>

<style>
@import url('primevue/resources/themes/aura-light-blue/theme.css');

html {
  font-size: 14px;
}
</style>
