import { z } from 'zod'

const querySchema = z.object({
  collectionId: z.string(),
})

const bodySchema = z.object({
  specimenIds: z.array(z.string()),
})

export default defineEventHandler(async (event) => {
  console.log(event)
  console.log('test')

  const query = await getValidatedQuery(event, (body) =>
    querySchema.safeParse(body),
  )

  if (!query.success) throw query.error.issues

  console.log(query)

  const body = await readValidatedBody(event, (body) => {
    return bodySchema.safeParse(body)
  })
  if (!body.success) throw body.error.issues

  console.log({ query, body })

  const fullList = await $fetch('/api/cc/specimens/fullList', {
    query: { collectionId: query.data.collectionId },
  })
  console.log(fullList)
  const specimenMap = fullList?.specimenMap

  const filteredList = Object.entries(specimenMap).filter(([id]) =>
    body.data.specimenIds.includes(id),
  )
  console.log(filteredList)
  return {
    filteredList,
  }
})
