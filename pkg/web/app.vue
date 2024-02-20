<script lang="ts" setup>
// import 'primevue/resources/themes/aura-light-green/theme.css'

const visible = ref(true)
</script>

<template>
  <div class="w-full h-full min-h-screen flex flex-col gap-2">
    <NavHeader />

    <div class="bg-slate-100 flex h-30 flex-grow flex-row gap-2">
      <NavSidebar class="bg-red-100" />

      <div v-show="visible" />

      <div class="bg-amber-100 flex-grow">
        <NuxtErrorBoundary>
          <NuxtLoadingIndicator />
          <NuxtPage />
        </NuxtErrorBoundary>
      </div>
    </div>
  </div>
</template>

<style>
@import url('primevue/resources/themes/aura-light-blue/theme.css');

html {
  font-size: 12px;
}
</style>
