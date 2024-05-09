#!/bin/bash
#
    docker stop peer0org1_saacs_ccaas
    docker run --rm -d --name peer0org1_saacs_ccaas                      --network fabric_test                     -e CHAINCODE_SERVER_ADDRESS=0.0.0.0:9999                     -e CHAINCODE_ID=saacs_1.0.1:332926c2b0c063d7352f58f0344153997b9d24fbf368f16207ab1a212c9938bc -e CORE_CHAINCODE_ID_NAME=saacs_1.0.1:332926c2b0c063d7352f58f0344153997b9d24fbf368f16207ab1a212c9938bc                     -e AUTH_MODE=saacs                     saacs_ccaas:latest

    docker stop peer0org2_saacs_ccaas
    docker run --rm -d --name peer0org2_saacs_ccaas                     --network fabric_test                     -e CHAINCODE_SERVER_ADDRESS=0.0.0.0:9999                     -e CHAINCODE_ID=saacs_1.0.1:332926c2b0c063d7352f58f0344153997b9d24fbf368f16207ab1a212c9938bc -e CORE_CHAINCODE_ID_NAME=saacs_1.0.1:332926c2b0c063d7352f58f0344153997b9d24fbf368f16207ab1a212c9938bc                     -e AUTH_MODE=saacs                     saacs_ccaas:latest
