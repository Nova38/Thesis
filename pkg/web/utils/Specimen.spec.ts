import { writeFileSync } from 'node:fs'
import { ccbio } from 'saacs-es'
import { beforeAll, describe, expect, it } from 'vitest'
import { createRegistry } from '@bufbuild/protobuf'

import { construct, crush } from 'radash'
import { Specimen } from './zSpecimen'

const raw = {
  collectionId: 'KU-Zoology',
  specimenId: '32b21bdb-b493-411e-af78-6cba731e3780',
  primary: {
    catalogNumber: '',
    accessionNumber: '',
    fieldNumber: '',
    tissueNumber: '',
    cataloger: '',
    collector: '',
    determiner: '',
    fieldDate: {
      verbatim: '',
      year: 0,
      month: '',
      day: 0,
    },
    catalogDate: {
      verbatim: '',
      year: 0,
      month: '',
      day: 0,
    },
    determinedDate: {
      verbatim: '',
      year: 0,
      month: '',
      day: 0,
    },
    determinedReason: '',
    originalDate: {
      verbatim: '',
      year: 0,
      month: '',
      day: 0,
    },
    lastModified: {
      txId: '280ac7973a8a4aa00d50363a8651a713fef87683f9f194bcd414f6ff5ff5f68b',
      mspId: 'Org1MSP',
      userId:
        'eDUwOTo6Q049b3JnMWFkbWluLE9VPWFkbWluLE89SHlwZXJsZWRnZXIsU1Q9Tm9ydGggQ2Fyb2xpbmEsQz1VUzo6Q049Y2Eub3JnMS5leGFtcGxlLmNvbSxPPW9yZzEuZXhhbXBsZS5jb20sTD1EdXJoYW0sU1Q9Tm9ydGggQ2Fyb2xpbmEsQz1VUw==',
      timestamp: '1970-01-01T00:00:00Z',
      note: '',
    },
  },
  secondary: {
    sex: 0,
    age: 0,
    weight: 0,
    weightUnits: '',
    preparations: {},
    condition: '',
    molt: '',
    notes: '',
    lastModified: {
      txId: '280ac7973a8a4aa00d50363a8651a713fef87683f9f194bcd414f6ff5ff5f68b',
      mspId: 'Org1MSP',
      userId:
        'eDUwOTo6Q049b3JnMWFkbWluLE9VPWFkbWluLE89SHlwZXJsZWRnZXIsU1Q9Tm9ydGggQ2Fyb2xpbmEsQz1VUzo6Q049Y2Eub3JnMS5leGFtcGxlLmNvbSxPPW9yZzEuZXhhbXBsZS5jb20sTD1EdXJoYW0sU1Q9Tm9ydGggQ2Fyb2xpbmEsQz1VUw==',
      timestamp: '1970-01-01T00:00:00Z',
      note: '',
    },
  },
  taxon: {
    kingdom: '',
    phylum: '',
    class: '',
    order: '',
    family: 'a',
    genus: 'a',
    species: 'asdf',
    subspecies: '',
    lastModified: {
      txId: '280ac7973a8a4aa00d50363a8651a713fef87683f9f194bcd414f6ff5ff5f68b',
      mspId: 'Org1MSP',
      userId:
        'eDUwOTo6Q049b3JnMWFkbWluLE9VPWFkbWluLE89SHlwZXJsZWRnZXIsU1Q9Tm9ydGggQ2Fyb2xpbmEsQz1VUzo6Q049Y2Eub3JnMS5leGFtcGxlLmNvbSxPPW9yZzEuZXhhbXBsZS5jb20sTD1EdXJoYW0sU1Q9Tm9ydGggQ2Fyb2xpbmEsQz1VUw==',
      timestamp: '1970-01-01T00:00:00Z',
      note: '',
    },
  },
  georeference: {
    country: '',
    stateProvince: '',
    county: '',
    locality: '',
    latitude: 0,
    longitude: 3,
    habitat: '',
    continent: '',
    locationRemarks: '',
    coordinateUncertaintyInMeters: 0,
    georeferenceBy: '',
    georeferenceDate: {
      verbatim: '',
      year: 0,
      month: '',
      day: 0,
    },
    georeferenceProtocol: '',
    geodeticDatum: '',
    footprintWkt: '',
    notes: '',
    lastModified: {
      txId: '280ac7973a8a4aa00d50363a8651a713fef87683f9f194bcd414f6ff5ff5f68b   ',
      mspId: 'Org1MSP',
      userId:
        'eDUwOTo6Q049b3JnMWFkbWluLE9VPWFkbWluLE89SHlwZXJsZWRnZXIsU1Q9Tm9ydGggQ2Fyb2xpbmEsQz1VUzo6Q049Y2Eub3JnMS5leGFtcGxlLmNvbSxPPW9yZzEuZXhhbXBsZS5jb20sTD1EdXJoYW0sU1Q9Tm9ydGggQ2Fyb2xpbmEsQz1VUw==',
      timestamp: '1970-01-01T00:00:00Z',
      note: '',
    },
  },
  images: {},
  loans: {
    loney: {
      id: 'loney',
      description: 'hi',
      loanedBy: 'you',
      loanedTo: 'me',
    },
  },
  grants: {},
  lastModified: {
    txId: 'cdd1f61d7e2a8a7d2519d732b40390a2ec50f4713e053d470984ddeb51907d45',
    mspId: 'Org1MSP',
    userId:
      'eDUwOTo6Q049b3JnMWFkbWluLE9VPWFkbWluLE89SHlwZXJsZWRnZXIsU1Q9Tm9ydGggQ2Fyb2xpbmEsQz1VUzo6Q049Y2Eub3JnMS5leGFtcGxlLmNvbSxPPW9yZzEuZXhhbXBsZS5jb20sTD1EdXJoYW0sU1Q9Tm9ydGggQ2Fyb2xpbmEsQz1VUw==',
    timestamp: '2024-01-07T22:07:34.861Z',
    note: '',
  },
}

const csvFlat = {
  'index': '1',
  'primary.catalogNumber': '2',
  'primary.accessionNumber': 'NA',
  'primary.fieldNumber': '123',
  'primary.tissueNumber': '',
  'primary.collector': 'Osburn, W.',
  'primary.originalDate': '18 February 1889',
  'primary.catalogDate': 'NA',
  'secondary.preparation': 'Skin - 1',
  'secondary.condition': '',
  'secondary.notes': '',
  'secondary.sex': 'M',
  'secondary.age': '',
  'secondary.molt': '',
  'secondary.weight': '',
  'secondary.weightUnits': '',
  'taxon.kingdom': 'Animalia',
  'taxon.phylum': 'Chordata',
  'taxon.class': 'Aves',
  'taxon.order': 'Passeriformes',
  'taxon.family': 'Fringillidae',
  'taxon.genus': 'Coccothraustes',
  'taxon.species': 'vespertinus',
  'taxon.subspecies': 'brooksi',
  'georeference.country': 'USA',
  'georeference.stateProvince': 'Colorado',
  'georeference.county': 'Loveland',
  'georeference.locality': 'N/A',
  'georeference.latitude': 'NA',
  'georeference.longitude': 'NA',
  'georeference.coordinateUncercaintyInMeters': 'NA',
  'georeference.georeferencedBy': ' ',
  'georeference.georeferencedDate': '',
  'georeference.georeferenceProtocol': '',
  'georeference.geodeticDatum': '',
  'georeference.notes': '',
}

describe('suite name', () => {
  beforeAll(() => {})
  it('simpleParse', () => {
    expect(() => Specimen.parse(raw)).not.toThrowError()
    expect(() => {
      new ccbio.Specimen(Specimen.parse(raw))
    }).not.toThrowError()

    // console.log(new ccbio.Specimen(Specimen.parse(raw)));
  })
  it('parseFlat', () => {
    const specimen = new ccbio.Specimen(raw)
    const flat = crush(raw)
    console.log(flat)
    const unFlat = construct(flat)
    console.log(unFlat)

    expect(Specimen.parse(unFlat)).toEqual(Specimen.parse(raw))
  })

  it('parseOutput', () => {
    const full_path
      = 'Z:/source/repos/Thesis/pkg/biochain/import/ku_orn_database_great_plains_pre_1970_NoDups.json'
    const out_path
      = 'Z:/source/repos/Thesis/pkg/biochain/import/ku_orn_cov.json'
    // read json from file
    const json = require(full_path)
    const output = []
    json.items.forEach((item: any) => {
      expect(() => {
        const s = new ccbio.Specimen(Specimen.parse(construct(item)))
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
