// FILEPATH: /z:/source/repos/Thesis/pkg/web/utils/flatten.spec.ts
import { beforeEach, describe, expect, it } from 'vitest'

import type { FieldMapping, ObjectMapping, SpecimenMapping } from './Mapping'

import { TransformObject, TransformRecordToFlatSpecimen } from './Mapping'

// import { type Mapping, transformObject } from "./"

describe('transformObject', () => {
  it('should transform object keys based on mappings', () => {
    const obj = {
      age: 30,
      name: 'John Doe',
      occupation: 'Engineer',
    }
    type OldType = typeof obj
    interface NewType {
      fullName: string
      job: string
      years: number
    }

    const mappings: ObjectMapping<NewType, OldType> = [
      { newKey: 'fullName', oldKey: 'name' },
      { newKey: 'years', oldKey: 'age' },
      { newKey: 'job', oldKey: 'occupation' },
    ]

    // ^?

    const result = TransformObject(obj, mappings)
    expect(result).toEqual({
      fullName: 'John Doe',
      job: 'Engineer',
      years: 30,
    })
  })

  it('should apply transformation function if provided', () => {
    const obj = {
      age: '30',
      name: 'John Doe',
      occupation: 'Engineer',
    }

    const mappings: FieldMapping<
      { fullName: string, job: string, years: number },
      typeof obj
    >[] = [
      { newKey: 'fullName', oldKey: 'name' },
      {
        newKey: 'years',
        oldKey: 'age',
        transform: (value: string) => Number.parseInt(value),
      },
      { newKey: 'job', oldKey: 'occupation' },
    ]

    const result = TransformObject(obj, mappings)
    expect(result).toEqual({
      fullName: 'John Doe',
      job: 'Engineer',
      years: 30,
    })
  })

  it('should apply general transformation function if provided', () => {
    const obj = {
      age: '30',
      name: 'John Doe',
      occupation: 'Engineer',
    }

    const mappings: ObjectMapping<
      { fullName: string, job: string, years: number },
      typeof obj
    > = [
      {
        newKey: 'fullName',
        oldKey: 'name',
        transform(value) {
          return value.toUpperCase()
        },
      },
      { newKey: 'years', oldKey: 'age' },
      {
        newKey: 'job',
        oldKey: 'occupation',
        transform(value) {
          return value.toUpperCase()
        },
      },
    ]

    const result = TransformObject(obj, mappings)
    expect(result).toEqual({
      fullName: 'JOHN DOE',
      job: 'ENGINEER',
      years: '30',
    })
  })
  describe('transformObject', () => {
    beforeEach(() => {})
    it('should transform object keys based on mappings', () => {
      const obj: Record<string, string> = {
        'accessionNumber': '456',
        'catalogNumber': '123',
        'fieldNumber': '789',
        'georeference.continent': '',
        'georeference.coordinateUncertaintyInMeters': '',
        'georeference.country': '',
        'georeference.county': '',
        'georeference.footprintWkt': '',
        'georeference.geodeticDatum': '',
        'georeference.georeferenceBy': '',
        'georeference.georeferenceDate.day': '',
        'georeference.georeferenceDate.month': '',
        'georeference.georeferenceDate.verbatim': '',
        'georeference.georeferenceDate.year': '',
        'georeference.georeferenceProtocol': '',
        'georeference.habitat': '',
        'georeference.latitude': '',
        'georeference.locality': '',
        'georeference.locationRemarks': '',
        'georeference.longitude': '',
        'georeference.notes': '',
        'georeference.stateProvince': '',
        'primary.catalogDate.day': '',
        'primary.catalogDate.month': '',
        'primary.catalogDate.verbatim': '',
        'primary.catalogDate.year': '',
        'primary.cataloger': '',
        'primary.collector': '',
        'primary.determinedDate.day': '',
        'primary.determinedDate.month': '',
        'primary.determinedDate.verbatim': '',
        'primary.determinedDate.year': '',
        'primary.determinedReason': '',
        'primary.determiner': '',
        'primary.fieldDate.day': '',
        'primary.fieldDate.month': '',
        'primary.fieldDate.verbatim': '',
        'primary.fieldDate.year': '',
        'primary.originalDate.day': '',
        'primary.originalDate.month': '',
        'primary.originalDate.verbatim': '',
        'primary.originalDate.year': '',
        'primary.tissueNumber': '',
        'secondary.age': '',
        'secondary.condition': '',
        'secondary.molt': '',
        'secondary.notes': '',
        'secondary.sex': '',
        'secondary.weight': '',
        'secondary.weightUnits': '',
        'taxon.class': '',
        'taxon.family': '',
        'taxon.genus': '',
        'taxon.kingdom': '',
        'taxon.order': '',
        'taxon.phylum': '',
        'taxon.species': '',
        'taxon.subspecies': '',
      }

      const mappings: SpecimenMapping = [
        { newKey: 'primary.catalogNumber', oldKey: 'catalogNumber' },
        { newKey: 'primary.accessionNumber', oldKey: 'accessionNumber' },
        // Swap oldKey and newKey for all other items in the array
        { newKey: 'primary.tissueNumber', oldKey: 'tissueNumber' },
        {
          newKey: 'primary.cataloger',
          oldKey: 'cataloger',
          transform: (value: string) => value.toUpperCase(),
        },
        {
          defaultValue: 'John Doe',
          newKey: 'primary.collector',
          oldKey: 'collector',
        },
        {
          defaultValue() {
            return 'Jane Doe'
          },
          newKey: 'primary.determiner',
          oldKey: 'determiner',
        },
        { newKey: 'primary.fieldNumber', oldKey: 'fieldNumber' },
      ]

      TransformRecordToFlatSpecimen(obj, mappings)
      // ^?

      const result = TransformRecordToFlatSpecimen(obj, mappings)
      expect(result).toEqual({
        'primary.accessionNumber': '456',
        'primary.catalogNumber': '123',
        'primary.collector': 'John Doe',
        'primary.determiner': 'Jane Doe',
        'primary.fieldNumber': '789',
        // ... rest of the transformed properties
      })
    })
  })
  it('not', () => {
    const obj: Record<string, string> = {
      'accessionNumber': '456',
      'catalogNumber': '123',
      'fieldNumber': '789',
      'georeference.continent': '',
      'georeference.coordinateUncertaintyInMeters': '',
      'georeference.country': '',
      'georeference.county': '',
      'georeference.footprintWkt': '',
      'georeference.geodeticDatum': '',
      'georeference.georeferenceBy': '',
      'georeference.georeferenceDate.day': '',
      'georeference.georeferenceDate.month': '',
      'georeference.georeferenceDate.verbatim': '',
      'georeference.georeferenceDate.year': '',
      'georeference.georeferenceProtocol': '',
      'georeference.habitat': '',
      'georeference.latitude': '',
      'georeference.locality': '',
      'georeference.locationRemarks': '',
      'georeference.longitude': '',
      'georeference.notes': '',
      'georeference.stateProvince': '',
      'primary.catalogDate.day': '',
      'primary.catalogDate.month': '',
      'primary.catalogDate.verbatim': '',
      'primary.catalogDate.year': '',
      'primary.cataloger': '',
      'primary.collector': '',
      'primary.determinedDate.day': '',
      'primary.determinedDate.month': '',
      'primary.determinedDate.verbatim': '',
      'primary.determinedDate.year': '',
      'primary.determinedReason': '',
      'primary.determiner': '',
      'primary.fieldDate.day': '',
      'primary.fieldDate.month': '',
      'primary.fieldDate.verbatim': '',
      'primary.fieldDate.year': '',
      'primary.originalDate.day': '',
      'primary.originalDate.month': '',
      'primary.originalDate.verbatim': '',
      'primary.originalDate.year': '',
      'primary.tissueNumber': '',
      'secondary.age': '',
      'secondary.condition': '',
      'secondary.molt': '',
      'secondary.notes': '',
      'secondary.sex': '',
      'secondary.weight': '',
      'secondary.weightUnits': '',
      'taxon.class': '',
      'taxon.family': '',
      'taxon.genus': '',
      'taxon.kingdom': '',
      'taxon.order': '',
      'taxon.phylum': '',
      'taxon.species': '',
      'taxon.subspecies': '',
    }

    const mappings: SpecimenMapping = [
      { newKey: 'primary.catalogNumber', oldKey: 'catalogNumber' },
      { newKey: 'primary.accessionNumber', oldKey: 'accessionNumber' },
      // Swap oldKey and newKey for all other items in the array
      { newKey: 'primary.tissueNumber', oldKey: 'tissueNumber' },
      {
        newKey: 'primary.cataloger',
        oldKey: 'cataloger',
        transform: (value: string) => value.toUpperCase(),
      },
      {
        defaultValue: 'John Doe',
        newKey: 'primary.collector',
        oldKey: 'collector',
      },
      {
        defaultValue() {
          return 'Jane Doe'
        },
        newKey: 'primary.determiner',
        oldKey: 'determiner',
      },
      { newKey: 'primary.fieldNumber', oldKey: 'fieldNumber' },
    ]

    TransformRecordToFlatSpecimen(obj, mappings)
    // ^?

    const result = TransformRecordToFlatSpecimen(obj, mappings)
    expect(result).toEqual({
      'primary.accessionNumber': '456',
      'primary.catalogNumber': '123',
      'primary.collector': 'John Doe',
      'primary.determiner': 'Jane Doe',
      'primary.fieldNumber': '789',
      // ... rest of the transformed properties
    })
  })
})
