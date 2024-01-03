<template>
  <div>hello</div>
</template>

<script lang="ts" setup>
definePageMeta({
  layout: "auth",
  key(route) {
    return route.fullPath;
  },
});
</script>

<style></style>
