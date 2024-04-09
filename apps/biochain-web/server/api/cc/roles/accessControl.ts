import { z } from 'zod'
// import { auth, common } from 'saacs'

import { useChaincode } from '~/server/utils/useChaincode'

const querySchema = z.object({
  collectionId: z.string(),
})

export default defineEventHandler(async (event) => {
  const cc = await useChaincode(event)

  const query = await getValidatedQuery(event, (body) =>
    querySchema.safeParse(body),
  )
  if (!query.success) throw query.error.issues
  console.log({ data: query.data })

  const r1 = await cc.service.listByAttrs(
    new common.generic.ListByAttrsRequest({
      key: new auth.objects.ItemKey({
        collectionId: query.data.collectionId,
        itemKeyParts: [query.data.collectionId],
        itemType: auth.models.UserCollectionRoles.typeName,
      }),
      numAttrs: 1,
    }),
  )

  const UserRoles = r1.items.map((i) => {
    const s = new auth.models.UserCollectionRoles()
    i.value?.unpackTo(s)
    return s.toJson({ emitDefaultValues: true })
  })

  const r2 = await cc.service.listByAttrs(
    new common.generic.ListByAttrsRequest({
      key: new auth.objects.ItemKey({
        collectionId: query.data.collectionId,
        itemKeyParts: [query.data.collectionId],
        itemType: auth.models.Role.typeName,
      }),
      numAttrs: 0,
    }),
  )
  const Roles = r2.items.map((i) => {
    const o = i.toJsonString({ typeRegistry: cc.service.registry })
    console.log(o)

    const s = new auth.models.Role()
    i.value?.unpackTo(s)
    return s.toJson({ emitDefaultValues: true })
  })

  const res = { Roles, UserRoles }

  console.log(res)
  return res
})
