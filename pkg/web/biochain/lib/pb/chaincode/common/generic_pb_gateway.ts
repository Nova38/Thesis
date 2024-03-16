import type { Contract } from '@hyperledger/fabric-gateway'
import type {
  IMessageTypeRegistry,
  JsonWriteStringOptions,
  PartialMessage,
} from '@bufbuild/protobuf'
import {
  AuthorizeOperationRequest,
  AuthorizeOperationResponse,
  BootstrapRequest,
  BootstrapResponse,
  CreateRequest,
  CreateResponse,
  DeleteRequest,
  DeleteResponse,
  GetCollectionsListResponse,
  GetCurrentUserResponse,
  GetFullRequest,
  GetFullResponse,
  GetHiddenTxRequest,
  GetHiddenTxResponse,
  GetHistoryRequest,
  GetHistoryResponse,
  GetRequest,
  GetResponse,
  GetSuggestionRequest,
  GetSuggestionResponse,
  HideTxRequest,
  HideTxResponse,
  ListByAttrsRequest,
  ListByAttrsResponse,
  ListByCollectionRequest,
  ListByCollectionResponse,
  ListRequest,
  ListResponse,
  SuggestionApproveRequest,
  SuggestionApproveResponse,
  SuggestionByPartialKeyRequest,
  SuggestionByPartialKeyResponse,
  SuggestionCreateRequest,
  SuggestionCreateResponse,
  SuggestionDeleteRequest,
  SuggestionDeleteResponse,
  SuggestionListByCollectionRequest,
  SuggestionListByCollectionResponse,
  UnHideTxRequest,
  UnHideTxResponse,
  UpdateRequest,
  UpdateResponse,
} from './generic_pb.js'

const utf8Decoder = new TextDecoder()
/**
 * @generated from service auth.common.GenericService
 */
export class GenericServiceClient {
  private contract: Contract
  private jsonWriteOptions: Partial<JsonWriteStringOptions> = {}
  registry: IMessageTypeRegistry

  constructor(contract: Contract, registry: IMessageTypeRegistry) {
    this.contract = contract
    this.registry = registry
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
      await this.contract.evaluateTransaction('GetCurrentUser'),
    )
    return GetCurrentUserResponse.fromJsonString(results, {
      typeRegistry: this.registry,
    })
  }

  /**
   * ──────────────────────────────── Invoke ───────────────────────────────────────
   *
   * @generated from rpc auth.common.GenericService.Bootstrap
   */
  async bootstrap(
    req: PartialMessage<BootstrapRequest>,
  ): Promise<BootstrapResponse> {
    const msg
      = req instanceof BootstrapRequest ? req : new BootstrapRequest(req)
    const results = utf8Decoder.decode(
      await this.contract.submitTransaction(
        'Bootstrap',
        msg.toJsonString(this.jsonWriteOptions),
      ),
    )
    return BootstrapResponse.fromJsonString(results, {
      typeRegistry: this.registry,
    })
  }

  /**
   * @generated from rpc auth.common.GenericService.AuthorizeOperation
   */
  async authorizeOperation(
    req: PartialMessage<AuthorizeOperationRequest>,
  ): Promise<AuthorizeOperationResponse> {
    const msg
      = req instanceof AuthorizeOperationRequest
        ? req
        : new AuthorizeOperationRequest(req)
    const results = utf8Decoder.decode(
      await this.contract.submitTransaction(
        'AuthorizeOperation',
        msg.toJsonString(this.jsonWriteOptions),
      ),
    )
    return AuthorizeOperationResponse.fromJsonString(results, {
      typeRegistry: this.registry,
    })
  }

  /**
   * @generated from rpc auth.common.GenericService.GetCollectionsList
   */
  async getCollectionsList(): Promise<GetCollectionsListResponse> {
    const results = utf8Decoder.decode(
      await this.contract.evaluateTransaction('GetCollectionsList'),
    )
    return GetCollectionsListResponse.fromJsonString(results, {
      typeRegistry: this.registry,
    })
  }

  /**
   * @generated from rpc auth.common.GenericService.Get
   */
  async get(req: PartialMessage<GetRequest>): Promise<GetResponse> {
    const msg = req instanceof GetRequest ? req : new GetRequest(req)
    const results = utf8Decoder.decode(
      await this.contract.evaluateTransaction(
        'Get',
        msg.toJsonString(this.jsonWriteOptions),
      ),
    )
    return GetResponse.fromJsonString(results, { typeRegistry: this.registry })
  }

  /**
   * @generated from rpc auth.common.GenericService.GetFull
   */
  async getFull(req: PartialMessage<GetFullRequest>): Promise<GetFullResponse> {
    const msg = req instanceof GetFullRequest ? req : new GetFullRequest(req)
    const results = utf8Decoder.decode(
      await this.contract.evaluateTransaction(
        'GetFull',
        msg.toJsonString(this.jsonWriteOptions),
      ),
    )
    return GetFullResponse.fromJsonString(results, {
      typeRegistry: this.registry,
    })
  }

  /**
   * @generated from rpc auth.common.GenericService.List
   */
  async list(req: PartialMessage<ListRequest>): Promise<ListResponse> {
    const msg = req instanceof ListRequest ? req : new ListRequest(req)
    const results = utf8Decoder.decode(
      await this.contract.evaluateTransaction(
        'List',
        msg.toJsonString(this.jsonWriteOptions),
      ),
    )
    return ListResponse.fromJsonString(results, { typeRegistry: this.registry })
  }

  /**
   * @generated from rpc auth.common.GenericService.ListByCollection
   */
  async listByCollection(
    req: PartialMessage<ListByCollectionRequest>,
  ): Promise<ListByCollectionResponse> {
    const msg
      = req instanceof ListByCollectionRequest
        ? req
        : new ListByCollectionRequest(req)
    const results = utf8Decoder.decode(
      await this.contract.evaluateTransaction(
        'ListByCollection',
        msg.toJsonString(this.jsonWriteOptions),
      ),
    )
    return ListByCollectionResponse.fromJsonString(results, {
      typeRegistry: this.registry,
    })
  }

  /**
   * @generated from rpc auth.common.GenericService.ListByAttrs
   */
  async listByAttrs(
    req: PartialMessage<ListByAttrsRequest>,
  ): Promise<ListByAttrsResponse> {
    const msg
      = req instanceof ListByAttrsRequest ? req : new ListByAttrsRequest(req)
    const results = utf8Decoder.decode(
      await this.contract.evaluateTransaction(
        'ListByAttrs',
        msg.toJsonString(this.jsonWriteOptions),
      ),
    )
    return ListByAttrsResponse.fromJsonString(results, {
      typeRegistry: this.registry,
    })
  }

  /**
   * @generated from rpc auth.common.GenericService.Create
   */
  async create(req: PartialMessage<CreateRequest>): Promise<CreateResponse> {
    const msg = req instanceof CreateRequest ? req : new CreateRequest(req)
    const results = utf8Decoder.decode(
      await this.contract.submitTransaction(
        'Create',
        msg.toJsonString(this.jsonWriteOptions),
      ),
    )
    return CreateResponse.fromJsonString(results, {
      typeRegistry: this.registry,
    })
  }

  /**
   * @generated from rpc auth.common.GenericService.Update
   */
  async update(req: PartialMessage<UpdateRequest>): Promise<UpdateResponse> {
    const msg = req instanceof UpdateRequest ? req : new UpdateRequest(req)
    const results = utf8Decoder.decode(
      await this.contract.submitTransaction(
        'Update',
        msg.toJsonString(this.jsonWriteOptions),
      ),
    )
    return UpdateResponse.fromJsonString(results, {
      typeRegistry: this.registry,
    })
  }

  /**
   * @generated from rpc auth.common.GenericService.Delete
   */
  async delete(req: PartialMessage<DeleteRequest>): Promise<DeleteResponse> {
    const msg = req instanceof DeleteRequest ? req : new DeleteRequest(req)
    const results = utf8Decoder.decode(
      await this.contract.submitTransaction(
        'Delete',
        msg.toJsonString(this.jsonWriteOptions),
      ),
    )
    return DeleteResponse.fromJsonString(results, {
      typeRegistry: this.registry,
    })
  }

  /**
   * @generated from rpc auth.common.GenericService.GetHistory
   */
  async getHistory(
    req: PartialMessage<GetHistoryRequest>,
  ): Promise<GetHistoryResponse> {
    const msg
      = req instanceof GetHistoryRequest ? req : new GetHistoryRequest(req)
    const results = utf8Decoder.decode(
      await this.contract.evaluateTransaction(
        'GetHistory',
        msg.toJsonString(this.jsonWriteOptions),
      ),
    )
    return GetHistoryResponse.fromJsonString(results, {
      typeRegistry: this.registry,
    })
  }

  /**
   * @generated from rpc auth.common.GenericService.GetHiddenTx
   */
  async getHiddenTx(
    req: PartialMessage<GetHiddenTxRequest>,
  ): Promise<GetHiddenTxResponse> {
    const msg
      = req instanceof GetHiddenTxRequest ? req : new GetHiddenTxRequest(req)
    const results = utf8Decoder.decode(
      await this.contract.evaluateTransaction(
        'GetHiddenTx',
        msg.toJsonString(this.jsonWriteOptions),
      ),
    )
    return GetHiddenTxResponse.fromJsonString(results, {
      typeRegistry: this.registry,
    })
  }

  /**
   * @generated from rpc auth.common.GenericService.HideTx
   */
  async hideTx(req: PartialMessage<HideTxRequest>): Promise<HideTxResponse> {
    const msg = req instanceof HideTxRequest ? req : new HideTxRequest(req)
    const results = utf8Decoder.decode(
      await this.contract.submitTransaction(
        'HideTx',
        msg.toJsonString(this.jsonWriteOptions),
      ),
    )
    return HideTxResponse.fromJsonString(results, {
      typeRegistry: this.registry,
    })
  }

  /**
   * @generated from rpc auth.common.GenericService.UnHideTx
   */
  async unHideTx(
    req: PartialMessage<UnHideTxRequest>,
  ): Promise<UnHideTxResponse> {
    const msg = req instanceof UnHideTxRequest ? req : new UnHideTxRequest(req)
    const results = utf8Decoder.decode(
      await this.contract.submitTransaction(
        'UnHideTx',
        msg.toJsonString(this.jsonWriteOptions),
      ),
    )
    return UnHideTxResponse.fromJsonString(results, {
      typeRegistry: this.registry,
    })
  }

  /**
   * @generated from rpc auth.common.GenericService.GetSuggestion
   */
  async getSuggestion(
    req: PartialMessage<GetSuggestionRequest>,
  ): Promise<GetSuggestionResponse> {
    const msg
      = req instanceof GetSuggestionRequest ? req : new GetSuggestionRequest(req)
    const results = utf8Decoder.decode(
      await this.contract.evaluateTransaction(
        'GetSuggestion',
        msg.toJsonString(this.jsonWriteOptions),
      ),
    )
    return GetSuggestionResponse.fromJsonString(results, {
      typeRegistry: this.registry,
    })
  }

  /**
   * @generated from rpc auth.common.GenericService.SuggestionListByCollection
   */
  async suggestionListByCollection(
    req: PartialMessage<SuggestionListByCollectionRequest>,
  ): Promise<SuggestionListByCollectionResponse> {
    const msg
      = req instanceof SuggestionListByCollectionRequest
        ? req
        : new SuggestionListByCollectionRequest(req)
    const results = utf8Decoder.decode(
      await this.contract.evaluateTransaction(
        'SuggestionListByCollection',
        msg.toJsonString(this.jsonWriteOptions),
      ),
    )
    return SuggestionListByCollectionResponse.fromJsonString(results, {
      typeRegistry: this.registry,
    })
  }

  /**
   * @generated from rpc auth.common.GenericService.SuggestionByPartialKey
   */
  async suggestionByPartialKey(
    req: PartialMessage<SuggestionByPartialKeyRequest>,
  ): Promise<SuggestionByPartialKeyResponse> {
    const msg
      = req instanceof SuggestionByPartialKeyRequest
        ? req
        : new SuggestionByPartialKeyRequest(req)
    const results = utf8Decoder.decode(
      await this.contract.evaluateTransaction(
        'SuggestionByPartialKey',
        msg.toJsonString(this.jsonWriteOptions),
      ),
    )
    return SuggestionByPartialKeyResponse.fromJsonString(results, {
      typeRegistry: this.registry,
    })
  }

  /**
   * ──────────────────────────────── Invoke ───────────────────────────────────────
   *
   * @generated from rpc auth.common.GenericService.SuggestionCreate
   */
  async suggestionCreate(
    req: PartialMessage<SuggestionCreateRequest>,
  ): Promise<SuggestionCreateResponse> {
    const msg
      = req instanceof SuggestionCreateRequest
        ? req
        : new SuggestionCreateRequest(req)
    const results = utf8Decoder.decode(
      await this.contract.submitTransaction(
        'SuggestionCreate',
        msg.toJsonString(this.jsonWriteOptions),
      ),
    )
    return SuggestionCreateResponse.fromJsonString(results, {
      typeRegistry: this.registry,
    })
  }

  /**
   * @generated from rpc auth.common.GenericService.SuggestionDelete
   */
  async suggestionDelete(
    req: PartialMessage<SuggestionDeleteRequest>,
  ): Promise<SuggestionDeleteResponse> {
    const msg
      = req instanceof SuggestionDeleteRequest
        ? req
        : new SuggestionDeleteRequest(req)
    const results = utf8Decoder.decode(
      await this.contract.submitTransaction(
        'SuggestionDelete',
        msg.toJsonString(this.jsonWriteOptions),
      ),
    )
    return SuggestionDeleteResponse.fromJsonString(results, {
      typeRegistry: this.registry,
    })
  }

  /**
   * @generated from rpc auth.common.GenericService.SuggestionApprove
   */
  async suggestionApprove(
    req: PartialMessage<SuggestionApproveRequest>,
  ): Promise<SuggestionApproveResponse> {
    const msg
      = req instanceof SuggestionApproveRequest
        ? req
        : new SuggestionApproveRequest(req)
    const results = utf8Decoder.decode(
      await this.contract.submitTransaction(
        'SuggestionApprove',
        msg.toJsonString(this.jsonWriteOptions),
      ),
    )
    return SuggestionApproveResponse.fromJsonString(results, {
      typeRegistry: this.registry,
    })
  }
}
