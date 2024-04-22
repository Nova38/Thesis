<script setup lang="ts">
import { ccbio } from '#imports'

// const items = [
//   {
//     label: 'Getting Started',
//     icon: 'i-heroicons-information-circle',
//     tx: '33434343434',
//   },
//   {
//     label: 'Installation',
//     icon: 'i-heroicons-arrow-down-tray',
//     tx: '33434343434',
//   },
//   {
//     label: 'Theming',
//     icon: 'i-heroicons-eye-dropper',
//     tx: '33434343434',
//   },
//   {
//     label: 'Layouts',
//     icon: 'i-heroicons-rectangle-group',
//     tx: '33434343434',
//   },
//   {
//     label: 'Components',
//     icon: 'i-heroicons-square-3-stack-3d',
//     tx: '33434343434',
//   },
//   {
//     label: 'Utilities',
//     icon: 'i-heroicons-wrench-screwdriver',
//     tx: '33434343434',
//   },
// ]
const props = withDefaults(
  defineProps<{
    history: ccbio.SpecimenHistory
  }>(),
  {
    history: () => new ccbio.SpecimenHistory(),
  },
)

const items = computed(() => {
  return props.history.entries.map((entry) => {
    return {
      label: entry.timestamp?.toDate().toDateString() ?? 'Unknown',
      icon: 'i-heroicons-calendar',
      tx: entry.txId,
      specimen: entry.value,
    }
  })
})

const hideTransaction = (tx: string) => {
  console.log('hide', tx)
}
</script>

<template>
  <UAccordion
    :items="items"
    :ui="{ wrapper: 'flex flex-col w-full' }"
    multiple
  >
    <template #default="{ item, index, open }">
      <UButton
        color="gray"
        variant="ghost"
        class="border-b border-gray-200 dark:border-gray-700"
        :ui="{ rounded: 'rounded-none', padding: { sm: 'p-3' } }"
      >
        <template #leading>
          <div
            class="bg-primary-500 dark:bg-primary-400 -my-1 flex h-6 w-6 items-center justify-center rounded-full"
          >
            <UIcon
              :name="item.icon"
              class="h-4 w-4 text-white dark:text-gray-900"
            />
          </div>
        </template>

        <span class="truncate">{{ index + 1 }}. {{ item.label }}</span>

        <template #trailing>
          <UIcon
            name="i-heroicons-chevron-right-20-solid"
            class="ms-auto h-5 w-5 transform transition-transform duration-200"
            :class="[open && 'rotate-90']"
          />
        </template>
      </UButton>
    </template>
    <template #item="{ item }">
      <UCard
        class="min-w-2xl mt-4 max-w-3xl"
        :ui="{ root: 'bg-surface-100 dark:bg-surface-800' }"
      >
        <FormKit
          disabled
          type="form"
          :actions="false"
        >
          <SpecimenForm
            :specimen="item.specimen"
            :form-prefix="`specimen-${item.tx}`"
          />
        </FormKit>
        <template
          v-if="$auth.loggedIn"
          #footer
        >
          <UButton
            color="red"
            variant="ghost"
            class="w-full"
            @click="hideTransaction(item.tx)"
          >
            Hide Transaction
          </UButton>
        </template>
      </UCard>
    </template>
  </UAccordion>
</template>
