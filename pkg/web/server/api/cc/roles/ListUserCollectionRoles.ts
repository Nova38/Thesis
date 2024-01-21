import { auth, common } from 'saacs-es'
import { z } from 'zod'

import { useChaincode } from '~/server/utils/useChaincode'

const querySchema = z.object({
  collectionId: z.string(),
})

export default defineEventHandler(async (event) => {
  const cc = await useChaincode(event)

  const query = await getValidatedQuery(event, body =>
    querySchema.safeParse(body))
  if (!query.success)
    throw query.error.issues
  console.log({ data: query.data })

  const result = await cc.service.listByAttrs(
    new common.generic.ListByAttrsRequest({
      key: new auth.objects.ItemKey({
        collectionId: query.data.collectionId,
        itemKeyParts: [query.data.collectionId],
        itemType: auth.objects.UserCollectionRoles.typeName,
      }),
      numAttrs: 1,
    }),
  )

  // console.log(result);
  const UserRoles = result.items.map((i) => {
    const s = new auth.objects.UserCollectionRoles()
    i.value?.unpackTo(s)
    return s.toJson({ emitDefaultValues: true })
  })

  console.log({ UserRoles })
  return UserRoles
})
