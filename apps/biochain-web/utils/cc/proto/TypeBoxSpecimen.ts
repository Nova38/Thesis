import { Type, type Static } from '@sinclair/typebox'

export type FieldMask = Static<typeof FieldMask>
export const FieldMask = Type.Object({
  paths: Type.Array(Type.String()),
})
const Timestamp = Type.Transform(Type.Number())
  .Decode((value) => new Date(value)) // decode: number to Date
  .Encode((value) => value.getTime())
type T = Static<typeof Timestamp>

// export type Timestamp = Static<typeof Timestamp>
// export const Timestamp = Type.Object({
//   seconds: Type.Number(),
//   nanos: Type.Number(),
// })

export type ProtoAny = Static<typeof ProtoAny>
export const ProtoAny = Type.Object({
  typeUrl: Type.String(),
  value: Type.Uint8Array(),
})

type ProtoDate = Static<typeof ProtoDate>
const ProtoDate = Type.Object({
  verbatim: Type.String(),
  timestamp: Type.Optional(Type.Date()),
  year: Type.Number(),
  month: Type.String(),
  day: Type.Number(),
})

type StateActivity = Static<typeof StateActivity>
const StateActivity = Type.Object({
  txId: Type.String(),
  mspId: Type.String(),
  userId: Type.String(),
  timestamp: Type.Optional(Type.String({ format: 'date-time' })),
  note: Type.String(),
})

enum EnumTransactionType {
  UNSPECIFIED = 0,
  INVOKE = 1,
  QUERY = 2,
}

type TransactionType = Static<typeof TransactionType>
const TransactionType = Type.Enum(EnumTransactionType)

enum EnumAuthType {
  UNSPECIFIED = 0,
  NONE = 1,
  ROLE = 2,
  IDENTITY = 3,
  EMBEDDED_ROLE = 4,
}

type AuthType = Static<typeof AuthType>
const AuthType = Type.Enum(EnumAuthType)

enum EnumItemKind {
  UNSPECIFIED = 0,
  PRIMARY_ITEM = 2,
  SUB_ITEM = 3,
  REFERENCE = 4,
}

type ItemKind = Static<typeof ItemKind>
const ItemKind = Type.Enum(EnumItemKind)

enum EnumAction {
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

type Action = Static<typeof Action>
const Action = Type.Enum(EnumAction)

enum EnumTxError {
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

type TxError = Static<typeof TxError>
const TxError = Type.Enum(EnumTxError)

type User = Static<typeof User>
const User = Type.Object({
  mspId: Type.String(),
  userId: Type.String(),
})

type KeySchema = Static<typeof KeySchema>
const KeySchema = Type.Object({
  itemType: Type.String(),
  itemKind: ItemKind,
  properties: Type.Optional(FieldMask),
})

// type StateActivity = Static<typeof StateActivity>
// const StateActivity = Type.Object({
//   txId: Type.String(),
//   mspId: Type.String(),
//   userId: Type.String(),
//   timestamp: Type.Optional(Timestamp),
//   note: Type.String(),
// })

// type ProtoDate = Static<typeof ProtoDate>
// const ProtoDate = Type.Object({
//   verbatim: Type.String(),
//   timestamp: Type.Optional(Timestamp),
//   year: Type.Number(),
//   month: Type.String(),
//   day: Type.Number(),
// })

// type Operation = Static<typeof Operation>
// const Operation = Type.Object({
//   action: Action,
//   collectionId: Type.String(),
//   itemType: Type.String(),
//   paths: Type.Optional(FieldMask),
// })

type ItemKey = Static<typeof ItemKey>
const ItemKey = Type.Object({
  collectionId: Type.String(),
  itemType: Type.String(),
  itemKind: ItemKind,
  itemKeyParts: Type.Array(Type.String()),
})

type Item = Static<typeof Item>
const Item = Type.Object({
  key: Type.Optional(ItemKey),
  value: Type.Optional(ProtoAny),
})

// type ReferenceKey = Static<typeof ReferenceKey>
// const ReferenceKey = Type.Object({
//   key1: Type.Optional(ItemKey),
//   key2: Type.Optional(ItemKey),
// })

// type Collection = Static<typeof Collection>
// const Collection = Type.Object({
//   collectionId: Type.String(),
//   name: Type.String(),
//   authType: AuthType,
//   itemTypes: Type.Array(Type.String()),
//   default: Type.Optional(Polices),
// })

// type Role = Static<typeof Role>
// const Role = Type.Object({
//   collectionId: Type.String(),
//   roleId: Type.String(),
//   polices: Type.Optional(Polices),
//   note: Type.String(),
//   parentRoleIds: Type.Array(Type.String()),
// })

// type Attribute = Static<typeof Attribute>
// const Attribute = Type.Object({
//   collectionId: Type.String(),
//   mspId: Type.String(),
//   oid: Type.String(),
//   value: Type.String(),
//   polices: Type.Optional(Polices),
//   note: Type.String(),
// })

// type UserDirectMembership = Static<typeof UserDirectMembership>
// const UserDirectMembership = Type.Object({
//   collectionId: Type.String(),
//   mspId: Type.String(),
//   userId: Type.String(),
//   polices: Type.Optional(Polices),
//   note: Type.String(),
// })

// type UserEmbeddedRoles = Static<typeof UserEmbeddedRoles>
// const UserEmbeddedRoles = Type.Object({
//   collectionId: Type.String(),
//   mspId: Type.String(),
//   userId: Type.String(),
//   roles: Type.Object(
//     {},
//     {
//       additionalProperties: RoleList,
//     },
//   ),
// })

// type RoleList = Static<typeof RoleList>
// const RoleList = Type.Object({
//   roleId: Type.Array(Type.String()),
// })

// type UserCollectionRoles = Static<typeof UserCollectionRoles>
// const UserCollectionRoles = Type.Object({
//   collectionId: Type.String(),
//   mspId: Type.String(),
//   userId: Type.String(),
//   roleIds: Type.Array(Type.String()),
//   note: Type.String(),
// })

// type Suggestion = Static<typeof Suggestion>
// const Suggestion = Type.Object({
//   primaryKey: Type.Optional(ItemKey),
//   suggestionId: Type.String(),
//   paths: Type.Optional(FieldMask),
//   value: Type.Optional(ProtoAny),
// })

type HiddenTx = Static<typeof HiddenTx>
const HiddenTx = Type.Object({
  txId: Type.String(),
  mspId: Type.String(),
  userId: Type.String(),
  timestamp: Type.Optional(Timestamp),
  note: Type.String(),
})

type HiddenTxList = Static<typeof HiddenTxList>
const HiddenTxList = Type.Object({
  primaryKey: Type.Optional(ItemKey),
  txs: Type.Array(HiddenTx),
})

// type Reference = Static<typeof Reference>
// const Reference = Type.Object({
//   reference: Type.Optional(ReferenceKey),
//   item1: Type.Optional(Item),
//   item2: Type.Optional(Item),
// })

type HistoryEntry = Static<typeof HistoryEntry>
const HistoryEntry = Type.Object({
  txId: Type.String(),
  isDelete: Type.Boolean(),
  isHidden: Type.Boolean(),
  timestamp: Type.Optional(Timestamp),
  note: Type.String(),
  value: Type.Optional(ProtoAny),
})

type History = Static<typeof History>
const History = Type.Object({
  entries: Type.Array(HistoryEntry),
  hiddenTxs: Type.Optional(HiddenTxList),
})

// type FullItem = Static<typeof FullItem>
// const FullItem = Type.Object({
//   item: Type.Optional(Item),
//   history: Type.Optional(History),
//   suggestions: Type.Array(Suggestion),
//   references: Type.Array(Reference),
// })

// type PathPolicy = Static<typeof PathPolicy>
// const PathPolicy = Type.Recursive((This) =>
//   Type.Object({
//     path: Type.String(),
//     fullPath: Type.String(),
//     allowSubPaths: Type.Boolean(),
//     subPaths: Type.Object(
//       {},
//       {
//         additionalProperties: This,
//       },
//     ),
//     actions: Type.Array(Action),
//   }),
// )

// type Polices = Static<typeof Polices>
// const Polices = Type.Object({
//   itemPolicies: Type.Object(
//     {},
//     {
//       additionalProperties: PathPolicy,
//     },
//   ),
//   defaultPolicy: Type.Optional(PathPolicy),
//   defaultExcludedTypes: Type.Array(Type.String()),
// })

// type Researcher = Static<typeof Researcher>
// const Researcher = Type.Object({
//   firstName: Type.String(),
//   lastName: Type.String(),
//   middleName: Type.String(),
// })

enum EnumSpecimen_Secondary_SEX {
  SEX_UNDEFINED = 0,
  SEX_UNKNOWN = 1,
  SEX_ATYPICAL = 2,
  SEX_MALE = 3,
  SEX_FEMALE = 4,
}

type Specimen_SEX = Static<typeof Specimen_SEX>
const Specimen_SEX = Type.Enum(EnumSpecimen_Secondary_SEX)

enum EnumSpecimen_Secondary_AGE {
  AGE_UNDEFINED = 0,
  AGE_UNKNOWN = 1,
  AGE_NEST = 2,
  AGE_EMBRYO_EGG = 3,
  AGE_CHICK_SUBADULT = 4,
  AGE_ADULT = 5,
  AGE_CONTINGENT = 6,
}

type Specimen_Secondary_AGE = Static<typeof Specimen_Secondary_AGE>
const Specimen_Secondary_AGE = Type.Enum(EnumSpecimen_Secondary_AGE)

type Specimen = Static<typeof Specimen>
const Specimen = Type.Object({
  collectionId: Type.String(),
  specimenId: Type.String(),
  primary: Type.Object({
    catalogNumber: Type.String(),
    accessionNumber: Type.String(),
    fieldNumber: Type.String(),
    tissueNumber: Type.String(),
    cataloger: Type.String(),
    collector: Type.String(),
    determiner: Type.String(),
    fieldDate: Type.Optional(ProtoDate),
    catalogDate: Type.Optional(ProtoDate),
    determinedDate: Type.Optional(ProtoDate),
    determinedReason: Type.String(),
    originalDate: Type.Optional(ProtoDate),
    lastModified: Type.Optional(StateActivity),
  }),
  secondary: Type.Object({
    sex: Specimen_SEX,
    age: Specimen_Secondary_AGE,
    weight: Type.Number(),
    weightUnits: Type.String(),
    preparations: Type.Record(
      Type.String(),
      Type.Object({
        verbatim: Type.String(),
      }),
    ),
    condition: Type.String(),
    molt: Type.String(),
    notes: Type.String(),
    lastModified: Type.Optional(StateActivity),
  }),
  taxon: Type.Object({
    kingdom: Type.String(),
    phylum: Type.String(),
    class: Type.String(),
    order: Type.String(),
    family: Type.String(),
    genus: Type.String(),
    species: Type.String(),
    subspecies: Type.String(),
    lastModified: Type.Optional(StateActivity),
  }),
  georeference: Type.Object({
    country: Type.String(),
    stateProvince: Type.String(),
    county: Type.String(),
    locality: Type.String(),
    latitude: Type.Number(),
    longitude: Type.Number(),
    habitat: Type.String(),
    continent: Type.String(),
    locationRemarks: Type.String(),
    coordinateUncertaintyInMeters: Type.Number(),
    georeferenceBy: Type.String(),
    georeferenceDate: Type.Optional(Type.Date()),
    georeferenceProtocol: Type.String(),
    geodeticDatum: Type.String(),
    footprintWkt: Type.String(),
    notes: Type.String(),
    lastModified: Type.Optional(StateActivity),
  }),

  images: Type.Record(
    Type.String(),
    Type.Object({
      id: Type.String(),
      url: Type.String(),
      notes: Type.String(),
      hash: Type.String(),
      lastModified: Type.Optional(StateActivity),
    }),
  ),
  loans: Type.Record(
    Type.String(),
    Type.Object({
      id: Type.String(),
      description: Type.String(),
      loanedBy: Type.String(),
      loanedTo: Type.String(),
      loanedDate: Type.Optional(Type.Date()),
      lastModified: Type.Optional(StateActivity),
    }),
  ),
  grants: Type.Record(
    Type.String(),
    Type.Object({
      id: Type.String(),
      description: Type.String(),
      grantedBy: Type.String(),
      grantedTo: Type.String(),
      grantedDate: Type.Optional(Type.Date()),
      lastModified: Type.Optional(StateActivity),
    }),
  ),
  lastModified: Type.Optional(StateActivity),
})

type SpecimenHistoryEntry = Static<typeof SpecimenHistoryEntry>
const SpecimenHistoryEntry = Type.Object({
  txId: Type.String(),
  isDelete: Type.Boolean(),
  isHidden: Type.Boolean(),
  timestamp: Type.Optional(Timestamp),
  note: Type.String(),
  value: Type.Optional(Specimen),
})

type SpecimenHistory = Static<typeof SpecimenHistory>
const SpecimenHistory = Type.Object({
  entries: Type.Array(SpecimenHistoryEntry),
  hiddenTxs: Type.Optional(HiddenTxList),
})

export const Schema = {
  Specimen: Specimen,
}
