#!/bin/bash
#
    docker stop peer0org1_identity_ccaas
    docker run --rm -d --name peer0org1_identity_ccaas                      --network fabric_test                     -e CHAINCODE_SERVER_ADDRESS=0.0.0.0:9999                     -e CHAINCODE_ID=identity_1.0.1:bd3542e10c21718d3acb4c15f9ade733bb6ce3e3c0b64cafe5f379d0dace0e04 -e CORE_CHAINCODE_ID_NAME=identity_1.0.1:bd3542e10c21718d3acb4c15f9ade733bb6ce3e3c0b64cafe5f379d0dace0e04                     -e AUTH_MODE=identity                     saacs_ccaas:latest

    docker stop peer0org2_identity_ccaas
    docker run --rm -d --name peer0org2_identity_ccaas                     --network fabric_test                     -e CHAINCODE_SERVER_ADDRESS=0.0.0.0:9999                     -e CHAINCODE_ID=identity_1.0.1:bd3542e10c21718d3acb4c15f9ade733bb6ce3e3c0b64cafe5f379d0dace0e04 -e CORE_CHAINCODE_ID_NAME=identity_1.0.1:bd3542e10c21718d3acb4c15f9ade733bb6ce3e3c0b64cafe5f379d0dace0e04                     -e AUTH_MODE=identity                     saacs_ccaas:latest
