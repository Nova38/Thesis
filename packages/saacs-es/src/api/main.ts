import users from "./users.json" assert { type: "json" };
import connectionProfile from "../../../../network/organizations/peerOrganizations/org1.example.com/connection-org1.json" assert { type: "json" };
import * as grpc from "@grpc/grpc-js";
import * as crypto from "crypto";
import { connect, Identity, signers } from "@hyperledger/fabric-gateway";
import { promises as fs } from "fs";
import { TextDecoder } from "util";
import { GenericServiceClient } from "../gen/chaincode/common/generic_pb_gateway.js";
import { createRegistry } from "@bufbuild/protobuf";

const certPath = (path: string) => `../../${path}`;

async function newGRPCClient() {
    const peer = connectionProfile.peers["peer0.org1.example.com"];
    const tlsCredentials = grpc.credentials.createSsl(
        Buffer.from(peer.tlsCACerts.pem),
    );

    return new grpc.Client(peer.url.split("//")[1], tlsCredentials, {
        "grpc.ssl_target_name_override":
            peer.grpcOptions["ssl-target-name-override"],
    });
}

const client = await newGRPCClient();

async function BuildGateway(userIdex: number) {
    const credentials = await fs.readFile(
        certPath(users.identities[0].clientSignedCert),
    );
    console.log(credentials.toString());
    const identity: Identity = { mspId: "Org1MSP", credentials };

    const privateKeyPem = await fs.readFile(
        certPath(users.identities[0].clientPrivateKey),
    );
    const privateKey = crypto.createPrivateKey(privateKeyPem);
    const signer = signers.newPrivateKeySigner(privateKey);

    console.log("client", client);

    const gateway = connect({ client, identity, signer });

    return gateway;
}

function BuildContract(contractName: string, contract: any) {
    const service = new GenericServiceClient(contract, createRegistry());
    return service;
}

const gateway = await BuildGateway(0);
try {
    const network = gateway.getNetwork("mychannel");
    const contract = network.getContract("roles");

    const service = new GenericServiceClient(contract, createRegistry());

    const result = await service.getCurrentUser();

    console.log(result);

    // service.create()
} finally {
    gateway.close();
    client.close();
}
