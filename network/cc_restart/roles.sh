#!/bin/bash
#
    docker stop peer0org1_roles_ccaas
    docker run --rm -d --name peer0org1_roles_ccaas                      --network fabric_test                     -e CHAINCODE_SERVER_ADDRESS=0.0.0.0:9999                     -e CHAINCODE_ID=roles_1.0.1:15b1436d0ee895ac4ca6380ec15fe73af096242304cf3ccf1cb989e2152d175c -e CORE_CHAINCODE_ID_NAME=roles_1.0.1:15b1436d0ee895ac4ca6380ec15fe73af096242304cf3ccf1cb989e2152d175c                     -e AUTH_MODE=roles                     saacs_ccaas:latest

    docker stop peer0org2_roles_ccaas
    docker run --rm -d --name peer0org2_roles_ccaas                     --network fabric_test                     -e CHAINCODE_SERVER_ADDRESS=0.0.0.0:9999                     -e CHAINCODE_ID=roles_1.0.1:15b1436d0ee895ac4ca6380ec15fe73af096242304cf3ccf1cb989e2152d175c -e CORE_CHAINCODE_ID_NAME=roles_1.0.1:15b1436d0ee895ac4ca6380ec15fe73af096242304cf3ccf1cb989e2152d175c                     -e AUTH_MODE=roles                     saacs_ccaas:latest
