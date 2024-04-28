<script lang="ts" setup>
import type { VerticalNavigationLink } from '#ui/types'

const { data, error } = await useCustomFetch(
  '/api/cc/collections/listCollections',
  {
    transform: (data) => {
      return data?.collections.map((col) => {
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
        label: 'Bulk Process',
        to: `${specimenBase}/bulk`,
        icon: 'carbon:data-bin',
      },

      {
        label: 'New Specimen',
        to: `${specimenBase}/New`,
        icon: 'carbon:document-add',
      },
    ] as VerticalNavigationLink[]
  })
})

const groups = computed(() => {
  if (!data.value || data.value.length === 0 || error.value) {
    return []
  }

  return data.value?.map((col) => {
    return {
      label: col.id,
      links: [
        {
          label: 'Specimen Table',
          to: `/collection/${col.id}/SpecimenTable`,
          icon: 'carbon:data-table',
        },
        {
          label: 'Access Control',
          to: `/collection/${col.id}/AccessControl`,
          icon: 'carbon:pedestrian',
        },
        {
          label: 'Bulk Process',
          to: `/collection/${col.id}/Specimen/bulk`,
          icon: 'carbon:data-bin',
        },
        {
          label: 'New Specimen',
          to: `/collection/${col.id}/Specimen/New`,
          icon: 'carbon:document-add',
        },
      ],
    }
  })
})

const hide = useNuxtApp().$auth.isAdmin
</script>

<template>
  <UContainer
    :ui="{
      padding: 'px-0 md:px-0 lg:px-0 sm:px-0',
      constrained: 'max-w-7xl lg:min-w-40 md:min-w-35',
      base: 'border-current bg-surface-200 dark:bg-surface-800 ',
    }"
  >
    <UAccordion
      multiple
      :items="groups"
      :ui="{
        item: {
          padding: 'pt-0 ',

          icon: 'text-gray-700 dark:text-gray-200',
        },
        default: {
          class: 'mb-0 mt-1.5 mr-0.5 w-full ',
        },
      }"
    >
      <template #item="{ item }">
        <UVerticalNavigation
          :links="item.links"
          :ui="{
            wrapper: 'ml-3 border-l-2 border-slate-600 dark:border-slate-200',
          }"
        />
      </template>
    </UAccordion>
  </UContainer>
</template>

<style></style>

<!-- <template v-if="hide">
      <UButton
        block
        label="Create Collection"
        to="/NewCollection"
      />
    </template> -->
