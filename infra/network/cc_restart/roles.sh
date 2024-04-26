#!/bin/bash
#
    docker stop peer0org1_roles_ccaas
    docker run --rm -d --name peer0org1_roles_ccaas                      --network fabric_test                     -e CHAINCODE_SERVER_ADDRESS=0.0.0.0:9999                     -e CHAINCODE_ID=roles_1.0.1:bf5142bd61da768bfe40c7cab73da20e12a90825e66f9b09420d43dc5e1bf469 -e CORE_CHAINCODE_ID_NAME=roles_1.0.1:bf5142bd61da768bfe40c7cab73da20e12a90825e66f9b09420d43dc5e1bf469                     -e AUTH_MODE=roles                     saacs_ccaas:latest

    docker stop peer0org2_roles_ccaas
    docker run --rm -d --name peer0org2_roles_ccaas                     --network fabric_test                     -e CHAINCODE_SERVER_ADDRESS=0.0.0.0:9999                     -e CHAINCODE_ID=roles_1.0.1:bf5142bd61da768bfe40c7cab73da20e12a90825e66f9b09420d43dc5e1bf469 -e CORE_CHAINCODE_ID_NAME=roles_1.0.1:bf5142bd61da768bfe40c7cab73da20e12a90825e66f9b09420d43dc5e1bf469                     -e AUTH_MODE=roles                     saacs_ccaas:latest
