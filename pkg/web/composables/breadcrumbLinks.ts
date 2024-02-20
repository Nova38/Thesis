export function useBreadcrumbLinks() {
  const links = computed(() => {
    const route = useRoute()

    const items = [
      {
        icon: 'i-heroicons-home',
        label: 'Home',
        to: '/',
      },
    ]

    if (
      route.name === 'collection-collectionId'
      || route.name === 'collection-collectionId-AccessControl'
      || route.name === 'collection-collectionId-SpecimenTable'
      || route.name === 'collection-collectionId-Specimen-New'
      || route.name === 'collection-collectionId-Specimen-import'
      || route.name === 'collection-collectionId-Specimen-update'
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
  return links
}
