import { v5 } from 'uuid'

const DNS_UUID = '6ba7b810-9dad-11d1-80b4-00c04fd430c8'
export const SpecimenNamespace = v5(
  'ku_orn.specimen.biochain.iitc.ku.edu',
  DNS_UUID,
)

export function CatNumToUUID(specimenId: string) {
  return v5(specimenId, SpecimenNamespace)
}
