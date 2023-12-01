// @generated by protoc-gen-reg v1 with parameter "target=ts"
// @generated from file chaincode/auth/common/reference.proto (package auth.common, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import { Contract } from "@hyperledger/fabric-gateway";
import type { JsonValue } from "@bufbuild/protobuf";
import { IMessageTypeRegistry, JsonWriteStringOptions } from "@bufbuild/protobuf";
import { ReferenceByItemRequest, ReferenceByItemResponse, ReferenceByPartialKeyRequest, ReferenceByPartialKeyResponse, ReferenceCreateRequest, ReferenceCreateResponse, ReferenceDeleteRequest, ReferenceDeleteResponse, ReferenceRequest, ReferenceResponse } from "./reference_pb.js";

/**
 * ════════════════════════════════ References ══════════════════════════════════
 * ──────────────────────────────── Query ────────────────────────────────────────
 *
 * @generated from service auth.common.ReferenceService
 */
export class ReferenceServiceClient {
    private contract: Contract;
    private jsonWriteOptions:Partial<JsonWriteStringOptions> = {};

    constructor(contract: Contract, registry: IMessageTypeRegistry) {
        this.contract = contract;
    }


    /**
     * @generated from rpc auth.common.ReferenceService.Reference
     */
    async reference(request: ReferenceRequest, evaluate: bool ): Promise< ReferenceResponse> {
        if (evaluate) {
            const promise = this.contract.evaluate(
                "Reference",
                $request.toJsonString(this.jsonWriteOptions)
            )
        } else {
            const promise = this.contract.submit(
                "Reference",
                $request.toJsonString(this.jsonWriteOptions)
            )
        }
        return promise.then(async (data) =>
             ReferenceResponse.fromJson(data as JsonValue)
        );
    }

    /**
     * @generated from rpc auth.common.ReferenceService.ReferenceByItem
     */
    async referenceByItem(request: ReferenceByItemRequest, evaluate: bool ): Promise< ReferenceByItemResponse> {
        if (evaluate) {
            const promise = this.contract.evaluate(
                "ReferenceByItem",
                $request.toJsonString(this.jsonWriteOptions)
            )
        } else {
            const promise = this.contract.submit(
                "ReferenceByItem",
                $request.toJsonString(this.jsonWriteOptions)
            )
        }
        return promise.then(async (data) =>
             ReferenceByItemResponse.fromJson(data as JsonValue)
        );
    }

    /**
     * @generated from rpc auth.common.ReferenceService.ReferenceByPartialKey
     */
    async referenceByPartialKey(request: ReferenceByPartialKeyRequest, evaluate: bool ): Promise< ReferenceByPartialKeyResponse> {
        if (evaluate) {
            const promise = this.contract.evaluate(
                "ReferenceByPartialKey",
                $request.toJsonString(this.jsonWriteOptions)
            )
        } else {
            const promise = this.contract.submit(
                "ReferenceByPartialKey",
                $request.toJsonString(this.jsonWriteOptions)
            )
        }
        return promise.then(async (data) =>
             ReferenceByPartialKeyResponse.fromJson(data as JsonValue)
        );
    }

    /**
     * @generated from rpc auth.common.ReferenceService.ReferenceCreate
     */
    async referenceCreate(request: ReferenceCreateRequest, evaluate: bool ): Promise< ReferenceCreateResponse> {
        if (evaluate) {
            const promise = this.contract.evaluate(
                "ReferenceCreate",
                $request.toJsonString(this.jsonWriteOptions)
            )
        } else {
            const promise = this.contract.submit(
                "ReferenceCreate",
                $request.toJsonString(this.jsonWriteOptions)
            )
        }
        return promise.then(async (data) =>
             ReferenceCreateResponse.fromJson(data as JsonValue)
        );
    }

    /**
     * @generated from rpc auth.common.ReferenceService.ReferenceDelete
     */
    async referenceDelete(request: ReferenceDeleteRequest, evaluate: bool ): Promise< ReferenceDeleteResponse> {
        if (evaluate) {
            const promise = this.contract.evaluate(
                "ReferenceDelete",
                $request.toJsonString(this.jsonWriteOptions)
            )
        } else {
            const promise = this.contract.submit(
                "ReferenceDelete",
                $request.toJsonString(this.jsonWriteOptions)
            )
        }
        return promise.then(async (data) =>
             ReferenceDeleteResponse.fromJson(data as JsonValue)
        );
    }
}