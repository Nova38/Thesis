export const useCollectionId = computed<string | null>(() => {
  if (useRoute().params.collectionId) {
    return useRoute().params.collectionId;
  }
  return null;
});

export const useUseCollection = () => {
  return ref();
};
