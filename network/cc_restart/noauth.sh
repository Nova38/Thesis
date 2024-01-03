#!/bin/bash
#
    docker stop peer0org1_noauth_ccaas
    docker run --rm -d --name peer0org1_noauth_ccaas                      --network fabric_test                     -e CHAINCODE_SERVER_ADDRESS=0.0.0.0:9999                     -e CHAINCODE_ID=noauth_1.0.1:54bf5d6e88eb59dc42e3682cabf142fee653e84b8c094796915def24b09a75f4 -e CORE_CHAINCODE_ID_NAME=noauth_1.0.1:54bf5d6e88eb59dc42e3682cabf142fee653e84b8c094796915def24b09a75f4                     -e AUTH_MODE=noauth                     saacs_ccaas:latest

    docker stop peer0org2_noauth_ccaas
    docker run --rm -d --name peer0org2_noauth_ccaas                     --network fabric_test                     -e CHAINCODE_SERVER_ADDRESS=0.0.0.0:9999                     -e CHAINCODE_ID=noauth_1.0.1:54bf5d6e88eb59dc42e3682cabf142fee653e84b8c094796915def24b09a75f4 -e CORE_CHAINCODE_ID_NAME=noauth_1.0.1:54bf5d6e88eb59dc42e3682cabf142fee653e84b8c094796915def24b09a75f4                     -e AUTH_MODE=noauth                     saacs_ccaas:latest
