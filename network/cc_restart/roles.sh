#!/bin/bash
#
    docker stop peer0org1_roles_ccaas
    docker run --rm -d --name peer0org1_roles_ccaas                      --network fabric_test                     -e CHAINCODE_SERVER_ADDRESS=0.0.0.0:9999                     -e CHAINCODE_ID=roles_1.0.1:e34f56ff150d7f448c3b3e278236a3ab8c5aee94e7a6c833d54fffe4481fcfd9 -e CORE_CHAINCODE_ID_NAME=roles_1.0.1:e34f56ff150d7f448c3b3e278236a3ab8c5aee94e7a6c833d54fffe4481fcfd9                     -e AUTH_MODE=roles                     saacs_ccaas:latest

    docker stop peer0org2_roles_ccaas
    docker run --rm -d --name peer0org2_roles_ccaas                     --network fabric_test                     -e CHAINCODE_SERVER_ADDRESS=0.0.0.0:9999                     -e CHAINCODE_ID=roles_1.0.1:e34f56ff150d7f448c3b3e278236a3ab8c5aee94e7a6c833d54fffe4481fcfd9 -e CORE_CHAINCODE_ID_NAME=roles_1.0.1:e34f56ff150d7f448c3b3e278236a3ab8c5aee94e7a6c833d54fffe4481fcfd9                     -e AUTH_MODE=roles                     saacs_ccaas:latest
