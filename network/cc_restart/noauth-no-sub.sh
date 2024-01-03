#!/bin/bash
#
    docker stop peer0org1_noauth-no-sub_ccaas
    docker run --rm -d --name peer0org1_noauth-no-sub_ccaas                      --network fabric_test                     -e CHAINCODE_SERVER_ADDRESS=0.0.0.0:9999                     -e CHAINCODE_ID=noauth-no-sub_1.0.1:3a2c992e1c80a957ffcbc598d74339aa32c2960686c15c72405f198dc4694953 -e CORE_CHAINCODE_ID_NAME=noauth-no-sub_1.0.1:3a2c992e1c80a957ffcbc598d74339aa32c2960686c15c72405f198dc4694953                     -e AUTH_MODE=noauth-no-sub                     saacs_ccaas:latest

    docker stop peer0org2_noauth-no-sub_ccaas
    docker run --rm -d --name peer0org2_noauth-no-sub_ccaas                     --network fabric_test                     -e CHAINCODE_SERVER_ADDRESS=0.0.0.0:9999                     -e CHAINCODE_ID=noauth-no-sub_1.0.1:3a2c992e1c80a957ffcbc598d74339aa32c2960686c15c72405f198dc4694953 -e CORE_CHAINCODE_ID_NAME=noauth-no-sub_1.0.1:3a2c992e1c80a957ffcbc598d74339aa32c2960686c15c72405f198dc4694953                     -e AUTH_MODE=noauth-no-sub                     saacs_ccaas:latest
