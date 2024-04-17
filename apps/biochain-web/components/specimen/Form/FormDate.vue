<script lang="ts" setup>
import { titleCase } from 'scule'

import { type FormKitNode } from '@formkit/core'

const props = defineProps<{
  name: string
}>()

const collapsed = defineModel<boolean>('collapsed', { default: true })

const c = ref(false)
const wrapperClass = computed(() => {
  return {
    'grid-col-span-3 col-start-1': c.value,
    'grid-col-span-1': !c.value,
  }
})

const nameAsTitle = computed(() => {
  return titleCase(props.name)
})

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
  <div
    :class="wrapperClass"
    @click="
      (e) => {
        console.log(e)
      }
    "
  >
    <PFieldset
      :legend="nameAsTitle"
      :toggleable="true"
      :collapsed
      class="max-w-fit"
      :pt="{
        root: {
          class: [
            'block',

            // Spacing
            'px-5 md:px-6 py-5',

            // Shape
            'rounded-lg',

            // Color
            'bg-surface-100 dark:bg-surface-900',
            'text-surface-700 dark:text-surface-0/80',
            'ring-1 ring-inset ring-surface-300 dark:ring-surface-700 ring-offset-0',
          ],
        },
      }"
      @toggle="
        (e) => {
          console.log('toggle', e)
        }
      "
    >
      <FormKit
        :id="props.name"
        type="group"
        :name="props.name"
        :plugins="[formPlug]"
      >
        <div class="inline-flex flex-wrap gap-2">
          <FormKit
            id="verbatim"
            type="text"
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
            id="timestamp"
            type="text"
            name="timestamp"
            label="timestamp"
            outer-class="min-w-20"
          />
          <FormKit
            id="year"
            type="number"
            name="year"
            number
            label="year"
            outer-class="min-w-20"
          />
          <FormKit
            id="month"
            type="text"
            name="month"
            label="month"
            outer-class="min-w-20"
          />
          <FormKit
            id="day"
            type="number"
            name="day"
            number
            label="day"
            outer-class="min-w-20"
          />
        </div>
      </FormKit>
    </PFieldset>
  </div>
</template>

<style scoped></style>
