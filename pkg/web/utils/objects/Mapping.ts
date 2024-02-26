import { reactify } from '@vueuse/core'
import type { FlatSpecimen } from './flatten'

export interface FieldMapping<
  N,
  O,
  NewKey extends keyof N = keyof N,
  OldKey extends keyof O = keyof O,
> {
  defaultValue?: (() => N[NewKey]) | N[NewKey]
  newKey: NewKey
  oldKey: OldKey
  transform?: (value: O[OldKey]) => N[NewKey]
}

export type ObjectMapping<N, O> = FieldMapping<N, O>[]
export type SpecimenMapping = ObjectMapping<
  FlatSpecimen,
  Record<string, number | string>
>

// export function TransformObject1<OldType, NewType>(
//   obj: OldType,
//   mappings: ObjectMapping<OldType, NewType>,
// ): NewType {
//   return mappings.reduce((newObj, mapping) => {
//     const { defaultValue, newKey, oldKey, transform } = mapping
//     const value = obj[oldKey]

//     newObj[newKey]
//       = value === undefined
//         ? defaultValue?.call ? defaultValue() : defaultValue
//         : transform?.(value) ?? (value as NewType[keyof NewType])

//     return newObj
//   }, {} as NewType)
// }

export function TransformObject<O, N>(obj: O, mappings: ObjectMapping<N, O>) {
  return mappings.reduce((newObj, mapping) => {
    const { defaultValue, newKey, oldKey, transform } = mapping
    if (oldKey === '') {
      if (defaultValue !== undefined) {
        newObj[newKey] =
          typeof defaultValue === 'function'
            ? (defaultValue as () => N[keyof N])()
            : defaultValue
      }
      return newObj
    }

    const value = obj[oldKey]

    if (value === undefined) {
      if (defaultValue !== undefined) {
        newObj[newKey] =
          typeof defaultValue === 'function'
            ? (defaultValue as () => N[keyof N])()
            : defaultValue
      }
    } else {
      newObj[newKey] = transform ? transform(value) : (value as N[keyof N])
    }

    return newObj
  }, {} as N)
}

export const t = reactify(TransformObject)

// export type SpecimenMapping = ObjectMapping<Record<string, string>, FlatSpecimen>

export function TransformRecordToFlatSpecimen(
  record: Record<string, string>,
  mappings: SpecimenMapping,
) {
  return TransformObject(record, mappings)
}

export function ClearMapping<N, O>(mapping: ObjectMapping<N, O>, key: keyof N) {
  return mapping.map((m) => {
    if (m.newKey === key) m.transform = undefined
    return m
  })
}

export function EmptySpecimenMapping(): SpecimenMapping {
  const mapping: SpecimenMapping = [
    {
      newKey: 'georeference.continent',
      oldKey: '',
    },
    { newKey: 'georeference.coordinateUncertaintyInMeters', oldKey: '' },
    {
      newKey: 'georeference.country',
      oldKey: '',
    },
    {
      newKey: 'georeference.county',
      oldKey: '',
    },
    {
      newKey: 'georeference.footprintWkt',
      oldKey: '',
    },
    {
      newKey: 'georeference.geodeticDatum',
      oldKey: '',
    },
    {
      newKey: 'georeference.georeferenceBy',
      oldKey: '',
    },
    {
      newKey: 'georeference.georeferenceDate.day',
      oldKey: '',
    },
    {
      newKey: 'georeference.georeferenceDate.month',
      oldKey: '',
    },
    {
      newKey: 'georeference.georeferenceDate.verbatim',
      oldKey: '',
    },
    {
      newKey: 'georeference.georeferenceDate.year',
      oldKey: '',
    },
    {
      newKey: 'georeference.georeferenceProtocol',
      oldKey: '',
    },
    {
      newKey: 'georeference.habitat',
      oldKey: '',
    },
    {
      newKey: 'georeference.latitude',
      oldKey: '',
    },
    {
      newKey: 'georeference.locality',
      oldKey: '',
    },
    {
      newKey: 'georeference.locationRemarks',
      oldKey: '',
    },
    {
      newKey: 'georeference.longitude',
      oldKey: '',
    },
    {
      newKey: 'georeference.notes',
      oldKey: '',
    },
    {
      newKey: 'georeference.stateProvince',
      oldKey: '',
    },
    {
      newKey: 'primary.accessionNumber',
      oldKey: '',
    },
    {
      newKey: 'primary.catalogDate.day',
      oldKey: '',
    },
    {
      newKey: 'primary.catalogDate.month',
      oldKey: '',
    },
    {
      newKey: 'primary.catalogDate.verbatim',
      oldKey: '',
    },
    {
      newKey: 'primary.catalogDate.year',
      oldKey: '',
    },
    {
      newKey: 'primary.catalogNumber',
      oldKey: '',
    },
    {
      newKey: 'primary.cataloger',
      oldKey: '',
    },
    {
      newKey: 'primary.collector',
      oldKey: '',
    },
    {
      newKey: 'primary.determinedDate.day',
      oldKey: '',
    },
    {
      newKey: 'primary.determinedDate.month',
      oldKey: '',
    },
    {
      newKey: 'primary.determinedDate.verbatim',
      oldKey: '',
    },
    {
      newKey: 'primary.determinedDate.year',
      oldKey: '',
    },
    {
      newKey: 'primary.determinedReason',
      oldKey: '',
    },
    {
      newKey: 'primary.determiner',
      oldKey: '',
    },
    {
      newKey: 'primary.fieldDate.day',
      oldKey: '',
    },
    {
      newKey: 'primary.fieldDate.month',
      oldKey: '',
    },
    {
      newKey: 'primary.fieldDate.verbatim',
      oldKey: '',
    },
    {
      newKey: 'primary.fieldDate.year',
      oldKey: '',
    },
    {
      newKey: 'primary.fieldNumber',
      oldKey: '',
    },
    {
      newKey: 'primary.originalDate.day',
      oldKey: '',
    },
    {
      newKey: 'primary.originalDate.month',
      oldKey: '',
    },
    {
      newKey: 'primary.originalDate.verbatim',
      oldKey: '',
    },
    {
      newKey: 'primary.originalDate.year',
      oldKey: '',
    },
    {
      newKey: 'primary.tissueNumber',
      oldKey: '',
    },
    {
      newKey: 'secondary.age',
      oldKey: '',
    },
    {
      newKey: 'secondary.condition',
      oldKey: '',
    },
    {
      newKey: 'secondary.molt',
      oldKey: '',
    },
    {
      newKey: 'secondary.notes',
      oldKey: '',
    },
    {
      newKey: 'secondary.preparations.imported.verbatim',
      oldKey: '',
    },
    {
      newKey: 'secondary.sex',
      oldKey: '',
    },
    {
      // defaultValue: 0,
      newKey: 'secondary.weight',
      oldKey: '',
    },
    {
      newKey: 'secondary.weightUnits',
      oldKey: '',
    },
    {
      newKey: 'taxon.class',
      oldKey: '',
    },
    {
      newKey: 'taxon.family',
      oldKey: '',
    },
    {
      newKey: 'taxon.genus',
      oldKey: '',
    },
    {
      newKey: 'taxon.kingdom',
      oldKey: '',
    },
    {
      newKey: 'taxon.order',
      oldKey: '',
    },
    {
      newKey: 'taxon.phylum',
      oldKey: '',
    },
    {
      newKey: 'taxon.species',
      oldKey: '',
    },
    {
      newKey: 'taxon.subspecies',
      oldKey: '',
    },
  ]

  return mapping
}
