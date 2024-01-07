#!/bin/bash
#
    docker stop peer0org1_roles_ccaas
    docker run --rm -d --name peer0org1_roles_ccaas                      --network fabric_test                     -e CHAINCODE_SERVER_ADDRESS=0.0.0.0:9999                     -e CHAINCODE_ID=roles_1.0.1:b523386d655ba46a7124a813e36a4ac5bafc5aac033edaff111f5dcd0b008b14 -e CORE_CHAINCODE_ID_NAME=roles_1.0.1:b523386d655ba46a7124a813e36a4ac5bafc5aac033edaff111f5dcd0b008b14                     -e AUTH_MODE=roles                     saacs_ccaas:latest

    docker stop peer0org2_roles_ccaas
    docker run --rm -d --name peer0org2_roles_ccaas                     --network fabric_test                     -e CHAINCODE_SERVER_ADDRESS=0.0.0.0:9999                     -e CHAINCODE_ID=roles_1.0.1:b523386d655ba46a7124a813e36a4ac5bafc5aac033edaff111f5dcd0b008b14 -e CORE_CHAINCODE_ID_NAME=roles_1.0.1:b523386d655ba46a7124a813e36a4ac5bafc5aac033edaff111f5dcd0b008b14                     -e AUTH_MODE=roles                     saacs_ccaas:latest
