#!/bin/bash
#
    docker stop peer0org1_noauth-no-sub_ccaas
    docker run --rm -d --name peer0org1_noauth-no-sub_ccaas                      --network fabric_test                     -e CHAINCODE_SERVER_ADDRESS=0.0.0.0:9999                     -e CHAINCODE_ID=noauth-no-sub_1.0.1:5415773e3e829d2f27875ba17b4fff7f3e6b46733e02698d8da468be74eae7ae -e CORE_CHAINCODE_ID_NAME=noauth-no-sub_1.0.1:5415773e3e829d2f27875ba17b4fff7f3e6b46733e02698d8da468be74eae7ae                     -e AUTH_MODE=noauth-no-sub                     saacs_ccaas:latest

    docker stop peer0org2_noauth-no-sub_ccaas
    docker run --rm -d --name peer0org2_noauth-no-sub_ccaas                     --network fabric_test                     -e CHAINCODE_SERVER_ADDRESS=0.0.0.0:9999                     -e CHAINCODE_ID=noauth-no-sub_1.0.1:5415773e3e829d2f27875ba17b4fff7f3e6b46733e02698d8da468be74eae7ae -e CORE_CHAINCODE_ID_NAME=noauth-no-sub_1.0.1:5415773e3e829d2f27875ba17b4fff7f3e6b46733e02698d8da468be74eae7ae                     -e AUTH_MODE=noauth-no-sub                     saacs_ccaas:latest
