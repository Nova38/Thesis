import { ccbio } from "saacs-es";
import {
  createRegistry,
  Timestamp,
  type PlainMessage,
} from "@bufbuild/protobuf";

export type status = "new" | "loading" | "success" | "error" | "pre-existing";

function ConvertSpecimen(row, meta: RowMeta, Mappings: { [x: string]: any }) {
  const specimen = new ccbio.Specimen();
}

export const FlatSpecimenKeys = [
  "collectionId",
  "specimenId",
  "primary.catalogNumber",
  "primary.accessionNumber",
  "primary.fieldNumber",
  "primary.tissueNumber",
  "primary.cataloger",
  "primary.collector",
  "primary.determiner",
  "primary.fieldDate.verbatim",
  "primary.fieldDate.year",
  "primary.fieldDate.month",
  "primary.fieldDate.day",
  "primary.catalogDate.verbatim",
  "primary.catalogDate.year",
  "primary.catalogDate.month",
  "primary.catalogDate.day",
  "primary.determinedDate.verbatim",
  "primary.determinedDate.year",
  "primary.determinedDate.month",
  "primary.determinedDate.day",
  "primary.determinedReason",
  "primary.originalDate.verbatim",
  "primary.originalDate.year",
  "primary.originalDate.month",
  "primary.originalDate.day",
  "secondary.sex",
  "secondary.age",
  "secondary.weight",
  "secondary.weightUnits",
  "secondary.condition",
  "secondary.molt",
  "secondary.notes",
  "taxon.kingdom",
  "taxon.phylum",
  "taxon.class",
  "taxon.order",
  "taxon.family",
  "taxon.genus",
  "taxon.species",
  "taxon.subspecies",
  "georeference.country",
  "georeference.stateProvince",
  "georeference.county",
  "georeference.locality",
  "georeference.latitude",
  "georeference.longitude",
  "georeference.habitat",
  "georeference.continent",
  "georeference.locationRemarks",
  "georeference.coordinateUncertaintyInMeters",
  "georeference.georeferenceBy",
  "georeference.georeferenceDate.verbatim",
  "georeference.georeferenceDate.year",
  "georeference.georeferenceDate.month",
  "georeference.georeferenceDate.day",
  "georeference.georeferenceProtocol",
  "georeference.geodeticDatum",
  "georeference.footprintWkt",
  "georeference.notes",
];

export const ImportableSpecimenKeys = [
  "primary.catalogNumber",
  "primary.accessionNumber",
  "primary.fieldNumber",
  "primary.tissueNumber",
  "primary.cataloger",
  "primary.collector",
  "primary.determiner",
  "primary.fieldDate.verbatim",
  "primary.fieldDate.year",
  "primary.fieldDate.month",
  "primary.fieldDate.day",
  "primary.catalogDate.verbatim",
  "primary.catalogDate.year",
  "primary.catalogDate.month",
  "primary.catalogDate.day",
  "primary.determinedDate.verbatim",
  "primary.determinedDate.year",
  "primary.determinedDate.month",
  "primary.determinedDate.day",
  "primary.determinedReason",
  "primary.originalDate.verbatim",
  "primary.originalDate.year",
  "primary.originalDate.month",
  "primary.originalDate.day",
  "secondary.sex",
  "secondary.age",
  "secondary.weight",
  "secondary.weightUnits",
  "secondary.condition",
  "secondary.molt",
  "secondary.notes",
  "taxon.kingdom",
  "taxon.phylum",
  "taxon.class",
  "taxon.order",
  "taxon.family",
  "taxon.genus",
  "taxon.species",
  "taxon.subspecies",
  "georeference.country",
  "georeference.stateProvince",
  "georeference.county",
  "georeference.locality",
  "georeference.latitude",
  "georeference.longitude",
  "georeference.habitat",
  "georeference.continent",
  "georeference.locationRemarks",
  "georeference.coordinateUncertaintyInMeters",
  "georeference.georeferenceBy",
  "georeference.georeferenceDate.verbatim",
  "georeference.georeferenceDate.year",
  "georeference.georeferenceDate.month",
  "georeference.georeferenceDate.day",
  "georeference.georeferenceProtocol",
  "georeference.geodeticDatum",
  "georeference.footprintWkt",
  "georeference.notes",
];

interface FlatSpecimen {
  "primary.catalogNumber": string;
  "primary.accessionNumber": string;
  "primary.fieldNumber": string;
  "primary.tissueNumber": string;
  "primary.cataloger": string;
  "primary.collector": string;
  "primary.determiner": string;
  "primary.fieldDate.verbatim": string;
  "primary.fieldDate.year": string;
  "primary.fieldDate.month": string;
  "primary.fieldDate.day": string;
  "primary.catalogDate.verbatim": string;
  "primary.catalogDate.year": string;
  "primary.catalogDate.month": string;
  "primary.catalogDate.day": string;
  "primary.determinedDate.verbatim": string;
  "primary.determinedDate.year": string;
  "primary.determinedDate.month": string;
  "primary.determinedDate.day": string;
  "primary.determinedReason": string;
  "primary.originalDate.verbatim": string;
  "primary.originalDate.year": string;
  "primary.originalDate.month": string;
  "primary.originalDate.day": string;
  "secondary.sex": string;
  "secondary.age": string;
  "secondary.weight": string;
  "secondary.weightUnits": string;
  "secondary.condition": string;
  "secondary.molt": string;
  "secondary.notes": string;
  "taxon.kingdom": string;
  "taxon.phylum": string;
  "taxon.class": string;
  "taxon.order": string;
  "taxon.family": string;
  "taxon.genus": string;
  "taxon.species": string;
  "taxon.subspecies": string;
  "georeference.country": string;
  "georeference.stateProvince": string;
  "georeference.county": string;
  "georeference.locality": string;
  "georeference.latitude": string;
  "georeference.longitude": string;
  "georeference.habitat": string;
  "georeference.continent": string;
  "georeference.locationRemarks": string;
  "georeference.coordinateUncertaintyInMeters": string;
  "georeference.georeferenceBy": string;
  "georeference.georeferenceDate.verbatim": string;
  "georeference.georeferenceDate.year": string;
  "georeference.georeferenceDate.month": string;
  "georeference.georeferenceDate.day": string;
  "georeference.georeferenceProtocol": string;
  "georeference.geodeticDatum": string;
  "georeference.footprintWkt": string;
  "georeference.notes": string;
}
