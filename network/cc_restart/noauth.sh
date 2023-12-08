#!/bin/bash
#
    docker stop peer0org1_noauth_ccaas
    docker run --rm -d --name peer0org1_noauth_ccaas                      --network fabric_test                     -e CHAINCODE_SERVER_ADDRESS=0.0.0.0:9999                     -e CHAINCODE_ID=noauth_1.0.1:3dd1b370597bf6feffa71f1802b2455332b68fdd5e61c2f516b95a13d09fc803 -e CORE_CHAINCODE_ID_NAME=noauth_1.0.1:3dd1b370597bf6feffa71f1802b2455332b68fdd5e61c2f516b95a13d09fc803                     -e AUTH_MODE=noauth                     saacs_ccaas:latest

    docker stop peer0org2_noauth_ccaas
    docker run --rm -d --name peer0org2_noauth_ccaas                     --network fabric_test                     -e CHAINCODE_SERVER_ADDRESS=0.0.0.0:9999                     -e CHAINCODE_ID=noauth_1.0.1:3dd1b370597bf6feffa71f1802b2455332b68fdd5e61c2f516b95a13d09fc803 -e CORE_CHAINCODE_ID_NAME=noauth_1.0.1:3dd1b370597bf6feffa71f1802b2455332b68fdd5e61c2f516b95a13d09fc803                     -e AUTH_MODE=noauth                     saacs_ccaas:latest
