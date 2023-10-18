// @generated by protoc-gen-es v1.3.1 with parameter "target=ts"
// @generated from file rbac/service.proto (package rbac, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import type { BinaryReadOptions, FieldList, JsonReadOptions, JsonValue, PartialMessage, PlainMessage } from "@bufbuild/protobuf";
import { FieldMask, Message, proto3 } from "@bufbuild/protobuf";
import { Collection, Collection_Id, User, User_Id } from "./rbac_pb.js";

/**
 * @generated from message rbac.GetUserRequest
 */
export class GetUserRequest extends Message<GetUserRequest> {
  /**
   * @generated from field: rbac.User.Id id = 1;
   */
  id?: User_Id;

  constructor(data?: PartialMessage<GetUserRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "rbac.GetUserRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "id", kind: "message", T: User_Id },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): GetUserRequest {
    return new GetUserRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): GetUserRequest {
    return new GetUserRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): GetUserRequest {
    return new GetUserRequest().fromJsonString(jsonString, options);
  }

  static equals(a: GetUserRequest | PlainMessage<GetUserRequest> | undefined, b: GetUserRequest | PlainMessage<GetUserRequest> | undefined): boolean {
    return proto3.util.equals(GetUserRequest, a, b);
  }
}

/**
 * @generated from message rbac.UserRegisterRequest
 */
export class UserRegisterRequest extends Message<UserRegisterRequest> {
  /**
   * @generated from field: rbac.User user = 1;
   */
  user?: User;

  constructor(data?: PartialMessage<UserRegisterRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "rbac.UserRegisterRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "user", kind: "message", T: User },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): UserRegisterRequest {
    return new UserRegisterRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): UserRegisterRequest {
    return new UserRegisterRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): UserRegisterRequest {
    return new UserRegisterRequest().fromJsonString(jsonString, options);
  }

  static equals(a: UserRegisterRequest | PlainMessage<UserRegisterRequest> | undefined, b: UserRegisterRequest | PlainMessage<UserRegisterRequest> | undefined): boolean {
    return proto3.util.equals(UserRegisterRequest, a, b);
  }
}

/**
 * @generated from message rbac.UpdateMembershipRequest
 */
export class UpdateMembershipRequest extends Message<UpdateMembershipRequest> {
  /**
   * @generated from field: rbac.User.Id id = 1;
   */
  id?: User_Id;

  /**
   * @generated from field: rbac.Collection.Id collection_id = 2;
   */
  collectionId?: Collection_Id;

  /**
   * @generated from field: int32 role = 3;
   */
  role = 0;

  constructor(data?: PartialMessage<UpdateMembershipRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "rbac.UpdateMembershipRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "id", kind: "message", T: User_Id },
    { no: 2, name: "collection_id", kind: "message", T: Collection_Id },
    { no: 3, name: "role", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): UpdateMembershipRequest {
    return new UpdateMembershipRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): UpdateMembershipRequest {
    return new UpdateMembershipRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): UpdateMembershipRequest {
    return new UpdateMembershipRequest().fromJsonString(jsonString, options);
  }

  static equals(a: UpdateMembershipRequest | PlainMessage<UpdateMembershipRequest> | undefined, b: UpdateMembershipRequest | PlainMessage<UpdateMembershipRequest> | undefined): boolean {
    return proto3.util.equals(UpdateMembershipRequest, a, b);
  }
}

/**
 * @generated from message rbac.GetCollectionRequest
 */
export class GetCollectionRequest extends Message<GetCollectionRequest> {
  /**
   * @generated from field: rbac.Collection.Id id = 1;
   */
  id?: Collection_Id;

  constructor(data?: PartialMessage<GetCollectionRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "rbac.GetCollectionRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "id", kind: "message", T: Collection_Id },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): GetCollectionRequest {
    return new GetCollectionRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): GetCollectionRequest {
    return new GetCollectionRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): GetCollectionRequest {
    return new GetCollectionRequest().fromJsonString(jsonString, options);
  }

  static equals(a: GetCollectionRequest | PlainMessage<GetCollectionRequest> | undefined, b: GetCollectionRequest | PlainMessage<GetCollectionRequest> | undefined): boolean {
    return proto3.util.equals(GetCollectionRequest, a, b);
  }
}

/**
 * @generated from message rbac.CollectionCreateRequest
 */
export class CollectionCreateRequest extends Message<CollectionCreateRequest> {
  /**
   * @generated from field: rbac.Collection collection = 1;
   */
  collection?: Collection;

  constructor(data?: PartialMessage<CollectionCreateRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "rbac.CollectionCreateRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "collection", kind: "message", T: Collection },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): CollectionCreateRequest {
    return new CollectionCreateRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): CollectionCreateRequest {
    return new CollectionCreateRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): CollectionCreateRequest {
    return new CollectionCreateRequest().fromJsonString(jsonString, options);
  }

  static equals(a: CollectionCreateRequest | PlainMessage<CollectionCreateRequest> | undefined, b: CollectionCreateRequest | PlainMessage<CollectionCreateRequest> | undefined): boolean {
    return proto3.util.equals(CollectionCreateRequest, a, b);
  }
}

/**
 * @generated from message rbac.CollectionUpdateRequest
 */
export class CollectionUpdateRequest extends Message<CollectionUpdateRequest> {
  /**
   * @generated from field: rbac.Collection collection = 1;
   */
  collection?: Collection;

  /**
   * @generated from field: google.protobuf.FieldMask update_mask = 2;
   */
  updateMask?: FieldMask;

  constructor(data?: PartialMessage<CollectionUpdateRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "rbac.CollectionUpdateRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "collection", kind: "message", T: Collection },
    { no: 2, name: "update_mask", kind: "message", T: FieldMask },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): CollectionUpdateRequest {
    return new CollectionUpdateRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): CollectionUpdateRequest {
    return new CollectionUpdateRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): CollectionUpdateRequest {
    return new CollectionUpdateRequest().fromJsonString(jsonString, options);
  }

  static equals(a: CollectionUpdateRequest | PlainMessage<CollectionUpdateRequest> | undefined, b: CollectionUpdateRequest | PlainMessage<CollectionUpdateRequest> | undefined): boolean {
    return proto3.util.equals(CollectionUpdateRequest, a, b);
  }
}

