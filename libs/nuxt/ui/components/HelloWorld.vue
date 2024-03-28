<script setup lang="ts">
const { myLayer } = useAppConfig()
console.log('Hello World')
</script>

<template>
  <div>
    <h1>Hello World!</h1>
    <pre>{{ myLayer }}</pre>
    <ColorModeSelector />
    <FormKit name="csvFile" type="file" label="Upload CSV" accept=".csv" />
  </div>
</template>
