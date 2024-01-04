#!/bin/bash
#
    docker stop peer0org1_noauth_ccaas
    docker run --rm -d --name peer0org1_noauth_ccaas                      --network fabric_test                     -e CHAINCODE_SERVER_ADDRESS=0.0.0.0:9999                     -e CHAINCODE_ID=noauth_1.0.1:ce4bfe74d203041e200203adfd41701e45c6746d2aebc350c823d0b459f03aa1 -e CORE_CHAINCODE_ID_NAME=noauth_1.0.1:ce4bfe74d203041e200203adfd41701e45c6746d2aebc350c823d0b459f03aa1                     -e AUTH_MODE=noauth                     saacs_ccaas:latest

    docker stop peer0org2_noauth_ccaas
    docker run --rm -d --name peer0org2_noauth_ccaas                     --network fabric_test                     -e CHAINCODE_SERVER_ADDRESS=0.0.0.0:9999                     -e CHAINCODE_ID=noauth_1.0.1:ce4bfe74d203041e200203adfd41701e45c6746d2aebc350c823d0b459f03aa1 -e CORE_CHAINCODE_ID_NAME=noauth_1.0.1:ce4bfe74d203041e200203adfd41701e45c6746d2aebc350c823d0b459f03aa1                     -e AUTH_MODE=noauth                     saacs_ccaas:latest
