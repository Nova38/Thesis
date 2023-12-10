import { Contract } from "@hyperledger/fabric-gateway";
import { Empty, IMessageTypeRegistry, JsonWriteStringOptions } from "@bufbuild/protobuf";
import { AuthorizeOperationRequest, AuthorizeOperationResponse, BootstrapRequest, BootstrapResponse, CreateRequest, CreateResponse, DeleteRequest, DeleteResponse, GetCurrentUserResponse, GetFullRequest, GetFullResponse, GetHiddenTxRequest, GetHiddenTxResponse, GetHistoryRequest, GetHistoryResponse, GetRequest, GetResponse, GetSuggestionRequest, GetSuggestionResponse, HideTxRequest, HideTxResponse, ListByAttrsRequest, ListByAttrsResponse, ListByCollectionRequest, ListByCollectionResponse, ListRequest, ListResponse, SuggestionApproveRequest, SuggestionApproveResponse, SuggestionByPartialKeyRequest, SuggestionByPartialKeyResponse, SuggestionCreateRequest, SuggestionCreateResponse, SuggestionDeleteRequest, SuggestionDeleteResponse, SuggestionListByCollectionRequest, SuggestionListByCollectionResponse, UnHideTxRequest, UnHideTxResponse, UpdateRequest, UpdateResponse } from "./generic_pb.js";

const utf8Decoder = new TextDecoder();
/**
 * @generated from service auth.common.GenericService
 */
export class GenericServiceClient {
    private contract: Contract;
    private jsonWriteOptions:Partial<JsonWriteStringOptions> = {};
    private registry: IMessageTypeRegistry;

    constructor(contract: Contract, registry: IMessageTypeRegistry) {
        this.contract = contract;
        this.registry = registry;
    }


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
    async getCurrentUser(request: Empty, evaluate: boolean ): Promise< GetCurrentUserResponse> {
        if (evaluate) {
            const results = utf8Decoder.decode(
                await this.contract.evaluateTransaction(
                "GetCurrentUser",
                request.toJsonString(this.jsonWriteOptions)
            )
            )
            return GetCurrentUserResponse.fromJsonString(results, {typeRegistry: this.registry});
        } else {
            const results = utf8Decoder.decode(
                    await this.contract.submitTransaction(
                "GetCurrentUser",
                request.toJsonString(this.jsonWriteOptions)
            )
            )
            return GetCurrentUserResponse.fromJsonString(results, {typeRegistry: this.registry});
        }
    }

    /**
     * ──────────────────────────────── Invoke ───────────────────────────────────────
     *
     * @generated from rpc auth.common.GenericService.Bootstrap
     */
    async bootstrap(request: BootstrapRequest, evaluate: boolean ): Promise< BootstrapResponse> {
        if (evaluate) {
            const results = utf8Decoder.decode(
                await this.contract.evaluateTransaction(
                "Bootstrap",
                request.toJsonString(this.jsonWriteOptions)
            )
            )
            return BootstrapResponse.fromJsonString(results, {typeRegistry: this.registry});
        } else {
            const results = utf8Decoder.decode(
                    await this.contract.submitTransaction(
                "Bootstrap",
                request.toJsonString(this.jsonWriteOptions)
            )
            )
            return BootstrapResponse.fromJsonString(results, {typeRegistry: this.registry});
        }
    }

    /**
     * @generated from rpc auth.common.GenericService.AuthorizeOperation
     */
    async authorizeOperation(request: AuthorizeOperationRequest, evaluate: boolean ): Promise< AuthorizeOperationResponse> {
        if (evaluate) {
            const results = utf8Decoder.decode(
                await this.contract.evaluateTransaction(
                "AuthorizeOperation",
                request.toJsonString(this.jsonWriteOptions)
            )
            )
            return AuthorizeOperationResponse.fromJsonString(results, {typeRegistry: this.registry});
        } else {
            const results = utf8Decoder.decode(
                    await this.contract.submitTransaction(
                "AuthorizeOperation",
                request.toJsonString(this.jsonWriteOptions)
            )
            )
            return AuthorizeOperationResponse.fromJsonString(results, {typeRegistry: this.registry});
        }
    }

    /**
     * @generated from rpc auth.common.GenericService.Get
     */
    async get(request: GetRequest, evaluate: boolean ): Promise< GetResponse> {
        if (evaluate) {
            const results = utf8Decoder.decode(
                await this.contract.evaluateTransaction(
                "Get",
                request.toJsonString(this.jsonWriteOptions)
            )
            )
            return GetResponse.fromJsonString(results, {typeRegistry: this.registry});
        } else {
            const results = utf8Decoder.decode(
                    await this.contract.submitTransaction(
                "Get",
                request.toJsonString(this.jsonWriteOptions)
            )
            )
            return GetResponse.fromJsonString(results, {typeRegistry: this.registry});
        }
    }

    /**
     * @generated from rpc auth.common.GenericService.GetFull
     */
    async getFull(request: GetFullRequest, evaluate: boolean ): Promise< GetFullResponse> {
        if (evaluate) {
            const results = utf8Decoder.decode(
                await this.contract.evaluateTransaction(
                "GetFull",
                request.toJsonString(this.jsonWriteOptions)
            )
            )
            return GetFullResponse.fromJsonString(results, {typeRegistry: this.registry});
        } else {
            const results = utf8Decoder.decode(
                    await this.contract.submitTransaction(
                "GetFull",
                request.toJsonString(this.jsonWriteOptions)
            )
            )
            return GetFullResponse.fromJsonString(results, {typeRegistry: this.registry});
        }
    }

    /**
     * @generated from rpc auth.common.GenericService.List
     */
    async list(request: ListRequest, evaluate: boolean ): Promise< ListResponse> {
        if (evaluate) {
            const results = utf8Decoder.decode(
                await this.contract.evaluateTransaction(
                "List",
                request.toJsonString(this.jsonWriteOptions)
            )
            )
            return ListResponse.fromJsonString(results, {typeRegistry: this.registry});
        } else {
            const results = utf8Decoder.decode(
                    await this.contract.submitTransaction(
                "List",
                request.toJsonString(this.jsonWriteOptions)
            )
            )
            return ListResponse.fromJsonString(results, {typeRegistry: this.registry});
        }
    }

    /**
     * @generated from rpc auth.common.GenericService.ListByCollection
     */
    async listByCollection(request: ListByCollectionRequest, evaluate: boolean ): Promise< ListByCollectionResponse> {
        if (evaluate) {
            const results = utf8Decoder.decode(
                await this.contract.evaluateTransaction(
                "ListByCollection",
                request.toJsonString(this.jsonWriteOptions)
            )
            )
            return ListByCollectionResponse.fromJsonString(results, {typeRegistry: this.registry});
        } else {
            const results = utf8Decoder.decode(
                    await this.contract.submitTransaction(
                "ListByCollection",
                request.toJsonString(this.jsonWriteOptions)
            )
            )
            return ListByCollectionResponse.fromJsonString(results, {typeRegistry: this.registry});
        }
    }

    /**
     * @generated from rpc auth.common.GenericService.ListByAttrs
     */
    async listByAttrs(request: ListByAttrsRequest, evaluate: boolean ): Promise< ListByAttrsResponse> {
        if (evaluate) {
            const results = utf8Decoder.decode(
                await this.contract.evaluateTransaction(
                "ListByAttrs",
                request.toJsonString(this.jsonWriteOptions)
            )
            )
            return ListByAttrsResponse.fromJsonString(results, {typeRegistry: this.registry});
        } else {
            const results = utf8Decoder.decode(
                    await this.contract.submitTransaction(
                "ListByAttrs",
                request.toJsonString(this.jsonWriteOptions)
            )
            )
            return ListByAttrsResponse.fromJsonString(results, {typeRegistry: this.registry});
        }
    }

    /**
     * @generated from rpc auth.common.GenericService.Create
     */
    async create(request: CreateRequest, evaluate: boolean ): Promise< CreateResponse> {
        if (evaluate) {
            const results = utf8Decoder.decode(
                await this.contract.evaluateTransaction(
                "Create",
                request.toJsonString(this.jsonWriteOptions)
            )
            )
            return CreateResponse.fromJsonString(results, {typeRegistry: this.registry});
        } else {
            const results = utf8Decoder.decode(
                    await this.contract.submitTransaction(
                "Create",
                request.toJsonString(this.jsonWriteOptions)
            )
            )
            return CreateResponse.fromJsonString(results, {typeRegistry: this.registry});
        }
    }

    /**
     * @generated from rpc auth.common.GenericService.Update
     */
    async update(request: UpdateRequest, evaluate: boolean ): Promise< UpdateResponse> {
        if (evaluate) {
            const results = utf8Decoder.decode(
                await this.contract.evaluateTransaction(
                "Update",
                request.toJsonString(this.jsonWriteOptions)
            )
            )
            return UpdateResponse.fromJsonString(results, {typeRegistry: this.registry});
        } else {
            const results = utf8Decoder.decode(
                    await this.contract.submitTransaction(
                "Update",
                request.toJsonString(this.jsonWriteOptions)
            )
            )
            return UpdateResponse.fromJsonString(results, {typeRegistry: this.registry});
        }
    }

    /**
     * @generated from rpc auth.common.GenericService.Delete
     */
    async delete(request: DeleteRequest, evaluate: boolean ): Promise< DeleteResponse> {
        if (evaluate) {
            const results = utf8Decoder.decode(
                await this.contract.evaluateTransaction(
                "Delete",
                request.toJsonString(this.jsonWriteOptions)
            )
            )
            return DeleteResponse.fromJsonString(results, {typeRegistry: this.registry});
        } else {
            const results = utf8Decoder.decode(
                    await this.contract.submitTransaction(
                "Delete",
                request.toJsonString(this.jsonWriteOptions)
            )
            )
            return DeleteResponse.fromJsonString(results, {typeRegistry: this.registry});
        }
    }

    /**
     * @generated from rpc auth.common.GenericService.GetHistory
     */
    async getHistory(request: GetHistoryRequest, evaluate: boolean ): Promise< GetHistoryResponse> {
        if (evaluate) {
            const results = utf8Decoder.decode(
                await this.contract.evaluateTransaction(
                "GetHistory",
                request.toJsonString(this.jsonWriteOptions)
            )
            )
            return GetHistoryResponse.fromJsonString(results, {typeRegistry: this.registry});
        } else {
            const results = utf8Decoder.decode(
                    await this.contract.submitTransaction(
                "GetHistory",
                request.toJsonString(this.jsonWriteOptions)
            )
            )
            return GetHistoryResponse.fromJsonString(results, {typeRegistry: this.registry});
        }
    }

    /**
     * @generated from rpc auth.common.GenericService.GetHiddenTx
     */
    async getHiddenTx(request: GetHiddenTxRequest, evaluate: boolean ): Promise< GetHiddenTxResponse> {
        if (evaluate) {
            const results = utf8Decoder.decode(
                await this.contract.evaluateTransaction(
                "GetHiddenTx",
                request.toJsonString(this.jsonWriteOptions)
            )
            )
            return GetHiddenTxResponse.fromJsonString(results, {typeRegistry: this.registry});
        } else {
            const results = utf8Decoder.decode(
                    await this.contract.submitTransaction(
                "GetHiddenTx",
                request.toJsonString(this.jsonWriteOptions)
            )
            )
            return GetHiddenTxResponse.fromJsonString(results, {typeRegistry: this.registry});
        }
    }

    /**
     * @generated from rpc auth.common.GenericService.HideTx
     */
    async hideTx(request: HideTxRequest, evaluate: boolean ): Promise< HideTxResponse> {
        if (evaluate) {
            const results = utf8Decoder.decode(
                await this.contract.evaluateTransaction(
                "HideTx",
                request.toJsonString(this.jsonWriteOptions)
            )
            )
            return HideTxResponse.fromJsonString(results, {typeRegistry: this.registry});
        } else {
            const results = utf8Decoder.decode(
                    await this.contract.submitTransaction(
                "HideTx",
                request.toJsonString(this.jsonWriteOptions)
            )
            )
            return HideTxResponse.fromJsonString(results, {typeRegistry: this.registry});
        }
    }

    /**
     * @generated from rpc auth.common.GenericService.UnHideTx
     */
    async unHideTx(request: UnHideTxRequest, evaluate: boolean ): Promise< UnHideTxResponse> {
        if (evaluate) {
            const results = utf8Decoder.decode(
                await this.contract.evaluateTransaction(
                "UnHideTx",
                request.toJsonString(this.jsonWriteOptions)
            )
            )
            return UnHideTxResponse.fromJsonString(results, {typeRegistry: this.registry});
        } else {
            const results = utf8Decoder.decode(
                    await this.contract.submitTransaction(
                "UnHideTx",
                request.toJsonString(this.jsonWriteOptions)
            )
            )
            return UnHideTxResponse.fromJsonString(results, {typeRegistry: this.registry});
        }
    }

    /**
     * @generated from rpc auth.common.GenericService.GetSuggestion
     */
    async getSuggestion(request: GetSuggestionRequest, evaluate: boolean ): Promise< GetSuggestionResponse> {
        if (evaluate) {
            const results = utf8Decoder.decode(
                await this.contract.evaluateTransaction(
                "GetSuggestion",
                request.toJsonString(this.jsonWriteOptions)
            )
            )
            return GetSuggestionResponse.fromJsonString(results, {typeRegistry: this.registry});
        } else {
            const results = utf8Decoder.decode(
                    await this.contract.submitTransaction(
                "GetSuggestion",
                request.toJsonString(this.jsonWriteOptions)
            )
            )
            return GetSuggestionResponse.fromJsonString(results, {typeRegistry: this.registry});
        }
    }

    /**
     * @generated from rpc auth.common.GenericService.SuggestionListByCollection
     */
    async suggestionListByCollection(request: SuggestionListByCollectionRequest, evaluate: boolean ): Promise< SuggestionListByCollectionResponse> {
        if (evaluate) {
            const results = utf8Decoder.decode(
                await this.contract.evaluateTransaction(
                "SuggestionListByCollection",
                request.toJsonString(this.jsonWriteOptions)
            )
            )
            return SuggestionListByCollectionResponse.fromJsonString(results, {typeRegistry: this.registry});
        } else {
            const results = utf8Decoder.decode(
                    await this.contract.submitTransaction(
                "SuggestionListByCollection",
                request.toJsonString(this.jsonWriteOptions)
            )
            )
            return SuggestionListByCollectionResponse.fromJsonString(results, {typeRegistry: this.registry});
        }
    }

    /**
     * @generated from rpc auth.common.GenericService.SuggestionByPartialKey
     */
    async suggestionByPartialKey(request: SuggestionByPartialKeyRequest, evaluate: boolean ): Promise< SuggestionByPartialKeyResponse> {
        if (evaluate) {
            const results = utf8Decoder.decode(
                await this.contract.evaluateTransaction(
                "SuggestionByPartialKey",
                request.toJsonString(this.jsonWriteOptions)
            )
            )
            return SuggestionByPartialKeyResponse.fromJsonString(results, {typeRegistry: this.registry});
        } else {
            const results = utf8Decoder.decode(
                    await this.contract.submitTransaction(
                "SuggestionByPartialKey",
                request.toJsonString(this.jsonWriteOptions)
            )
            )
            return SuggestionByPartialKeyResponse.fromJsonString(results, {typeRegistry: this.registry});
        }
    }

    /**
     * ──────────────────────────────── Invoke ───────────────────────────────────────
     *
     * @generated from rpc auth.common.GenericService.SuggestionCreate
     */
    async suggestionCreate(request: SuggestionCreateRequest, evaluate: boolean ): Promise< SuggestionCreateResponse> {
        if (evaluate) {
            const results = utf8Decoder.decode(
                await this.contract.evaluateTransaction(
                "SuggestionCreate",
                request.toJsonString(this.jsonWriteOptions)
            )
            )
            return SuggestionCreateResponse.fromJsonString(results, {typeRegistry: this.registry});
        } else {
            const results = utf8Decoder.decode(
                    await this.contract.submitTransaction(
                "SuggestionCreate",
                request.toJsonString(this.jsonWriteOptions)
            )
            )
            return SuggestionCreateResponse.fromJsonString(results, {typeRegistry: this.registry});
        }
    }

    /**
     * @generated from rpc auth.common.GenericService.SuggestionDelete
     */
    async suggestionDelete(request: SuggestionDeleteRequest, evaluate: boolean ): Promise< SuggestionDeleteResponse> {
        if (evaluate) {
            const results = utf8Decoder.decode(
                await this.contract.evaluateTransaction(
                "SuggestionDelete",
                request.toJsonString(this.jsonWriteOptions)
            )
            )
            return SuggestionDeleteResponse.fromJsonString(results, {typeRegistry: this.registry});
        } else {
            const results = utf8Decoder.decode(
                    await this.contract.submitTransaction(
                "SuggestionDelete",
                request.toJsonString(this.jsonWriteOptions)
            )
            )
            return SuggestionDeleteResponse.fromJsonString(results, {typeRegistry: this.registry});
        }
    }

    /**
     * @generated from rpc auth.common.GenericService.SuggestionApprove
     */
    async suggestionApprove(request: SuggestionApproveRequest, evaluate: boolean ): Promise< SuggestionApproveResponse> {
        if (evaluate) {
            const results = utf8Decoder.decode(
                await this.contract.evaluateTransaction(
                "SuggestionApprove",
                request.toJsonString(this.jsonWriteOptions)
            )
            )
            return SuggestionApproveResponse.fromJsonString(results, {typeRegistry: this.registry});
        } else {
            const results = utf8Decoder.decode(
                    await this.contract.submitTransaction(
                "SuggestionApprove",
                request.toJsonString(this.jsonWriteOptions)
            )
            )
            return SuggestionApproveResponse.fromJsonString(results, {typeRegistry: this.registry});
        }
    }
}
