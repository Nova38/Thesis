<script lang="ts" setup>
const shown = ref(true)
const dropdown = ref() // we need a DOM node
</script>

<template>
  <div>
    <fieldset
      v-auto-animate
      class="relative rounded border border-gray-300 p-4"
    >
      <legend
        class="absolute left-2 top-0 -mt-2 bg-white px-2 text-sm text-gray-700"
        @click="
          () => {
            shown = !shown
          }
        "
      >
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

<style scoped></style>
