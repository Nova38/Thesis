#!/bin/bash
#
    docker stop peer0org1_identity_ccaas
    docker run --rm -d --name peer0org1_identity_ccaas                      --network fabric_test                     -e CHAINCODE_SERVER_ADDRESS=0.0.0.0:9999                     -e CHAINCODE_ID=identity_1.0.1:8b55e14c96bf02c7370ad6102e5e1b27d266525fc5169bcf3d1a57ae4d0483b5 -e CORE_CHAINCODE_ID_NAME=identity_1.0.1:8b55e14c96bf02c7370ad6102e5e1b27d266525fc5169bcf3d1a57ae4d0483b5                     -e AUTH_MODE=identity                     saacs_ccaas:latest

    docker stop peer0org2_identity_ccaas
    docker run --rm -d --name peer0org2_identity_ccaas                     --network fabric_test                     -e CHAINCODE_SERVER_ADDRESS=0.0.0.0:9999                     -e CHAINCODE_ID=identity_1.0.1:8b55e14c96bf02c7370ad6102e5e1b27d266525fc5169bcf3d1a57ae4d0483b5 -e CORE_CHAINCODE_ID_NAME=identity_1.0.1:8b55e14c96bf02c7370ad6102e5e1b27d266525fc5169bcf3d1a57ae4d0483b5                     -e AUTH_MODE=identity                     saacs_ccaas:latest
