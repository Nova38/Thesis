// import z from "zod";

// export type FieldMask = z.infer<typeof FieldMask>;
// export const FieldMask = z.object({
//   paths: z.array(z.string()),
// });

// export type Timestamp = z.infer<typeof Timestamp>;
// export const Timestamp = z.object({
//   seconds: z.number(),
//   nanos: z.number(),
// });

// export type ProtoAny = z.infer<typeof ProtoAny>;
// export const ProtoAny = z.object({
//   typeUrl: z.string(),
//   value: z.instanceof(Uint8Array),
// });

// // enums
// declare enum r_TransactionType {
//   UNSPECIFIED = 0,
//   INVOKE = 1,
//   QUERY = 2,
// }
// export type TransactionType = z.infer<typeof TransactionType>;
// export const TransactionType = z.nativeEnum(r_TransactionType);

// declare enum r_AuthType {
//   UNSPECIFIED = 0,
//   NONE = 1,
//   ROLE = 2,
//   IDENTITY = 3,
//   EMBEDDED_ROLE = 4,
// }
// export type AuthType = z.infer<typeof AuthType>;
// export const AuthType = z.nativeEnum(r_AuthType);

// declare enum r_ItemKind {
//   UNSPECIFIED = 0,
//   PRIMARY_ITEM = 2,
//   SUB_ITEM = 3,
//   REFERENCE = 4,
// }
// export type ItemKind = z.infer<typeof ItemKind>;
// export const ItemKind = z.nativeEnum(r_ItemKind);

// declare enum r_Action {
//   UNSPECIFIED = 0,
//   UTILITY = 1,
//   VIEW = 10,
//   CREATE = 11,
//   UPDATE = 12,
//   DELETE = 13,
//   SUGGEST_VIEW = 14,
//   SUGGEST_CREATE = 15,
//   SUGGEST_DELETE = 16,
//   SUGGEST_APPROVE = 17,
//   VIEW_HISTORY = 18,
//   VIEW_HIDDEN_TXS = 19,
//   HIDE_TX = 20,
//   UNHIDE_TX = 21,
//   REFERENCE_CREATE = 30,
//   REFERENCE_DELETE = 31,
//   REFERENCE_VIEW = 32,
// }

// export type Action = z.infer<typeof Action>;
// export const Action = z.nativeEnum(r_Action);

// declare enum r_TxError {
//   UNSPECIFIED = 0,
//   REQUEST_INVALID = 1,
//   RUNTIME = 2,
//   RUNTIME_BAD_OPS = 3,
//   KEY_NOT_FOUND = 4,
//   KEY_ALREADY_EXISTS = 5,
//   COLLECTION_INVALID_ID = 11,
//   COLLECTION_UNREGISTERED = 12,
//   COLLECTION_ALREADY_REGISTERED = 13,
//   COLLECTION_INVALID = 14,
//   COLLECTION_INVALID_ITEM_TYPE = 15,
//   COLLECTION_INVALID_ROLE_ID = 16,
//   USER_INVALID_ID = 20,
//   USER_UNREGISTERED = 21,
//   USER_ALREADY_REGISTERED = 22,
//   USER_INVALID = 23,
//   USER_NO_ROLE = 24,
//   USER_PERMISSION_DENIED = 26,
//   ITEM_INVALID_ID = 31,
//   ITEM_UNREGISTERED = 32,
//   ITEM_ALREADY_REGISTERED = 33,
//   ITEM_INVALID = 34,
//   INVALID_ITEM_FIELD_PATH = 35,
//   INVALID_ITEM_FIELD_VALUE = 36,
// }
// export type TxError = z.infer<typeof TxError>;
// export const TxError = z.nativeEnum(r_TxError);

// declare enum r_Specimen_Secondary_SEX {
//   SEX_UNDEFINED = 0,
//   SEX_UNKNOWN = 1,
//   SEX_ATYPICAL = 2,
//   SEX_MALE = 3,
//   SEX_FEMALE = 4,
// }
// export type Specimen_Secondary_SEX = z.infer<typeof Specimen_Secondary_SEX>;
// export const Specimen_Secondary_SEX = z.nativeEnum(r_Specimen_Secondary_SEX);

// declare enum r_Specimen_Secondary_AGE {
//   AGE_UNDEFINED = 0,
//   AGE_UNKNOWN = 1,
//   AGE_NEST = 2,
//   AGE_EMBRYO_EGG = 3,
//   AGE_CHICK_SUBADULT = 4,
//   AGE_ADULT = 5,
//   AGE_CONTINGENT = 6,
// }
// export type Specimen_Secondary_AGE = z.infer<typeof Specimen_Secondary_AGE>;
// export const Specimen_Secondary_AGE = z.nativeEnum(r_Specimen_Secondary_AGE);

// // objects

// export type User = z.infer<typeof User>;
// export const User = z.object({
//   mspId: z.string(),
//   userId: z.string(),
// });

// export type KeySchema = z.infer<typeof KeySchema>;
// export const KeySchema = z.object({
//   itemType: z.string(),
//   itemKind: z.nativeEnum(r_ItemKind),
//   properties: FieldMask.optional(),
// });

// export type StateActivity = z.infer<typeof StateActivity>;
// export const StateActivity = z.object({
//   txId: z.string(),
//   mspId: z.string(),
//   userId: z.string(),
//   timestamp: Timestamp.optional(),
//   note: z.string(),
// });

// export type StateActivity2 = z.infer<typeof StateActivity2>;
// export const StateActivity2 = z.object({
//   txId: z.string(),
//   mspId: z.string(),
//   userId: z.string(),
//   timestamp: Timestamp.optional(),
//   note: z.string(),
// });

// export type Operation = z.infer<typeof Operation>;
// export const Operation = z.object({
//   action: z.nativeEnum(r_Action),
//   collectionId: z.string(),
//   itemType: z.string(),
//   paths: FieldMask.optional(),
// });

// export type ItemKey = z.infer<typeof ItemKey>;
// export const ItemKey = z.object({
//   collectionId: z.string(),
//   itemType: z.string(),
//   itemKind: z.nativeEnum(r_ItemKind),
//   itemKeyParts: z.array(z.string()),
// });

// export type Item = z.infer<typeof Item>;
// export const Item = z.object({
//   key: ItemKey.optional(),
//   value: ProtoAny.optional(),
// });

// export type ReferenceKey = z.infer<typeof ReferenceKey>;
// export const ReferenceKey = z.object({
//   key1: ItemKey.optional(),
//   key2: ItemKey.optional(),
// });

// export type UserEmbeddedRoles = z.infer<typeof UserEmbeddedRoles>;
// export const UserEmbeddedRoles = z.object({
//   collectionId: z.string(),
//   mspId: z.string(),
//   userId: z.string(),
//   roles: z.object({}),
// });

// export type RoleList = z.infer<typeof RoleList>;
// export const RoleList = z.object({
//   roleId: z.array(z.string()),
// });

// export type UserCollectionRoles = z.infer<typeof UserCollectionRoles>;
// export const UserCollectionRoles = z.object({
//   collectionId: z.string(),
//   mspId: z.string(),
//   userId: z.string(),
//   roleIds: z.array(z.string()),
//   note: z.string(),
// });

// export type Suggestion = z.infer<typeof Suggestion>;
// export const Suggestion = z.object({
//   primaryKey: ItemKey.optional(),
//   suggestionId: z.string(),
//   paths: FieldMask.optional(),
//   value: ProtoAny.optional(),
// });

// export type HiddenTx = z.infer<typeof HiddenTx>;
// export const HiddenTx = z.object({
//   txId: z.string(),
//   mspId: z.string(),
//   userId: z.string(),
//   timestamp: Timestamp.optional(),
//   note: z.string(),
// });

// export type HiddenTxList = z.infer<typeof HiddenTxList>;
// export const HiddenTxList = z.object({
//   primaryKey: ItemKey.optional(),
//   txs: z.array(HiddenTx),
// });

// export type Reference = z.infer<typeof Reference>;
// export const Reference = z.object({
//   reference: ReferenceKey.optional(),
//   item1: Item.optional(),
//   item2: Item.optional(),
// });

// export type HistoryEntry = z.infer<typeof HistoryEntry>;
// export const HistoryEntry = z.object({
//   txId: z.string(),
//   isDelete: z.boolean(),
//   isHidden: z.boolean(),
//   timestamp: Timestamp.optional(),
//   note: z.string(),
//   value: ProtoAny.optional(),
// });

// export type History = z.infer<typeof History>;
// export const History = z.object({
//   entries: z.array(HistoryEntry),
//   hiddenTxs: HiddenTxList.optional(),
// });

// export type PathPolicy = z.infer<typeof PathPolicy>;
// export const PathPolicy = z.object({
//   path: z.string(),
//   fullPath: z.string(),
//   allowSubPaths: z.boolean(),
//   subPaths: z.object({}),
//   actions: z.nativeEnum(r_Action),
// });

// export type Polices = z.infer<typeof Polices>;
// export const Polices = z.object({
//   itemPolicies: z.object({}),
//   defaultPolicy: PathPolicy.optional(),
//   defaultExcludedTypes: z.array(z.string()),
// });

// export type Date = z.infer<typeof Date>;
// export const Date = z.object({
//   verbatim: z.string(),
//   timestamp: Timestamp.optional(),
//   year: z.number(),
//   month: z.string(),
//   day: z.number(),
// });

// export type Researcher = z.infer<typeof Researcher>;
// export const Researcher = z.object({
//   firstName: z.string(),
//   lastName: z.string(),
//   middleName: z.string(),
// });

// export const zSpecimen = z.object({
//   collectionId: z.string(),
//   specimenId: z.string(),
//   primary: z.object({
//     catalogNumber: z.string(),
//     accessionNumber: z.string(),
//     fieldNumber: z.string(),
//     tissueNumber: z.string(),
//     cataloger: z.string(),
//     collector: z.string(),
//     determiner: z.string(),
//     fieldDate: z.date().optional(),
//     catalogDate: z.date().optional(),
//     determinedDate: z.date().optional(),
//     determinedReason: z.string(),
//     originalDate: z.date().optional(),
//     lastModified: StateActivity.optional(),
//   }),
//   secondary: z.record(
//     z.object({
//       verbatim: z.string(),
//     }),
//   ),
//   taxon: z.object({
//     kingdom: z.string(),
//     phylum: z.string(),
//     class: z.string(),
//     order: z.string(),
//     family: z.string(),
//     genus: z.string(),
//     species: z.string(),
//     subspecies: z.string(),
//     lastModified: StateActivity.optional(),
//   }),
//   georeference: z.object({
//     country: z.string(),
//     stateProvince: z.string(),
//     county: z.string(),
//     locality: z.string(),
//     latitude: z.number(),
//     longitude: z.number(),
//     habitat: z.string(),
//     continent: z.string(),
//     locationRemarks: z.string(),
//     coordinateUncertaintyInMeters: z.number(),
//     georeferenceBy: z.string(),
//     georeferenceDate: z.date().optional(),
//     georeferenceProtocol: z.string(),
//     geodeticDatum: z.string(),
//     footprintWkt: z.string(),
//     notes: z.string(),
//     lastModified: StateActivity.optional(),
//   }),
//   images: z.record(
//     z.object({
//       id: z.string(),
//       url: z.string(),
//       notes: z.string(),
//       hash: z.string(),
//       lastModified: StateActivity.optional(),
//     }),
//   ),
//   loans: z.record(
//     z.object({
//       id: z.string(),
//       description: z.string(),
//       loanedBy: z.string(),
//       loanedTo: z.string(),
//       loanedDate: z.date().optional(),
//       lastModified: StateActivity.optional(),
//     }),
//   ),
//   grants: z.record(
//     z.object({
//       id: z.string(),
//       description: z.string(),
//       grantedBy: z.string(),
//       grantedTo: z.string(),
//       grantedDate: z.date().optional(),
//       lastModified: StateActivity.optional(),
//     }),
//   ),
//   lastModified: StateActivity.optional(),
// });
