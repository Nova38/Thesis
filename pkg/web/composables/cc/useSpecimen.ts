// export const useRouteSpecimenId = () => {
//   route.params.specimenId.toString()

import type { PlainMessage } from "@bufbuild/protobuf";
import { ccbio } from "saacs-es";

export const makeSpecimenKey = (
  collectionId: string | Ref<string>,
  specimenId: string | Ref<string>,
) => {
  return `collectionId:${toValue(collectionId)}-specimenId:${toValue(
    specimenId,
  )}`;
};
export const makeSpecimenHistoryKey = (
  collectionId: string | Ref<string>,
  specimenId: string | Ref<string>,
) => {
  return `collectionId:${toValue(collectionId)}-specimenId:${toValue(
    specimenId,
  )}-history`;
};
// export const specimenIdKey = useState("currentSpecimen", () => {
//   return `collectionId:${toValue(useRouteCollectionId)}-specimenId:${toValue(
//     useRouteSpecimenId,
//   )}`;
// });

export const useRouteSpecimenId = useState("RouteSpecimenId", () => {
  const route = useRoute();
  return route.params?.specimenId?.toString();
});

// export const useRouteSpecimenId = useState(
//   "currentSpecimen",
//   () => useRoute().params?.specimenId?.toString(),
// );

export const useGetRouteSpecimen = () => {
  return useCustomFetch<ccbio.Specimen>(`/api/cc/specimens/get`, {
    key: makeSpecimenKey(useRouteSpecimenId, useRouteCollectionId),

    query: {
      collectionId: toValue(useRouteSpecimenId),
      specimenId: toValue(useRouteCollectionId),
    },
  });
};

export const useGetSpecimenHistory = () => {
  return useCustomFetch(`/api/cc/specimens/history`, {
    key: makeSpecimenHistoryKey(useRouteSpecimenId, useRouteCollectionId),

    query: {
      collectionId: toValue(useRouteCollectionId),
      specimenId: toValue(useRouteSpecimenId),
    },
  });
};

export const useCreateSpecimen = (specimen: PlainMessage<ccbio.Specimen>) => {
  return useCustomFetch(`/api/cc/specimens/create`, {
    method: "POST",
    body: new ccbio.Specimen(specimen).toJsonString(),
  });
};
