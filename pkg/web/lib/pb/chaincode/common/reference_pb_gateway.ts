import type { Contract } from '@hyperledger/fabric-gateway'
import type {
  IMessageTypeRegistry,
  JsonWriteStringOptions,
  PartialMessage,
} from '@bufbuild/protobuf'
import {
  ReferenceByItemRequest,
  ReferenceByItemResponse,
  ReferenceByPartialKeyRequest,
  ReferenceByPartialKeyResponse,
  ReferenceCreateRequest,
  ReferenceCreateResponse,
  ReferenceDeleteRequest,
  ReferenceDeleteResponse,
  ReferenceRequest,
  ReferenceResponse,
} from './reference_pb.js'

const utf8Decoder = new TextDecoder()
/**
 * ════════════════════════════════ References ══════════════════════════════════
 * ──────────────────────────────── Query ────────────────────────────────────────
 *
 * @generated from service auth.common.ReferenceService
 */
export class ReferenceServiceClient {
  private contract: Contract
  private jsonWriteOptions: Partial<JsonWriteStringOptions> = {}
  registry: IMessageTypeRegistry

  constructor(contract: Contract, registry: IMessageTypeRegistry) {
    this.contract = contract
    this.registry = registry
    this.jsonWriteOptions.typeRegistry = registry
  }

  /**
   * @generated from rpc auth.common.ReferenceService.Reference
   */
  async reference(
    req: PartialMessage<ReferenceRequest>,
  ): Promise<ReferenceResponse> {
    const msg
      = req instanceof ReferenceRequest ? req : new ReferenceRequest(req)
    const results = utf8Decoder.decode(
      await this.contract.evaluateTransaction(
        'Reference',
        msg.toJsonString(this.jsonWriteOptions),
      ),
    )
    return ReferenceResponse.fromJsonString(results, {
      typeRegistry: this.registry,
    })
  }

  /**
   * @generated from rpc auth.common.ReferenceService.ReferenceByItem
   */
  async referenceByItem(
    req: PartialMessage<ReferenceByItemRequest>,
  ): Promise<ReferenceByItemResponse> {
    const msg
      = req instanceof ReferenceByItemRequest
        ? req
        : new ReferenceByItemRequest(req)
    const results = utf8Decoder.decode(
      await this.contract.evaluateTransaction(
        'ReferenceByItem',
        msg.toJsonString(this.jsonWriteOptions),
      ),
    )
    return ReferenceByItemResponse.fromJsonString(results, {
      typeRegistry: this.registry,
    })
  }

  /**
   * @generated from rpc auth.common.ReferenceService.ReferenceByPartialKey
   */
  async referenceByPartialKey(
    req: PartialMessage<ReferenceByPartialKeyRequest>,
  ): Promise<ReferenceByPartialKeyResponse> {
    const msg
      = req instanceof ReferenceByPartialKeyRequest
        ? req
        : new ReferenceByPartialKeyRequest(req)
    const results = utf8Decoder.decode(
      await this.contract.evaluateTransaction(
        'ReferenceByPartialKey',
        msg.toJsonString(this.jsonWriteOptions),
      ),
    )
    return ReferenceByPartialKeyResponse.fromJsonString(results, {
      typeRegistry: this.registry,
    })
  }

  /**
   * @generated from rpc auth.common.ReferenceService.ReferenceCreate
   */
  async referenceCreate(
    req: PartialMessage<ReferenceCreateRequest>,
  ): Promise<ReferenceCreateResponse> {
    const msg
      = req instanceof ReferenceCreateRequest
        ? req
        : new ReferenceCreateRequest(req)
    const results = utf8Decoder.decode(
      await this.contract.submitTransaction(
        'ReferenceCreate',
        msg.toJsonString(this.jsonWriteOptions),
      ),
    )
    return ReferenceCreateResponse.fromJsonString(results, {
      typeRegistry: this.registry,
    })
  }

  /**
   * @generated from rpc auth.common.ReferenceService.ReferenceDelete
   */
  async referenceDelete(
    req: PartialMessage<ReferenceDeleteRequest>,
  ): Promise<ReferenceDeleteResponse> {
    const msg
      = req instanceof ReferenceDeleteRequest
        ? req
        : new ReferenceDeleteRequest(req)
    const results = utf8Decoder.decode(
      await this.contract.submitTransaction(
        'ReferenceDelete',
        msg.toJsonString(this.jsonWriteOptions),
      ),
    )
    return ReferenceDeleteResponse.fromJsonString(results, {
      typeRegistry: this.registry,
    })
  }
}
