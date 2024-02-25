<script lang="ts" setup>
// import 'primevue/resources/themes/aura-light-green/theme.css'
</script>

<template>
  <div>
    <div class="w-full h-full min-h-screen flex flex-col gap-2">
      <NavHeader />

      <div class="flex flex-grow flex-row gap-2 flex-nowrap max-w-full ">
        <NavSidebar class=" " />

        <div class="bg-slate-100 flex-grow overflow-scroll 	">
          <NuxtLoadingIndicator />

          <NuxtErrorBoundary>
            <NuxtPage />
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
