// @generated by protoc-gen-connect-es v0.13.0
// @generated from file chaincode/auth/common/generic.proto (package auth.common, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import { Empty, MethodKind } from "@bufbuild/protobuf";
import { AuthorizeOperationRequest, AuthorizeOperationResponse, BootstrapRequest, BootstrapResponse, CreateRequest, CreateResponse, CreateUserResponse, DeleteRequest, DeleteResponse, GetCurrentUserResponse, GetRequest, GetResponse, HiddenTxRequest, HiddenTxResponse, HideTxRequest, HideTxResponse, HistoryRequest, HistoryResponse, ListByAttrsRequest, ListByAttrsResponse, ListByCollectionRequest, ListByCollectionResponse, ListRequest, ListResponse, ReferenceByCollectionRequest, ReferenceByCollectionResponse, ReferenceByObjectRequest, ReferenceByObjectResponse, ReferenceCreateRequest, ReferenceCreateResponse, ReferenceDeleteRequest, ReferenceDeleteResponse, ReferenceListByTypeRequest, ReferenceListByTypeResponse, ReferenceRequest, ReferenceResponse, SuggestionApproveRequest, SuggestionApproveResponse, SuggestionByPartialKeyRequest, SuggestionByPartialKeyResponse, SuggestionCreateRequest, SuggestionCreateResponse, SuggestionDeleteRequest, SuggestionDeleteResponse, SuggestionListByCollectionRequest, SuggestionListByCollectionResponse, SuggestionListRequest, SuggestionListResponse, SuggestionRequest, SuggestionResponse, UnHideTxRequest, UnHideTxResponse, UpdateRequest, UpdateResponse } from "./generic_pb.js";

/**
 * @generated from service auth.common.GenericService
 */
export declare const GenericService: {
  readonly typeName: "auth.common.GenericService",
  readonly methods: {
    /**
     * ══════════════════════════════════ Helper ═════════════════════════════════════
     * ────────────────────────────────── Query ──────────────────────────────────────
     * rpc GetAllTypes(google.protobuf.Empty) returns (GetAllTypesResponse) {
     *   option (auth.transaction_type) = TRANSACTION_TYPE_QUERY;
     *   option (auth.operation) = {action: ACTION_UTILITY};
     * }
     *
     * @generated from rpc auth.common.GenericService.GetCurrentUser
     */
    readonly getCurrentUser: {
      readonly name: "GetCurrentUser",
      readonly I: typeof Empty,
      readonly O: typeof GetCurrentUserResponse,
      readonly kind: MethodKind.Unary,
    },
    /**
     * ──────────────────────────────── Invoke ───────────────────────────────────────
     *
     * @generated from rpc auth.common.GenericService.Bootstrap
     */
    readonly bootstrap: {
      readonly name: "Bootstrap",
      readonly I: typeof BootstrapRequest,
      readonly O: typeof BootstrapResponse,
      readonly kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc auth.common.GenericService.AuthorizeOperation
     */
    readonly authorizeOperation: {
      readonly name: "AuthorizeOperation",
      readonly I: typeof AuthorizeOperationRequest,
      readonly O: typeof AuthorizeOperationResponse,
      readonly kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc auth.common.GenericService.CreateUser
     */
    readonly createUser: {
      readonly name: "CreateUser",
      readonly I: typeof Empty,
      readonly O: typeof CreateUserResponse,
      readonly kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc auth.common.GenericService.Get
     */
    readonly get: {
      readonly name: "Get",
      readonly I: typeof GetRequest,
      readonly O: typeof GetResponse,
      readonly kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc auth.common.GenericService.List
     */
    readonly list: {
      readonly name: "List",
      readonly I: typeof ListRequest,
      readonly O: typeof ListResponse,
      readonly kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc auth.common.GenericService.ListByCollection
     */
    readonly listByCollection: {
      readonly name: "ListByCollection",
      readonly I: typeof ListByCollectionRequest,
      readonly O: typeof ListByCollectionResponse,
      readonly kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc auth.common.GenericService.ListByAttrs
     */
    readonly listByAttrs: {
      readonly name: "ListByAttrs",
      readonly I: typeof ListByAttrsRequest,
      readonly O: typeof ListByAttrsResponse,
      readonly kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc auth.common.GenericService.Create
     */
    readonly create: {
      readonly name: "Create",
      readonly I: typeof CreateRequest,
      readonly O: typeof CreateResponse,
      readonly kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc auth.common.GenericService.Update
     */
    readonly update: {
      readonly name: "Update",
      readonly I: typeof UpdateRequest,
      readonly O: typeof UpdateResponse,
      readonly kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc auth.common.GenericService.Delete
     */
    readonly delete: {
      readonly name: "Delete",
      readonly I: typeof DeleteRequest,
      readonly O: typeof DeleteResponse,
      readonly kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc auth.common.GenericService.History
     */
    readonly history: {
      readonly name: "History",
      readonly I: typeof HistoryRequest,
      readonly O: typeof HistoryResponse,
      readonly kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc auth.common.GenericService.HiddenTx
     */
    readonly hiddenTx: {
      readonly name: "HiddenTx",
      readonly I: typeof HiddenTxRequest,
      readonly O: typeof HiddenTxResponse,
      readonly kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc auth.common.GenericService.HideTx
     */
    readonly hideTx: {
      readonly name: "HideTx",
      readonly I: typeof HideTxRequest,
      readonly O: typeof HideTxResponse,
      readonly kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc auth.common.GenericService.UnHideTx
     */
    readonly unHideTx: {
      readonly name: "UnHideTx",
      readonly I: typeof UnHideTxRequest,
      readonly O: typeof UnHideTxResponse,
      readonly kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc auth.common.GenericService.Reference
     */
    readonly reference: {
      readonly name: "Reference",
      readonly I: typeof ReferenceRequest,
      readonly O: typeof ReferenceResponse,
      readonly kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc auth.common.GenericService.ReferenceListByType
     */
    readonly referenceListByType: {
      readonly name: "ReferenceListByType",
      readonly I: typeof ReferenceListByTypeRequest,
      readonly O: typeof ReferenceListByTypeResponse,
      readonly kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc auth.common.GenericService.ReferenceByCollection
     */
    readonly referenceByCollection: {
      readonly name: "ReferenceByCollection",
      readonly I: typeof ReferenceByCollectionRequest,
      readonly O: typeof ReferenceByCollectionResponse,
      readonly kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc auth.common.GenericService.ReferenceByObject
     */
    readonly referenceByObject: {
      readonly name: "ReferenceByObject",
      readonly I: typeof ReferenceByObjectRequest,
      readonly O: typeof ReferenceByObjectResponse,
      readonly kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc auth.common.GenericService.ReferenceCreate
     */
    readonly referenceCreate: {
      readonly name: "ReferenceCreate",
      readonly I: typeof ReferenceCreateRequest,
      readonly O: typeof ReferenceCreateResponse,
      readonly kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc auth.common.GenericService.ReferenceDelete
     */
    readonly referenceDelete: {
      readonly name: "ReferenceDelete",
      readonly I: typeof ReferenceDeleteRequest,
      readonly O: typeof ReferenceDeleteResponse,
      readonly kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc auth.common.GenericService.Suggestion
     */
    readonly suggestion: {
      readonly name: "Suggestion",
      readonly I: typeof SuggestionRequest,
      readonly O: typeof SuggestionResponse,
      readonly kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc auth.common.GenericService.SuggestionList
     */
    readonly suggestionList: {
      readonly name: "SuggestionList",
      readonly I: typeof SuggestionListRequest,
      readonly O: typeof SuggestionListResponse,
      readonly kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc auth.common.GenericService.SuggestionListByCollection
     */
    readonly suggestionListByCollection: {
      readonly name: "SuggestionListByCollection",
      readonly I: typeof SuggestionListByCollectionRequest,
      readonly O: typeof SuggestionListByCollectionResponse,
      readonly kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc auth.common.GenericService.SuggestionByPartialKey
     */
    readonly suggestionByPartialKey: {
      readonly name: "SuggestionByPartialKey",
      readonly I: typeof SuggestionByPartialKeyRequest,
      readonly O: typeof SuggestionByPartialKeyResponse,
      readonly kind: MethodKind.Unary,
    },
    /**
     * ──────────────────────────────── Invoke ───────────────────────────────────────
     *
     * @generated from rpc auth.common.GenericService.SuggestionCreate
     */
    readonly suggestionCreate: {
      readonly name: "SuggestionCreate",
      readonly I: typeof SuggestionCreateRequest,
      readonly O: typeof SuggestionCreateResponse,
      readonly kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc auth.common.GenericService.SuggestionDelete
     */
    readonly suggestionDelete: {
      readonly name: "SuggestionDelete",
      readonly I: typeof SuggestionDeleteRequest,
      readonly O: typeof SuggestionDeleteResponse,
      readonly kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc auth.common.GenericService.SuggestionApprove
     */
    readonly suggestionApprove: {
      readonly name: "SuggestionApprove",
      readonly I: typeof SuggestionApproveRequest,
      readonly O: typeof SuggestionApproveResponse,
      readonly kind: MethodKind.Unary,
    },
  }
};
