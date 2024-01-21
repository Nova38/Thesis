import type { BreadcrumbLink } from '@nuxt/ui/dist/runtime/types'

export function useBreadcrumbLinks() {
  const route = useRoute()
  const links = computed(() => {
    const paths = route.path.split('/').filter(p => p)

    const items: BreadcrumbLink[] = [
      {
        icon: 'i-heroicons-home',
        label: 'Home',
        to: '/',
      },
    ]

    // Check for the collection route
    if (route.params?.collectionId) {
      items.push({
        // material-symbols:collections-bookmark-outline-rounded
        icon: 'i-material-symbols-collections-bookmark-outline-rounded',
        label: `Collection: ${route.params.collectionId}`,
        to: `/collection/${route.params.collectionId}/SpecimenTable`,
      })
    }
    // if (paths.find((p) => p === "SpecimenTable")) {
    //   // items.push({
    //   //     label: `Specimen`,
    //   //     to: `/specimen`,
    //   //     icon: "i-heroicons-moon",
    //   // });

    //   items.push({
    //     label: `Specimen Table`,
    //     to: `/collection/${route.params.collectionId}/SpecimenTable`,
    //     // material-symbols:collections-bookmark-outline-rounded
    //     icon: "i-material-symbols-collections-bookmark-outline-rounded",
    //   });
    // }

    // Handle the Specimen route
    if (paths.find(p => p === 'Specimen')) {
      // items.push({
      //     label: `Specimen`,
      //     to: `/specimen`,
      //     icon: "i-heroicons-moon",
      // });

      if (route.params?.specimenId) {
        // See if we have a speciemen id
        items.push({
          icon: 'i-heroicons-moon',
          label: `Specimen: ${route.params.specimenId}`,
          to: `/collection/${route.params.collectionId}/specimen/view-${route.params.specimenId}`,
        })
      }
    }

    // items.push({

    console.log('items', items)
    console.log('paths', paths)

    return items
    //
    // const links = paths.map((p, i) => {
    //     const link = {
    //         title: p,
    //         to: "/" + paths.slice(0, i + 1).join("/"),
    //     };
    //     return link;
    // });
    // return links;
  })

  return links
}
