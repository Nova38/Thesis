#!/bin/bash
#
    docker stop peer0org1_roles_ccaas
    docker run --rm -d --name peer0org1_roles_ccaas                      --network fabric_test                     -e CHAINCODE_SERVER_ADDRESS=0.0.0.0:9999                     -e CHAINCODE_ID=roles_1.0.1:8453c5b6a14706822d093cb70f0520ef0ad8e47a21f2ff1d2140ed893cf5b4e5 -e CORE_CHAINCODE_ID_NAME=roles_1.0.1:8453c5b6a14706822d093cb70f0520ef0ad8e47a21f2ff1d2140ed893cf5b4e5                     -e AUTH_MODE=roles                     saacs_ccaas:latest

    docker stop peer0org2_roles_ccaas
    docker run --rm -d --name peer0org2_roles_ccaas                     --network fabric_test                     -e CHAINCODE_SERVER_ADDRESS=0.0.0.0:9999                     -e CHAINCODE_ID=roles_1.0.1:8453c5b6a14706822d093cb70f0520ef0ad8e47a21f2ff1d2140ed893cf5b4e5 -e CORE_CHAINCODE_ID_NAME=roles_1.0.1:8453c5b6a14706822d093cb70f0520ef0ad8e47a21f2ff1d2140ed893cf5b4e5                     -e AUTH_MODE=roles                     saacs_ccaas:latest
