#!/bin/bash
#
    docker stop peer0org1_identity_ccaas
    docker run --rm -d --name peer0org1_identity_ccaas                      --network fabric_test                     -e CHAINCODE_SERVER_ADDRESS=0.0.0.0:9999                     -e CHAINCODE_ID=identity_1.0.1:d6654f37bd6197d946df60401b93efda967f16bf387fb046a4047a8e6c0b5bbd -e CORE_CHAINCODE_ID_NAME=identity_1.0.1:d6654f37bd6197d946df60401b93efda967f16bf387fb046a4047a8e6c0b5bbd                     -e AUTH_MODE=identity                     saacs_ccaas:latest

    docker stop peer0org2_identity_ccaas
    docker run --rm -d --name peer0org2_identity_ccaas                     --network fabric_test                     -e CHAINCODE_SERVER_ADDRESS=0.0.0.0:9999                     -e CHAINCODE_ID=identity_1.0.1:d6654f37bd6197d946df60401b93efda967f16bf387fb046a4047a8e6c0b5bbd -e CORE_CHAINCODE_ID_NAME=identity_1.0.1:d6654f37bd6197d946df60401b93efda967f16bf387fb046a4047a8e6c0b5bbd                     -e AUTH_MODE=identity                     saacs_ccaas:latest
