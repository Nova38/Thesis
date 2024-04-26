import { Attribute } from "./saacs/auth/v0/attribute_pb.js";
import { Collection } from "./saacs/auth/v0/collection_pb.js";
import { UserDirectMembership } from "./saacs/auth/v0/identity_pb.js";
import { Role, UserCollectionRoles, UserGlobalRoles } from "./saacs/auth/v0/roles_pb.js";
import { HiddenTxList } from "./saacs/common/v0/history_pb.js";
import { Specimen } from "./saacs/biochain/v0/state_pb.js";
import { Suggestion } from "./saacs/common/v0/suggestion_pb.js";
import { Book } from "./saacs/example/v0/book_pb.js";
import { Group, SimpleItem } from "./saacs/example/v0/items_pb.js";
import { ItemWithNestedKey } from "./saacs/example/v0/nested_pb.js";
import { ItemKey, KeySchema } from "./saacs/common/v0/item_pb.js";
import { isMessage  } from "@bufbuild/protobuf";

export type ItemTypeMessage =
  | Attribute
  | Collection
  | UserDirectMembership
  | Role
  | UserCollectionRoles
  | UserGlobalRoles
  | HiddenTxList
  | Specimen
  | Suggestion
  | Book
  | SimpleItem
  | Group
  | ItemWithNestedKey

export const ItemKeySchema : Record<string, KeySchema> = {
  'saacs.auth.v0.Attribute': KeySchema.fromJson({
    "itemType": "saacs.auth.v0.Attribute",
    "itemKind": 2,
    "properties": "mspId,oid,value"
}),
  'saacs.auth.v0.Collection': KeySchema.fromJson({
    "itemType": "saacs.auth.v0.Collection",
    "itemKind": 2,
    "properties": "collectionId"
}),
  'saacs.auth.v0.UserDirectMembership': KeySchema.fromJson({
    "itemType": "saacs.auth.v0.UserDirectMembership",
    "itemKind": 2,
    "properties": "mspId,userId"
}),
  'saacs.auth.v0.Role': KeySchema.fromJson({
    "itemType": "saacs.auth.v0.Role",
    "itemKind": 2,
    "properties": "roleId"
}),
  'saacs.auth.v0.UserCollectionRoles': KeySchema.fromJson({
    "itemType": "saacs.auth.v0.UserCollectionRoles",
    "itemKind": 2,
    "properties": "mspId,userId"
}),
  'saacs.auth.v0.UserGlobalRoles': KeySchema.fromJson({
    "itemType": "saacs.auth.v0.UserGlobalRoles",
    "itemKind": 2,
    "properties": "mspId,userId"
}),
  'saacs.common.v0.HiddenTxList': KeySchema.fromJson({
    "itemType": "saacs.common.v0.HiddenTxList",
    "itemKind": 3,
    "properties": "mspId"
}),
  'saacs.biochain.v0.Specimen': KeySchema.fromJson({
    "itemType": "saacs.biochain.v0.Specimen",
    "itemKind": 2,
    "properties": "specimenId"
}),
  'saacs.common.v0.Suggestion': KeySchema.fromJson({
    "itemType": "saacs.common.v0.Suggestion",
    "itemKind": 3,
    "properties": "suggestionId"
}),
  'saacs.sample.v0.Book': KeySchema.fromJson({
    "itemType": "saacs.sample.v0.Book",
    "itemKind": 2,
    "properties": "isbn"
}),
  'saacs.sample.v0.SimpleItem': KeySchema.fromJson({
    "itemType": "saacs.sample.v0.SimpleItem",
    "itemKind": 2,
    "properties": "id"
}),
  'saacs.sample.v0.Group': KeySchema.fromJson({
    "itemType": "saacs.sample.v0.Group",
    "itemKind": 2,
    "properties": "groupId"
}),
  'saacs.sample.v0.ItemWithNestedKey': KeySchema.fromJson({
    "itemType": "saacs.sample.v0.ItemWithNestedKey",
    "itemKind": 2,
    "properties": "id,nested.part1,nested.part2"
}),
}

export type PrimaryItemTypeMessage =
  | Attribute
  | Collection
  | UserDirectMembership
  | Role
  | UserCollectionRoles
  | UserGlobalRoles
  | Specimen
  | Book
  | SimpleItem
  | Group
  | ItemWithNestedKey

export const PrimaryItemKeySchema : Record<string, KeySchema> = {
  'saacs.auth.v0.Attribute': KeySchema.fromJson({
    "itemType": "saacs.auth.v0.Attribute",
    "itemKind": 2,
    "properties": "mspId,oid,value"
}),
  'saacs.auth.v0.Collection': KeySchema.fromJson({
    "itemType": "saacs.auth.v0.Collection",
    "itemKind": 2,
    "properties": "collectionId"
}),
  'saacs.auth.v0.UserDirectMembership': KeySchema.fromJson({
    "itemType": "saacs.auth.v0.UserDirectMembership",
    "itemKind": 2,
    "properties": "mspId,userId"
}),
  'saacs.auth.v0.Role': KeySchema.fromJson({
    "itemType": "saacs.auth.v0.Role",
    "itemKind": 2,
    "properties": "roleId"
}),
  'saacs.auth.v0.UserCollectionRoles': KeySchema.fromJson({
    "itemType": "saacs.auth.v0.UserCollectionRoles",
    "itemKind": 2,
    "properties": "mspId,userId"
}),
  'saacs.auth.v0.UserGlobalRoles': KeySchema.fromJson({
    "itemType": "saacs.auth.v0.UserGlobalRoles",
    "itemKind": 2,
    "properties": "mspId,userId"
}),
  'saacs.biochain.v0.Specimen': KeySchema.fromJson({
    "itemType": "saacs.biochain.v0.Specimen",
    "itemKind": 2,
    "properties": "specimenId"
}),
  'saacs.sample.v0.Book': KeySchema.fromJson({
    "itemType": "saacs.sample.v0.Book",
    "itemKind": 2,
    "properties": "isbn"
}),
  'saacs.sample.v0.SimpleItem': KeySchema.fromJson({
    "itemType": "saacs.sample.v0.SimpleItem",
    "itemKind": 2,
    "properties": "id"
}),
  'saacs.sample.v0.Group': KeySchema.fromJson({
    "itemType": "saacs.sample.v0.Group",
    "itemKind": 2,
    "properties": "groupId"
}),
  'saacs.sample.v0.ItemWithNestedKey': KeySchema.fromJson({
    "itemType": "saacs.sample.v0.ItemWithNestedKey",
    "itemKind": 2,
    "properties": "id,nested.part1,nested.part2"
}),
}

export type SecondaryItemTypeMessage =
  | HiddenTxList
  | Suggestion

export const SecondaryItemKeySchema : Record<string, KeySchema> = {
  'saacs.common.v0.HiddenTxList': KeySchema.fromJson({
    "itemType": "saacs.common.v0.HiddenTxList",
    "itemKind": 3,
    "properties": "mspId"
}),
  'saacs.common.v0.Suggestion': KeySchema.fromJson({
    "itemType": "saacs.common.v0.Suggestion",
    "itemKind": 3,
    "properties": "suggestionId"
}),
}

export function PrimaryToKeySchema (item: PrimaryItemTypeMessage){
  switch (true) {
// msp_id, oid, value
// item.mspId, item.oid, item.value
    case isMessage (item, Attribute):
      return new ItemKey({
        itemType: 'saacs.auth.v0.Attribute',
        itemKind: 2,
        collectionId: item?.collectionId,
        itemKeyParts: [ item.mspId?? '', item.oid?? '', item.value ?? '' ]
      })
// collection_id
// item.collectionId
    case isMessage (item, Collection):
      return new ItemKey({
        itemType: 'saacs.auth.v0.Collection',
        itemKind: 2,
        collectionId: item?.collectionId,
        itemKeyParts: [ item.collectionId ?? '' ]
      })
// msp_id, user_id
// item.mspId, item.userId
    case isMessage (item, UserDirectMembership):
      return new ItemKey({
        itemType: 'saacs.auth.v0.UserDirectMembership',
        itemKind: 2,
        collectionId: item?.collectionId,
        itemKeyParts: [ item.mspId?? '', item.userId ?? '' ]
      })
// role_id
// item.roleId
    case isMessage (item, Role):
      return new ItemKey({
        itemType: 'saacs.auth.v0.Role',
        itemKind: 2,
        collectionId: item?.collectionId,
        itemKeyParts: [ item.roleId ?? '' ]
      })
// msp_id, user_id
// item.mspId, item.userId
    case isMessage (item, UserCollectionRoles):
      return new ItemKey({
        itemType: 'saacs.auth.v0.UserCollectionRoles',
        itemKind: 2,
        collectionId: item?.collectionId,
        itemKeyParts: [ item.mspId?? '', item.userId ?? '' ]
      })
// msp_id, user_id
// item.mspId, item.userId
    case isMessage (item, UserGlobalRoles):
      return new ItemKey({
        itemType: 'saacs.auth.v0.UserGlobalRoles',
        itemKind: 2,
        collectionId: item?.collectionId,
        itemKeyParts: [ item.mspId?? '', item.userId ?? '' ]
      })
// specimen_id
// item.specimenId
    case isMessage (item, Specimen):
      return new ItemKey({
        itemType: 'saacs.biochain.v0.Specimen',
        itemKind: 2,
        collectionId: item?.collectionId,
        itemKeyParts: [ item.specimenId ?? '' ]
      })
// isbn
// item.isbn
    case isMessage (item, Book):
      return new ItemKey({
        itemType: 'saacs.sample.v0.Book',
        itemKind: 2,
        collectionId: item?.collectionId,
        itemKeyParts: [ item.isbn ?? '' ]
      })
// id
// item.id
    case isMessage (item, SimpleItem):
      return new ItemKey({
        itemType: 'saacs.sample.v0.SimpleItem',
        itemKind: 2,
        collectionId: item?.collectionId,
        itemKeyParts: [ item.id ?? '' ]
      })
// group_id
// item.groupId
    case isMessage (item, Group):
      return new ItemKey({
        itemType: 'saacs.sample.v0.Group',
        itemKind: 2,
        collectionId: item?.collectionId,
        itemKeyParts: [ item.groupId ?? '' ]
      })
// id, nested.part1, nested.part2
// item.id, item.nested?.part1, item.nested?.part2
    case isMessage (item, ItemWithNestedKey):
      return new ItemKey({
        itemType: 'saacs.sample.v0.ItemWithNestedKey',
        itemKind: 2,
        collectionId: item?.collectionId,
        itemKeyParts: [ item.id?? '', item.nested?.part1?? '', item.nested?.part2 ?? '' ]
      })
    default:
      throw new Error('Unknown item type')
  }
}
