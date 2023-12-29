#!/bin/bash
#
    docker stop peer0org1_roles_ccaas
    docker run --rm -d --name peer0org1_roles_ccaas                      --network fabric_test                     -e CHAINCODE_SERVER_ADDRESS=0.0.0.0:9999                     -e CHAINCODE_ID=roles_1.0.1:bf087b459301f74668e4bb5726a77d9604d64d27b948ef5fb9120be2e348aace -e CORE_CHAINCODE_ID_NAME=roles_1.0.1:bf087b459301f74668e4bb5726a77d9604d64d27b948ef5fb9120be2e348aace                     -e AUTH_MODE=roles                     saacs_ccaas:latest

    docker stop peer0org2_roles_ccaas
    docker run --rm -d --name peer0org2_roles_ccaas                     --network fabric_test                     -e CHAINCODE_SERVER_ADDRESS=0.0.0.0:9999                     -e CHAINCODE_ID=roles_1.0.1:bf087b459301f74668e4bb5726a77d9604d64d27b948ef5fb9120be2e348aace -e CORE_CHAINCODE_ID_NAME=roles_1.0.1:bf087b459301f74668e4bb5726a77d9604d64d27b948ef5fb9120be2e348aace                     -e AUTH_MODE=roles                     saacs_ccaas:latest
