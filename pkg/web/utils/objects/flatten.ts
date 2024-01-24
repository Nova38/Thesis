// https://decipher.dev/30-seconds-of-typescript/docs/flattenObject/
export function FlattenObject(obj, prefix = '') {
  return Object.keys(obj).reduce((acc, k) => {
    const pre = prefix.length ? `${prefix}.` : ''
    if (typeof obj[k] === 'object')
      Object.assign(acc, FlattenObject(obj[k], pre + k))
    else acc[pre + k] = obj[k]
    return acc
  }, {})
}

export function FlattenEmptySpecimen() {
  return FlattenObject(
    JSON.parse(MakeEmptySpecimen().toJsonString({ emitDefaultValues: true })),
  )
}

export function FlattedSpecimenKeys() {
  return Object.keys(FlattenEmptySpecimen())
}

// export const FlatSpecimenKeys = [
//   'collectionId',
//   'specimenId',
//   'primary.catalogNumber',
//   'primary.accessionNumber',
//   'primary.fieldNumber',
//   'primary.tissueNumber',
//   'primary.cataloger',
//   'primary.collector',
//   'primary.determiner',
//   'primary.fieldDate.verbatim',
//   'primary.fieldDate.year',
//   'primary.fieldDate.month',
//   'primary.fieldDate.day',
//   'primary.catalogDate.verbatim',
//   'primary.catalogDate.year',
//   'primary.catalogDate.month',
//   'primary.catalogDate.day',
//   'primary.determinedDate.verbatim',
//   'primary.determinedDate.year',
//   'primary.determinedDate.month',
//   'primary.determinedDate.day',
//   'primary.determinedReason',
//   'primary.originalDate.verbatim',
//   'primary.originalDate.year',
//   'primary.originalDate.month',
//   'primary.originalDate.day',
//   'secondary.sex',
//   'secondary.age',
//   'secondary.weight',
//   'secondary.weightUnits',
//   'secondary.condition',
//   'secondary.molt',
//   'secondary.notes',
//   'taxon.kingdom',
//   'taxon.phylum',
//   'taxon.class',
//   'taxon.order',
//   'taxon.family',
//   'taxon.genus',
//   'taxon.species',
//   'taxon.subspecies',
//   'georeference.country',
//   'georeference.stateProvince',
//   'georeference.county',
//   'georeference.locality',
//   'georeference.latitude',
//   'georeference.longitude',
//   'georeference.habitat',
//   'georeference.continent',
//   'georeference.locationRemarks',
//   'georeference.coordinateUncertaintyInMeters',
//   'georeference.georeferenceBy',
//   'georeference.georeferenceDate.verbatim',
//   'georeference.georeferenceDate.year',
//   'georeference.georeferenceDate.month',
//   'georeference.georeferenceDate.day',
//   'georeference.georeferenceProtocol',
//   'georeference.geodeticDatum',
//   'georeference.footprintWkt',
//   'georeference.notes',
// ] as const

// export const ImportableSpecimenKeys = [
//   'primary.catalogNumber',
//   'primary.accessionNumber',
//   'primary.fieldNumber',
//   'primary.tissueNumber',
//   'primary.cataloger',
//   'primary.collector',
//   'primary.determiner',
//   'primary.fieldDate.verbatim',
//   'primary.fieldDate.year',
//   'primary.fieldDate.month',
//   'primary.fieldDate.day',
//   'primary.catalogDate.verbatim',
//   'primary.catalogDate.year',
//   'primary.catalogDate.month',
//   'primary.catalogDate.day',
//   'primary.determinedDate.verbatim',
//   'primary.determinedDate.year',
//   'primary.determinedDate.month',
//   'primary.determinedDate.day',
//   'primary.determinedReason',
//   'primary.originalDate.verbatim',
//   'primary.originalDate.year',
//   'primary.originalDate.month',
//   'primary.originalDate.day',
//   'secondary.sex',
//   'secondary.age',
//   'secondary.weight',
//   'secondary.weightUnits',
//   'secondary.condition',
//   'secondary.molt',
//   'secondary.notes',
//   'taxon.kingdom',
//   'taxon.phylum',
//   'taxon.class',
//   'taxon.order',
//   'taxon.family',
//   'taxon.genus',
//   'taxon.species',
//   'taxon.subspecies',
//   'georeference.country',
//   'georeference.stateProvince',
//   'georeference.county',
//   'georeference.locality',
//   'georeference.latitude',
//   'georeference.longitude',
//   'georeference.habitat',
//   'georeference.continent',
//   'georeference.locationRemarks',
//   'georeference.coordinateUncertaintyInMeters',
//   'georeference.georeferenceBy',
//   'georeference.georeferenceDate.verbatim',
//   'georeference.georeferenceDate.year',
//   'georeference.georeferenceDate.month',
//   'georeference.georeferenceDate.day',
//   'georeference.georeferenceProtocol',
//   'georeference.geodeticDatum',
//   'georeference.footprintWkt',
//   'georeference.notes',
// ]

export interface FlatSpecimen {
  'georeference.continent': string
  'georeference.coordinateUncertaintyInMeters': string
  'georeference.country': string
  'georeference.county': string
  'georeference.footprintWkt': string
  'georeference.geodeticDatum': string
  'georeference.georeferenceBy': string
  'georeference.georeferenceDate.day': string
  'georeference.georeferenceDate.month': string
  'georeference.georeferenceDate.verbatim': string
  'georeference.georeferenceDate.year': string
  'georeference.georeferenceProtocol': string
  'georeference.habitat': string
  'georeference.latitude': string
  'georeference.locality': string
  'georeference.locationRemarks': string
  'georeference.longitude': string
  'georeference.notes': string
  'georeference.stateProvince': string
  'primary.accessionNumber': string
  'primary.catalogDate.day': string
  'primary.catalogDate.month': string
  'primary.catalogDate.verbatim': string
  'primary.catalogDate.year': string
  'primary.catalogNumber': string
  'primary.cataloger': string
  'primary.collector': string
  'primary.determinedDate.day': string
  'primary.determinedDate.month': string
  'primary.determinedDate.verbatim': string
  'primary.determinedDate.year': string
  'primary.determinedReason': string
  'primary.determiner': string
  'primary.fieldDate.day': string
  'primary.fieldDate.month': string
  'primary.fieldDate.verbatim': string
  'primary.fieldDate.year': string
  'primary.fieldNumber': string
  'primary.originalDate.day': string
  'primary.originalDate.month': string
  'primary.originalDate.verbatim': string
  'primary.originalDate.year': string
  'primary.tissueNumber': string
  'secondary.age': string
  'secondary.condition': string
  'secondary.molt': string
  'secondary.notes': string
  'secondary.preparations.imported.verbatim': string
  'secondary.sex': string
  'secondary.weight': string
  'secondary.weightUnits': string
  'taxon.class': string
  'taxon.family': string
  'taxon.genus': string
  'taxon.kingdom': string
  'taxon.order': string
  'taxon.phylum': string
  'taxon.species': string
  'taxon.subspecies': string
}

export type FlatSpecimenKeys = keyof FlatSpecimen
export function isFlatSpecimenKey(key: string): key is FlatSpecimenKeys {
  return key in FlattenEmptySpecimen()
}
