<script lang="ts" setup>
// const { data, pending, error } = await useFetch(
//   "/api/cc/collections/listCollections",
// );

import type { VerticalNavigationLink } from '#ui/types'

const { data, error, pending } = await useCustomFetch(
  '/api/cc/collections/listCollections',
  {
    transform: (data) => {
      return data.collections.map((col) => {
        return {
          col,
          id: col.collectionId,
          to: `/collection-${col.collectionId}`,
        }
      })
    },
  },
)

const links = computed(() => {
  if (!data.value || data.value.length === 0 || error.value) {
    return [
      {
        label: 'Home.',
        to: '/',
      },
    ] as VerticalNavigationLink[]
  }

  return data.value?.map((col) => {
    const base = `/collection/${col.id}`
    const specimenBase = `${base}/Specimen`

    return [
      {
        label: col.id,
        to: base,
      },
      {
        label: 'Specimen Table',
        to: `${base}/SpecimenTable`,
        icon: 'carbon:data-table',
      },
      {
        label: 'Access Control',
        to: `${base}/AccessControl`,
        icon: 'carbon:pedestrian',
      },
      {
        label: 'Bulk Import',
        to: `${specimenBase}/bulk/import`,
        icon: 'carbon:data-bin',
      },
      {
        label: 'Bulk Update',
        to: `${specimenBase}/bulk/import`,
        icon: 'carbon:data-backup',
      },
      {
        label: 'New Specimen',
        to: `${specimenBase}/New`,
        icon: 'carbon:document-add',
      },
    ] as VerticalNavigationLink[]
  })
})
</script>

<template>
  <div class="min-w-40">
    <UVerticalNavigation
      :links="links"
      :ui="{
        wrapper: 'border-s border-gray-200 dark:border-gray-800 space-y-2',
        base: 'group block border-s -ms-px leading-6 before:hidden',
        padding: 'p-0 ps-4',
        rounded: '',
        font: '',
        active:
          'text-primary-500 dark:text-primary-400 border-current font-semibold',
        inactive:
          'border-transparent hover:border-gray-400 dark:hover:border-gray-500 text-gray-700 hover:text-gray-900 dark:text-gray-400 dark:hover:text-gray-300',
      }"
    />
  </div>
</template>

<style></style>
