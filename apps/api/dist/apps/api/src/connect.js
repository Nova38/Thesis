"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
var chaincodeConfig = {
    peerEndpoint: 'peer0-org1.129.237.123.11.nip.io:443',
    channel: 'demo',
    contract: 'biochain',
};
var utf8Decoder = new TextDecoder();
var client;
