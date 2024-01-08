// extends Message[<a-zA-Z_]*>

export declare interface FieldMask {
  paths: string[];
}

export declare interface Timestamp {
  seconds: Number;
  nanos: number;
}

export declare interface ProtoAny {
  typeUrl: string;
  value: Uint8Array;
}

// AUTH

declare enum TransactionType {
  UNSPECIFIED = 0,
  INVOKE = 1,
  QUERY = 2,
}
declare enum AuthType {
  UNSPECIFIED = 0,
  NONE = 1,
  ROLE = 2,
  IDENTITY = 3,
  EMBEDDED_ROLE = 4,
}
declare enum ItemKind {
  UNSPECIFIED = 0,
  PRIMARY_ITEM = 2,
  SUB_ITEM = 3,
  REFERENCE = 4,
}

declare enum Action {
  UNSPECIFIED = 0,
  UTILITY = 1,
  VIEW = 10,
  CREATE = 11,
  UPDATE = 12,
  DELETE = 13,
  SUGGEST_VIEW = 14,
  SUGGEST_CREATE = 15,
  SUGGEST_DELETE = 16,
  SUGGEST_APPROVE = 17,
  VIEW_HISTORY = 18,
  VIEW_HIDDEN_TXS = 19,
  HIDE_TX = 20,
  UNHIDE_TX = 21,
  REFERENCE_CREATE = 30,
  REFERENCE_DELETE = 31,
  REFERENCE_VIEW = 32,
}
declare enum TxError {
  UNSPECIFIED = 0,
  REQUEST_INVALID = 1,
  RUNTIME = 2,
  RUNTIME_BAD_OPS = 3,
  KEY_NOT_FOUND = 4,
  KEY_ALREADY_EXISTS = 5,
  COLLECTION_INVALID_ID = 11,
  COLLECTION_UNREGISTERED = 12,
  COLLECTION_ALREADY_REGISTERED = 13,
  COLLECTION_INVALID = 14,
  COLLECTION_INVALID_ITEM_TYPE = 15,
  COLLECTION_INVALID_ROLE_ID = 16,
  USER_INVALID_ID = 20,
  USER_UNREGISTERED = 21,
  USER_ALREADY_REGISTERED = 22,
  USER_INVALID = 23,
  USER_NO_ROLE = 24,
  USER_PERMISSION_DENIED = 26,
  ITEM_INVALID_ID = 31,
  ITEM_UNREGISTERED = 32,
  ITEM_ALREADY_REGISTERED = 33,
  ITEM_INVALID = 34,
  INVALID_ITEM_FIELD_PATH = 35,
  INVALID_ITEM_FIELD_VALUE = 36,
}
declare interface User {
  mspId: string;
  userId: string;
}
interface KeySchema {
  itemType: string;
  itemKind: ItemKind;
  properties?: FieldMask;
}

interface StateActivity {
  txId: string;
  mspId: string;
  userId: string;
  timestamp?: Timestamp;
  note: string;
}
declare interface StateActivity2 {
  txId: string;

  mspId: string;

  userId: string;

  timestamp?: Timestamp;
  note: string;
}
declare interface Operation {
  action: Action;
  collectionId: string;
  itemType: string;
  paths?: FieldMask;
}

declare interface ItemKey {
  collectionId: string;
  itemType: string;
  itemKind: ItemKind;
  itemKeyParts: string[];
}
declare interface Item {
  key?: ItemKey;
  value?: ProtoAny;
}
declare interface ReferenceKey {
  key1?: ItemKey;
  key2?: ItemKey;
}
declare interface Collection {
  collectionId: string;
  name: string;
  authType: AuthType;
  itemTypes: string[];
  default?: Polices;
}
declare interface Role {
  collectionId: string;
  roleId: string;
  polices?: Polices;
  note: string;
  parentRoleIds: string[];
}
declare interface Attribute {
  collectionId: string;
  mspId: string;
  oid: string;
  value: string;
  polices?: Polices;
  note: string;
}
declare interface UserDirectMembership {
  collectionId: string;
  mspId: string;
  userId: string;
  polices?: Polices;
  note: string;
}
declare interface UserEmbeddedRoles {
  collectionId: string;
  mspId: string;
  userId: string;
  roles: {
    [key: string]: RoleList;
  };
}
declare interface RoleList {
  roleId: string[];
}
declare interface UserCollectionRoles {
  collectionId: string;
  mspId: string;
  userId: string;
  roleIds: string[];
  note: string;
}
declare interface Suggestion {
  primaryKey?: ItemKey;
  suggestionId: string;
  paths?: FieldMask;
  value?: ProtoAny;
}
declare interface HiddenTx {
  txId: string;
  mspId: string;
  userId: string;
  timestamp?: Timestamp;
  note: string;
}
declare interface HiddenTxList {
  primaryKey?: ItemKey;
  txs: HiddenTx[];
}
declare interface Reference {
  reference?: ReferenceKey;
  item1?: Item;
  item2?: Item;
}
declare interface FullItem {
  item?: Item;
  history?: History;
  suggestions: Suggestion[];
  references: Reference[];
}
declare interface HistoryEntry {
  txId: string;
  isDelete: boolean;
  isHidden: boolean;
  timestamp?: Timestamp;
  note: string;
  value?: ProtoAny;
}
declare interface History {
  entries: HistoryEntry[];
  hiddenTxs?: HiddenTxList;
}
declare interface PathPolicy {
  path: string;
  fullPath: string;
  allowSubPaths: boolean;
  subPaths: {
    [key: string]: PathPolicy;
  };
  actions: Action[];
}
declare interface Polices {
  itemPolicies: {
    [key: string]: PathPolicy;
  };
  defaultPolicy?: PathPolicy;
  defaultExcludedTypes: string[];
}

// CCBIO

declare interface SpecimenHistory {
  entries: SpecimenHistoryEntry[];
  hiddenTxs?: HiddenTxList;
}
declare interface SpecimenHistoryEntry {
  txId: string;
  isDelete: boolean;
  isHidden: boolean;
  timestamp?: Timestamp;
  note: string;
  value?: Specimen;
}
declare interface SpecimenUpdate {
  specimen?: Specimen;
  mask?: FieldMask;
}
declare interface Date {
  verbatim: string;
  timestamp?: Timestamp;
  year: number;
  month: string;
  day: number;
}
declare interface Researcher {
  firstName: string;
  lastName: string;
  middleName: string;
}
declare interface Specimen {
  collectionId: string;
  specimenId: string;
  primary?: Specimen_Primary;
  secondary?: Specimen_Secondary;
  taxon?: Specimen_Taxon;
  georeference?: Specimen_Georeference;
  images: {
    [key: string]: Specimen_Image;
  };
  loans: {
    [key: string]: Specimen_Loan;
  };
  grants: {
    [key: string]: Specimen_Grant;
  };
  lastModified?: StateActivity;
}
declare interface Specimen_Primary {
  catalogNumber: string;
  accessionNumber: string;
  fieldNumber: string;
  tissueNumber: string;
  cataloger: string;
  collector: string;
  determiner: string;
  fieldDate?: Date;
  catalogDate?: Date;
  determinedDate?: Date;
  determinedReason: string;
  originalDate?: Date;
  lastModified?: StateActivity;
}
declare interface Specimen_Secondary {
  sex: Specimen_Secondary_SEX;
  age: Specimen_Secondary_AGE;
  weight: number;
  weightUnits: string;
  preparations: {
    [key: string]: Specimen_Secondary_Preparation;
  };
  condition: string;
  molt: string;
  notes: string;
  lastModified?: StateActivity;
}
declare enum Specimen_Secondary_SEX {
  SEX_UNDEFINED = 0,
  SEX_UNKNOWN = 1,
  SEX_ATYPICAL = 2,
  SEX_MALE = 3,
  SEX_FEMALE = 4,
}
declare enum Specimen_Secondary_AGE {
  AGE_UNDEFINED = 0,
  AGE_UNKNOWN = 1,
  AGE_NEST = 2,
  AGE_EMBRYO_EGG = 3,
  AGE_CHICK_SUBADULT = 4,
  AGE_ADULT = 5,
  AGE_CONTINGENT = 6,
}
declare interface Specimen_Secondary_Preparation {
  verbatim: string;
}
declare interface Specimen_Taxon {
  kingdom: string;
  phylum: string;
  class: string;
  order: string;
  family: string;
  genus: string;
  species: string;
  subspecies: string;
  lastModified?: StateActivity;
}
declare interface Specimen_Georeference {
  country: string;
  stateProvince: string;
  county: string;
  locality: string;
  latitude: number;
  longitude: number;
  habitat: string;
  continent: string;
  locationRemarks: string;
  coordinateUncertaintyInMeters: number;
  georeferenceBy: string;
  georeferenceDate?: Date;
  georeferenceProtocol: string;
  geodeticDatum: string;
  footprintWkt: string;
  notes: string;
  lastModified?: StateActivity;
}
declare interface Specimen_Image {
  id: string;
  url: string;
  notes: string;
  hash: string;
  lastModified?: StateActivity;
}
declare interface Specimen_Loan {
  id: string;
  description: string;
  loanedBy: string;
  loanedTo: string;
  loanedDate?: Date;
  lastModified?: StateActivity;
}
declare interface Specimen_Grant {
  id: string;
  description: string;
  grantedBy: string;
  grantedTo: string;
  grantedDate?: Date;
  lastModified?: StateActivity;
}
