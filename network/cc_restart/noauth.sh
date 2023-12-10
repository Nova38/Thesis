#!/bin/bash
#
    docker stop peer0org1_noauth_ccaas
    docker run --rm -d --name peer0org1_noauth_ccaas                      --network fabric_test                     -e CHAINCODE_SERVER_ADDRESS=0.0.0.0:9999                     -e CHAINCODE_ID=noauth_1.0.1:f6828dbb5063bbaa687c7570e80b73a4e7895ceed19fff2dc905ea3bdb478d38 -e CORE_CHAINCODE_ID_NAME=noauth_1.0.1:f6828dbb5063bbaa687c7570e80b73a4e7895ceed19fff2dc905ea3bdb478d38                     -e AUTH_MODE=noauth                     saacs_ccaas:latest

    docker stop peer0org2_noauth_ccaas
    docker run --rm -d --name peer0org2_noauth_ccaas                     --network fabric_test                     -e CHAINCODE_SERVER_ADDRESS=0.0.0.0:9999                     -e CHAINCODE_ID=noauth_1.0.1:f6828dbb5063bbaa687c7570e80b73a4e7895ceed19fff2dc905ea3bdb478d38 -e CORE_CHAINCODE_ID_NAME=noauth_1.0.1:f6828dbb5063bbaa687c7570e80b73a4e7895ceed19fff2dc905ea3bdb478d38                     -e AUTH_MODE=noauth                     saacs_ccaas:latest
