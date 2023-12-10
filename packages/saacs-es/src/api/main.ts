import users from "./users.json";

import * as grpc from "@grpc/grpc-js";
import * as crypto from "crypto";
import { connect, Identity, signers } from "@hyperledger/fabric-gateway";
import { promises as fs } from "fs";
import { TextDecoder } from "util";
import { GenericServiceClient } from "../gen/chaincode/common/generic_pb_gateway.js";
import { createRegistry } from "@bufbuild/protobuf";

const certPath = (path: string) => `../../../../${path}`;
const client = new grpc.Client(
    "gateway.example.org:1337",
    grpc.credentials.createInsecure(),
);

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

    const gateway = connect({ identity, signer, client });

    return gateway;

}

function BuildContract(contractName: string, contract: any) {
    const service = new GenericServiceClient(contract, createRegistry());
    return service;
}
// console.log();


const gateway = await BuildGateway(0);
try {
    const network = gateway.getNetwork("channelName");
    const contract = network.getContract("chaincodeName");

    const service = new GenericServiceClient(contract, createRegistry());

    // service.create()
} finally {
    gateway.close();
    client.close();
}
