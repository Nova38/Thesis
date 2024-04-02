<script lang="ts" setup>
import { titleCase } from 'scule'

const props = defineProps<{
  name: string
}>()

const nameAsTitle = computed(() => {
  return titleCase(props.name)
})

import { type FormKitNode } from '@formkit/core'

const formPlug = (node: FormKitNode) => {
  if (!node.props.id) return
  console.log('AutoPropsFromIdPlugin', node.props.id)
  node.name = node.props.id // auto set name to id
  node.props.label = titleCase(node.props.id) // auto set label
  if (!['button', 'submit'].includes(node.props.type)) {
    // automatically set help text, but exclude buttons
  }
}
</script>

<template>
  <div>
    <PFieldset
      :legend="nameAsTitle"
      :toggleable="true"
      class="max-w-fit"
      :pt="{
        root: {
          class: [
            'block',

            // Spacing
            'px-5 md:px-6 py-5',

            // Shape
            'rounded-md rounded-lg',

            // Color
            'bg-surface-100 dark:bg-surface-900',
            'text-surface-700 dark:text-surface-0/80',
            'ring-1 ring-inset ring-surface-300 dark:ring-surface-700 ring-offset-0',
          ],
        },
      }"
    >
      <FormKit
        type="group"
        :id="props.name"
        :name="props.name"
        :plugins="[formPlug]"
      >
        <div class="inline-flex flex-wrap gap-2">
          <FormKit
            type="text"
            id="verbatim"
            outer-class="min-w-20"
          />
          <!-- <FormKit
            type="text"
            name="verbatim"
            id="verbatim"
            label="verbatim"
            outer-class="min-w-20"
          /> -->

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
