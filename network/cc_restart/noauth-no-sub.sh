#!/bin/bash
#
    docker stop peer0org1_noauth-no-sub_ccaas
    docker run --rm -d --name peer0org1_noauth-no-sub_ccaas                      --network fabric_test                     -e CHAINCODE_SERVER_ADDRESS=0.0.0.0:9999                     -e CHAINCODE_ID=noauth-no-sub_1.0.1:2c9eb8ebd0db11f91fa4b8fe228a361c29268e99404ef0d5b24b09fea0ff6f58 -e CORE_CHAINCODE_ID_NAME=noauth-no-sub_1.0.1:2c9eb8ebd0db11f91fa4b8fe228a361c29268e99404ef0d5b24b09fea0ff6f58                     -e AUTH_MODE=noauth-no-sub                     saacs_ccaas:latest

    docker stop peer0org2_noauth-no-sub_ccaas
    docker run --rm -d --name peer0org2_noauth-no-sub_ccaas                     --network fabric_test                     -e CHAINCODE_SERVER_ADDRESS=0.0.0.0:9999                     -e CHAINCODE_ID=noauth-no-sub_1.0.1:2c9eb8ebd0db11f91fa4b8fe228a361c29268e99404ef0d5b24b09fea0ff6f58 -e CORE_CHAINCODE_ID_NAME=noauth-no-sub_1.0.1:2c9eb8ebd0db11f91fa4b8fe228a361c29268e99404ef0d5b24b09fea0ff6f58                     -e AUTH_MODE=noauth-no-sub                     saacs_ccaas:latest
