#!/bin/bash
#
    docker stop peer0org1_noauth-no-sub_ccaas
    docker run --rm -d --name peer0org1_noauth-no-sub_ccaas                      --network fabric_test                     -e CHAINCODE_SERVER_ADDRESS=0.0.0.0:9999                     -e CHAINCODE_ID=noauth-no-sub_1.0.1:40145881166bae862b71329662a751ee7c674999e99e4d6f0421d64714fbe1ec -e CORE_CHAINCODE_ID_NAME=noauth-no-sub_1.0.1:40145881166bae862b71329662a751ee7c674999e99e4d6f0421d64714fbe1ec                     -e AUTH_MODE=noauth-no-sub                     saacs_ccaas:latest

    docker stop peer0org2_noauth-no-sub_ccaas
    docker run --rm -d --name peer0org2_noauth-no-sub_ccaas                     --network fabric_test                     -e CHAINCODE_SERVER_ADDRESS=0.0.0.0:9999                     -e CHAINCODE_ID=noauth-no-sub_1.0.1:40145881166bae862b71329662a751ee7c674999e99e4d6f0421d64714fbe1ec -e CORE_CHAINCODE_ID_NAME=noauth-no-sub_1.0.1:40145881166bae862b71329662a751ee7c674999e99e4d6f0421d64714fbe1ec                     -e AUTH_MODE=noauth-no-sub                     saacs_ccaas:latest
