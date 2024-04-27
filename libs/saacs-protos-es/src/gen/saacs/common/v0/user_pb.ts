// @generated by protoc-gen-es v1.9.0 with parameter "target=ts"
// @generated from file saacs/common/v0/user.proto (package saacs.common.v0, syntax proto3)
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

/**
 * @generated from message saacs.common.v0.User
 */
export class User extends Message<User> {
  /**
   * @generated from field: string msp_id = 1;
   */
  mspId = ''

  /**
   * @generated from field: string user_id = 2;
   */
  userId = ''

  constructor(data?: PartialMessage<User>) {
    super()
    proto3.util.initPartial(data, this)
  }

  static readonly runtime: typeof proto3 = proto3
  static readonly typeName = 'saacs.common.v0.User'
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: 'msp_id', kind: 'scalar', T: 9 /* ScalarType.STRING */ },
    { no: 2, name: 'user_id', kind: 'scalar', T: 9 /* ScalarType.STRING */ },
  ])

  static fromBinary(
    bytes: Uint8Array,
    options?: Partial<BinaryReadOptions>,
  ): User {
    return new User().fromBinary(bytes, options)
  }

  static fromJson(
    jsonValue: JsonValue,
    options?: Partial<JsonReadOptions>,
  ): User {
    return new User().fromJson(jsonValue, options)
  }

  static fromJsonString(
    jsonString: string,
    options?: Partial<JsonReadOptions>,
  ): User {
    return new User().fromJsonString(jsonString, options)
  }

  static equals(
    a: User | PlainMessage<User> | undefined,
    b: User | PlainMessage<User> | undefined,
  ): boolean {
    return proto3.util.equals(User, a, b)
  }
}
