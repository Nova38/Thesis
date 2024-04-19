import { Attribute, Role, UserCollectionRoles, UserDirectMembership, UserEmbeddedRoles } from "./auth/v1/models_pb.js";
import { Collection, HiddenTxList, ItemKey, ReferenceKey, Suggestion } from "./auth/v1/objects_pb.js";
import { Specimen } from "./biochain/v1/state_pb.js";
import { Book, Group, SimpleItem } from "./sample/v0/items_pb.js";
import { KeySchema } from "./auth/v1/auth_pb.js";
import { isMessage  } from "@bufbuild/protobuf";

export type ItemType = 
  | UserDirectMembership
  | Role
  | UserCollectionRoles
  | UserEmbeddedRoles
  | Attribute
  | ReferenceKey
  | Collection
  | Suggestion
  | HiddenTxList
  | Specimen
  | SimpleItem
  | Group
  | Book

export const ItemKeySchema : Record<string, KeySchema> = {
  'auth.UserDirectMembership': KeySchema.fromJson({
    "itemType": "auth.UserDirectMembership",
    "itemKind": 2,
    "properties": "mspId,userId"
}),
  'auth.Role': KeySchema.fromJson({
    "itemType": "auth.Role",
    "itemKind": 2,
    "properties": "roleId"
}),
  'auth.UserCollectionRoles': KeySchema.fromJson({
    "itemType": "auth.UserCollectionRoles",
    "itemKind": 2,
    "properties": "mspId,userId"
}),
  'auth.UserEmbeddedRoles': KeySchema.fromJson({
    "itemType": "auth.UserEmbeddedRoles",
    "itemKind": 2,
    "properties": "mspId,userId"
}),
  'auth.Attribute': KeySchema.fromJson({
    "itemType": "auth.Attribute",
    "itemKind": 2,
    "properties": "mspId,oidvalue"
}),
  'auth.ReferenceKey': KeySchema.fromJson({
    "itemType": "auth.ReferenceKey",
    "itemKind": 4
}),
  'auth.Collection': KeySchema.fromJson({
    "itemType": "auth.Collection",
    "itemKind": 2,
    "properties": "collectionId"
}),
  'auth.Suggestion': KeySchema.fromJson({
    "itemType": "auth.Suggestion",
    "itemKind": 3,
    "properties": "suggestionId"
}),
  'auth.HiddenTxList': KeySchema.fromJson({
    "itemType": "auth.HiddenTxList",
    "itemKind": 3,
    "properties": ""
}),
  'biochain.v1.Specimen': KeySchema.fromJson({
    "itemType": "biochain.v1.Specimen",
    "itemKind": 2,
    "properties": "specimenId"
}),
  'sample.SimpleItem': KeySchema.fromJson({
    "itemType": "sample.SimpleItem",
    "itemKind": 2,
    "properties": "id"
}),
  'sample.Group': KeySchema.fromJson({
    "itemType": "sample.Group",
    "itemKind": 2,
    "properties": "groupId"
}),
  'sample.Book': KeySchema.fromJson({
    "itemType": "sample.Book",
    "itemKind": 2,
    "properties": "isbn"
}),
}
export type PrimaryItemType = 
  | UserDirectMembership
  | Role
  | UserCollectionRoles
  | UserEmbeddedRoles
  | Attribute
  | Collection
  | Specimen
  | SimpleItem
  | Group
  | Book

export const PrimaryItemKeySchema : Record<string, KeySchema> = {
  'auth.UserDirectMembership': KeySchema.fromJson({
    "itemType": "auth.UserDirectMembership",
    "itemKind": 2,
    "properties": "mspId,userId"
}),
  'auth.Role': KeySchema.fromJson({
    "itemType": "auth.Role",
    "itemKind": 2,
    "properties": "roleId"
}),
  'auth.UserCollectionRoles': KeySchema.fromJson({
    "itemType": "auth.UserCollectionRoles",
    "itemKind": 2,
    "properties": "mspId,userId"
}),
  'auth.UserEmbeddedRoles': KeySchema.fromJson({
    "itemType": "auth.UserEmbeddedRoles",
    "itemKind": 2,
    "properties": "mspId,userId"
}),
  'auth.Attribute': KeySchema.fromJson({
    "itemType": "auth.Attribute",
    "itemKind": 2,
    "properties": "mspId,oidvalue"
}),
  'auth.Collection': KeySchema.fromJson({
    "itemType": "auth.Collection",
    "itemKind": 2,
    "properties": "collectionId"
}),
  'biochain.v1.Specimen': KeySchema.fromJson({
    "itemType": "biochain.v1.Specimen",
    "itemKind": 2,
    "properties": "specimenId"
}),
  'sample.SimpleItem': KeySchema.fromJson({
    "itemType": "sample.SimpleItem",
    "itemKind": 2,
    "properties": "id"
}),
  'sample.Group': KeySchema.fromJson({
    "itemType": "sample.Group",
    "itemKind": 2,
    "properties": "groupId"
}),
  'sample.Book': KeySchema.fromJson({
    "itemType": "sample.Book",
    "itemKind": 2,
    "properties": "isbn"
}),
}
export type SecondaryItemType = 
  | Suggestion
  | HiddenTxList

export const SecondaryItemKeySchema : Record<string, KeySchema> = {
  'auth.Suggestion': KeySchema.fromJson({
    "itemType": "auth.Suggestion",
    "itemKind": 3,
    "properties": "suggestionId"
}),
  'auth.HiddenTxList': KeySchema.fromJson({
    "itemType": "auth.HiddenTxList",
    "itemKind": 3,
    "properties": ""
}),
}
export function PrimaryToKeySchema (item: PrimaryItemType){
  switch (true) {
// msp_id, user_id
// item.mspId, item.userId
    case isMessage (item, UserDirectMembership):
      return new ItemKey({
        itemType: 'auth.UserDirectMembership',
        itemKind: 2,
        collectionId: item?.collectionId,
        itemKeyParts: [ item.mspId, item.userId ]
      })
// role_id
// item.roleId
    case isMessage (item, Role):
      return new ItemKey({
        itemType: 'auth.Role',
        itemKind: 2,
        collectionId: item?.collectionId,
        itemKeyParts: [ item.roleId ]
      })
// msp_id, user_id
// item.mspId, item.userId
    case isMessage (item, UserCollectionRoles):
      return new ItemKey({
        itemType: 'auth.UserCollectionRoles',
        itemKind: 2,
        collectionId: item?.collectionId,
        itemKeyParts: [ item.mspId, item.userId ]
      })
// msp_id, user_id
// item.mspId, item.userId
    case isMessage (item, UserEmbeddedRoles):
      return new ItemKey({
        itemType: 'auth.UserEmbeddedRoles',
        itemKind: 2,
        collectionId: item?.collectionId,
        itemKeyParts: [ item.mspId, item.userId ]
      })
// msp_id, oidvalue
// item.mspId, item.oidvalue
    case isMessage (item, Attribute):
      return new ItemKey({
        itemType: 'auth.Attribute',
        itemKind: 2,
        collectionId: item?.collectionId,
        itemKeyParts: [ item.mspId, item.oidvalue ]
      })
// collection_id
// item.collectionId
    case isMessage (item, Collection):
      return new ItemKey({
        itemType: 'auth.Collection',
        itemKind: 2,
        collectionId: item?.collectionId,
        itemKeyParts: [ item.collectionId ]
      })
// specimen_id
// item.specimenId
    case isMessage (item, Specimen):
      return new ItemKey({
        itemType: 'biochain.v1.Specimen',
        itemKind: 2,
        collectionId: item?.collectionId,
        itemKeyParts: [ item.specimenId ]
      })
// id
// item.id
    case isMessage (item, SimpleItem):
      return new ItemKey({
        itemType: 'sample.SimpleItem',
        itemKind: 2,
        collectionId: item?.collectionId,
        itemKeyParts: [ item.id ]
      })
// group_id
// item.groupId
    case isMessage (item, Group):
      return new ItemKey({
        itemType: 'sample.Group',
        itemKind: 2,
        collectionId: item?.collectionId,
        itemKeyParts: [ item.groupId ]
      })
// isbn
// item.isbn
    case isMessage (item, Book):
      return new ItemKey({
        itemType: 'sample.Book',
        itemKind: 2,
        collectionId: item?.collectionId,
        itemKeyParts: [ item.isbn ]
      })
    default:
      throw new Error('Unknown item type')
  }
}
