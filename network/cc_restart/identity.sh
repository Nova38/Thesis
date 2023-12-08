#!/bin/bash
#
    docker stop peer0org1_identity_ccaas
    docker run --rm -d --name peer0org1_identity_ccaas                      --network fabric_test                     -e CHAINCODE_SERVER_ADDRESS=0.0.0.0:9999                     -e CHAINCODE_ID=identity_1.0.1:5ef6fca1fc146e0d0c3d79f1116cd24f905e8805394ee8f8c4f0dfd0896ac042 -e CORE_CHAINCODE_ID_NAME=identity_1.0.1:5ef6fca1fc146e0d0c3d79f1116cd24f905e8805394ee8f8c4f0dfd0896ac042                     -e AUTH_MODE=identity                     saacs_ccaas:latest

    docker stop peer0org2_identity_ccaas
    docker run --rm -d --name peer0org2_identity_ccaas --network fabric_test -e CHAINCODE_SERVER_ADDRESS=0.0.0.0:9999 -e CHAINCODE_ID=identity_1.0.1:5ef6fca1fc146e0d0c3d79f1116cd24f905e8805394ee8f8c4f0dfd0896ac042 -e CORE_CHAINCODE_ID_NAME=identity_1.0.1:5ef6fca1fc146e0d0c3d79f1116cd24f905e8805394ee8f8c4f0dfd0896ac042 -e AUTH_MODE=identity saacs_ccaas:latest
