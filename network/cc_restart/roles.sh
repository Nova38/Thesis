#!/bin/bash
#
    docker stop peer0org1_roles_ccaas
    docker run --rm -d --name peer0org1_roles_ccaas                      --network fabric_test                     -e CHAINCODE_SERVER_ADDRESS=0.0.0.0:9999                     -e CHAINCODE_ID=roles_1.0.1:557391d8d56c8d5ce82bcf8005a00a71a3982f587fefa06f21f371a9a84d22bd -e CORE_CHAINCODE_ID_NAME=roles_1.0.1:557391d8d56c8d5ce82bcf8005a00a71a3982f587fefa06f21f371a9a84d22bd                     -e AUTH_MODE=roles                     saacs_ccaas:latest

    docker stop peer0org2_roles_ccaas
    docker run --rm -d --name peer0org2_roles_ccaas                     --network fabric_test                     -e CHAINCODE_SERVER_ADDRESS=0.0.0.0:9999                     -e CHAINCODE_ID=roles_1.0.1:557391d8d56c8d5ce82bcf8005a00a71a3982f587fefa06f21f371a9a84d22bd -e CORE_CHAINCODE_ID_NAME=roles_1.0.1:557391d8d56c8d5ce82bcf8005a00a71a3982f587fefa06f21f371a9a84d22bd                     -e AUTH_MODE=roles                     saacs_ccaas:latest
