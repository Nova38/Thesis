<script lang="ts" setup>
const value = ref()
</script>

<template>
  <div>
    <FormKit
      v-model="value"
      type="select"
      label="Where would you prefer to live?"
      name="planet"
    >
      <optgroup label="Inner Planets">
        <option value="mercury">
          Mercury
        </option>
        <option value="venus">
          Venus
        </option>
        <option value="earth">
          Earth
        </option>
        <option value="mars">
          Mars
        </option>
      </optgroup>
      <optgroup label="Outer planets">
        <option value="jupiter">
          Jupiter
        </option>
        <option value="saturn">
          Saturn
        </option>
        <option value="uranus">
          Uranus
        </option>
        <option value="neptune">
          Neptune
        </option>
      </optgroup>
    </FormKit>
    <pre wrap>{{ value }}</pre>
  </div>
</template>

<style scoped></style>
