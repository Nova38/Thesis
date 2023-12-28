#!/bin/bash
#
    docker stop peer0org1_noauth-no-sub_ccaas
    docker run --rm -d --name peer0org1_noauth-no-sub_ccaas                      --network fabric_test                     -e CHAINCODE_SERVER_ADDRESS=0.0.0.0:9999                     -e CHAINCODE_ID=noauth-no-sub_1.0.1:e66d808ef6f76020c20dde89690b453a75147ff332fa2bbc0e53e677f0cb2321 -e CORE_CHAINCODE_ID_NAME=noauth-no-sub_1.0.1:e66d808ef6f76020c20dde89690b453a75147ff332fa2bbc0e53e677f0cb2321                     -e AUTH_MODE=noauth-no-sub                     saacs_ccaas:latest

    docker stop peer0org2_noauth-no-sub_ccaas
    docker run --rm -d --name peer0org2_noauth-no-sub_ccaas                     --network fabric_test                     -e CHAINCODE_SERVER_ADDRESS=0.0.0.0:9999                     -e CHAINCODE_ID=noauth-no-sub_1.0.1:e66d808ef6f76020c20dde89690b453a75147ff332fa2bbc0e53e677f0cb2321 -e CORE_CHAINCODE_ID_NAME=noauth-no-sub_1.0.1:e66d808ef6f76020c20dde89690b453a75147ff332fa2bbc0e53e677f0cb2321                     -e AUTH_MODE=noauth-no-sub                     saacs_ccaas:latest
