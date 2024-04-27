<script lang="ts" setup>
const crumb = computed(() => {
  const route = useRoute()

  const items = [
    {
      icon: 'i-heroicons-home',
      label: 'Home',
      to: '/',
    },
  ]

  if (
    route.name === 'collection-collectionId' ||
    route.name === 'collection-collectionId-AccessControl' ||
    route.name === 'collection-collectionId-SpecimenTable' ||
    route.name === 'collection-collectionId-Specimen-New' ||
    route.name === 'collection-collectionId-Specimen-import' ||
    route.name === 'collection-collectionId-Specimen-update'
  ) {
    items.push({
      // material-symbols:collections-bookmark-outline-rounded
      icon: 'i-material-symbols-collections-bookmark-outline-rounded',
      label: `Collection: ${route.params.collectionId.toString()}`,
      to: `/collection/${route.params.collectionId.toString()}/SpecimenTable`,
    })
  }

  if (route.name === 'collection-collectionId-Specimen-View-specimenId') {
    items.push({
      icon: 'i-heroicons-moon',
      label: `Specimen: ${route.params.specimenId.toString()}`,
      to: `/collection/${route.params.collectionId.toString()}/specimen/view-${route.params.specimenId.toString()}`,
    })
  }

  return items
})
</script>

<template>
  <PToolbar
    class="rounded-[3rem] bg-surface-900 bg-gradient-to-r from-sky-500/70 to-sky-500/80 shadow-md"
  >
    <template #start>
      <!-- <UButton label="Click Me" /> -->
      <PButton class="mr-2">
        <Icon
          name="carbon:menu"
          size="1em"
        />
      </PButton>
      <div class="flex flex-row items-center">
        <!-- <UBreadcrumb :links="crumb" /> -->
        <template
          v-for="item in crumb"
          :key="item.label"
        >
          <span class="px-2">
            <Icon
              name="uil:angle-right"
              size="2em"
            />
          </span>
          <NuxtLink
            :to="item.to"
            active-class="font-bold"
            exact-active-class="font-bold"
          >
            {{ item.label }} &nbsp;
          </NuxtLink>
        </template>
      </div>
    </template>

    <template #center />

    <template #end>
      <ColorModeSelector />
      <AuthUserMenu class="ml-auto" />
    </template>
  </PToolbar>
</template>

<style></style>
