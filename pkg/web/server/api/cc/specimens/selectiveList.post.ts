import { z } from 'zod'

const querySchema = z.object({
  collectionId: z.string(),
  specimenIds: z.array(z.string()),
})

export default defineEventHandler(async (event) => {
  return 'Hello Nitro'
})
