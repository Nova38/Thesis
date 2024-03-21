<script lang="ts" setup>
const props = defineProps<{
  row: string
}>()

const bulk = useBulkStore()

const exists = computed(() => {
  return (
    bulk.RawRowsMeta.get(props.row)?.exist ?? {
      exist: 'unknown',
    }
  )
})

const severity = computed(() => {
  switch (bulk.Mode) {
    case 'update':
      switch (exists.value) {
        case 'new':
          return 'error'
        case 'pre-existing':
          return 'success'
        case 'unknown':
          return 'warn'
        default:
          return 'error'
      }
    case 'import':
      switch (exists.value) {
        case 'new':
          return 'success'
        case 'pre-existing':
          return 'error'
        case 'unknown':
          return 'warn'
        default:
          return 'error'
      }
    case 'hybrid':
      switch (exists.value) {
        case 'new':
          return 'success'
        case 'pre-existing':
          return 'info'
        case 'unknown':
          return 'warn'
        default:
          return 'error'
      }
      break
    default:
      return 'error'
  }
})
</script>

<template>
  <div>
    <PInlineMessage :severity class="text-nowrap">
      {{ exists }}
    </PInlineMessage>
  </div>
</template>

<style scoped></style>
