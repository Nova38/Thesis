#!/bin/bash
#
    docker stop peer0org1_roles_ccaas
    docker run --rm -d --name peer0org1_roles_ccaas                      --network fabric_test                     -e CHAINCODE_SERVER_ADDRESS=0.0.0.0:9999                     -e CHAINCODE_ID=roles_1.0.1:8f0206932d1996fc9a3a78c352bfefda0397a5b9fa55a1d19ca452de34033d59 -e CORE_CHAINCODE_ID_NAME=roles_1.0.1:8f0206932d1996fc9a3a78c352bfefda0397a5b9fa55a1d19ca452de34033d59                     -e AUTH_MODE=roles                     saacs_ccaas:latest

    docker stop peer0org2_roles_ccaas
    docker run --rm -d --name peer0org2_roles_ccaas                     --network fabric_test                     -e CHAINCODE_SERVER_ADDRESS=0.0.0.0:9999                     -e CHAINCODE_ID=roles_1.0.1:8f0206932d1996fc9a3a78c352bfefda0397a5b9fa55a1d19ca452de34033d59 -e CORE_CHAINCODE_ID_NAME=roles_1.0.1:8f0206932d1996fc9a3a78c352bfefda0397a5b9fa55a1d19ca452de34033d59                     -e AUTH_MODE=roles                     saacs_ccaas:latest
