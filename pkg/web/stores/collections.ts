// import { defineStore } from "pinia";

import { auth, ccbio } from "saacs-es";

export const useCollectionsStore = defineStore("Collections", () => {
  const CollectionId =
    useNuxtApp().$router.currentRoute.value.params.collectionId;
  const c = computed(() => {
    return useRoute().params.collectionId ?? "";
  });
  const cN = computed(() => {
    return useNuxtApp().$router.currentRoute.value.params.collectionId ?? "";
  });

  const Collection = computed(() => {});

  const Bookmark = ref<string>();
  const Loading = ref(true);

  const SpecimenList = ref<PlainSpecimen[]>([]);
  const SpecimenMap = ref<Record<string, PlainSpecimen>>();
  const SpecimenCatalogNumbers = ref<string[]>();

  const SpecimenUUIDs = ref<string[]>();

  function GetSpecimenFromCatalogNumber(catalogNumber: string) {
    for (const s of SpecimenList.value) {
      if (s.primary.catalogNumber === catalogNumber) {
        return s;
      }
    }
  }

  async function FilterSpecimenList({
    uuids,
    catalogNumbers,
  }: {
    uuids?: string[];
    catalogNumbers?: string[];
  }) {
    if (uuids) {
      SpecimenList.value = SpecimenList.value?.filter((s) =>
        uuids.includes(s.specimenId),
      );
    }
    if (catalogNumbers) {
      SpecimenList.value = SpecimenList.value?.filter((s) =>
        catalogNumbers.includes(s.primary.catalogNumber),
      );
    }
  }

  async function LoadRows() {
    return await useCustomFetch<PlainSpecimen[]>(`/api/cc/specimens/list`, {
      key: `collectionId${CollectionId}-bookmark${Bookmark.value}`,
      query: {
        collectionId: CollectionId,
        bookmark: Bookmark.value,
      },
      onResponse: async ({ response }) => {
        console.log("history", response._data);
        Bookmark.value = response._data?.bookmark ?? "";
        console.log(Bookmark.value);
        SpecimenMap.value = defu(
          SpecimenMap.value,
          response._data?.specimenMap,
        );
        SpecimenList.value = Object.values(SpecimenMap.value);
        SpecimenCatalogNumbers.value = SpecimenList.value.map(
          (s) => s.primary.catalogNumber,
        );
        SpecimenUUIDs.value = SpecimenList.value.map((s) => s.specimenId);
      },
    });
  }

  async function LoadFull() {
    Loading.value = true;
    let lastBookmark = "-";
    while (Bookmark.value !== lastBookmark && Bookmark.value) {
      console.log("loadBatch", Bookmark.value, lastBookmark);
      lastBookmark = Bookmark.value;
      const v = await LoadRows();
      if (v.error.value) {
        console.log(v.error);
        break;
      }
      if (lastBookmark === Bookmark.value) {
        console.log("no change", Bookmark.value, lastBookmark);
        break;
      }
    }
    Loading.value = false;
  }

  async function isUsedUUID(uuid: string) {
    return SpecimenUUIDs?.value?.includes(uuid) || false;
  }

  async function Reload() {
    Bookmark.value = "";
    SpecimenMap.value = {};
    SpecimenList.value = [];
    await LoadFull();
  }

  return {
    c,
    cN,
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
  };
});
