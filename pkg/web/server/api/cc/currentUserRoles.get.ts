import { z } from 'zod'
import { auth, common } from '~/lib'

const querySchema = z.object({
  collectionId: z.string(),
})
export default defineEventHandler(async (event) => {
  const cc = await useChaincode(event)
  const query = await getValidatedQuery(event, body =>
    querySchema.safeParse(body))
  console.log(query)
  if (!query.success)
    throw query.error.issues

  const { user } = await cc.service.getCurrentUser()
  if (!user)
    throw new Error('User not found')

  const result = await cc.service.get(
    new common.generic.GetRequest({
      key: new auth.objects.ItemKey({
        collectionId: query.data.collectionId,
        itemKeyParts: [user.mspId, user.userId],
        itemType: auth.objects.UserCollectionRoles.typeName,
      }),
    }),
  )
  console.log(result)

  const role = new auth.objects.UserCollectionRoles()
  result.item?.value?.unpackTo(role)

  return role
})
