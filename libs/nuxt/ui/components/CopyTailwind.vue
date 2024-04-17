<script lang="ts" setup>
import tailwindConfig from '#tailwind-config'
import { useClipboard } from '@vueuse/core'

const config = JSON.stringify(tailwindConfig, null, 2)

const { text, copy, copied, isSupported } = useClipboard()
</script>

<template>
  <div>
    <UButton
      label="Copy Tailwind Config"
      @click="copy(config)"
    />
  </div>
</template>

<style scoped></style>
