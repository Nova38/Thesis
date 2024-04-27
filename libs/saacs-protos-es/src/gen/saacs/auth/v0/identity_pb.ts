// @generated by protoc-gen-es v1.9.0 with parameter "target=ts"
// @generated from file saacs/auth/v0/identity.proto (package saacs.auth.v0, syntax proto3)
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
import { Polices } from './policy_pb.js'

/**
 * Identity Auth Object
 *
 * @generated from message saacs.auth.v0.UserDirectMembership
 */
export class UserDirectMembership extends Message<UserDirectMembership> {
  /**
   * The collection that the user is a member of
   *
   * @generated from field: string collection_id = 1;
   */
  collectionId = ''

  /**
   * The msp of the organization that the user's certificate is from
   *
   * @generated from field: string msp_id = 2;
   */
  mspId = ''

  /**
   * The id of the user from the certificate
   *
   * @generated from field: string user_id = 3;
   */
  userId = ''

  /**
   * The Permissions that the user will have
   *
   * @generated from field: saacs.auth.v0.Polices polices = 4;
   */
  polices?: Polices

  /**
   * @generated from field: string note = 6;
   */
  note = ''

  constructor(data?: PartialMessage<UserDirectMembership>) {
    super()
    proto3.util.initPartial(data, this)
  }

  static readonly runtime: typeof proto3 = proto3
  static readonly typeName = 'saacs.auth.v0.UserDirectMembership'
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    {
      no: 1,
      name: 'collection_id',
      kind: 'scalar',
      T: 9 /* ScalarType.STRING */,
    },
    { no: 2, name: 'msp_id', kind: 'scalar', T: 9 /* ScalarType.STRING */ },
    { no: 3, name: 'user_id', kind: 'scalar', T: 9 /* ScalarType.STRING */ },
    { no: 4, name: 'polices', kind: 'message', T: Polices },
    { no: 6, name: 'note', kind: 'scalar', T: 9 /* ScalarType.STRING */ },
  ])

  static fromBinary(
    bytes: Uint8Array,
    options?: Partial<BinaryReadOptions>,
  ): UserDirectMembership {
    return new UserDirectMembership().fromBinary(bytes, options)
  }

  static fromJson(
    jsonValue: JsonValue,
    options?: Partial<JsonReadOptions>,
  ): UserDirectMembership {
    return new UserDirectMembership().fromJson(jsonValue, options)
  }

  static fromJsonString(
    jsonString: string,
    options?: Partial<JsonReadOptions>,
  ): UserDirectMembership {
    return new UserDirectMembership().fromJsonString(jsonString, options)
  }

  static equals(
    a: UserDirectMembership | PlainMessage<UserDirectMembership> | undefined,
    b: UserDirectMembership | PlainMessage<UserDirectMembership> | undefined,
  ): boolean {
    return proto3.util.equals(UserDirectMembership, a, b)
  }
}
