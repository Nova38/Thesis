<script lang="ts" setup>
// import 'primevue/resources/themes/aura-light-green/theme.css'

// import { devtools } from '@vue/devtools'
// if (process.env.NODE_ENV === 'development')
//   devtools.connect(/* host (the default is "http://localhost"), port (the default is 8090) */)

// Import fully resolved config
</script>

<template>
  <div>
    <div class="flex min-h-screen w-full flex-col">
      <NavHeader class="sticky" />
      <div class="flex max-w-full flex-grow flex-row flex-nowrap p-0.5">
        <NavSidebar class="bg-zinc-100 dark:bg-zinc-800" />

        <div class="flex-grow overflow-x-scroll">
          <NuxtLoadingIndicator />

          <!-- <NuxtErrorBoundary> -->
          <NuxtPage class="p-2" />
          <!-- </NuxtErrorBoundary> -->
        </div>
        <UNotifications />
      </div>
    </div>
  </div>
</template>
<!-- <NuxtLayout>
    <NuxtLoadingIndicator />

    <NuxtPage />
  </NuxtLayout>
</template> -->

<style>
/* html {
  font-size: 12px;
} */
</style>
