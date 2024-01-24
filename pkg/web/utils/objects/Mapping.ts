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

export interface FieldMapping<N, O> {
  defaultValue?: (() => N[keyof N]) | N[keyof N]
  newKey: keyof N
  oldKey: keyof O
  transform?: (value: O[keyof O]) => N[keyof N]
}

export type ObjectMapping<N, O> = FieldMapping<N, O>[]
export type SpecimenMapping = ObjectMapping<FlatSpecimen, Record<string, string>>

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

export function TransformObject<O, N>(
  obj: O,
  mappings: ObjectMapping<N, O>,
) {
  return mappings.reduce((newObj, mapping) => {
    const { defaultValue, newKey, oldKey, transform } = mapping
    const value = obj[oldKey]

    if (value === undefined) {
      if (defaultValue !== undefined) {
        newObj[newKey]
          = typeof defaultValue === 'function'
            ? (defaultValue as () => N[keyof N])()
            : defaultValue
      }
    }
    else {
      newObj[newKey] = transform
        ? (transform(value))
        : (value as N[keyof N])
    }

    return newObj
  }, {} as N)
}

// export type SpecimenMapping = ObjectMapping<Record<string, string>, FlatSpecimen>

export function TransformRecordToFlatSpecimen(
  record: Record<string, string>,
  mappings: SpecimenMapping,
) {
  return TransformObject(record, mappings)
}

export function ClearMapping<N, O>(mapping: ObjectMapping<N, O>, key: keyof N) {
  return mapping.map((m) => {
    if (m.newKey === key)
      m.transform = undefined
    return m
  })
}
