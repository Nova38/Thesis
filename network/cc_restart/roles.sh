#!/bin/bash
#
    docker stop peer0org1_roles_ccaas
    docker run --rm -d --name peer0org1_roles_ccaas                      --network fabric_test                     -e CHAINCODE_SERVER_ADDRESS=0.0.0.0:9999                     -e CHAINCODE_ID=roles_1.0.1:e844f6bb1845b27efd6d42f9f646fb09f27ea6d3de5b20098ccf8b7454a23a2e -e CORE_CHAINCODE_ID_NAME=roles_1.0.1:e844f6bb1845b27efd6d42f9f646fb09f27ea6d3de5b20098ccf8b7454a23a2e                     -e AUTH_MODE=roles                     saacs_ccaas:latest

    docker stop peer0org2_roles_ccaas
    docker run --rm -d --name peer0org2_roles_ccaas                     --network fabric_test                     -e CHAINCODE_SERVER_ADDRESS=0.0.0.0:9999                     -e CHAINCODE_ID=roles_1.0.1:e844f6bb1845b27efd6d42f9f646fb09f27ea6d3de5b20098ccf8b7454a23a2e -e CORE_CHAINCODE_ID_NAME=roles_1.0.1:e844f6bb1845b27efd6d42f9f646fb09f27ea6d3de5b20098ccf8b7454a23a2e                     -e AUTH_MODE=roles                     saacs_ccaas:latest
