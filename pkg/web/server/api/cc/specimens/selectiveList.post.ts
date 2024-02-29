import { z } from 'zod'

const querySchema = z.object({
  collectionId: z.string(),
  specimenIds: z.array(z.string()),
})

export default defineEventHandler(async (event) => {
  const query = await getValidatedQuery(event, body =>
    querySchema.safeParse(body))

  if (!query.success)
    throw query.error.issues

  const fullList = await $fetch('/api/cc/specimens/fullList')
  const specimenMap = fullList?.specimenMap

  if (!specimenMap)
    throw new Error('Failed to fetch full list')

  const filteredList = Object.entries(specimenMap).filter(
    ([key]) => key in query.data.specimenIds,
  )

  return filteredList
})
