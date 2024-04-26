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
  try {
    const result = await cc.service.suggestionByPartialKey(
      new pb.SuggestionByPartialKeyRequest({
        pagination: {
          bookmark: '',
          pageSize: 1000,
        },
        itemKey: {
          collectionId: r.data.collectionId,
          itemKeyParts: [r.data.specimenId],
          itemType: ccbio.Specimen.typeName,
        },
        numAttrs: 2,
      }),
    )
    console.log(result)
    return result
  } catch (error) {
    console.error(error)
    throw new Error('Failed to fetch suggestions')
  }

  // console.log(result.toJsonString({ typeRegistry: GlobalRegistry }));
  const history = new ccbio.SpecimenHistory({})

  for (const item of result.history.entries) {
    const s = new ccbio.Specimen()
    item.value?.unpackTo(s)

    history.entries.push(
      new ccbio.SpecimenHistoryEntry({
        ...item,
        value: s,
      }),
    )
  }

  // console.log({ history });
  return history
})
