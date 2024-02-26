<script lang="ts" setup>
// import 'primevue/resources/themes/aura-light-green/theme.css'
</script>

<template>
  <div>
    <div class="w-full h-full min-h-screen flex flex-col gap-2">
      <NavHeader />

      <div class="flex flex-grow flex-row flex-nowrap max-w-full ">
        <NavSidebar class="bg-zinc-100 dark:bg-zinc-800 " />

        <div class="flex-grow overflow-scroll 	">
          <NuxtLoadingIndicator />

          <NuxtErrorBoundary>
            <NuxtPage class="bg-zinc-50 dark:bg-zinc-900"/>
          </NuxtErrorBoundary>
        </div>
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
