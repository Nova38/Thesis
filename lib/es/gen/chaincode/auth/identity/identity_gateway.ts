// @generated by protoc-gen-reg v1 with parameter "target=ts"
// @generated from file chaincode/auth/identity/identity.proto (package chaincode.auth.identity, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import { Contract } from "@hyperledger/fabric-gateway";
import type { JsonValue } from "@bufbuild/protobuf";
import { IMessageTypeRegistry, JsonWriteStringOptions } from "@bufbuild/protobuf";
import { IdentityBootstrapRequest, IdentityBootstrapResponse } from "./identity_pb.js";

/**
 * @generated from service chaincode.auth.identity.IdentityService
 */
export class IdentityServiceClient {
    private contract: Contract;
    private jsonWriteOptions:Partial<JsonWriteStringOptions> = {};

    constructor(contract: Contract, registry: IMessageTypeRegistry) {
        this.contract = contract;
    }


    /**
     * @generated from rpc chaincode.auth.identity.IdentityService.IdentityBootstrap
     */
    async identityBootstrap(request: IdentityBootstrapRequest, evaluate: bool ): Promise< IdentityBootstrapResponse> {
        if (evaluate) {
            const promise = this.contract.evaluate(
                "IdentityBootstrap",
                $request.toJsonString(this.jsonWriteOptions)
            )
        } else {
            const promise = this.contract.submit(
                "IdentityBootstrap",
                $request.toJsonString(this.jsonWriteOptions)
            )
        }
        return promise.then(async (data) =>
             IdentityBootstrapResponse.fromJson(data as JsonValue)
        );
    }
}