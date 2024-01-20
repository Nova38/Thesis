export default defineEventHandler(async (event) => {
  const cc = await useChaincode(event)

  const result = await cc.service.getCollectionsList()

  return result
})
