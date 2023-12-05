// @generated by protoc-gen-es v1.3.1 with parameter "target=ts"
// @generated from file biochain/v1/state.proto (package ccbio.schema.v0, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import type { BinaryReadOptions, FieldList, JsonReadOptions, JsonValue, PartialMessage, PlainMessage } from "@bufbuild/protobuf";
import { Message, proto3, Timestamp } from "@bufbuild/protobuf";
import { StateActivity } from "../../auth/v1/auth_pb.js";

/**
 * @generated from message ccbio.schema.v0.Specimen
 */
export class Specimen extends Message<Specimen> {
  /**
   * @generated from field: string collection_id = 1;
   */
  collectionId = "";

  /**
   * @generated from field: string specimen_id = 2;
   */
  specimenId = "";

  /**
   * @generated from field: ccbio.schema.v0.Specimen.Primary primary = 3;
   */
  primary?: Specimen_Primary;

  /**
   * @generated from field: ccbio.schema.v0.Specimen.Secondary secondary = 4;
   */
  secondary?: Specimen_Secondary;

  /**
   * @generated from field: ccbio.schema.v0.Specimen.Taxon taxon = 5;
   */
  taxon?: Specimen_Taxon;

  /**
   * @generated from field: ccbio.schema.v0.Specimen.Georeference georeference = 6;
   */
  georeference?: Specimen_Georeference;

  /**
   * @generated from field: map<string, ccbio.schema.v0.Specimen.Image> images = 7;
   */
  images: { [key: string]: Specimen_Image } = {};

  /**
   * @generated from field: map<string, ccbio.schema.v0.Specimen.Loan> loans = 10;
   */
  loans: { [key: string]: Specimen_Loan } = {};

  /**
   * @generated from field: map<string, ccbio.schema.v0.Specimen.Grant> grants = 11;
   */
  grants: { [key: string]: Specimen_Grant } = {};

  constructor(data?: PartialMessage<Specimen>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "ccbio.schema.v0.Specimen";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "collection_id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "specimen_id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 3, name: "primary", kind: "message", T: Specimen_Primary },
    { no: 4, name: "secondary", kind: "message", T: Specimen_Secondary },
    { no: 5, name: "taxon", kind: "message", T: Specimen_Taxon },
    { no: 6, name: "georeference", kind: "message", T: Specimen_Georeference },
    { no: 7, name: "images", kind: "map", K: 9 /* ScalarType.STRING */, V: {kind: "message", T: Specimen_Image} },
    { no: 10, name: "loans", kind: "map", K: 9 /* ScalarType.STRING */, V: {kind: "message", T: Specimen_Loan} },
    { no: 11, name: "grants", kind: "map", K: 9 /* ScalarType.STRING */, V: {kind: "message", T: Specimen_Grant} },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): Specimen {
    return new Specimen().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): Specimen {
    return new Specimen().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): Specimen {
    return new Specimen().fromJsonString(jsonString, options);
  }

  static equals(a: Specimen | PlainMessage<Specimen> | undefined, b: Specimen | PlainMessage<Specimen> | undefined): boolean {
    return proto3.util.equals(Specimen, a, b);
  }
}

/**
 * @generated from message ccbio.schema.v0.Specimen.Primary
 */
export class Specimen_Primary extends Message<Specimen_Primary> {
  /**
   * @generated from field: string catalog_number = 1;
   */
  catalogNumber = "";

  /**
   * @generated from field: string accession_number = 2;
   */
  accessionNumber = "";

  /**
   * @generated from field: string field_number = 3;
   */
  fieldNumber = "";

  /**
   * @generated from field: string tissue_number = 4;
   */
  tissueNumber = "";

  /**
   * @generated from field: string cataloger = 5;
   */
  cataloger = "";

  /**
   * @generated from field: string collector = 6;
   */
  collector = "";

  /**
   * @generated from field: string determiner = 7;
   */
  determiner = "";

  /**
   * @generated from field: google.protobuf.Timestamp field_date = 8;
   */
  fieldDate?: Timestamp;

  /**
   * @generated from field: google.protobuf.Timestamp catalog_date = 9;
   */
  catalogDate?: Timestamp;

  /**
   * @generated from field: google.protobuf.Timestamp determined_date = 10;
   */
  determinedDate?: Timestamp;

  /**
   * @generated from field: string determined_reason = 11;
   */
  determinedReason = "";

  /**
   * @generated from field: auth.StateActivity last_modified = 20;
   */
  lastModified?: StateActivity;

  constructor(data?: PartialMessage<Specimen_Primary>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "ccbio.schema.v0.Specimen.Primary";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "catalog_number", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "accession_number", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 3, name: "field_number", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 4, name: "tissue_number", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 5, name: "cataloger", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 6, name: "collector", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 7, name: "determiner", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 8, name: "field_date", kind: "message", T: Timestamp },
    { no: 9, name: "catalog_date", kind: "message", T: Timestamp },
    { no: 10, name: "determined_date", kind: "message", T: Timestamp },
    { no: 11, name: "determined_reason", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 20, name: "last_modified", kind: "message", T: StateActivity },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): Specimen_Primary {
    return new Specimen_Primary().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): Specimen_Primary {
    return new Specimen_Primary().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): Specimen_Primary {
    return new Specimen_Primary().fromJsonString(jsonString, options);
  }

  static equals(a: Specimen_Primary | PlainMessage<Specimen_Primary> | undefined, b: Specimen_Primary | PlainMessage<Specimen_Primary> | undefined): boolean {
    return proto3.util.equals(Specimen_Primary, a, b);
  }
}

/**
 * @generated from message ccbio.schema.v0.Specimen.Secondary
 */
export class Specimen_Secondary extends Message<Specimen_Secondary> {
  /**
   * @generated from field: string preparation = 3;
   */
  preparation = "";

  /**
   * @generated from field: string condition = 4;
   */
  condition = "";

  /**
   * @generated from field: string notes = 5;
   */
  notes = "";

  /**
   * @generated from field: auth.StateActivity last_modified = 20;
   */
  lastModified?: StateActivity;

  constructor(data?: PartialMessage<Specimen_Secondary>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "ccbio.schema.v0.Specimen.Secondary";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 3, name: "preparation", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 4, name: "condition", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 5, name: "notes", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 20, name: "last_modified", kind: "message", T: StateActivity },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): Specimen_Secondary {
    return new Specimen_Secondary().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): Specimen_Secondary {
    return new Specimen_Secondary().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): Specimen_Secondary {
    return new Specimen_Secondary().fromJsonString(jsonString, options);
  }

  static equals(a: Specimen_Secondary | PlainMessage<Specimen_Secondary> | undefined, b: Specimen_Secondary | PlainMessage<Specimen_Secondary> | undefined): boolean {
    return proto3.util.equals(Specimen_Secondary, a, b);
  }
}

/**
 * @generated from message ccbio.schema.v0.Specimen.Taxon
 */
export class Specimen_Taxon extends Message<Specimen_Taxon> {
  /**
   * @generated from field: string kingdom = 1;
   */
  kingdom = "";

  /**
   * @generated from field: string phylum = 2;
   */
  phylum = "";

  /**
   * @generated from field: string class = 3;
   */
  class = "";

  /**
   * @generated from field: string order = 4;
   */
  order = "";

  /**
   * @generated from field: string family = 5;
   */
  family = "";

  /**
   * @generated from field: string genus = 6;
   */
  genus = "";

  /**
   * @generated from field: string species = 7;
   */
  species = "";

  /**
   * @generated from field: string subspecies = 8;
   */
  subspecies = "";

  /**
   * @generated from field: auth.StateActivity last_modified = 20;
   */
  lastModified?: StateActivity;

  constructor(data?: PartialMessage<Specimen_Taxon>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "ccbio.schema.v0.Specimen.Taxon";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "kingdom", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "phylum", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 3, name: "class", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 4, name: "order", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 5, name: "family", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 6, name: "genus", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 7, name: "species", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 8, name: "subspecies", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 20, name: "last_modified", kind: "message", T: StateActivity },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): Specimen_Taxon {
    return new Specimen_Taxon().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): Specimen_Taxon {
    return new Specimen_Taxon().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): Specimen_Taxon {
    return new Specimen_Taxon().fromJsonString(jsonString, options);
  }

  static equals(a: Specimen_Taxon | PlainMessage<Specimen_Taxon> | undefined, b: Specimen_Taxon | PlainMessage<Specimen_Taxon> | undefined): boolean {
    return proto3.util.equals(Specimen_Taxon, a, b);
  }
}

/**
 * @generated from message ccbio.schema.v0.Specimen.Georeference
 */
export class Specimen_Georeference extends Message<Specimen_Georeference> {
  /**
   * @generated from field: string country = 1;
   */
  country = "";

  /**
   * @generated from field: string state_province = 2;
   */
  stateProvince = "";

  /**
   * @generated from field: string county = 3;
   */
  county = "";

  /**
   * @generated from field: string locality = 4;
   */
  locality = "";

  /**
   * @generated from field: string latitude = 5;
   */
  latitude = "";

  /**
   * @generated from field: string longitude = 6;
   */
  longitude = "";

  /**
   * @generated from field: string habitat = 7;
   */
  habitat = "";

  /**
   * @generated from field: repeated string notes = 8;
   */
  notes: string[] = [];

  /**
   * @generated from field: auth.StateActivity last_modified = 20;
   */
  lastModified?: StateActivity;

  constructor(data?: PartialMessage<Specimen_Georeference>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "ccbio.schema.v0.Specimen.Georeference";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "country", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "state_province", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 3, name: "county", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 4, name: "locality", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 5, name: "latitude", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 6, name: "longitude", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 7, name: "habitat", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 8, name: "notes", kind: "scalar", T: 9 /* ScalarType.STRING */, repeated: true },
    { no: 20, name: "last_modified", kind: "message", T: StateActivity },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): Specimen_Georeference {
    return new Specimen_Georeference().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): Specimen_Georeference {
    return new Specimen_Georeference().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): Specimen_Georeference {
    return new Specimen_Georeference().fromJsonString(jsonString, options);
  }

  static equals(a: Specimen_Georeference | PlainMessage<Specimen_Georeference> | undefined, b: Specimen_Georeference | PlainMessage<Specimen_Georeference> | undefined): boolean {
    return proto3.util.equals(Specimen_Georeference, a, b);
  }
}

/**
 * Mapped Types
 *
 * @generated from message ccbio.schema.v0.Specimen.Image
 */
export class Specimen_Image extends Message<Specimen_Image> {
  /**
   * @generated from field: string id = 1;
   */
  id = "";

  /**
   * @generated from field: string url = 2;
   */
  url = "";

  /**
   * @generated from field: string notes = 3;
   */
  notes = "";

  /**
   * @generated from field: string hash = 4;
   */
  hash = "";

  /**
   * @generated from field: auth.StateActivity last_modified = 20;
   */
  lastModified?: StateActivity;

  constructor(data?: PartialMessage<Specimen_Image>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "ccbio.schema.v0.Specimen.Image";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "url", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 3, name: "notes", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 4, name: "hash", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 20, name: "last_modified", kind: "message", T: StateActivity },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): Specimen_Image {
    return new Specimen_Image().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): Specimen_Image {
    return new Specimen_Image().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): Specimen_Image {
    return new Specimen_Image().fromJsonString(jsonString, options);
  }

  static equals(a: Specimen_Image | PlainMessage<Specimen_Image> | undefined, b: Specimen_Image | PlainMessage<Specimen_Image> | undefined): boolean {
    return proto3.util.equals(Specimen_Image, a, b);
  }
}

/**
 * @generated from message ccbio.schema.v0.Specimen.Loan
 */
export class Specimen_Loan extends Message<Specimen_Loan> {
  /**
   * @generated from field: string id = 1;
   */
  id = "";

  /**
   * @generated from field: string description = 2;
   */
  description = "";

  /**
   * @generated from field: string loaned_by = 3;
   */
  loanedBy = "";

  /**
   * @generated from field: string loaned_to = 4;
   */
  loanedTo = "";

  /**
   * @generated from field: google.protobuf.Timestamp loaned_date = 5;
   */
  loanedDate?: Timestamp;

  /**
   * @generated from field: auth.StateActivity last_modified = 20;
   */
  lastModified?: StateActivity;

  constructor(data?: PartialMessage<Specimen_Loan>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "ccbio.schema.v0.Specimen.Loan";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "description", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 3, name: "loaned_by", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 4, name: "loaned_to", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 5, name: "loaned_date", kind: "message", T: Timestamp },
    { no: 20, name: "last_modified", kind: "message", T: StateActivity },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): Specimen_Loan {
    return new Specimen_Loan().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): Specimen_Loan {
    return new Specimen_Loan().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): Specimen_Loan {
    return new Specimen_Loan().fromJsonString(jsonString, options);
  }

  static equals(a: Specimen_Loan | PlainMessage<Specimen_Loan> | undefined, b: Specimen_Loan | PlainMessage<Specimen_Loan> | undefined): boolean {
    return proto3.util.equals(Specimen_Loan, a, b);
  }
}

/**
 * @generated from message ccbio.schema.v0.Specimen.Grant
 */
export class Specimen_Grant extends Message<Specimen_Grant> {
  /**
   * @generated from field: string id = 1;
   */
  id = "";

  /**
   * @generated from field: string description = 2;
   */
  description = "";

  /**
   * @generated from field: string granted_by = 3;
   */
  grantedBy = "";

  /**
   * @generated from field: string granted_to = 4;
   */
  grantedTo = "";

  /**
   * @generated from field: google.protobuf.Timestamp granted_date = 5;
   */
  grantedDate?: Timestamp;

  /**
   * @generated from field: auth.StateActivity last_modified = 20;
   */
  lastModified?: StateActivity;

  constructor(data?: PartialMessage<Specimen_Grant>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "ccbio.schema.v0.Specimen.Grant";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "description", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 3, name: "granted_by", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 4, name: "granted_to", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 5, name: "granted_date", kind: "message", T: Timestamp },
    { no: 20, name: "last_modified", kind: "message", T: StateActivity },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): Specimen_Grant {
    return new Specimen_Grant().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): Specimen_Grant {
    return new Specimen_Grant().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): Specimen_Grant {
    return new Specimen_Grant().fromJsonString(jsonString, options);
  }

  static equals(a: Specimen_Grant | PlainMessage<Specimen_Grant> | undefined, b: Specimen_Grant | PlainMessage<Specimen_Grant> | undefined): boolean {
    return proto3.util.equals(Specimen_Grant, a, b);
  }
}

