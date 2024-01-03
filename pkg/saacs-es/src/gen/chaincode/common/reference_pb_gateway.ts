import { Contract } from "@hyperledger/fabric-gateway";
import { IMessageTypeRegistry, JsonWriteStringOptions } from "@bufbuild/protobuf";
import { ReferenceByItemRequest, ReferenceByItemResponse, ReferenceByPartialKeyRequest, ReferenceByPartialKeyResponse, ReferenceCreateRequest, ReferenceCreateResponse, ReferenceDeleteRequest, ReferenceDeleteResponse, ReferenceRequest, ReferenceResponse } from "./reference_pb.js";

const utf8Decoder = new TextDecoder();
/**
 * ════════════════════════════════ References ══════════════════════════════════
 * ──────────────────────────────── Query ────────────────────────────────────────
 *
 * @generated from service auth.common.ReferenceService
 */
export class ReferenceServiceClient {
    private contract: Contract;
    private jsonWriteOptions:Partial<JsonWriteStringOptions> = {};
    registry: IMessageTypeRegistry;

    constructor(contract: Contract, registry: IMessageTypeRegistry) {
        this.contract = contract;
        this.registry = registry;
        this.jsonWriteOptions.typeRegistry = registry
    }


    /**
     * @generated from rpc auth.common.ReferenceService.Reference
     */
    async reference(request: ReferenceRequest ): Promise<ReferenceResponse> {
            const results = utf8Decoder.decode(
                await this.contract.evaluateTransaction(
                "Reference",
                request.toJsonString(this.jsonWriteOptions)
            )
            )
            return ReferenceResponse.fromJsonString(results, {typeRegistry: this.registry});
    }

    /**
     * @generated from rpc auth.common.ReferenceService.ReferenceByItem
     */
    async referenceByItem(request: ReferenceByItemRequest ): Promise<ReferenceByItemResponse> {
            const results = utf8Decoder.decode(
                await this.contract.evaluateTransaction(
                "ReferenceByItem",
                request.toJsonString(this.jsonWriteOptions)
            )
            )
            return ReferenceByItemResponse.fromJsonString(results, {typeRegistry: this.registry});
    }

    /**
     * @generated from rpc auth.common.ReferenceService.ReferenceByPartialKey
     */
    async referenceByPartialKey(request: ReferenceByPartialKeyRequest ): Promise<ReferenceByPartialKeyResponse> {
            const results = utf8Decoder.decode(
                await this.contract.evaluateTransaction(
                "ReferenceByPartialKey",
                request.toJsonString(this.jsonWriteOptions)
            )
            )
            return ReferenceByPartialKeyResponse.fromJsonString(results, {typeRegistry: this.registry});
    }

    /**
     * @generated from rpc auth.common.ReferenceService.ReferenceCreate
     */
    async referenceCreate(request: ReferenceCreateRequest ): Promise<ReferenceCreateResponse> {
            const results = utf8Decoder.decode(
                await this.contract.submitTransaction(
                "ReferenceCreate",
                request.toJsonString(this.jsonWriteOptions)
            )
            )
            return ReferenceCreateResponse.fromJsonString(results, {typeRegistry: this.registry});
    }

    /**
     * @generated from rpc auth.common.ReferenceService.ReferenceDelete
     */
    async referenceDelete(request: ReferenceDeleteRequest ): Promise<ReferenceDeleteResponse> {
            const results = utf8Decoder.decode(
                await this.contract.submitTransaction(
                "ReferenceDelete",
                request.toJsonString(this.jsonWriteOptions)
            )
            )
            return ReferenceDeleteResponse.fromJsonString(results, {typeRegistry: this.registry});
    }
}
