#!/bin/bash
#
    docker stop peer0org1_caliper-saacs-binary_ccaas
    docker run --rm -d --name peer0org1_caliper-saacs-binary_ccaas                      --network fabric_test                     -e CHAINCODE_SERVER_ADDRESS=0.0.0.0:9999                     -e CHAINCODE_ID=caliper-saacs-binary_1.0.1:2628165e449fad65998f740170c4a559b595843f376a6b1f8fe4374410913f4c -e CORE_CHAINCODE_ID_NAME=caliper-saacs-binary_1.0.1:2628165e449fad65998f740170c4a559b595843f376a6b1f8fe4374410913f4c                     -e AUTH_MODE=caliper-saacs-binary                     -e=SERIALIZATION=proto                     saacs_ccaas:latest

    docker stop peer0org2_caliper-saacs-binary_ccaas
    docker run --rm -d --name peer0org2_caliper-saacs-binary_ccaas                     --network fabric_test                     -e CHAINCODE_SERVER_ADDRESS=0.0.0.0:9999                     -e CHAINCODE_ID=caliper-saacs-binary_1.0.1:2628165e449fad65998f740170c4a559b595843f376a6b1f8fe4374410913f4c -e CORE_CHAINCODE_ID_NAME=caliper-saacs-binary_1.0.1:2628165e449fad65998f740170c4a559b595843f376a6b1f8fe4374410913f4c                     -e AUTH_MODE=caliper-saacs-binary                     -e=SERIALIZATION=proto                     saacs_ccaas:latest
