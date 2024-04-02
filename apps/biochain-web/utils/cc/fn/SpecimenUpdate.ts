import * as ccbio from '~/lib/pb/biochain/v1'
import type { UpdateRowMeta } from '~/utils/types'

export async function DoSpecimenUpdate(
  specimen: PlainSpecimen,
  curSpecimen: PlainSpecimen,
): Promise<UpdateRowMeta> {
  console.log({ curSpecimen, specimen })

  const { differences, mask } = diffCrush(
    toValue(specimen),
    toValue(curSpecimen),
    ['lastModified'],
  )

  const req = new ccbio.SpecimenUpdate({
    mask,
    specimen: new ccbio.Specimen(ZSpecimen.parse(specimen)),
  })

  try {
    const res = await useCustomFetch(`/api/cc/specimens/update`, {
      body: req.toJsonString({
        typeRegistry: GlobalRegistry,
      }),
      method: 'POST',
    })
    console.log(res)
    return {
      differences,
      exist: 'new',
      status: 'success',
      statusMessage: 'Uploaded Successfully',
      uuid: specimen.specimenId,
    }
  } catch (err) {
    switch (true) {
      case err instanceof Error:
        return {
          differences,
          exist: 'new',
          status: 'error',
          statusMessage: err.message,
          uuid: specimen.specimenId,
        }
      case typeof err === 'string':
        return {
          differences,
          exist: 'new',
          status: 'error',
          statusMessage: err,
          uuid: specimen.specimenId,
        }

      default:
        return {
          differences,
          exist: 'new',
          status: 'error',
          statusMessage: 'Unknown Error',
          uuid: specimen.specimenId,
        }
    }
  }
}
