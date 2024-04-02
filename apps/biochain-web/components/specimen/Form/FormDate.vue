<script lang="ts" setup>
import { titleCase } from 'scule'

const props = defineProps<{
  name: string
}>()

const nameAsTitle = computed(() => {
  return titleCase(props.name)
})
</script>

<template>
  <div>
    <PFieldset
      :legend="nameAsTitle"
      :toggleable="true"
      class="max-w-fit"
      :pt="{
        // root: 'bg-surface-50',
      }"
    >
      <FormKit
        type="group"
        :id="props.name"
        :name="props.name"
      >
        <div class="inline-flex flex-wrap gap-2">
          <FormKit
            type="text"
            name="verbatim"
            id="verbatim"
            label="verbatim"
            outer-class="min-w-20"
          />
          <FormKit
            type="text"
            name="timestamp"
            id="timestamp"
            label="timestamp"
            outer-class="min-w-20"
          />
          <FormKit
            type="text"
            name="year"
            id="year"
            label="year"
            outer-class="min-w-20"
          />
          <FormKit
            type="text"
            name="month"
            id="month"
            label="month"
            outer-class="min-w-20"
          />
          <FormKit
            type="text"
            name="day"
            id="day"
            label="day"
            outer-class="min-w-20"
          />
        </div>
      </FormKit>
    </PFieldset>
  </div>
</template>

<style scoped></style>
