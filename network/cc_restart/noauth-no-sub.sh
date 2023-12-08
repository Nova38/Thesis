#!/bin/bash
#
    docker stop peer0org1_noauth-no-sub_ccaas
    docker run --rm -d --name peer0org1_noauth-no-sub_ccaas                      --network fabric_test                     -e CHAINCODE_SERVER_ADDRESS=0.0.0.0:9999                     -e CHAINCODE_ID=noauth-no-sub_1.0.1:5a9192dcf5648f57237bde96dc72626974a96321120618367dd104b06b39c1ed -e CORE_CHAINCODE_ID_NAME=noauth-no-sub_1.0.1:5a9192dcf5648f57237bde96dc72626974a96321120618367dd104b06b39c1ed                     -e AUTH_MODE=noauth-no-sub                     saacs_ccaas:latest

    docker stop peer0org2_noauth-no-sub_ccaas
    docker run --rm -d --name peer0org2_noauth-no-sub_ccaas                     --network fabric_test                     -e CHAINCODE_SERVER_ADDRESS=0.0.0.0:9999                     -e CHAINCODE_ID=noauth-no-sub_1.0.1:5a9192dcf5648f57237bde96dc72626974a96321120618367dd104b06b39c1ed -e CORE_CHAINCODE_ID_NAME=noauth-no-sub_1.0.1:5a9192dcf5648f57237bde96dc72626974a96321120618367dd104b06b39c1ed                     -e AUTH_MODE=noauth-no-sub                     saacs_ccaas:latest
