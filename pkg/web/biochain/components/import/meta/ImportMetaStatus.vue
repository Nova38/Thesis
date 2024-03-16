<script lang="ts" setup>
const props = defineProps<{
  row: string
}>()

const bulk = useBulkStore()

const meta = computed(() => {
  return bulk.RawRowsMeta.get(props.row) ?? {
    status: 'error' as ProcessingStatus,
    statusMessage: 'No meta found',
    exist: 'unknown' as ExistStatus,
  }
})
const severity = computed(() => {
  switch (meta.value.status) {
    case 'loading':
      return 'info'
    case 'new':
      return 'info'
    case 'parsing-error':
      return 'warn'
    case 'error':
      return 'error'
    default:
      return 'info'
  }
})
</script>

<template>
  <div>
    <UPopover
      :text="meta.statusMessage"
      mode="hover"
      :popper="{ arrow: true }"
    >
      <PInlineMessage :severity class="text-nowrap">
        {{ meta.status }}
      </PInlineMessage>

      <template #panel>
        <div>
          <div>Message: </div>
          <div>
            {{ severity }}
            <pre>{{ meta.statusMessage }}</pre>
          </div>
        </div>
      </template>
    </UPopover>
  </div>
</template>

<style scoped>

</style>
