// @generated by protoc-gen-es v1.9.0 with parameter "target=ts"
// @generated from file saacs/auth/v0/roles.proto (package saacs.auth.v0, syntax proto3)
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
 * Shared Auth Object for Role Based Authentication
 *
 * @generated from message saacs.auth.v0.Role
 */
export class Role extends Message<Role> {
  /**
   * @generated from field: string collection_id = 1;
   */
  collectionId = ''

  /**
   * @generated from field: string role_id = 2;
   */
  roleId = ''

  /**
   * @generated from field: saacs.auth.v0.Polices polices = 4;
   */
  polices?: Polices

  /**
   * @generated from field: string note = 5;
   */
  note = ''

  /**
   * @generated from field: repeated string parent_role_ids = 6;
   */
  parentRoleIds: string[] = []

  constructor(data?: PartialMessage<Role>) {
    super()
    proto3.util.initPartial(data, this)
  }

  static readonly runtime: typeof proto3 = proto3
  static readonly typeName = 'saacs.auth.v0.Role'
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    {
      no: 1,
      name: 'collection_id',
      kind: 'scalar',
      T: 9 /* ScalarType.STRING */,
    },
    { no: 2, name: 'role_id', kind: 'scalar', T: 9 /* ScalarType.STRING */ },
    { no: 4, name: 'polices', kind: 'message', T: Polices },
    { no: 5, name: 'note', kind: 'scalar', T: 9 /* ScalarType.STRING */ },
    {
      no: 6,
      name: 'parent_role_ids',
      kind: 'scalar',
      T: 9 /* ScalarType.STRING */,
      repeated: true,
    },
  ])

  static fromBinary(
    bytes: Uint8Array,
    options?: Partial<BinaryReadOptions>,
  ): Role {
    return new Role().fromBinary(bytes, options)
  }

  static fromJson(
    jsonValue: JsonValue,
    options?: Partial<JsonReadOptions>,
  ): Role {
    return new Role().fromJson(jsonValue, options)
  }

  static fromJsonString(
    jsonString: string,
    options?: Partial<JsonReadOptions>,
  ): Role {
    return new Role().fromJsonString(jsonString, options)
  }

  static equals(
    a: Role | PlainMessage<Role> | undefined,
    b: Role | PlainMessage<Role> | undefined,
  ): boolean {
    return proto3.util.equals(Role, a, b)
  }
}

/**
 * @generated from message saacs.auth.v0.RoleIDList
 */
export class RoleIDList extends Message<RoleIDList> {
  /**
   * @generated from field: repeated string role_id = 1;
   */
  roleId: string[] = []

  constructor(data?: PartialMessage<RoleIDList>) {
    super()
    proto3.util.initPartial(data, this)
  }

  static readonly runtime: typeof proto3 = proto3
  static readonly typeName = 'saacs.auth.v0.RoleIDList'
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    {
      no: 1,
      name: 'role_id',
      kind: 'scalar',
      T: 9 /* ScalarType.STRING */,
      repeated: true,
    },
  ])

  static fromBinary(
    bytes: Uint8Array,
    options?: Partial<BinaryReadOptions>,
  ): RoleIDList {
    return new RoleIDList().fromBinary(bytes, options)
  }

  static fromJson(
    jsonValue: JsonValue,
    options?: Partial<JsonReadOptions>,
  ): RoleIDList {
    return new RoleIDList().fromJson(jsonValue, options)
  }

  static fromJsonString(
    jsonString: string,
    options?: Partial<JsonReadOptions>,
  ): RoleIDList {
    return new RoleIDList().fromJsonString(jsonString, options)
  }

  static equals(
    a: RoleIDList | PlainMessage<RoleIDList> | undefined,
    b: RoleIDList | PlainMessage<RoleIDList> | undefined,
  ): boolean {
    return proto3.util.equals(RoleIDList, a, b)
  }
}

/**
 * Auth Object For RBAC
 *
 * @generated from message saacs.auth.v0.UserCollectionRoles
 */
export class UserCollectionRoles extends Message<UserCollectionRoles> {
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
   * The roles that the user has in the collection
   *
   * @generated from field: repeated string role_ids = 4;
   */
  roleIds: string[] = []

  /**
   * @generated from field: string note = 6;
   */
  note = ''

  constructor(data?: PartialMessage<UserCollectionRoles>) {
    super()
    proto3.util.initPartial(data, this)
  }

  static readonly runtime: typeof proto3 = proto3
  static readonly typeName = 'saacs.auth.v0.UserCollectionRoles'
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    {
      no: 1,
      name: 'collection_id',
      kind: 'scalar',
      T: 9 /* ScalarType.STRING */,
    },
    { no: 2, name: 'msp_id', kind: 'scalar', T: 9 /* ScalarType.STRING */ },
    { no: 3, name: 'user_id', kind: 'scalar', T: 9 /* ScalarType.STRING */ },
    {
      no: 4,
      name: 'role_ids',
      kind: 'scalar',
      T: 9 /* ScalarType.STRING */,
      repeated: true,
    },
    { no: 6, name: 'note', kind: 'scalar', T: 9 /* ScalarType.STRING */ },
  ])

  static fromBinary(
    bytes: Uint8Array,
    options?: Partial<BinaryReadOptions>,
  ): UserCollectionRoles {
    return new UserCollectionRoles().fromBinary(bytes, options)
  }

  static fromJson(
    jsonValue: JsonValue,
    options?: Partial<JsonReadOptions>,
  ): UserCollectionRoles {
    return new UserCollectionRoles().fromJson(jsonValue, options)
  }

  static fromJsonString(
    jsonString: string,
    options?: Partial<JsonReadOptions>,
  ): UserCollectionRoles {
    return new UserCollectionRoles().fromJsonString(jsonString, options)
  }

  static equals(
    a: UserCollectionRoles | PlainMessage<UserCollectionRoles> | undefined,
    b: UserCollectionRoles | PlainMessage<UserCollectionRoles> | undefined,
  ): boolean {
    return proto3.util.equals(UserCollectionRoles, a, b)
  }
}

/**
 * Auth Object For Embedded RBAC
 *
 * @generated from message saacs.auth.v0.UserGlobalRoles
 */
export class UserGlobalRoles extends Message<UserGlobalRoles> {
  /**
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
   * The roles that the user has in the collection
   * key is the collection id
   * value is the list of rolesIds
   *
   * @generated from field: map<string, saacs.auth.v0.RoleIDList> roles = 4;
   */
  roles: { [key: string]: RoleIDList } = {}

  constructor(data?: PartialMessage<UserGlobalRoles>) {
    super()
    proto3.util.initPartial(data, this)
  }

  static readonly runtime: typeof proto3 = proto3
  static readonly typeName = 'saacs.auth.v0.UserGlobalRoles'
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    {
      no: 1,
      name: 'collection_id',
      kind: 'scalar',
      T: 9 /* ScalarType.STRING */,
    },
    { no: 2, name: 'msp_id', kind: 'scalar', T: 9 /* ScalarType.STRING */ },
    { no: 3, name: 'user_id', kind: 'scalar', T: 9 /* ScalarType.STRING */ },
    {
      no: 4,
      name: 'roles',
      kind: 'map',
      K: 9 /* ScalarType.STRING */,
      V: { kind: 'message', T: RoleIDList },
    },
  ])

  static fromBinary(
    bytes: Uint8Array,
    options?: Partial<BinaryReadOptions>,
  ): UserGlobalRoles {
    return new UserGlobalRoles().fromBinary(bytes, options)
  }

  static fromJson(
    jsonValue: JsonValue,
    options?: Partial<JsonReadOptions>,
  ): UserGlobalRoles {
    return new UserGlobalRoles().fromJson(jsonValue, options)
  }

  static fromJsonString(
    jsonString: string,
    options?: Partial<JsonReadOptions>,
  ): UserGlobalRoles {
    return new UserGlobalRoles().fromJsonString(jsonString, options)
  }

  static equals(
    a: UserGlobalRoles | PlainMessage<UserGlobalRoles> | undefined,
    b: UserGlobalRoles | PlainMessage<UserGlobalRoles> | undefined,
  ): boolean {
    return proto3.util.equals(UserGlobalRoles, a, b)
  }
}
