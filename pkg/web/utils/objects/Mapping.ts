import { ccbio } from 'saacs-es'
import { object } from 'zod'

import type { FlatSpecimen } from './flatten'

export type status = 'error' | 'loading' | 'new' | 'pre-existing' | 'success'

// function ConvertSpecimen(row, meta: RowMeta, Mappings: { [x: string]: any }) {
//   const specimen = new ccbio.Specimen()
// }

// export type ObjectMapping<OldType, NewType> = FieldMapping<OldType, NewType>[]

// export function TransformObject<OldType, NewType>(
//   obj: OldType,
//   mappings: ObjectMapping<OldType, NewType>,
// ): NewType {
//   return mappings.reduce((newObj, mapping) => {
//     const value = obj[mapping.oldKey]
//     if (value === undefined) {
//       if (mapping.defaultValue) {
//         newObj[mapping.newKey] = typeof mapping.defaultValue === 'function'
//           ? (mapping.defaultValue as () => NewType[keyof NewType])()
//           : mapping.defaultValue as NewType[keyof NewType]
//       }
//       return newObj
//     }
//     newObj[mapping.newKey] = mapping.transform
//       ? mapping.transform(value)
//       : value as NewType[keyof NewType]
//     return newObj
//   }, {} as NewType)
// }

// export interface FieldMapping<OldType, NewType> {
//   defaultValue?: (() => NewType[keyof NewType]) | NewType[keyof NewType]
//   newKey: keyof NewType
//   oldKey: keyof OldType
//   transform?: (value: OldType[keyof OldType]) => NewType[keyof NewType]
// }

export interface FieldMapping<OldType, NewType> {
  defaultValue?: (() => NewType[keyof NewType]) | NewType[keyof NewType]
  newKey: keyof NewType
  oldKey: keyof OldType
  transform?: (value: OldType[keyof OldType]) => NewType[keyof NewType]
}

export type ObjectMapping<OldType, NewType> = FieldMapping<OldType, NewType>[]

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

export function TransformObject<OldType, NewType>(
  obj: OldType,
  mappings: ObjectMapping<OldType, NewType>,
): NewType {
  return mappings.reduce((newObj, mapping) => {
    const { defaultValue, newKey, oldKey, transform } = mapping
    const value = obj[oldKey]

    if (value === undefined) {
      if (defaultValue !== undefined) {
        newObj[newKey]
          = typeof defaultValue === 'function'
            ? (defaultValue as () => NewType[keyof NewType])()
            : defaultValue
      }
    }
    else {
      newObj[newKey] = transform
        ? (transform(value))
        : (value as NewType[keyof NewType])
    }

    return newObj
  }, {} as NewType)
}

export type RecordMapping = ObjectMapping<FlatSpecimen, Record<string, string>>
export type SpecimenMapping = ObjectMapping<Record<string, string>, FlatSpecimen>

export function TransformRecordToFlatSpecimen(
  record: Record<string, string>,
  mappings: SpecimenMapping,
): FlatSpecimen {
  return TransformObject(record, mappings)
}
