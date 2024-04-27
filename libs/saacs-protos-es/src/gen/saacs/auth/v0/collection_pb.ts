// @generated by protoc-gen-es v1.9.0 with parameter "target=ts"
// @generated from file saacs/auth/v0/collection.proto (package saacs.auth.v0, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import type {
  BinaryReadOptions,
  FieldList,
  JsonReadOptions,
  JsonValue,
  PartialMessage,
  PlainMessage,
} from '@bufbuild/protobuf'
import { Message, proto3 } from '@bufbuild/protobuf'
import { AuthType } from './type_pb.js'
import { Polices } from './policy_pb.js'

/**
 * Collection
 * ┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈
 * Note that the types of items are stored in the default ACLEntry
 *
 * key := {COLLECTION}{COLLECTION_ID}
 *
 * @generated from message saacs.auth.v0.Collection
 */
export class Collection extends Message<Collection> {
  /**
   * The key for the ledger
   *
   * @generated from field: string collection_id = 1;
   */
  collectionId = ''

  /**
   * @generated from field: string name = 2;
   */
  name = ''

  /**
   * @generated from field: saacs.auth.v0.AuthType auth_type = 3;
   */
  authType = AuthType.UNSPECIFIED

  /**
   * @generated from field: repeated string item_types = 4;
   */
  itemTypes: string[] = []

  /**
   * @generated from field: saacs.auth.v0.Polices default = 5;
   */
  default?: Polices

  /**
   * @generated from field: bool use_auth_parents = 6;
   */
  useAuthParents = false

  constructor(data?: PartialMessage<Collection>) {
    super()
    proto3.util.initPartial(data, this)
  }

  static readonly runtime: typeof proto3 = proto3
  static readonly typeName = 'saacs.auth.v0.Collection'
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    {
      no: 1,
      name: 'collection_id',
      kind: 'scalar',
      T: 9 /* ScalarType.STRING */,
    },
    { no: 2, name: 'name', kind: 'scalar', T: 9 /* ScalarType.STRING */ },
    { no: 3, name: 'auth_type', kind: 'enum', T: proto3.getEnumType(AuthType) },
    {
      no: 4,
      name: 'item_types',
      kind: 'scalar',
      T: 9 /* ScalarType.STRING */,
      repeated: true,
    },
    { no: 5, name: 'default', kind: 'message', T: Polices },
    {
      no: 6,
      name: 'use_auth_parents',
      kind: 'scalar',
      T: 8 /* ScalarType.BOOL */,
    },
  ])

  static fromBinary(
    bytes: Uint8Array,
    options?: Partial<BinaryReadOptions>,
  ): Collection {
    return new Collection().fromBinary(bytes, options)
  }

  static fromJson(
    jsonValue: JsonValue,
    options?: Partial<JsonReadOptions>,
  ): Collection {
    return new Collection().fromJson(jsonValue, options)
  }

  static fromJsonString(
    jsonString: string,
    options?: Partial<JsonReadOptions>,
  ): Collection {
    return new Collection().fromJsonString(jsonString, options)
  }

  static equals(
    a: Collection | PlainMessage<Collection> | undefined,
    b: Collection | PlainMessage<Collection> | undefined,
  ): boolean {
    return proto3.util.equals(Collection, a, b)
  }
}
