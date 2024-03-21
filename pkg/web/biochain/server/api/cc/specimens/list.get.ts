import { z } from 'zod'
import { auth, ccbio, common } from 'saacs'

const querySchema = z.object({
  bookmark: z.string().optional(),
  collectionId: z.string(),
  limit: z.number().optional(),
})

export default defineEventHandler(async (event) => {
  const cc = await useChaincode(event)

  const query = await getValidatedQuery(event, (body) =>
    querySchema.safeParse(body),
  )
  console.log(query)

  if (!query.success) throw query.error.issues

  // console.log("1");

  const result = await cc.service.listByAttrs(
    new common.generic.ListByAttrsRequest({
      bookmark: query.data.bookmark ?? '',
      key: new auth.objects.ItemKey({
        collectionId: query.data.collectionId,
        itemKeyParts: [query.data.collectionId],
        itemType: ccbio.Specimen.typeName,
      }),
      limit: query.data.limit ?? 1000,
      numAttrs: 0,
    }),
  )
  // console.log("2");

  // console.log(result);
  const specimenMap: Record<string, any> = {}

  result.items.forEach((i) => {
    const s = new ccbio.Specimen()
    i.value?.unpackTo(s)
    specimenMap[s.specimenId] = s.toJson({ emitDefaultValues: true })
  })

  return { bookmark: result.bookmark, specimenMap }
})
