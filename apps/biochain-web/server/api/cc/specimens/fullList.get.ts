import { z } from 'zod'
import type { PlainSpecimen } from '~/utils/types'
// import { auth, ccbio, common } from 'saacs'

const querySchema = z.object({
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

  let bookmark = ''
  let lastBookmark = '-'
  const SpecimenMap: Record<string, PlainSpecimen> = {}

  while (bookmark !== lastBookmark) {
    const result = await cc.service.listByAttrs(
      new pb.ListByAttrsRequest({
        pagination: {
          bookmark: bookmark ?? '',
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
    lastBookmark = bookmark
    bookmark = result.pagination?.bookmark ?? ''

    result.items.forEach((i) => {
      const s = new ccbio.Specimen()

      if (!i.value?.unpackTo(s)) {
        console.log('unpack failed', i)
        return
      }
      SpecimenMap[s.specimenId] = s
    })

    if (lastBookmark === bookmark) {
      console.log('no change', bookmark, lastBookmark)
      break
    }
  }
  // console.log("1");

  // console.log("2");

  // console.log(result);

  return new ccbio.SpecimenMap({
    specimens: SpecimenMap,
  })
})
