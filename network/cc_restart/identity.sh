#!/bin/bash
#
    docker stop peer0org1_identity_ccaas
    docker run --rm -d --name peer0org1_identity_ccaas                      --network fabric_test                     -e CHAINCODE_SERVER_ADDRESS=0.0.0.0:9999                     -e CHAINCODE_ID=identity_1.0.1:6a048a51bfc1d85c1c46a884a266174c522f22debccbd712c70a959e293b5a08 -e CORE_CHAINCODE_ID_NAME=identity_1.0.1:6a048a51bfc1d85c1c46a884a266174c522f22debccbd712c70a959e293b5a08                     -e AUTH_MODE=identity                     saacs_ccaas:latest

    docker stop peer0org2_identity_ccaas
    docker run --rm -d --name peer0org2_identity_ccaas                     --network fabric_test                     -e CHAINCODE_SERVER_ADDRESS=0.0.0.0:9999                     -e CHAINCODE_ID=identity_1.0.1:6a048a51bfc1d85c1c46a884a266174c522f22debccbd712c70a959e293b5a08 -e CORE_CHAINCODE_ID_NAME=identity_1.0.1:6a048a51bfc1d85c1c46a884a266174c522f22debccbd712c70a959e293b5a08                     -e AUTH_MODE=identity                     saacs_ccaas:latest
