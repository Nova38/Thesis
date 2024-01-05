// export const useRouteSpecimenId = () => {
//   route.params.specimenId.toString()

export const specimenIdKey = () => {
  const route = useRoute();

  return `collectionId:${route.params.collectionId.toString()}-specimenId:${route.params.specimenId.toString()}`;
};

export const useGetSpecimen = async () => {
  const route = useRoute();
  return await useCustomFetch(`/api/cc/specimens/get`, {
    key: specimenIdKey(),

    query: {
      collectionId: route.params.collectionId,
      specimenId: route.params.specimenId,
    },
  });
};

export const useGetSpecimenHistory = async () => {
  const route = useRoute();
  return await useCustomFetch(`/api/cc/specimens/history`, {
    key: `${specimenIdKey()}-history`,

    query: {
      collectionId: route.params.collectionId,
      specimenId: route.params.specimenId,
    },
  });
};
