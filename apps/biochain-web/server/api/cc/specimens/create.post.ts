import { Any } from '@bufbuild/protobuf'
// import { auth, ccbio, common } from 'saacs'
import { ZSpecimen } from '~/utils/cc/proto/Specimen'

const bodySchema = ZSpecimen

// z.object({
//   collectionId: z.string(),
//   specimenId: z.string(),
// })

export default defineEventHandler(async (event) => {
  console.log(event)
  const cc = await useChaincode(event)

  const body = await readBody(event)
  console.log({ body })

  const r = await readValidatedBody(event, (body) => bodySchema.safeParse(body))
  if (!r.success) {
    console.error(r.error.issues)
    throw r.error.issues
  }
  console.log({ data: r.data })

  // const r = bodySchema.parse(b);
  // console.log(r);

  const specimen = new ccbio.Specimen(body)
  console.log(specimen.toJsonString())

  const v = Any.pack(specimen)
  const req = new common.generic.CreateRequest({
    item: {
      key: new auth.objects.ItemKey({
        collectionId: specimen.collectionId,
        itemKeyParts: [specimen.specimenId],
        itemType: ccbio.Specimen.typeName,
      }),
      value: v,
    },
  })

  try {
    const result = await cc.service.create(req)
    console.debug(result)
    console.log(result)
    const unpacked = new ccbio.Specimen()
    result.item?.value?.unpackTo(unpacked)

    // console.log({ unpacked });

    return {
      unpacked,
    }
  } catch (error: any) {
    console.debug(error)

    console.log(error)
    console.log(error?.details[0].message)
    throw error
  }
})
