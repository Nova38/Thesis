<script lang="ts" setup>

const shown = ref(true)
const dropdown = ref() // we need a DOM node

</script>

<template>
  <div>
    <fieldset class="border border-gray-300 rounded p-4 relative" v-auto-animate>
    <legend class="text-sm text-gray-700 px-2 bg-white absolute top-0 left-2 -mt-2" @click="()=>{shown = !shown}">
      hi this is a legend
    </legend>
    <div v-if="shown">
          <FormKit
      type="text"
      name="weightUnits"
      label="Weight Units"
    />
    </div>
  </fieldset>
  </div>
</template>

<style scoped>

</style>
