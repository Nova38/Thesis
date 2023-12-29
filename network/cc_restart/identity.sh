#!/bin/bash
#
    docker stop peer0org1_identity_ccaas
    docker run --rm -d --name peer0org1_identity_ccaas                      --network fabric_test                     -e CHAINCODE_SERVER_ADDRESS=0.0.0.0:9999                     -e CHAINCODE_ID=identity_1.0.1:4fd7f83dada4cd353b6318fb16419d4574a1cfa366679c80cee320e53ac59bc3 -e CORE_CHAINCODE_ID_NAME=identity_1.0.1:4fd7f83dada4cd353b6318fb16419d4574a1cfa366679c80cee320e53ac59bc3                     -e AUTH_MODE=identity                     saacs_ccaas:latest

    docker stop peer0org2_identity_ccaas
    docker run --rm -d --name peer0org2_identity_ccaas                     --network fabric_test                     -e CHAINCODE_SERVER_ADDRESS=0.0.0.0:9999                     -e CHAINCODE_ID=identity_1.0.1:4fd7f83dada4cd353b6318fb16419d4574a1cfa366679c80cee320e53ac59bc3 -e CORE_CHAINCODE_ID_NAME=identity_1.0.1:4fd7f83dada4cd353b6318fb16419d4574a1cfa366679c80cee320e53ac59bc3                     -e AUTH_MODE=identity                     saacs_ccaas:latest
