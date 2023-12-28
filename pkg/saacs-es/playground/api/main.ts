import users from "./users.json" assert { type: 'json' };

import * as grpc from "@grpc/grpc-js";
import * as crypto from "crypto";
import { connect, Identity, signers } from "@hyperledger/fabric-gateway";
import { promises as fs } from "fs";
import { TextDecoder } from "util";
import { pb } from "../../src/index.js";
import { createRegistry } from "@bufbuild/protobuf";

async function  main() {
    console.log("Hello world");
    const GenericServiceClient = pb.common.generic.GenericServiceClient;

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



        let x = await service.getCurrentUser(true);
        console.log(x);

        // service.create()
    } finally {
        gateway.close();
        client.close();
    }

}
