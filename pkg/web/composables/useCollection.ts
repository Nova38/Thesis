export const useCollectionId = () => {
  return computed(() => useRoute().params?.collectionId ?? null);
};

export const useUseCollection = () => {
  return ref();
};
