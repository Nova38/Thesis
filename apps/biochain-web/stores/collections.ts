// import { defineStore } from "pinia";

export const useCollectionsStore = defineStore('Collections', () => {
  const CollectionId = (): string => {
    return useRoute().params?.collectionId.toString() ?? ''
  }

  const Collection = computed(() => {})

  const Bookmark = ref<string>()
  const Loading = ref(true)

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
    return await useCustomFetch(`/api/cc/specimens/list`, {
      key: `collectionId${CollectionId()}-bookmark${Bookmark.value}`,
      query: {
        collectionId: CollectionId(),
        bookmark: Bookmark.value,
      },
      onResponse: async ({ response }) => {
        console.log('history', response._data)
        Bookmark.value = response._data?.bookmark ?? ''
        console.log(Bookmark.value)
        SpecimenMap.value = defu(SpecimenMap.value, response._data?.specimenMap)
        SpecimenList.value = Object.values(SpecimenMap.value)
        SpecimenCatalogNumbers.value = SpecimenList.value.map(
          (s) => s.primary?.catalogNumber || '',
        )
        SpecimenUUIDs.value = SpecimenList.value.map((s) => s.specimenId)
      },
    })
  }

  async function LoadFull() {
    Loading.value = true
    let lastBookmark = '-'
    while (Bookmark.value !== lastBookmark && Bookmark.value) {
      console.log('loadBatch', Bookmark.value, lastBookmark)
      lastBookmark = Bookmark.value
      const v = await LoadRows()
      if (v.error.value) {
        console.log(v.error)
        break
      }
      if (lastBookmark === Bookmark.value) {
        console.log('no change', Bookmark.value, lastBookmark)
        break
      }
    }
    Loading.value = false
  }

  async function FullListLoad() {
    Loading.value = true

    const response = await $fetch(`/api/cc/specimens/fullList`, {
      key: `collectionId${CollectionId()}-bookmark${Bookmark.value}`,
      query: {
        collectionId: CollectionId(),
        bookmark: Bookmark.value,
      },
    })

    console.log('history', response)
    Bookmark.value = response.bookmark ?? ''
    console.log(Bookmark.value)
    SpecimenMap.value = defu(SpecimenMap.value, response.specimens)
    SpecimenList.value = Object.values(SpecimenMap.value)
    SpecimenCatalogNumbers.value = SpecimenList.value.map(
      (s) => s.primary.catalogNumber,
    )
    SpecimenUUIDs.value = SpecimenList.value.map((s) => s.specimenId)
    Loading.value = false
  }

  async function isUsedUUID(uuid: string) {
    return SpecimenUUIDs?.value?.includes(uuid) || false
  }

  async function Reload() {
    console.log('reloading')
    console.log(CollectionId())
    Bookmark.value = ''
    SpecimenMap.value = {}
    SpecimenList.value = []
    console.log(useRoute().params.collectionId)
    await FullListLoad()
  }

  return {
    CollectionId,
    Collection,
    Bookmark,
    SpecimenList,
    SpecimenMap,
    Loading,
    SpecimenCatalogNumbers,
    isUsedUUID,
    FilterSpecimenList,
    GetSpecimenFromCatalogNumber,
    LoadRows,
    Reload,
    FullListLoad,
    LoadFull,
  }
})
