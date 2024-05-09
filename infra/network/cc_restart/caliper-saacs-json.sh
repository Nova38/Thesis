#!/bin/bash
#
    docker stop peer0org1_caliper-saacs-json_ccaas
    docker run --rm -d --name peer0org1_caliper-saacs-json_ccaas                      --network fabric_test                     -e CHAINCODE_SERVER_ADDRESS=0.0.0.0:9999                     -e CHAINCODE_ID=caliper-saacs-json_1.0.1:3c8974170e09cb77040e7a98ce9b9262f584d1cbed385cd49a38c47962bb5df7 -e CORE_CHAINCODE_ID_NAME=caliper-saacs-json_1.0.1:3c8974170e09cb77040e7a98ce9b9262f584d1cbed385cd49a38c47962bb5df7                     -e AUTH_MODE=caliper-saacs-json                     -e=SERIALIZATION=json                     saacs_ccaas:latest

    docker stop peer0org2_caliper-saacs-json_ccaas
    docker run --rm -d --name peer0org2_caliper-saacs-json_ccaas                     --network fabric_test                     -e CHAINCODE_SERVER_ADDRESS=0.0.0.0:9999                     -e CHAINCODE_ID=caliper-saacs-json_1.0.1:3c8974170e09cb77040e7a98ce9b9262f584d1cbed385cd49a38c47962bb5df7 -e CORE_CHAINCODE_ID_NAME=caliper-saacs-json_1.0.1:3c8974170e09cb77040e7a98ce9b9262f584d1cbed385cd49a38c47962bb5df7                     -e AUTH_MODE=caliper-saacs-json                     -e=SERIALIZATION=json                     saacs_ccaas:latest
