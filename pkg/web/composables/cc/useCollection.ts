export const useRouteCollectionId = useRoute().params?.collectionId.toString();

export const useRouteCollection = useState("RouteCollection", () => {
  useRoute().params?.collectionId.toString();
});
