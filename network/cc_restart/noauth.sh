#!/bin/bash
#
    docker stop peer0org1_noauth_ccaas
    docker run --rm -d --name peer0org1_noauth_ccaas                      --network fabric_test                     -e CHAINCODE_SERVER_ADDRESS=0.0.0.0:9999                     -e CHAINCODE_ID=noauth_1.0.1:80d2ea0195c271b52ce4179e4c42769b0235e9df3e60d755f1e356317ab04391 -e CORE_CHAINCODE_ID_NAME=noauth_1.0.1:80d2ea0195c271b52ce4179e4c42769b0235e9df3e60d755f1e356317ab04391                     -e AUTH_MODE=noauth                     saacs_ccaas:latest

    docker stop peer0org2_noauth_ccaas
    docker run --rm -d --name peer0org2_noauth_ccaas                     --network fabric_test                     -e CHAINCODE_SERVER_ADDRESS=0.0.0.0:9999                     -e CHAINCODE_ID=noauth_1.0.1:80d2ea0195c271b52ce4179e4c42769b0235e9df3e60d755f1e356317ab04391 -e CORE_CHAINCODE_ID_NAME=noauth_1.0.1:80d2ea0195c271b52ce4179e4c42769b0235e9df3e60d755f1e356317ab04391                     -e AUTH_MODE=noauth                     saacs_ccaas:latest
