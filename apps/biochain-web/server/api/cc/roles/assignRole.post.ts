import { saacs } from '@saacs/client'
import { z } from 'zod'
// import { auth, common } from 'saacs'

import { useChaincode } from '~/server/utils/useChaincode'

const bodySchema = z.object({
  collectionId: z.string(),
  userName: z.string(),
  role: z.string(),
})

export default defineEventHandler(async (event) => {
  const cc = await useChaincode(event)

  const query = await readValidatedBody(event, bodySchema.parse)

  const user = await findUserByUsername(query.userName)

  if (user.userId === undefined || user.userId === '')
    return createError("User's UserID is not valid or doesn't have one")
  if (user.mspId === undefined || user.mspId === '')
    return createError("User's mspId is not valid or doesn't exist")

  const userRole = new pb.UserCollectionRoles({
    collectionId: query.collectionId,
    mspId: user.mspId,
    userId: user.userId,
    roleIds: ['role'],
  })

  // console.log(result);
  const result = await cc.service.create({
    item: saacs.PrimaryToItem(userRole),
  })

  console.log(result)
  return { result }
})
