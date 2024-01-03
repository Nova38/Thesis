#!/bin/bash
#
    docker stop peer0org1_noauth_ccaas
    docker run --rm -d --name peer0org1_noauth_ccaas                      --network fabric_test                     -e CHAINCODE_SERVER_ADDRESS=0.0.0.0:9999                     -e CHAINCODE_ID=noauth_1.0.1:131fd9de451901ce3375eb01fba37a862735a0ae79d6553ac6b3a83ec65f6564 -e CORE_CHAINCODE_ID_NAME=noauth_1.0.1:131fd9de451901ce3375eb01fba37a862735a0ae79d6553ac6b3a83ec65f6564                     -e AUTH_MODE=noauth                     saacs_ccaas:latest

    docker stop peer0org2_noauth_ccaas
    docker run --rm -d --name peer0org2_noauth_ccaas                     --network fabric_test                     -e CHAINCODE_SERVER_ADDRESS=0.0.0.0:9999                     -e CHAINCODE_ID=noauth_1.0.1:131fd9de451901ce3375eb01fba37a862735a0ae79d6553ac6b3a83ec65f6564 -e CORE_CHAINCODE_ID_NAME=noauth_1.0.1:131fd9de451901ce3375eb01fba37a862735a0ae79d6553ac6b3a83ec65f6564                     -e AUTH_MODE=noauth                     saacs_ccaas:latest
