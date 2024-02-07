export default defineNuxtRouteMiddleware((to, _) => {
  type RouteParams = Exclude<typeof to.params, Record<never, never>>

  if (to.params as RouteParams) {
    console.log('collectionId', to.params)
    // useBulkUpdate.CollectionId.value = to.params.collectionId
  }
})
