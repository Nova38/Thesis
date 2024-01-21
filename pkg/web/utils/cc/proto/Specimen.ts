import { type PlainMessage, Timestamp } from '@bufbuild/protobuf'
import { ccbio } from 'saacs-es'
import z from 'zod'

export type PlainSpecimen = PlainMessage<ccbio.Specimen>

export function MakeEmptySpecimen() {
  return new ccbio.Specimen({
    collectionId: '',
    georeference: {
      georeferenceDate: {},
    },
    grants: {},
    loans: {},
    primary: {
      catalogDate: {},
      determinedDate: {},
      fieldDate: {},
      originalDate: {},
    },
    secondary: {
      preparations: {},
    },
    taxon: {},
  })
}

// export type ProtoDate = z.infer<typeof ProtoDate>
export const ProtoDate = z.object({
  day: z.coerce.number().optional(),
  month: z.string().trim().optional(),
  timestamp: z.coerce
    .date()
    .transform(d => Timestamp.fromDate(d))
    .optional(),
  verbatim: z.string().trim().optional(),
  year: z.coerce.number().optional(),
})

export type LastModified = z.infer<typeof LastModified>
export const LastModified = z.object({
  mspId: z.string().trim().optional(),
  note: z.string().trim().optional(),
  timestamp: z.coerce.date().transform(d => Timestamp.fromDate(d)),
  txId: z.string().trim().optional(),
  userId: z.string().trim().optional(),
})

const sex = z.enum([
  'SEX_UNDEFINED',
  'SEX_UNKNOWN',
  'SEX_ATYPICAL',
  'SEX_MALE',
  'SEX_FEMALE',
])

const age = z.enum([
  'AGE_UNDEFINED',
  'AGE_UNKNOWN',
  'AGE_NEST',
  'AGE_EMBRYO_EGG',
  'AGE_CHICK_SUBADULT',
  'AGE_ADULT',
  'AGE_CONTINGENT',
])

export const Specimen = z.object({
  collectionId: z.string().trim(),
  georeference: z.object({
    continent: z.string().trim().optional(),
    coordinateUncertaintyInMeters: z.number().optional(),
    country: z.string().trim().optional(),
    county: z.string().trim().optional(),
    footprintWkt: z.string().trim().optional(),
    geodeticDatum: z.string().trim().optional(),
    georeferenceBy: z.string().trim().optional(),
    georeferenceDate: ProtoDate.optional(),
    georeferenceProtocol: z.string().trim().optional(),
    habitat: z.string().trim().optional(),
    lastModified: LastModified.optional(),
    latitude: z.coerce.number().optional(),
    locality: z.coerce.string().trim().optional(),
    locationRemarks: z.string().trim().optional(),
    longitude: z.coerce.number().optional(),
    notes: z.string().trim().optional(),
    stateProvince: z.string().trim().optional(),
  }),
  grants: z
    .record(
      z.object({
        description: z.string().trim().optional(),
        grantedBy: z.string().trim().optional(),
        grantedDate: ProtoDate.optional(),
        grantedTo: z.string().trim().optional(),
        id: z.string().trim().optional(),
        lastModified: LastModified.optional(),
      }),
    )
    .default({}),
  images: z
    .record(
      z.object({
        hash: z.string().trim().optional(),
        id: z.string().trim().optional(),
        lastModified: LastModified.optional(),
        notes: z.string().trim().optional(),
        url: z.string().trim().optional(),
      }),
    )
    .default({}),
  lastModified: LastModified.optional(),
  loans: z
    .record(
      z.object({
        description: z.string().trim().optional(),
        id: z.string().trim().optional(),
        lastModified: LastModified.optional(),
        loanedBy: z.string().trim().optional(),
        loanedDate: ProtoDate.optional(),
        loanedTo: z.string().trim().optional(),
      }),
    )
    .default({}),
  primary: z.object({
    accessionNumber: z.string().trim().optional(),
    catalogDate: ProtoDate.optional(),
    catalogNumber: z.string().trim().optional(),
    cataloger: z.string().trim().optional(),
    collector: z.string().trim().optional(),
    determinedDate: ProtoDate.optional(),
    determinedReason: z.string().trim().optional(),
    determiner: z.string().trim().optional(),
    fieldDate: ProtoDate.optional(),
    fieldNumber: z.string().trim().optional(),
    lastModified: LastModified.optional(),
    originalDate: ProtoDate.optional(),
    tissueNumber: z.string().trim().optional(),
  }),
  secondary: z.object({
    age: z.nativeEnum(ccbio.Specimen_Secondary_AGE).default(0),
    condition: z.string().trim().optional(),
    lastModified: LastModified.optional(),
    molt: z.string().trim().optional(),
    notes: z.string().trim().optional(),
    preparations: z
      .record(
        z.object({
          verbatim: z.string().trim().optional(),
        }),
      )
      .default({}),
    sex: z.nativeEnum(ccbio.Specimen_Secondary_SEX).default(0),
    weight: z.coerce.number().optional(),
    weightUnits: z.string().trim().optional(),
  }),
  specimenId: z.string().uuid(),
  taxon: z
    .object({
      class: z.string().trim().default(''),
      family: z.string().trim().default(''),
      genus: z.string().trim().default(''),
      kingdom: z.string().trim().default(''),
      lastModified: LastModified.optional(),
      order: z.string().trim().default(''),
      phylum: z.string().trim().default(''),
      species: z.string().trim().default(''),
      subspecies: z.string().trim().default(''),
    })
    .default({}),
})

export const reqSpeciemn = Specimen.required()

export function zodToProto(raw: z.infer<typeof Specimen>) {
  // const parsed = ;

  const proto = new ccbio.Specimen()
}

export function parseFromFlat(raw: object) {
  // crush(raw);
  // const unFlat =
}
