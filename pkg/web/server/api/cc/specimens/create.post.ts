import { Any } from '@bufbuild/protobuf'
import { ccbio, common } from 'saacs-es'
import { z } from 'zod'

const bodySchema = z.object({
  collectionId: z.string(),
  specimenId: z.string(),
})

export default defineEventHandler(async (event) => {
  const cc = await useChaincode(event)

  const body = await readBody(event)
  console.log({ body })

  const r = await readValidatedBody(event, body =>
    bodySchema.safeParse(body))
  if (!r.success)
    throw r.error.issues
  console.log({ data: r.data })

  // const r = bodySchema.parse(b);
  // console.log(r);

  const specimen = new ccbio.Specimen(body)
  console.log(specimen.toJsonString())

  const v = Any.pack(specimen)
  const req = new common.generic.CreateRequest({
    item: {
      value: v,
    },
  })
  console.log(req.toJsonString({ typeRegistry: GlobalRegistry }))

  try {
    const result = await cc.service.create(req)

    console.log(result)
    const unpacked = new ccbio.Specimen()
    result.item?.value?.unpackTo(unpacked)

    // console.log({ unpacked });

    return {
      unpacked,
    }
  }
  catch (error: any) {
    console.log(error)
    console.log(error?.details[0].message)
    throw error
  }
})
