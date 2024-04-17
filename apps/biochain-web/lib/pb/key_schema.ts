import { KeySchema } from './auth/v1/auth_pb.js'
import {
  Attribute,
  Role,
  UserCollectionRoles,
  UserDirectMembership,
  UserEmbeddedRoles,
} from './auth/v1/models_pb.js'
import {
  Collection,
  HiddenTxList,
  ReferenceKey,
  Suggestion,
} from './auth/v1/objects_pb.js'
import { Specimen } from './biochain/v1/state_pb.js'
import { Book, Group, SimpleItem } from './sample/v0/items_pb.js'

export const ItemKeySchema: Record<string, KeySchema> = {
  UserDirectMembership: KeySchema.fromJson({
    itemType: 'auth.UserDirectMembership',
    itemKind: 2,
    properties: 'mspId,userId',
  }),
  Role: KeySchema.fromJson({
    itemType: 'auth.Role',
    itemKind: 2,
    properties: 'roleId',
  }),
  UserCollectionRoles: KeySchema.fromJson({
    itemType: 'auth.UserCollectionRoles',
    itemKind: 2,
    properties: 'mspId,userId',
  }),
  UserEmbeddedRoles: KeySchema.fromJson({
    itemType: 'auth.UserEmbeddedRoles',
    itemKind: 2,
    properties: 'mspId,userId',
  }),
  Attribute: KeySchema.fromJson({
    itemType: 'auth.Attribute',
    itemKind: 2,
    properties: 'mspId,oidvalue',
  }),
  ReferenceKey: KeySchema.fromJson({
    itemType: 'auth.ReferenceKey',
    itemKind: 4,
  }),
  Collection: KeySchema.fromJson({
    itemType: 'auth.Collection',
    itemKind: 2,
    properties: 'collectionId',
  }),
  Suggestion: KeySchema.fromJson({
    itemType: 'auth.Suggestion',
    itemKind: 3,
    properties: 'suggestionId',
  }),
  HiddenTxList: KeySchema.fromJson({
    itemType: 'auth.HiddenTxList',
    itemKind: 3,
    properties: '',
  }),
  Specimen: KeySchema.fromJson({
    itemType: 'biochain.v1.Specimen',
    itemKind: 2,
    properties: 'specimenId',
  }),
  SimpleItem: KeySchema.fromJson({
    itemType: 'sample.SimpleItem',
    itemKind: 2,
    properties: 'id',
  }),
  Group: KeySchema.fromJson({
    itemType: 'sample.Group',
    itemKind: 2,
    properties: 'groupId',
  }),
  Book: KeySchema.fromJson({
    itemType: 'sample.Book',
    itemKind: 2,
    properties: 'isbn',
  }),
}
