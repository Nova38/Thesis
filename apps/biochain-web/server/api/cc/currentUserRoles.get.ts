import { z } from 'zod'
import HandleFabricError from '~/server/utils/FabricErrors'
// import { auth, common } from 'saacs'

const querySchema = z.object({
  collectionId: z.string(),
})
export default defineEventHandler(async (event) => {
  const cc = await useChaincode(event)
  const query = await getValidatedQuery(event, querySchema.parse)

  const { user } = await cc.utilService.getCurrentUser({})
  if (!user) throw new Error('User not found')

  try {
    const result = await cc.service.get(
      new pb.GetRequest({
        key: new pb.ItemKey({
          collectionId: query.collectionId,
          itemKeyParts: [user.mspId, user.userId],
          itemType: pb.UserCollectionRoles.typeName,
        }),
      }),
    )
    console.log(result)

    const role = new pb.UserCollectionRoles()
    result.item?.value?.unpackTo(role)

    return role
  } catch (error) {
    const e = HandleFabricError(error)

    // console.log(e)
    return createError(e)
  }
})
