import { createPrivateKey } from 'crypto';
import * as grpc from '@grpc/grpc-js';
import {
  connect,
  Contract,
  Gateway,
  Network,
  signers,
} from '@hyperledger/fabric-gateway';

// import {GenericService} from "../../../lib/es/gen/connect/chaincode/auth/common/generic_connect"
// import {gen} from "es"


const chaincodeConfig = {
    peerEndpoint: 'peer0-org1.129.237.123.11.nip.io:443',
    channel: 'demo',
    contract: 'biochain',
  };
const utf8Decoder = new TextDecoder();

let client: grpc.Client;

