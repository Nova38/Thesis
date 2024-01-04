#!/bin/bash
#
    docker stop peer0org1_identity_ccaas
    docker run --rm -d --name peer0org1_identity_ccaas                      --network fabric_test                     -e CHAINCODE_SERVER_ADDRESS=0.0.0.0:9999                     -e CHAINCODE_ID=identity_1.0.1:2af223ea67aced0554a6b68f9d94a1e4538e648e653d34123e1836a3cafa3bc2 -e CORE_CHAINCODE_ID_NAME=identity_1.0.1:2af223ea67aced0554a6b68f9d94a1e4538e648e653d34123e1836a3cafa3bc2                     -e AUTH_MODE=identity                     saacs_ccaas:latest

    docker stop peer0org2_identity_ccaas
    docker run --rm -d --name peer0org2_identity_ccaas                     --network fabric_test                     -e CHAINCODE_SERVER_ADDRESS=0.0.0.0:9999                     -e CHAINCODE_ID=identity_1.0.1:2af223ea67aced0554a6b68f9d94a1e4538e648e653d34123e1836a3cafa3bc2 -e CORE_CHAINCODE_ID_NAME=identity_1.0.1:2af223ea67aced0554a6b68f9d94a1e4538e648e653d34123e1836a3cafa3bc2                     -e AUTH_MODE=identity                     saacs_ccaas:latest
