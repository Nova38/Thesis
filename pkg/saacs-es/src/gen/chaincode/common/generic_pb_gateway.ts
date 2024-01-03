import { Contract } from "@hyperledger/fabric-gateway";
import { IMessageTypeRegistry, JsonWriteStringOptions } from "@bufbuild/protobuf";
import { AuthorizeOperationRequest, AuthorizeOperationResponse, BootstrapRequest, BootstrapResponse, CreateRequest, CreateResponse, DeleteRequest, DeleteResponse, GetCollectionsListResponse, GetCurrentUserResponse, GetFullRequest, GetFullResponse, GetHiddenTxRequest, GetHiddenTxResponse, GetHistoryRequest, GetHistoryResponse, GetRequest, GetResponse, GetSuggestionRequest, GetSuggestionResponse, HideTxRequest, HideTxResponse, ListByAttrsRequest, ListByAttrsResponse, ListByCollectionRequest, ListByCollectionResponse, ListRequest, ListResponse, SuggestionApproveRequest, SuggestionApproveResponse, SuggestionByPartialKeyRequest, SuggestionByPartialKeyResponse, SuggestionCreateRequest, SuggestionCreateResponse, SuggestionDeleteRequest, SuggestionDeleteResponse, SuggestionListByCollectionRequest, SuggestionListByCollectionResponse, UnHideTxRequest, UnHideTxResponse, UpdateRequest, UpdateResponse } from "./generic_pb.js";

const utf8Decoder = new TextDecoder();
/**
 * @generated from service auth.common.GenericService
 */
export class GenericServiceClient {
    private contract: Contract;
    private jsonWriteOptions:Partial<JsonWriteStringOptions> = {};
    registry: IMessageTypeRegistry;

    constructor(contract: Contract, registry: IMessageTypeRegistry) {
        this.contract = contract;
        this.registry = registry;
        this.jsonWriteOptions.typeRegistry = registry
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
    async getCurrentUser(): Promise<GetCurrentUserResponse> {
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
    async bootstrap(request: BootstrapRequest ): Promise<BootstrapResponse> {
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
    async authorizeOperation(request: AuthorizeOperationRequest ): Promise<AuthorizeOperationResponse> {
            const results = utf8Decoder.decode(
                await this.contract.submitTransaction(
                "AuthorizeOperation",
                request.toJsonString(this.jsonWriteOptions)
            )
            )
            return AuthorizeOperationResponse.fromJsonString(results, {typeRegistry: this.registry});
    }

    /**
     * @generated from rpc auth.common.GenericService.GetCollectionsList
     */
    async getCollectionsList(): Promise<GetCollectionsListResponse> {
        const results = utf8Decoder.decode(
                await this.contract.evaluateTransaction(
                "GetCollectionsList"
            )
        )
            return GetCollectionsListResponse.fromJsonString(results, {typeRegistry: this.registry});
    }

    /**
     * @generated from rpc auth.common.GenericService.Get
     */
    async get(request: GetRequest ): Promise<GetResponse> {
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
    async getFull(request: GetFullRequest ): Promise<GetFullResponse> {
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
    async list(request: ListRequest ): Promise<ListResponse> {
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
    async listByCollection(request: ListByCollectionRequest ): Promise<ListByCollectionResponse> {
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
    async listByAttrs(request: ListByAttrsRequest ): Promise<ListByAttrsResponse> {
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
    async create(request: CreateRequest ): Promise<CreateResponse> {
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
    async update(request: UpdateRequest ): Promise<UpdateResponse> {
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
    async delete(request: DeleteRequest ): Promise<DeleteResponse> {
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
    async getHistory(request: GetHistoryRequest ): Promise<GetHistoryResponse> {
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
    async getHiddenTx(request: GetHiddenTxRequest ): Promise<GetHiddenTxResponse> {
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
    async hideTx(request: HideTxRequest ): Promise<HideTxResponse> {
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
    async unHideTx(request: UnHideTxRequest ): Promise<UnHideTxResponse> {
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
    async getSuggestion(request: GetSuggestionRequest ): Promise<GetSuggestionResponse> {
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
    async suggestionListByCollection(request: SuggestionListByCollectionRequest ): Promise<SuggestionListByCollectionResponse> {
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
    async suggestionByPartialKey(request: SuggestionByPartialKeyRequest ): Promise<SuggestionByPartialKeyResponse> {
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
    async suggestionCreate(request: SuggestionCreateRequest ): Promise<SuggestionCreateResponse> {
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
    async suggestionDelete(request: SuggestionDeleteRequest ): Promise<SuggestionDeleteResponse> {
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
    async suggestionApprove(request: SuggestionApproveRequest ): Promise<SuggestionApproveResponse> {
            const results = utf8Decoder.decode(
                await this.contract.submitTransaction(
                "SuggestionApprove",
                request.toJsonString(this.jsonWriteOptions)
            )
            )
            return SuggestionApproveResponse.fromJsonString(results, {typeRegistry: this.registry});
    }
}
