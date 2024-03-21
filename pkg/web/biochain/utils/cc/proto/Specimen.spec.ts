import { writeFileSync } from 'node:fs'
import { createRegistry } from '@bufbuild/protobuf'
import { construct, crush } from 'radash'
import { beforeAll, describe, expect, it } from 'vitest'

import { ZSpecimen } from './Specimen'
import { ccbio } from 'saacs'

const raw = {
  collectionId: 'KU-Zoology',
  georeference: {
    continent: '',
    coordinateUncertaintyInMeters: 0,
    country: '',
    county: '',
    footprintWkt: '',
    geodeticDatum: '',
    georeferenceBy: '',
    georeferenceDate: {
      day: 0,
      month: '',
      verbatim: '',
      year: 0,
    },
    georeferenceProtocol: '',
    habitat: '',
    lastModified: {
      mspId: 'Org1MSP',
      note: '',
      timestamp: '1970-01-01T00:00:00Z',
      txId: '280ac7973a8a4aa00d50363a8651a713fef87683f9f194bcd414f6ff5ff5f68b   ',
      userId:
        'eDUwOTo6Q049b3JnMWFkbWluLE9VPWFkbWluLE89SHlwZXJsZWRnZXIsU1Q9Tm9ydGggQ2Fyb2xpbmEsQz1VUzo6Q049Y2Eub3JnMS5leGFtcGxlLmNvbSxPPW9yZzEuZXhhbXBsZS5jb20sTD1EdXJoYW0sU1Q9Tm9ydGggQ2Fyb2xpbmEsQz1VUw==',
    },
    latitude: 0,
    locality: '',
    locationRemarks: '',
    longitude: 3,
    notes: '',
    stateProvince: '',
  },
  grants: {},
  images: {},
  lastModified: {
    mspId: 'Org1MSP',
    note: '',
    timestamp: '2024-01-07T22:07:34.861Z',
    txId: 'cdd1f61d7e2a8a7d2519d732b40390a2ec50f4713e053d470984ddeb51907d45',
    userId:
      'eDUwOTo6Q049b3JnMWFkbWluLE9VPWFkbWluLE89SHlwZXJsZWRnZXIsU1Q9Tm9ydGggQ2Fyb2xpbmEsQz1VUzo6Q049Y2Eub3JnMS5leGFtcGxlLmNvbSxPPW9yZzEuZXhhbXBsZS5jb20sTD1EdXJoYW0sU1Q9Tm9ydGggQ2Fyb2xpbmEsQz1VUw==',
  },
  loans: {
    loney: {
      description: 'hi',
      id: 'loney',
      loanedBy: 'you',
      loanedTo: 'me',
    },
  },
  primary: {
    accessionNumber: '',
    catalogDate: {
      day: 0,
      month: '',
      verbatim: '',
      year: 0,
    },
    catalogNumber: '',
    cataloger: '',
    collector: '',
    determinedDate: {
      day: 0,
      month: '',
      verbatim: '',
      year: 0,
    },
    determinedReason: '',
    determiner: '',
    fieldDate: {
      day: 0,
      month: '',
      verbatim: '',
      year: 0,
    },
    fieldNumber: '',
    lastModified: {
      mspId: 'Org1MSP',
      note: '',
      timestamp: '1970-01-01T00:00:00Z',
      txId: '280ac7973a8a4aa00d50363a8651a713fef87683f9f194bcd414f6ff5ff5f68b',
      userId:
        'eDUwOTo6Q049b3JnMWFkbWluLE9VPWFkbWluLE89SHlwZXJsZWRnZXIsU1Q9Tm9ydGggQ2Fyb2xpbmEsQz1VUzo6Q049Y2Eub3JnMS5leGFtcGxlLmNvbSxPPW9yZzEuZXhhbXBsZS5jb20sTD1EdXJoYW0sU1Q9Tm9ydGggQ2Fyb2xpbmEsQz1VUw==',
    },
    originalDate: {
      day: 0,
      month: '',
      verbatim: '',
      year: 0,
    },
    tissueNumber: '',
  },
  secondary: {
    age: 0,
    condition: '',
    lastModified: {
      mspId: 'Org1MSP',
      note: '',
      timestamp: '1970-01-01T00:00:00Z',
      txId: '280ac7973a8a4aa00d50363a8651a713fef87683f9f194bcd414f6ff5ff5f68b',
      userId:
        'eDUwOTo6Q049b3JnMWFkbWluLE9VPWFkbWluLE89SHlwZXJsZWRnZXIsU1Q9Tm9ydGggQ2Fyb2xpbmEsQz1VUzo6Q049Y2Eub3JnMS5leGFtcGxlLmNvbSxPPW9yZzEuZXhhbXBsZS5jb20sTD1EdXJoYW0sU1Q9Tm9ydGggQ2Fyb2xpbmEsQz1VUw==',
    },
    molt: '',
    notes: '',
    preparations: {},
    sex: 0,
    weight: 0,
    weightUnits: '',
  },
  specimenId: '32b21bdb-b493-411e-af78-6cba731e3780',
  taxon: {
    class: '',
    family: 'a',
    genus: 'a',
    kingdom: '',
    lastModified: {
      mspId: 'Org1MSP',
      note: '',
      timestamp: '1970-01-01T00:00:00Z',
      txId: '280ac7973a8a4aa00d50363a8651a713fef87683f9f194bcd414f6ff5ff5f68b',
      userId:
        'eDUwOTo6Q049b3JnMWFkbWluLE9VPWFkbWluLE89SHlwZXJsZWRnZXIsU1Q9Tm9ydGggQ2Fyb2xpbmEsQz1VUzo6Q049Y2Eub3JnMS5leGFtcGxlLmNvbSxPPW9yZzEuZXhhbXBsZS5jb20sTD1EdXJoYW0sU1Q9Tm9ydGggQ2Fyb2xpbmEsQz1VUw==',
    },
    order: '',
    phylum: '',
    species: 'asdf',
    subspecies: '',
  },
} as const

const csvFlat = {
  'georeference.coordinateUncercaintyInMeters': 'NA',
  'georeference.country': 'USA',
  'georeference.county': 'Loveland',
  'georeference.geodeticDatum': '',
  'georeference.georeferenceProtocol': '',
  'georeference.georeferencedBy': ' ',
  'georeference.georeferencedDate': '',
  'georeference.latitude': 'NA',
  'georeference.locality': 'N/A',
  'georeference.longitude': 'NA',
  'georeference.notes': '',
  'georeference.stateProvince': 'Colorado',
  index: '1',
  'primary.accessionNumber': 'NA',
  'primary.catalogDate': 'NA',
  'primary.catalogNumber': '2',
  'primary.collector': 'Osburn, W.',
  'primary.fieldNumber': '123',
  'primary.originalDate': '18 February 1889',
  'primary.tissueNumber': '',
  'secondary.age': '',
  'secondary.condition': '',
  'secondary.molt': '',
  'secondary.notes': '',
  'secondary.preparation': 'Skin - 1',
  'secondary.sex': 'M',
  'secondary.weight': '',
  'secondary.weightUnits': '',
  'taxon.class': 'Aves',
  'taxon.family': 'Fringillidae',
  'taxon.genus': 'Coccothraustes',
  'taxon.kingdom': 'Animalia',
  'taxon.order': 'Passeriformes',
  'taxon.phylum': 'Chordata',
  'taxon.species': 'vespertinus',
  'taxon.subspecies': 'brooksi',
}

describe('suite name', () => {
  beforeAll(() => {})
  it('simpleParse', () => {
    expect(() => ZSpecimen.parse(raw)).not.toThrowError()
    expect(() => {
      new ccbio.Specimen(ZSpecimen.parse(raw))
    }).not.toThrowError()

    // console.log(new ccbio.Specimen(Specimen.parse(raw)));
  })
  it('parseFlat', () => {
    const specimen = ccbio.Specimen.fromJson(raw)
    const flat = crush(raw)
    console.log(flat)
    const unFlat = construct(flat)
    console.log(unFlat)

    expect(ZSpecimen.parse(unFlat)).toEqual(ZSpecimen.parse(raw))
  })

  it('parseOutput', () => {
    const full_path =
      'Z:/source/repos/Thesis/pkg/biochain/import/ku_orn_database_great_plains_pre_1970_NoDups.json'
    const out_path =
      'Z:/source/repos/Thesis/pkg/biochain/import/ku_orn_cov.json'
    // read json from file
    const json = require(full_path)
    const output: ccbio.Specimen[] = []
    json.items.forEach((item: any) => {
      expect(() => {
        const s = new ccbio.Specimen(ZSpecimen.parse(construct(item)))
        output.push(s)
        s.toJsonString({ typeRegistry: createRegistry(ccbio.Specimen) })
        // new ccbio.Specimen(Specimen.parse(construct(item)));
      }).not.toThrowError()

      // console.log(item);
      // const specimen = new ccbio.Specimen();
    })

    writeFileSync(out_path, JSON.stringify(output))
  })
})
