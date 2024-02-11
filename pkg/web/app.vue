<script lang="ts" setup>
// import 'primevue/resources/themes/aura-light-green/theme.css'

const visible = ref(true)
</script>

<template>
  <div>
    <PToolbar class="bg-blue-600">
      <template #start>
        <UButton label="Click Me" />
        <PButton
          icon="pi pi-print"
          class="mr-2"
          severity="secondary"
        />
        <PButton
          icon="pi pi-upload"
          severity="secondary"
        />
      </template>

      <template #center />

      <template #end />
    </PToolbar>
    <div class="grid grid-flow-col gap-2 max-w-screen">
      <div v-show="visible">
        <NavSidebar />
      </div>

      <div class=" col-span-4	">
        <NuxtLoadingIndicator />
        <NuxtPage />
      </div>
    </div>

    <!-- <NuxtLayout>
      <template #content>

      </template>
    </NuxtLayout> -->
  </div>
</template>

<style>
@import url('primevue/resources/themes/aura-light-blue/theme.css');

html {
  font-size: 12px;
}
</style>
