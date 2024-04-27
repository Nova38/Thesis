import { Contract } from '@hyperledger/fabric-gateway'
import type { PartialMessage } from '@bufbuild/protobuf'
import {
  IMessageTypeRegistry,
  JsonWriteStringOptions,
} from '@bufbuild/protobuf'
import {
  CreateRequest,
  CreateResponse,
  DeleteRequest,
  DeleteResponse,
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
  SuggestionListByItemRequest,
  SuggestionListByItemResponse,
  UnHideTxRequest,
  UnHideTxResponse,
  UpdateRequest,
  UpdateResponse,
} from './chaincode_pb.js'

const utf8Decoder = new TextDecoder()
/**
 * rpc CreateCollection(CreateCollectionRequest) returns
 * (CreateCollectionResponse) {
 *   option (saacs.common.v0.transaction_type) = TRANSACTION_TYPE_INVOKE;
 *   option (saacs.common.v0.operation) = {action: ACTION_CREATE, item_type:
 *   "Collection"};
 * }
 * ══════════════════════════════════ Item
 * ═════════════════════════════════════
 *
 * @generated from service saacs.chaincode.v0.ItemService
 */
export class ItemServiceClient {
  private contract: Contract
  private jsonWriteOptions: Partial<JsonWriteStringOptions> = {}
  registry: IMessageTypeRegistry

  constructor(contract: Contract, registry: IMessageTypeRegistry) {
    this.contract = contract
    this.registry = registry
    this.jsonWriteOptions.typeRegistry = registry
  }

  /**
   * @generated from rpc saacs.chaincode.v0.ItemService.Get
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
   * @generated from rpc saacs.chaincode.v0.ItemService.GetFull
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
   * @generated from rpc saacs.chaincode.v0.ItemService.List
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
   * @generated from rpc saacs.chaincode.v0.ItemService.ListByAttrs
   */
  async listByAttrs(
    req: PartialMessage<ListByAttrsRequest>,
  ): Promise<ListByAttrsResponse> {
    const msg =
      req instanceof ListByAttrsRequest ? req : new ListByAttrsRequest(req)
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
   * @generated from rpc saacs.chaincode.v0.ItemService.Create
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
   * @generated from rpc saacs.chaincode.v0.ItemService.Update
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
   * @generated from rpc saacs.chaincode.v0.ItemService.Delete
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
   * @generated from rpc saacs.chaincode.v0.ItemService.GetHistory
   */
  async getHistory(
    req: PartialMessage<GetHistoryRequest>,
  ): Promise<GetHistoryResponse> {
    const msg =
      req instanceof GetHistoryRequest ? req : new GetHistoryRequest(req)
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
   * @generated from rpc saacs.chaincode.v0.ItemService.GetHiddenTx
   */
  async getHiddenTx(
    req: PartialMessage<GetHiddenTxRequest>,
  ): Promise<GetHiddenTxResponse> {
    const msg =
      req instanceof GetHiddenTxRequest ? req : new GetHiddenTxRequest(req)
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
   * @generated from rpc saacs.chaincode.v0.ItemService.HideTx
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
   * @generated from rpc saacs.chaincode.v0.ItemService.UnHideTx
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
   * @generated from rpc saacs.chaincode.v0.ItemService.GetSuggestion
   */
  async getSuggestion(
    req: PartialMessage<GetSuggestionRequest>,
  ): Promise<GetSuggestionResponse> {
    const msg =
      req instanceof GetSuggestionRequest ? req : new GetSuggestionRequest(req)
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
   * @generated from rpc saacs.chaincode.v0.ItemService.SuggestionListByCollection
   */
  async suggestionListByCollection(
    req: PartialMessage<SuggestionListByCollectionRequest>,
  ): Promise<SuggestionListByCollectionResponse> {
    const msg =
      req instanceof SuggestionListByCollectionRequest
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
   * @generated from rpc saacs.chaincode.v0.ItemService.SuggestionListByItem
   */
  async suggestionListByItem(
    req: PartialMessage<SuggestionListByItemRequest>,
  ): Promise<SuggestionListByItemResponse> {
    const msg =
      req instanceof SuggestionListByItemRequest
        ? req
        : new SuggestionListByItemRequest(req)
    const results = utf8Decoder.decode(
      await this.contract.evaluateTransaction(
        'SuggestionListByItem',
        msg.toJsonString(this.jsonWriteOptions),
      ),
    )
    return SuggestionListByItemResponse.fromJsonString(results, {
      typeRegistry: this.registry,
    })
  }

  /**
   * @generated from rpc saacs.chaincode.v0.ItemService.SuggestionByPartialKey
   */
  async suggestionByPartialKey(
    req: PartialMessage<SuggestionByPartialKeyRequest>,
  ): Promise<SuggestionByPartialKeyResponse> {
    const msg =
      req instanceof SuggestionByPartialKeyRequest
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
   * @generated from rpc saacs.chaincode.v0.ItemService.SuggestionCreate
   */
  async suggestionCreate(
    req: PartialMessage<SuggestionCreateRequest>,
  ): Promise<SuggestionCreateResponse> {
    const msg =
      req instanceof SuggestionCreateRequest
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
   * @generated from rpc saacs.chaincode.v0.ItemService.SuggestionDelete
   */
  async suggestionDelete(
    req: PartialMessage<SuggestionDeleteRequest>,
  ): Promise<SuggestionDeleteResponse> {
    const msg =
      req instanceof SuggestionDeleteRequest
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
   * @generated from rpc saacs.chaincode.v0.ItemService.SuggestionApprove
   */
  async suggestionApprove(
    req: PartialMessage<SuggestionApproveRequest>,
  ): Promise<SuggestionApproveResponse> {
    const msg =
      req instanceof SuggestionApproveRequest
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
