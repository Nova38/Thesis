import { z } from 'zod'
// import { auth, common } from 'saacs'

import { useChaincode } from '~/server/utils/useChaincode'

const querySchema = z.object({
  collectionId: z.string(),
  userName: z.string(),
  userId: z.string().optional(),
})

export default defineEventHandler(async (event) => {
  const cc = await useChaincode(event)

  const query = await getValidatedQuery(event, querySchema.parse)

  // console.log(result);
  const UserRoles = result.items.map((i) => {
    const s = new pb.UserCollectionRoles()
    i.value?.unpackTo(s)
    return s.toJson({ emitDefaultValues: true })
  })

  console.log({ UserRoles })
  return UserRoles
})
