import { ccbio } from '#imports'
import type { FieldMask } from '@bufbuild/protobuf'

import { Any } from '@bufbuild/protobuf'
// import { auth, ccbio, common } from 'saacs'

export interface bodySchema {
  mask: FieldMask
  specimen: ccbio.Specimen
}

export default defineEventHandler(async (event) => {
  const cc = await useChaincode(event)

  try {
    const body = await readBody(event)

    const specimen = new ccbio.Specimen().fromJson(body.specimen)
    const value = Any.pack(specimen)

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
  } catch (error) {
    console.log(error)
    throw error
  }
})
