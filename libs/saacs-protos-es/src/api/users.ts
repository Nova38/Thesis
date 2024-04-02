import * as grpc from "@grpc/grpc-js";

import remote from "./remote.json" assert { type: "json" };
import { ofetch } from "ofetch";
import { createRegistry, IMessageTypeRegistry } from "@bufbuild/protobuf";
import { GenericServiceClient } from "../gen/chaincode/common/generic_pb_gateway.js";
// import { auth, ccbio, common, sample } from "../index.js";
import {
    connect,
    ConnectOptions,
    Identity,
    signers,
} from "@hyperledger/fabric-gateway";
import { auth, objects } from "../gen/auth/v1/index.js";
import { generic, reference } from "../gen/chaincode/common/index.js";
import { ccbio } from "../gen/index.js";
import { sample } from "../gen/index.js";
import * as crypto from "crypto";

import { promises as fs } from "fs";

//

const baseurl = "https://api-biochain.ittc.ku.edu/api";

async function Signup(
    username: string,
    password: string,
    credentials: string,
    key: string,
    mspId: string,
): Promise<string> {
    const res = await ofetch(`${baseurl + "/auth/register"}`, {
        method: "POST",
        body: JSON.stringify({
            username: username,
            password: password,
            credentials: credentials,
            key: key,
            mspId: mspId,
        }),
    });

    return res;
}

async function setupUsers() {
    const users = remote.organizations.Org1MSP.users;
    console.log(users);
    for (const [k, value] of Object.entries(users)) {
        // console.log(`${key}: ${value}`);

        const username = k;
        const password = value.password;
        const credentials = value.cert.pem;
        const key = value.key.pem;
        const mspId = "Org1MSP";

        // const res = await console.log(
        //     username,
        //     password,
        //     credentials,
        //     key,
        //     mspId,
        // );
        const res = await Signup(username, password, credentials, key, mspId);
        console.log(res);
    }
}

async function main() {
    await setupUsers();
}

main();
