import * as hlf from '../dist/index.mjs'


;

import * as grpc from "@grpc/grpc-js";
import * as crypto from "crypto";
import { connect, Identity, signers } from "@hyperledger/fabric-gateway";
import { promises as fs } from "fs";
import { TextDecoder } from "util";
import { createRegistry } from "@bufbuild/protobuf";

async function  main() {
    console.log("Hello world");
    const GenericServiceClient = hlf.pb.common.generic.GenericServiceClient;
}


main()
