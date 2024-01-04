#!/bin/bash
#
    docker stop peer0org1_roles_ccaas
    docker run --rm -d --name peer0org1_roles_ccaas                      --network fabric_test                     -e CHAINCODE_SERVER_ADDRESS=0.0.0.0:9999                     -e CHAINCODE_ID=roles_1.0.1:24444e60c6def1a8ed62bf0f7912dd1fdae585be1ee02105e0a6af33123e398a -e CORE_CHAINCODE_ID_NAME=roles_1.0.1:24444e60c6def1a8ed62bf0f7912dd1fdae585be1ee02105e0a6af33123e398a                     -e AUTH_MODE=roles                     saacs_ccaas:latest

    docker stop peer0org2_roles_ccaas
    docker run --rm -d --name peer0org2_roles_ccaas                     --network fabric_test                     -e CHAINCODE_SERVER_ADDRESS=0.0.0.0:9999                     -e CHAINCODE_ID=roles_1.0.1:24444e60c6def1a8ed62bf0f7912dd1fdae585be1ee02105e0a6af33123e398a -e CORE_CHAINCODE_ID_NAME=roles_1.0.1:24444e60c6def1a8ed62bf0f7912dd1fdae585be1ee02105e0a6af33123e398a                     -e AUTH_MODE=roles                     saacs_ccaas:latest
