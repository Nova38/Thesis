// @generated by protoc-gen-connect-es v0.13.0
// @generated from file chaincode/auth/rbac/schema/v1/auth_service.proto (package rbac.schema.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import { Empty, MethodKind } from "@bufbuild/protobuf";
import { CollectionCreateRequest, CollectionCreateResponse, CollectionGetHistoryRequest, CollectionGetHistoryResponse, CollectionGetListResponse, CollectionGetRequest, CollectionGetResponse, CollectionUpdatePermissionRequest, CollectionUpdatePermissionResponse, CollectionUpdateRolesRequest, CollectionUpdateRolesResponse, UserGetCurrentIdResponse, UserGetCurrentResponse, UserGetHistoryRequest, UserGetHistoryResponse, UserGetListResponse, UserGetRequest, UserGetResponse, UserRegisterRequest, UserRegisterResponse, UserUpdateMembershipRequest, UserUpdateMembershipResponse } from "./auth_service_pb.js";

/**
 * buf:lint:ignore RPC_NO_DELETE
 *
 * @generated from service rbac.schema.v1.AuthService
 */
export const AuthService = {
  typeName: "rbac.schema.v1.AuthService",
  methods: {
    /**
     *
     * UserGetCurrent: Returns the current user.
     *
     * Returns the current user.
     * # Requires:
     *  - User submitting the transaction is a registered user.
     *
     * @generated from rpc rbac.schema.v1.AuthService.UserGetCurrent
     */
    userGetCurrent: {
      name: "UserGetCurrent",
      I: Empty,
      O: UserGetCurrentResponse,
      kind: MethodKind.Unary,
    },
    /**
     *
     * Returns the current user id.
     *
     * # Requires:
     *  - User submitting the transaction is a registered user.
     *
     * @generated from rpc rbac.schema.v1.AuthService.UserGetCurrentId
     */
    userGetCurrentId: {
      name: "UserGetCurrentId",
      I: Empty,
      O: UserGetCurrentIdResponse,
      kind: MethodKind.Unary,
    },
    /**
     *
     * UserGetList: Returns the list of users.
     *
     * # Requires:
     *  - Non-register users can call this method.
     *
     * @generated from rpc rbac.schema.v1.AuthService.UserGetList
     */
    userGetList: {
      name: "UserGetList",
      I: Empty,
      O: UserGetListResponse,
      kind: MethodKind.Unary,
    },
    /**
     *
     * UserGet: Returns the user.
     *
     * # Requires:
     *  - Non-register users can call this method.
     *
     * @generated from rpc rbac.schema.v1.AuthService.UserGet
     */
    userGet: {
      name: "UserGet",
      I: UserGetRequest,
      O: UserGetResponse,
      kind: MethodKind.Unary,
    },
    /**
     *
     * UserGetHistory: Returns the user history.
     *
     * # Requires:
     *  - Non-register users can call this method.
     *
     * @generated from rpc rbac.schema.v1.AuthService.UserGetHistory
     */
    userGetHistory: {
      name: "UserGetHistory",
      I: UserGetHistoryRequest,
      O: UserGetHistoryResponse,
      kind: MethodKind.Unary,
    },
    /**
     *
     * UserRegister: Registers the user.
     *
     * # Requires:
     *  - The certificate for the user submitting this request must not be already registered as a user.
     *
     * @generated from rpc rbac.schema.v1.AuthService.UserRegister
     */
    userRegister: {
      name: "UserRegister",
      I: UserRegisterRequest,
      O: UserRegisterResponse,
      kind: MethodKind.Unary,
    },
    /**
     *
     * UserUpdateMembership: Updates the user's membership.
     *
     * # Requires:
     *  - User submitting the transaction is a registered user.
     *  - The specified user id is a registered user.
     *  - The specified collection id is a registered collection.
     *  - The user submitting the transaction is a member of the specified collection.
     *  - The user submitting the transaction the a role who has permission
     *     to update the membership of the specified collection.
     *
     * @generated from rpc rbac.schema.v1.AuthService.UserUpdateMembership
     */
    userUpdateMembership: {
      name: "UserUpdateMembership",
      I: UserUpdateMembershipRequest,
      O: UserUpdateMembershipResponse,
      kind: MethodKind.Unary,
    },
    /**
     *
     * CollectionGetList: Returns the list of collections.
     *
     * # Requires:
     *  - Non-register users can call this method.
     *
     * @generated from rpc rbac.schema.v1.AuthService.CollectionGetList
     */
    collectionGetList: {
      name: "CollectionGetList",
      I: Empty,
      O: CollectionGetListResponse,
      kind: MethodKind.Unary,
    },
    /**
     *
     * CollectionGet: Returns the collection.
     *
     * # Requires:
     *  - Non-register users can call this method.
     *
     * @generated from rpc rbac.schema.v1.AuthService.CollectionGet
     */
    collectionGet: {
      name: "CollectionGet",
      I: CollectionGetRequest,
      O: CollectionGetResponse,
      kind: MethodKind.Unary,
    },
    /**
     *
     * CollectionGetHistory: Returns the collection history.
     *
     * # Requires:
     *  - Non-register users can call this method.
     *
     * @generated from rpc rbac.schema.v1.AuthService.CollectionGetHistory
     */
    collectionGetHistory: {
      name: "CollectionGetHistory",
      I: CollectionGetHistoryRequest,
      O: CollectionGetHistoryResponse,
      kind: MethodKind.Unary,
    },
    /**
     *
     * CollectionCreate: Creates the collection.
     *
     * # Requires:
     *  - User submitting the transaction is a registered user.
     *
     * @generated from rpc rbac.schema.v1.AuthService.CollectionCreate
     */
    collectionCreate: {
      name: "CollectionCreate",
      I: CollectionCreateRequest,
      O: CollectionCreateResponse,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc rbac.schema.v1.AuthService.CollectionUpdateRoles
     */
    collectionUpdateRoles: {
      name: "CollectionUpdateRoles",
      I: CollectionUpdateRolesRequest,
      O: CollectionUpdateRolesResponse,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc rbac.schema.v1.AuthService.CollectionUpdatePermission
     */
    collectionUpdatePermission: {
      name: "CollectionUpdatePermission",
      I: CollectionUpdatePermissionRequest,
      O: CollectionUpdatePermissionResponse,
      kind: MethodKind.Unary,
    },
  }
};
