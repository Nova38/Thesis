import { z } from 'zod'
// import { auth, ccbio, common } from 'saacs'

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
    new pb.ListByAttrsRequest({
      pagination: {
        bookmark: query.data.bookmark ?? '',
        pageSize: query.data.limit ?? 1000,
      },
      key: new pb.ItemKey({
        collectionId: query.data.collectionId,
        itemKeyParts: [query.data.collectionId],
        itemType: ccbio.Specimen.typeName,
      }),
      numAttrs: 0,
    }),
  )
  // console.log("2");

  // console.log(result);
  // const specimenMap: Record<string, any> = {}
  const response = new ccbio.SpecimenMap({
    bookmark: result.pagination?.bookmark ?? '',
    specimens: Object.fromEntries(
      result.items.map((i) => {
        const s = new ccbio.Specimen()
        i.value?.unpackTo(s)
        return [s.specimenId, s]
      }),
    ),
  })

  // result.items.forEach((i) => {
  //   const s = new ccbio.Specimen()
  //   i.value?.unpackTo(s)
  //   specimenMap[s.specimenId] = s.toJson({ emitDefaultValues: true })
  // })

  return response
})
