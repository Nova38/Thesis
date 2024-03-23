import { z } from 'zod'
// import { ccbio, common } from 'saacs'

const querySchema = z.object({
  collectionId: z.string(),
  specimenId: z.string(),
})

export default defineEventHandler(async (event) => {
  const cc = await useChaincode(event)

  const r = await getValidatedQuery(event, (body) =>
    querySchema.safeParse(body),
  )
  if (!r.success) throw r.error.issues
  // console.log({ data: r.data });

  // const r = bodySchema.parse(b);
  // console.log(r);

  const result = await cc.service.get(
    new common.generic.GetRequest({
      key: {
        collectionId: r.data.collectionId,
        itemKeyParts: [r.data.specimenId],
        itemType: ccbio.Specimen.typeName,
      },
    }),
  )

  // console.log(result);
  const unpacked = new ccbio.Specimen()
  result.item?.value?.unpackTo(unpacked)

  // console.log({ unpacked });
  return unpacked
})
