#!/bin/bash
#
    docker stop peer0org1_roles_ccaas
    docker run --rm -d --name peer0org1_roles_ccaas                      --network fabric_test                     -e CHAINCODE_SERVER_ADDRESS=0.0.0.0:9999                     -e CHAINCODE_ID=roles_1.0.1:015b075cd1ba7599b07d09c5efd4db6ef50af658b7a84bc033754c6da7dcfee4 -e CORE_CHAINCODE_ID_NAME=roles_1.0.1:015b075cd1ba7599b07d09c5efd4db6ef50af658b7a84bc033754c6da7dcfee4                     -e AUTH_MODE=roles                     saacs_ccaas:latest

    docker stop peer0org2_roles_ccaas
    docker run --rm -d --name peer0org2_roles_ccaas                     --network fabric_test                     -e CHAINCODE_SERVER_ADDRESS=0.0.0.0:9999                     -e CHAINCODE_ID=roles_1.0.1:015b075cd1ba7599b07d09c5efd4db6ef50af658b7a84bc033754c6da7dcfee4 -e CORE_CHAINCODE_ID_NAME=roles_1.0.1:015b075cd1ba7599b07d09c5efd4db6ef50af658b7a84bc033754c6da7dcfee4                     -e AUTH_MODE=roles                     saacs_ccaas:latest
