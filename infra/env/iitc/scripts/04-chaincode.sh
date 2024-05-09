#!/bin/bash
set -e

###
SEP="-------------------------------------------------------------------------"

### Variables
export DOMAIN=129.237.123.11.nip.io

### Environment Variables for AMD (Default)
export PEER_IMAGE=hyperledger/fabric-peer
export PEER_VERSION=2.5.5

export ORDERER_IMAGE=hyperledger/fabric-orderer
export ORDERER_VERSION=2.5.5

export CA_IMAGE=hyperledger/fabric-ca
export CA_VERSION=1.5.7

###

export CHANNEL=demo
export IMAGE_NAME="nova38/saacs:latest"
export CHAINCODE_NAME=biochain
export CHAINCODE_LABEL=biochain

###
export SEQUENCE=1
export VERSION="1.0"

##

echo
echo $SEP
echo "Installing Chaincode"
echo "Making Config"
echo $SEP
echo

#### 1. Get connection string without users for organization Org1MSP and OrdererMSP
kubectl hlf inspect --output org1.yaml -o Org1MSP -o OrdererMSP

### 2. Register a user in the certification authority for signing (register)
# kubectl hlf ca register \
#     --name=org1-ca \
#     --user=admin \
#     --secret=adminpw \
#     --type=admin \
#     --enroll-id enroll \
#     --enroll-secret=enrollpw \
#     --mspid Org1MSP

### 3. Obtain the certificates using the previously created user (enroll)
kubectl hlf ca enroll \
    --name=org1-ca \
    --user=admin \
    --secret=adminpw \
    --mspid Org1MSP \
    --ca-name ca \
    --output peer-org1.yaml

#### 4. Attach the user to the connection string
kubectl hlf utils adduser \
    --userPath=peer-org1.yaml \
    --config=org1.yaml \
    --username=admin \
    --mspid=Org1MSP

### Create metadata file
echo
echo $SEP
echo "Create metadata file"
echo $SEP
echo
#### remove the code.tar.gz chaincode.tgz if they exist
rm code.tar.gz chaincode.tgz || true

export CHAINCODE_NAME=biochain
export CHAINCODE_LABEL=biochain
cat <<METADATA-EOF >"metadata.json"
{
    "type": "ccaas",
    "label": "${CHAINCODE_LABEL}"
}
METADATA-EOF

### Prepare connection file
cat >"connection.json" <<CONN_EOF
{
  "address": "${CHAINCODE_NAME}:7052",
  "dial_timeout": "10s",
  "tls_required": false
}
CONN_EOF

cat >"connection.json" <<CONN_EOF
{
  "address": "${CHAINCODE_NAME}:7052",
  "dial_timeout": "10s",
  "tls_required": false
}
CONN_EOF

tar cfz code.tar.gz connection.json
tar cfz chaincode.tgz metadata.json code.tar.gz
export PACKAGE_ID=$(kubectl hlf chaincode calculatepackageid --path=chaincode.tgz --language=golang --label=$CHAINCODE_LABEL)
echo "PACKAGE_ID=$PACKAGE_ID"

# echo
# echo $SEP
# echo "Installing Chaincode"
# echo $SEP
# echo

kubectl hlf chaincode install --path=./chaincode.tgz \
    --config=org1.yaml \
    --language=golang \
    --label=$CHAINCODE_LABEL \
    --user=admin \
    --peer=org1-peer0.default
kubectl hlf chaincode install --path=./chaincode.tgz \
    --config=org1.yaml \
    --language=golang \
    --label=$CHAINCODE_LABEL \
    --user=admin \
    --peer=org1-peer1.default

# echo
# echo $SEP
# echo "Deploying Chaincode"
# echo $SEP
# echo

## Deploy chaincode container on cluster
kubectl hlf externalchaincode sync \
    --image=$IMAGE_NAME \
    --name=$CHAINCODE_NAME \
    --namespace=default \
    --package-id=$PACKAGE_ID \
    --tls-required=false \
    --replicas=1

# ## Check installed chaincodes

# kubectl hlf chaincode queryinstalled \
#     --config=org1.yaml \
#     --user=admin \
#     --peer=org1-peer0.default

sleep 5
echo
echo $SEP
echo "Approving Chaincode"
echo $SEP
echo
export SEQUENCE=3

kubectl hlf chaincode approveformyorg \
    --config=org1.yaml \
    --user=admin \
    --peer=org1-peer0.default \
    --package-id=$PACKAGE_ID \
    --version "$VERSION" \
    --sequence "$SEQUENCE" \
    --name=$CHAINCODE_NAME \
    --policy="OR('Org1MSP.member')" \
    --channel=$CHANNEL

echo
echo $SEP
echo "Commiting Chaincode"
echo $SEP
echo

## Commit chaincode
kubectl hlf chaincode commit \
    --config=org1.yaml \
    --user=admin \
    --mspid=Org1MSP \
    --version "$VERSION" \
    --sequence "$SEQUENCE" \
    --name=$CHAINCODE_NAME \
    --policy="OR('Org1MSP.member')" \
    --channel=$CHANNEL

echo
echo $SEP
echo "Invokeing Chaincode"
echo $SEP
echo

## Invoke a transaction on the channel
kubectl hlf chaincode invoke --config=org1.yaml \
    --user=admin \
    --peer=org1-peer0.default \
    --chaincode=$CHAINCODE_NAME \
    --channel=$CHANNEL \
    --fcn=getCurrentUser -a '[]'
