// import { defineStore } from "pinia";

import toast from '~/primevue/presets/wind/toast'

export const useCollectionsStore = defineStore('Collections', () => {
  const CollectionId = (): string => {
    return useRoute().params?.collectionId.toString() ?? ''
  }

  const Collection = computed(() => {})

  const Bookmark = ref<string>()

  // If the system is actively loading more values
  const Loading = ref(true)
  // Signal to stop loading on page or by button
  const AbortLoading = ref(false)
  // If the table has fully loaded
  const FullyLoaded = ref(false)

  const SpecimenList = ref<PlainSpecimen[]>([])
  const SpecimenMap = ref<Record<string, PlainSpecimen>>()
  const SpecimenCatalogNumbers = ref<string[]>()

  const SpecimenUUIDs = ref<string[]>()

  function GetSpecimenFromCatalogNumber(catalogNumber: string) {
    for (const s of SpecimenList.value) {
      if (s.primary?.catalogNumber === catalogNumber) return s
    }
  }

  async function FilterSpecimenList({
    catalogNumbers,
    uuids,
  }: {
    catalogNumbers?: string[]
    uuids?: string[]
  }) {
    if (uuids) {
      SpecimenList.value = SpecimenList.value?.filter((s) =>
        uuids.includes(s.specimenId),
      )
    }
    if (catalogNumbers) {
      SpecimenList.value = SpecimenList.value?.filter((s) =>
        catalogNumbers.includes(s.primary?.catalogNumber || ''),
      )
    }
  }

  async function LoadRows() {
    // return await useCustomFetch(`/api/cc/specimens/list`, {
    //   key: `collectionId${CollectionId()}-bookmark${Bookmark.value}`,
    //   query: {
    //     collectionId: CollectionId(),
    //     bookmark: Bookmark.value,
    //   },
    //   onResponse: async ({ response }) => {
    //     console.log('history', response._data)
    //     Bookmark.value = response._data?.bookmark ?? ''
    //     console.log(Bookmark.value)
    //     SpecimenMap.value = defu(SpecimenMap.value, response._data?.specimenMap)
    //     SpecimenList.value = Object.values(SpecimenMap.value)
    //     SpecimenCatalogNumbers.value = SpecimenList.value.map(
    //       (s) => s.primary?.catalogNumber || '',
    //     )
    //     SpecimenUUIDs.value = SpecimenList.value.map((s) => s.specimenId)
    //   },
    // })
  }

  async function LoadFull() {
    Loading.value = true
    let lastBookmark = '-'
    console.log('Starting loadFull', Bookmark.value, lastBookmark)

    const responses = []

    while (Bookmark.value !== lastBookmark) {
      console.log('loadBatch', Bookmark.value, lastBookmark)
      lastBookmark = Bookmark.value ?? ''

      useToast().add({
        title: 'Loading',
        description: `Loading more specimens...\n Current Bookmark: ${Bookmark.value}`,
        timeout: 5000,
      })

      const response = await $fetch(`/api/cc/specimens/list`, {
        key: `collectionId${CollectionId()}-bookmark${Bookmark.value}`,
        query: {
          collectionId: CollectionId(),
          bookmark: Bookmark.value,
        },
      })
      Bookmark.value = response.bookmark ?? ''

      responses.push(response.specimens)

      if (AbortLoading.value) {
        AbortLoading.value = false
        Loading.value = false

        useToast().add({
          title: 'Aborting Loading Operation',
          timeout: 5000,
          color: 'amber',
        })

        break
      }

      if (lastBookmark === Bookmark.value) {
        useToast().add({
          title: 'Done Fetching',
          description: 'No new specimens to fetch\nProcessing data...',
          timeout: 5000,
          color: 'cyan',
        })
        console.log('no change', Bookmark.value, lastBookmark)
        break
      }
    }

    const specimensResponse = responses.reduce((acc, r) => defu(acc, r), {})

    SpecimenMap.value = specimensResponse
    SpecimenList.value = Object.values(SpecimenMap.value)
    SpecimenCatalogNumbers.value = SpecimenList.value.map(
      (s) => s.primary?.catalogNumber || '',
    )
    SpecimenUUIDs.value = SpecimenList.value.map((s) => s.specimenId)

    FullyLoaded.value = true
    Loading.value = false

    useToast().add({
      title: 'Loading Complete',
      timeout: 5000,
      color: 'green',
    })
  }

  // async function FullListLoad() {
  //   Loading.value = true

  //   const response = await $fetch(`/api/cc/specimens/fullList`, {
  //     key: `collectionId${CollectionId()}-bookmark${Bookmark.value}`,
  //     query: {
  //       collectionId: CollectionId(),
  //       bookmark: Bookmark.value,
  //     },
  //   })

  //   console.log('history', response)
  //   Bookmark.value = response.bookmark ?? ''
  //   console.log(Bookmark.value)
  //   SpecimenMap.value = defu(SpecimenMap.value, response.specimens)
  //   SpecimenList.value = Object.values(SpecimenMap.value)
  //   SpecimenCatalogNumbers.value = SpecimenList.value.map(
  //     (s) => s.primary.catalogNumber,
  //   )
  //   SpecimenUUIDs.value = SpecimenList.value.map((s) => s.specimenId)
  //   Loading.value = false
  // }

  async function isUsedUUID(uuid: string) {
    return SpecimenUUIDs?.value?.includes(uuid) || false
  }

  async function ResumeLoading() {
    AbortLoading.value = false

    return await LoadFull
  }

  async function Reload() {
    console.log('reloading')
    console.log(CollectionId())
    Bookmark.value = ''
    SpecimenMap.value = {}
    SpecimenList.value = []
    console.log(useRoute().params.collectionId)
    // await FullListLoad()
    await LoadFull()
  }

  return {
    CollectionId,
    Collection,
    Bookmark,
    SpecimenList,
    SpecimenMap,
    SpecimenCatalogNumbers,
    isUsedUUID,
    FilterSpecimenList,
    GetSpecimenFromCatalogNumber,
    LoadRows,
    Reload,
    Loading,
    FullyLoaded,
    AbortLoading,
    LoadFull,
    ResumeLoading,
  }
})
