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

<script lang="ts" setup></script>
