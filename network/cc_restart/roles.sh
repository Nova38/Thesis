#!/bin/bash
#
    docker stop peer0org1_roles_ccaas
    docker run -d --name peer0org1_roles_ccaas                      --network fabric_test                     -e CHAINCODE_SERVER_ADDRESS=0.0.0.0:9999                     -e CHAINCODE_ID=roles_1.0.1:0a7c5ab3e16d83b8e9387fab83e345b8a7cd53aa0d31dc2b5ed1060757afa647 -e CORE_CHAINCODE_ID_NAME=roles_1.0.1:0a7c5ab3e16d83b8e9387fab83e345b8a7cd53aa0d31dc2b5ed1060757afa647                     -e AUTH_MODE=roles                     saacs_ccaas:latest

    docker stop peer0org2_roles_ccaas
    docker run -d --name peer0org2_roles_ccaas                     --network fabric_test                     -e CHAINCODE_SERVER_ADDRESS=0.0.0.0:9999                     -e CHAINCODE_ID=roles_1.0.1:0a7c5ab3e16d83b8e9387fab83e345b8a7cd53aa0d31dc2b5ed1060757afa647 -e CORE_CHAINCODE_ID_NAME=roles_1.0.1:0a7c5ab3e16d83b8e9387fab83e345b8a7cd53aa0d31dc2b5ed1060757afa647                     -e AUTH_MODE=roles                     saacs_ccaas:latest
