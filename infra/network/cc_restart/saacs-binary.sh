#!/bin/bash
#
docker stop peer0org1_saacs-binary_ccaas
docker run --rm -d --name peer0org1_saacs-binary_ccaas --network fabric_test -e CHAINCODE_SERVER_ADDRESS=0.0.0.0:9999 -e CHAINCODE_ID=saacs-binary_1.0.1:381fea7ed86f5ef2af4dba795da4ab4e18de533f211227b94f2fa8ba5486f83b -e CORE_CHAINCODE_ID_NAME=saacs-binary_1.0.1:381fea7ed86f5ef2af4dba795da4ab4e18de533f211227b94f2fa8ba5486f83b -e AUTH_MODE=saacs-binary -e=SERIALIZATION=proto saacs_ccaas:latest

docker stop peer0org2_saacs-binary_ccaas
docker run --rm -d --name peer0org2_saacs-binary_ccaas --network fabric_test -e CHAINCODE_SERVER_ADDRESS=0.0.0.0:9999 -e CHAINCODE_ID=saacs-binary_1.0.1:381fea7ed86f5ef2af4dba795da4ab4e18de533f211227b94f2fa8ba5486f83b -e CORE_CHAINCODE_ID_NAME=saacs-binary_1.0.1:381fea7ed86f5ef2af4dba795da4ab4e18de533f211227b94f2fa8ba5486f83b -e AUTH_MODE=saacs-binary -e=SERIALIZATION=proto saacs_ccaas:latest
