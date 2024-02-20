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
    const base = `/collection/${col.id}/`

    return [{
      label: col.id,
      to: base,
    }, {
      label: 'Specimen Table',
      to: `${base}/SpecimenTable`,
    }, {
      label: 'Access Control',
      to: `${base}/access`,
    }, {
      label: 'Bulk Import',
      to: `${base}/bulk/import`,
    }, {
      label: 'Bulk Update',
      to: `${base}/bulk/update`,
    }] as VerticalNavigationLink[]
  })
})
</script>

<template>
  <div class="min-w-40">
    <UVerticalNavigation
      :links="links"
      class=" "
    />
  </div>
</template>

<style></style>
