#!/bin/bash
#
    docker stop peer0org1_roles_ccaas
    docker run --rm -d --name peer0org1_roles_ccaas                      --network fabric_test                     -e CHAINCODE_SERVER_ADDRESS=0.0.0.0:9999                     -e CHAINCODE_ID=roles_1.0.1:8d9fe717bc0c6afcebdea667f5cbd761c31c61050c64f07247cab888e05264b3 -e CORE_CHAINCODE_ID_NAME=roles_1.0.1:8d9fe717bc0c6afcebdea667f5cbd761c31c61050c64f07247cab888e05264b3                     -e AUTH_MODE=roles                     saacs_ccaas:latest

    docker stop peer0org2_roles_ccaas
    docker run --rm -d --name peer0org2_roles_ccaas                     --network fabric_test                     -e CHAINCODE_SERVER_ADDRESS=0.0.0.0:9999                     -e CHAINCODE_ID=roles_1.0.1:8d9fe717bc0c6afcebdea667f5cbd761c31c61050c64f07247cab888e05264b3 -e CORE_CHAINCODE_ID_NAME=roles_1.0.1:8d9fe717bc0c6afcebdea667f5cbd761c31c61050c64f07247cab888e05264b3                     -e AUTH_MODE=roles                     saacs_ccaas:latest
