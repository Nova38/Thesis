#!/bin/bash
#
    docker stop peer0org1_roles_ccaas
    docker run --rm -d --name peer0org1_roles_ccaas                      --network fabric_test                     -e CHAINCODE_SERVER_ADDRESS=0.0.0.0:9999                     -e CHAINCODE_ID=roles_1.0.1:07de8a32693d71f568747648397aa7c6c58331c2e6a8033b76ce511b689f4cbf -e CORE_CHAINCODE_ID_NAME=roles_1.0.1:07de8a32693d71f568747648397aa7c6c58331c2e6a8033b76ce511b689f4cbf                     -e AUTH_MODE=roles                     saacs_ccaas:latest

    docker stop peer0org2_roles_ccaas
    docker run --rm -d --name peer0org2_roles_ccaas                     --network fabric_test                     -e CHAINCODE_SERVER_ADDRESS=0.0.0.0:9999                     -e CHAINCODE_ID=roles_1.0.1:07de8a32693d71f568747648397aa7c6c58331c2e6a8033b76ce511b689f4cbf -e CORE_CHAINCODE_ID_NAME=roles_1.0.1:07de8a32693d71f568747648397aa7c6c58331c2e6a8033b76ce511b689f4cbf                     -e AUTH_MODE=roles                     saacs_ccaas:latest
