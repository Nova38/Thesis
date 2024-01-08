import { ccbio } from "saacs-es";
import z from "zod";
import { Timestamp, FieldMask } from "@bufbuild/protobuf";

import { objectify } from "radash";

export type ProtoDate = z.infer<typeof ProtoDate>;
export const ProtoDate = z.object({
  verbatim: z.string().trim().optional(),
  timestamp: z.coerce
    .date()
    .transform((d) => Timestamp.fromDate(d))
    .optional(),
  year: z.number().optional(),
  month: z.string().trim().optional(),
  day: z.number().optional(),
});

export type LastModified = z.infer<typeof LastModified>;
export const LastModified = z.object({
  txId: z.string().trim().optional(),
  mspId: z.string().trim().optional(),
  userId: z.string().trim().optional(),
  timestamp: z.coerce.date().transform((d) => Timestamp.fromDate(d)),
  note: z.string().trim().optional(),
});

const sex = z.enum([
  "SEX_UNDEFINED",
  "SEX_UNKNOWN",
  "SEX_ATYPICAL",
  "SEX_MALE",
  "SEX_FEMALE",
]);

const age = z.enum([
  "AGE_UNDEFINED",
  "AGE_UNKNOWN",
  "AGE_NEST",
  "AGE_EMBRYO_EGG",
  "AGE_CHICK_SUBADULT",
  "AGE_ADULT",
  "AGE_CONTINGENT",
]);

const empty_taxon = new ccbio.Specimen_Taxon({});

export const Specimen = z.object({
  collectionId: z.string().trim(),
  specimenId: z.string().uuid(),
  primary: z.object({
    catalogNumber: z.string().trim().optional(),
    accessionNumber: z.string().trim().optional(),
    fieldNumber: z.string().trim().optional(),
    tissueNumber: z.string().trim().optional(),
    cataloger: z.string().trim().optional(),
    collector: z.string().trim().optional(),
    determiner: z.string().trim().optional(),
    fieldDate: ProtoDate.optional(),
    catalogDate: ProtoDate.optional(),
    determinedDate: ProtoDate.optional(),
    determinedReason: z.string().trim().optional(),
    originalDate: ProtoDate.optional(),
    lastModified: LastModified.optional(),
  }),
  secondary: z.object({
    sex: z.nativeEnum(ccbio.Specimen_Secondary_SEX),
    age: z.nativeEnum(ccbio.Specimen_Secondary_AGE),
    weight: z.number().optional(),
    weightUnits: z.string().trim().optional(),
    condition: z.string().trim().optional(),
    molt: z.string().trim().optional(),
    notes: z.string().trim().optional(),
    preparations: z
      .record(
        z.object({
          verbatim: z.string().trim().optional(),
        }),
      )
      .default({}),
    lastModified: LastModified.optional(),
  }),
  taxon: z
    .object({
      kingdom: z.string().trim().default(""),
      phylum: z.string().trim().default(""),
      class: z.string().trim().default(""),
      order: z.string().trim().default(""),
      family: z.string().trim().default(""),
      genus: z.string().trim().default(""),
      species: z.string().trim().default(""),
      subspecies: z.string().trim().default(""),
      lastModified: LastModified.optional(),
    })
    .default({}),
  georeference: z.object({
    country: z.string().trim().optional(),
    stateProvince: z.string().trim().optional(),
    county: z.string().trim().optional(),
    locality: z.string().trim().optional(),
    latitude: z.number().optional(),
    longitude: z.number().optional(),
    habitat: z.string().trim().optional(),
    continent: z.string().trim().optional(),
    locationRemarks: z.string().trim().optional(),
    coordinateUncertaintyInMeters: z.number().optional(),
    georeferenceBy: z.string().trim().optional(),
    georeferenceDate: ProtoDate.optional(),
    georeferenceProtocol: z.string().trim().optional(),
    geodeticDatum: z.string().trim().optional(),
    footprintWkt: z.string().trim().optional(),
    notes: z.string().trim().optional(),
    lastModified: LastModified.optional(),
  }),
  images: z
    .record(
      z.object({
        id: z.string().trim().optional(),
        url: z.string().trim().optional(),
        notes: z.string().trim().optional(),
        hash: z.string().trim().optional(),
        lastModified: LastModified.optional(),
      }),
    )
    .default({}),
  loans: z
    .record(
      z.object({
        id: z.string().trim().optional(),
        description: z.string().trim().optional(),
        loanedBy: z.string().trim().optional(),
        loanedTo: z.string().trim().optional(),
        loanedDate: ProtoDate.optional(),
        lastModified: LastModified.optional(),
      }),
    )
    .default({}),
  grants: z
    .record(
      z.object({
        id: z.string().trim().optional(),
        description: z.string().trim().optional(),
        grantedBy: z.string().trim().optional(),
        grantedTo: z.string().trim().optional(),
        grantedDate: ProtoDate.optional(),
        lastModified: LastModified.optional(),
      }),
    )
    .default({}),
  lastModified: LastModified.optional(),
});

export const reqSpeciemn = Specimen.required();

export function zodToProto(raw: z.infer<typeof Specimen>) {
  // const parsed = ;

  const proto = new ccbio.Specimen();
}

export function parseFromFlat(raw: object) {
  // crush(raw);
  // const unFlat =
}
