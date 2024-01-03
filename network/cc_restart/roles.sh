#!/bin/bash
#
    docker stop peer0org1_roles_ccaas
    docker run --rm -d --name peer0org1_roles_ccaas                      --network fabric_test                     -e CHAINCODE_SERVER_ADDRESS=0.0.0.0:9999                     -e CHAINCODE_ID=roles_1.0.1:11d48cab5b263af11dba2630e39c6b7c6811c03572b0d35449bfddec6552ff8a -e CORE_CHAINCODE_ID_NAME=roles_1.0.1:11d48cab5b263af11dba2630e39c6b7c6811c03572b0d35449bfddec6552ff8a                     -e AUTH_MODE=roles                     saacs_ccaas:latest

    docker stop peer0org2_roles_ccaas
    docker run --rm -d --name peer0org2_roles_ccaas                     --network fabric_test                     -e CHAINCODE_SERVER_ADDRESS=0.0.0.0:9999                     -e CHAINCODE_ID=roles_1.0.1:11d48cab5b263af11dba2630e39c6b7c6811c03572b0d35449bfddec6552ff8a -e CORE_CHAINCODE_ID_NAME=roles_1.0.1:11d48cab5b263af11dba2630e39c6b7c6811c03572b0d35449bfddec6552ff8a                     -e AUTH_MODE=roles                     saacs_ccaas:latest
