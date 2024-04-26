<script lang="ts" setup>
import { titleCase } from 'scule'

const props = defineProps<{
  name: string
}>()

const collapsed = defineModel<boolean>('collapsed', {
  default: true,
})

const c = ref(false)
const wrapperClass = computed(() => {
  return {
    'grid-col-span-3 col-start-1': !c.value,
    'grid-col-span-1': c.value,
  }
})

const nameAsTitle = computed(() => {
  return titleCase(props.name)
})
</script>

<template>
  <PFieldset
    :legend="nameAsTitle"
    :toggleable="true"
    :collapsed="collapsed"
    class="col-span-full col-start-1 max-w-fit"
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
    @update:collapsed="collapsed = $event"
    @toggle="
      (e) => {
        console.log('toggle', e)
      }
    "
  >
    <FormKit
      type="group"
      :name="props.name"
      :plugins="[]"
    >
      <div class="inline-flex flex-wrap gap-2">
        <FormKit
          name="verbatim"
          type="text"
          label="verbatim"
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
          label="timestamp"
          outer-class="min-w-20"
        />
        <FormKit
          type="number"
          name="year"
          number
          label="year"
          outer-class="min-w-20"
        />
        <FormKit
          type="text"
          name="month"
          label="month"
          outer-class="min-w-20"
        />
        <FormKit
          type="number"
          name="day"
          number
          label="day"
          outer-class="min-w-20"
        />
      </div>
    </FormKit>
  </PFieldset>
</template>

<style scoped></style>
