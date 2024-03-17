import type { FieldMask } from '@bufbuild/protobuf'

import { Any } from '@bufbuild/protobuf'
import { auth, ccbio, common } from 'saacs'

export interface bodySchema {
  mask: FieldMask
  specimen: ccbio.Specimen
}

export default defineEventHandler(async (event) => {
  const cc = await useChaincode(event)

  try {
    const body = await readBody(event)
    console.log()
    console.log({ body })

    const specimen = new ccbio.Specimen().fromJson(body.specimen)
    console.log({ specimen })
    const value = Any.pack(specimen)
    console.log({ value })

    const req = new common.generic.UpdateRequest({
      item: {
        key: new auth.objects.ItemKey({
          collectionId: specimen.collectionId,
          itemKeyParts: [specimen.specimenId],
          itemType: ccbio.Specimen.typeName,
        }),
        value,
      },
      updateMask: body.mask,
    })
    const result = await cc.service.update(req)

    console.log(result)
    const unpacked = new ccbio.Specimen()
    result.item?.value?.unpackTo(unpacked)

    return unpacked
    return req.toJson({ typeRegistry: cc.service.registry })
  }
  catch (error) {
    console.log(error)
    throw error
  }
})
