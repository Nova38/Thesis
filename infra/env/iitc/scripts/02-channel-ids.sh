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

echo
echo $SEP
echo "Making Identities"
echo $SEP
echo

### Register and enrolling OrdererMSP identity
kubectl hlf ca register \
    --name=ord-ca \
    --user=admin \
    --secret=adminpw \
    --type=admin \
    --enroll-id enroll \
    --enroll-secret=enrollpw \
    --mspid=OrdererMSP

# enroll

kubectl hlf ca enroll \
    --name=ord-ca \
    --namespace=default \
    --user=admin \
    --secret=adminpw \
    --mspid OrdererMSP \
    --ca-name tlsca \
    --output orderermsp.yaml

### Register and enrolling Org1MSP identity

#### register
kubectl hlf ca register \
    --name=org1-ca \
    --namespace=default \
    --user=admin --secret=adminpw \
    --type=admin \
    --enroll-id enroll \
    --enroll-secret=enrollpw \
    --mspid=Org1MSP

#### enroll
kubectl hlf ca enroll \
    --name=org1-ca \
    --namespace=default \
    --user=admin \
    --secret=adminpw \
    --mspid Org1MSP \
    --ca-name ca \
    --output org1msp.yaml

# enroll
kubectl hlf identity create --name org1-admin --namespace default \
    --ca-name org1-ca --ca-namespace default \
    --ca ca --mspid Org1MSP --enroll-id admin --enroll-secret adminpw

### Create the secret
kubectl delete secret wallet --ignore-not-found

kubectl create secret generic wallet \
    --namespace=default \
    --from-file=org1msp.yaml=$PWD/org1msp.yaml \
    --from-file=orderermsp.yaml=$PWD/orderermsp.yaml
