import { z } from 'zod'
// import { ccbio, common } from 'saacs'

const querySchema = z.object({
  collectionId: z.string(),
  showHidden: z.string().optional(),
  specimenId: z.string(),
})

export default defineEventHandler(async (event) => {
  const cc = await useChaincode(event)

  console.log(getQuery(event))

  const r = await getValidatedQuery(event, querySchema.parse)
  // if (!r.success) {
  //   return createError({ message: 'Invalid query', statusCode: 400 })
  // }
  // if (!r.data) {
  //   return createError({ message: 'Invalid query', statusCode: 400 })
  // }
  // console.log({ data: r.data });

  // const r = bodySchema.parse(b);
  // console.log(r);
  try {
    const result = await cc.service.getHistory(
      new pb.GetHistoryRequest({
        key: {
          collectionId: r.collectionId,
          itemKeyParts: [r.specimenId],
          itemType: ccbio.Specimen.typeName,
        },
        historyOptions: {
          include: r.showHidden === 'true',
        },
      }),
    )

    // console.log(result.toJsonString({ typeRegistry: GlobalRegistry }));
    const history = new ccbio.SpecimenHistory({})

    if (!result.history) return { results: [] }

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
  } catch (error) {
    console.error(error)
    createError({ message: 'Failed to fetch history', statusCode: 500 })
  }
})
