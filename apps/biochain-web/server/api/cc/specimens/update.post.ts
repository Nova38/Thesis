import { ccbio } from '#imports'
import type { FieldMask } from '@bufbuild/protobuf'
import { ZSpecimen } from '~/utils/cc/proto/Specimen'
import { z } from 'zod'
import { Any } from '@bufbuild/protobuf'
import { key } from '@saacs/saacs-pb'
import { saacs } from '@saacs/client'
// import { auth, ccbio, common } from 'saacs'

export interface bodySchema {
  mask: FieldMask
  specimen: ccbio.Specimen
}
const bodySchema = z.object({
  mask: z
    .object({
      paths: z.array(z.string()),
    })
    .optional(),
  specimen: ZSpecimen,
})

export default defineEventHandler(async (event) => {
  const cc = await useChaincode(event)

  try {
    const body = await readValidatedBody(event, bodySchema.parse)

    if (!body.mask) {
      const current = $fetch('/api/cc/specimens/get', {
        method: 'GET',
        query: {
          collectionId: body.specimen.collectionId,
          specimenId: body.specimen.specimenId,
        },
      })

      // todo: diff
    }

    const item = saacs.PrimaryToItem(new pb.Specimen(body.specimen))

    const result = await cc.service.update({
      item: item,
      updateMask: body.mask,
    })

    console.log(result)
    const unpacked = new ccbio.Specimen()
    result.item?.value?.unpackTo(unpacked)

    return unpacked
  } catch (error) {
    console.log(error)
    throw error
  }
})
