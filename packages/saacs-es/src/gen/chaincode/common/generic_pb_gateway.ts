import { Contract } from "@hyperledger/fabric-gateway";
import { IMessageTypeRegistry, JsonWriteStringOptions } from "@bufbuild/protobuf";
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
// 2
    async getCurrentUser(): Promise< GetCurrentUserResponse> {
        const results = utf8Decoder.decode(
                await this.contract.evaluateTransaction(
                "GetCurrentUser"
            )
        )
            return GetCurrentUserResponse.fromJsonString(results, {typeRegistry: this.registry});
    }

    /**
     * ──────────────────────────────── Invoke ───────────────────────────────────────
     *
     * @generated from rpc auth.common.GenericService.Bootstrap
     */
// 1
    async bootstrap(request: BootstrapRequest, evaluate: boolean ): Promise< BootstrapResponse> {
            const results = utf8Decoder.decode(
                await this.contract.submitTransaction(
                "Bootstrap",
                request.toJsonString(this.jsonWriteOptions)
            )
            )
            return BootstrapResponse.fromJsonString(results, {typeRegistry: this.registry});
    }

    /**
     * @generated from rpc auth.common.GenericService.AuthorizeOperation
     */
// 1
    async authorizeOperation(request: AuthorizeOperationRequest, evaluate: boolean ): Promise< AuthorizeOperationResponse> {
            const results = utf8Decoder.decode(
                await this.contract.submitTransaction(
                "AuthorizeOperation",
                request.toJsonString(this.jsonWriteOptions)
            )
            )
            return AuthorizeOperationResponse.fromJsonString(results, {typeRegistry: this.registry});
    }

    /**
     * @generated from rpc auth.common.GenericService.Get
     */
// 2
    async get(request: GetRequest, evaluate: boolean ): Promise< GetResponse> {
            const results = utf8Decoder.decode(
                await this.contract.evaluateTransaction(
                "Get",
                request.toJsonString(this.jsonWriteOptions)
            )
            )
            return GetResponse.fromJsonString(results, {typeRegistry: this.registry});
    }

    /**
     * @generated from rpc auth.common.GenericService.GetFull
     */
// 2
    async getFull(request: GetFullRequest, evaluate: boolean ): Promise< GetFullResponse> {
            const results = utf8Decoder.decode(
                await this.contract.evaluateTransaction(
                "GetFull",
                request.toJsonString(this.jsonWriteOptions)
            )
            )
            return GetFullResponse.fromJsonString(results, {typeRegistry: this.registry});
    }

    /**
     * @generated from rpc auth.common.GenericService.List
     */
// 2
    async list(request: ListRequest, evaluate: boolean ): Promise< ListResponse> {
            const results = utf8Decoder.decode(
                await this.contract.evaluateTransaction(
                "List",
                request.toJsonString(this.jsonWriteOptions)
            )
            )
            return ListResponse.fromJsonString(results, {typeRegistry: this.registry});
    }

    /**
     * @generated from rpc auth.common.GenericService.ListByCollection
     */
// 2
    async listByCollection(request: ListByCollectionRequest, evaluate: boolean ): Promise< ListByCollectionResponse> {
            const results = utf8Decoder.decode(
                await this.contract.evaluateTransaction(
                "ListByCollection",
                request.toJsonString(this.jsonWriteOptions)
            )
            )
            return ListByCollectionResponse.fromJsonString(results, {typeRegistry: this.registry});
    }

    /**
     * @generated from rpc auth.common.GenericService.ListByAttrs
     */
// 2
    async listByAttrs(request: ListByAttrsRequest, evaluate: boolean ): Promise< ListByAttrsResponse> {
            const results = utf8Decoder.decode(
                await this.contract.evaluateTransaction(
                "ListByAttrs",
                request.toJsonString(this.jsonWriteOptions)
            )
            )
            return ListByAttrsResponse.fromJsonString(results, {typeRegistry: this.registry});
    }

    /**
     * @generated from rpc auth.common.GenericService.Create
     */
// 1
    async create(request: CreateRequest, evaluate: boolean ): Promise< CreateResponse> {
            const results = utf8Decoder.decode(
                await this.contract.submitTransaction(
                "Create",
                request.toJsonString(this.jsonWriteOptions)
            )
            )
            return CreateResponse.fromJsonString(results, {typeRegistry: this.registry});
    }

    /**
     * @generated from rpc auth.common.GenericService.Update
     */
// 1
    async update(request: UpdateRequest, evaluate: boolean ): Promise< UpdateResponse> {
            const results = utf8Decoder.decode(
                await this.contract.submitTransaction(
                "Update",
                request.toJsonString(this.jsonWriteOptions)
            )
            )
            return UpdateResponse.fromJsonString(results, {typeRegistry: this.registry});
    }

    /**
     * @generated from rpc auth.common.GenericService.Delete
     */
// 1
    async delete(request: DeleteRequest, evaluate: boolean ): Promise< DeleteResponse> {
            const results = utf8Decoder.decode(
                await this.contract.submitTransaction(
                "Delete",
                request.toJsonString(this.jsonWriteOptions)
            )
            )
            return DeleteResponse.fromJsonString(results, {typeRegistry: this.registry});
    }

    /**
     * @generated from rpc auth.common.GenericService.GetHistory
     */
// 2
    async getHistory(request: GetHistoryRequest, evaluate: boolean ): Promise< GetHistoryResponse> {
            const results = utf8Decoder.decode(
                await this.contract.evaluateTransaction(
                "GetHistory",
                request.toJsonString(this.jsonWriteOptions)
            )
            )
            return GetHistoryResponse.fromJsonString(results, {typeRegistry: this.registry});
    }

    /**
     * @generated from rpc auth.common.GenericService.GetHiddenTx
     */
// 2
    async getHiddenTx(request: GetHiddenTxRequest, evaluate: boolean ): Promise< GetHiddenTxResponse> {
            const results = utf8Decoder.decode(
                await this.contract.evaluateTransaction(
                "GetHiddenTx",
                request.toJsonString(this.jsonWriteOptions)
            )
            )
            return GetHiddenTxResponse.fromJsonString(results, {typeRegistry: this.registry});
    }

    /**
     * @generated from rpc auth.common.GenericService.HideTx
     */
// 1
    async hideTx(request: HideTxRequest, evaluate: boolean ): Promise< HideTxResponse> {
            const results = utf8Decoder.decode(
                await this.contract.submitTransaction(
                "HideTx",
                request.toJsonString(this.jsonWriteOptions)
            )
            )
            return HideTxResponse.fromJsonString(results, {typeRegistry: this.registry});
    }

    /**
     * @generated from rpc auth.common.GenericService.UnHideTx
     */
// 1
    async unHideTx(request: UnHideTxRequest, evaluate: boolean ): Promise< UnHideTxResponse> {
            const results = utf8Decoder.decode(
                await this.contract.submitTransaction(
                "UnHideTx",
                request.toJsonString(this.jsonWriteOptions)
            )
            )
            return UnHideTxResponse.fromJsonString(results, {typeRegistry: this.registry});
    }

    /**
     * @generated from rpc auth.common.GenericService.GetSuggestion
     */
// 2
    async getSuggestion(request: GetSuggestionRequest, evaluate: boolean ): Promise< GetSuggestionResponse> {
            const results = utf8Decoder.decode(
                await this.contract.evaluateTransaction(
                "GetSuggestion",
                request.toJsonString(this.jsonWriteOptions)
            )
            )
            return GetSuggestionResponse.fromJsonString(results, {typeRegistry: this.registry});
    }

    /**
     * @generated from rpc auth.common.GenericService.SuggestionListByCollection
     */
// 2
    async suggestionListByCollection(request: SuggestionListByCollectionRequest, evaluate: boolean ): Promise< SuggestionListByCollectionResponse> {
            const results = utf8Decoder.decode(
                await this.contract.evaluateTransaction(
                "SuggestionListByCollection",
                request.toJsonString(this.jsonWriteOptions)
            )
            )
            return SuggestionListByCollectionResponse.fromJsonString(results, {typeRegistry: this.registry});
    }

    /**
     * @generated from rpc auth.common.GenericService.SuggestionByPartialKey
     */
// 2
    async suggestionByPartialKey(request: SuggestionByPartialKeyRequest, evaluate: boolean ): Promise< SuggestionByPartialKeyResponse> {
            const results = utf8Decoder.decode(
                await this.contract.evaluateTransaction(
                "SuggestionByPartialKey",
                request.toJsonString(this.jsonWriteOptions)
            )
            )
            return SuggestionByPartialKeyResponse.fromJsonString(results, {typeRegistry: this.registry});
    }

    /**
     * ──────────────────────────────── Invoke ───────────────────────────────────────
     *
     * @generated from rpc auth.common.GenericService.SuggestionCreate
     */
// 1
    async suggestionCreate(request: SuggestionCreateRequest, evaluate: boolean ): Promise< SuggestionCreateResponse> {
            const results = utf8Decoder.decode(
                await this.contract.submitTransaction(
                "SuggestionCreate",
                request.toJsonString(this.jsonWriteOptions)
            )
            )
            return SuggestionCreateResponse.fromJsonString(results, {typeRegistry: this.registry});
    }

    /**
     * @generated from rpc auth.common.GenericService.SuggestionDelete
     */
// 1
    async suggestionDelete(request: SuggestionDeleteRequest, evaluate: boolean ): Promise< SuggestionDeleteResponse> {
            const results = utf8Decoder.decode(
                await this.contract.submitTransaction(
                "SuggestionDelete",
                request.toJsonString(this.jsonWriteOptions)
            )
            )
            return SuggestionDeleteResponse.fromJsonString(results, {typeRegistry: this.registry});
    }

    /**
     * @generated from rpc auth.common.GenericService.SuggestionApprove
     */
// 1
    async suggestionApprove(request: SuggestionApproveRequest, evaluate: boolean ): Promise< SuggestionApproveResponse> {
            const results = utf8Decoder.decode(
                await this.contract.submitTransaction(
                "SuggestionApprove",
                request.toJsonString(this.jsonWriteOptions)
            )
            )
            return SuggestionApproveResponse.fromJsonString(results, {typeRegistry: this.registry});
    }
}
