#!/bin/bash
#
    docker stop peer0org1_roles_ccaas
    docker run --rm -d --name peer0org1_roles_ccaas                      --network fabric_test                     -e CHAINCODE_SERVER_ADDRESS=0.0.0.0:9999                     -e CHAINCODE_ID=roles_1.0.1:ffa5bbbc693a032adcf357af35cb6ddbb34d7b25d9578099f68e60c195a6c7df -e CORE_CHAINCODE_ID_NAME=roles_1.0.1:ffa5bbbc693a032adcf357af35cb6ddbb34d7b25d9578099f68e60c195a6c7df                     -e AUTH_MODE=roles saacs_ccaas:latest

    docker stop peer0org2_roles_ccaas
    docker run --rm -d --name peer0org2_roles_ccaas                     --network fabric_test                     -e CHAINCODE_SERVER_ADDRESS=0.0.0.0:9999                     -e CHAINCODE_ID=roles_1.0.1:ffa5bbbc693a032adcf357af35cb6ddbb34d7b25d9578099f68e60c195a6c7df -e CORE_CHAINCODE_ID_NAME=roles_1.0.1:ffa5bbbc693a032adcf357af35cb6ddbb34d7b25d9578099f68e60c195a6c7df -e AUTH_MODE=roles saacs_ccaas:latest
