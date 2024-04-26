import { Contract } from "@hyperledger/fabric-gateway";
import type { PartialMessage } from "@bufbuild/protobuf";
import { IMessageTypeRegistry, JsonWriteStringOptions } from "@bufbuild/protobuf";
import { AuthorizeOperationRequest, AuthorizeOperationResponse, BootstrapRequest, BootstrapResponse, GetCollectionsListRequest, GetCollectionsListResponse, GetCurrentUserRequest, GetCurrentUserResponse } from "./utils_pb.js";

const utf8Decoder = new TextDecoder();
/**
 * @generated from service saacs.chaincode.v0.UtilsService
 */
export class UtilsServiceClient  {
    private contract: Contract;
    private jsonWriteOptions:Partial<JsonWriteStringOptions> = {};
    registry: IMessageTypeRegistry;

    constructor(contract: Contract, registry: IMessageTypeRegistry) {
        this.contract = contract;
        this.registry = registry;
        this.jsonWriteOptions.typeRegistry = registry
    }


    /**
     * @generated from rpc saacs.chaincode.v0.UtilsService.GetCurrentUser
     */
  async getCurrentUser(req: PartialMessage<GetCurrentUserRequest>): Promise<GetCurrentUserResponse> {
    const msg = req instanceof GetCurrentUserRequest ? req : new GetCurrentUserRequest(req);
    const results = utf8Decoder.decode(
      await this.contract.evaluateTransaction(
        "GetCurrentUser",
        msg.toJsonString(this.jsonWriteOptions)
      ))
    return GetCurrentUserResponse.fromJsonString(results, {typeRegistry: this.registry});
    }


    /**
     * @generated from rpc saacs.chaincode.v0.UtilsService.Bootstrap
     */
  async bootstrap(req: PartialMessage<BootstrapRequest>): Promise<BootstrapResponse> {
    const msg = req instanceof BootstrapRequest ? req : new BootstrapRequest(req);
    const results = utf8Decoder.decode(
      await this.contract.submitTransaction(
        "Bootstrap",
        msg.toJsonString(this.jsonWriteOptions)
      ))
    return BootstrapResponse.fromJsonString(results, {typeRegistry: this.registry});
    }


    /**
     * @generated from rpc saacs.chaincode.v0.UtilsService.AuthorizeOperation
     */
  async authorizeOperation(req: PartialMessage<AuthorizeOperationRequest>): Promise<AuthorizeOperationResponse> {
    const msg = req instanceof AuthorizeOperationRequest ? req : new AuthorizeOperationRequest(req);
    const results = utf8Decoder.decode(
      await this.contract.submitTransaction(
        "AuthorizeOperation",
        msg.toJsonString(this.jsonWriteOptions)
      ))
    return AuthorizeOperationResponse.fromJsonString(results, {typeRegistry: this.registry});
    }


    /**
     * @generated from rpc saacs.chaincode.v0.UtilsService.GetCollectionsList
     */
  async getCollectionsList(req: PartialMessage<GetCollectionsListRequest>): Promise<GetCollectionsListResponse> {
    const msg = req instanceof GetCollectionsListRequest ? req : new GetCollectionsListRequest(req);
    const results = utf8Decoder.decode(
      await this.contract.evaluateTransaction(
        "GetCollectionsList",
        msg.toJsonString(this.jsonWriteOptions)
      ))
    return GetCollectionsListResponse.fromJsonString(results, {typeRegistry: this.registry});
    }

}
