#!/bin/bash
set -e

# Registering users with attributes#

CA_NAME=org1-ca
CA_NAMESPACE=default

ENROLL_ID=enroll       # enroll id for the CA, default `enroll`
ENROLL_SECRET=enrollpw # enroll secret for the CA, default `enrollpw`
USER_TYPE=client       # can be `peer`, `orderer`, `client` or `admin`
USER_SECRET=clientpw
MSP_ID=Org1MSP
CA_TYPE=ca # can be `ca` or `tlsca`

USER_FILE=users.txt

KEYSTORE_DIR=./keystore

cat $USER_FILE | xargs -I{} kubectl hlf ca register \
    --name=$CA_NAME \
    --namespace=$CA_NAMESPACE \
    --user {} \
    --secret=$USER_SECRET \
    --type=$USER_TYPE \
    --enroll-id=$ENROLL_ID \
    --enroll-secret=$ENROLL_SECRET \
    --mspid $MSP_ID \
    --attributes="isPublic=false"

cat $USER_FILE | xargs -I{} kubectl hlf ca enroll \
    --name=$CA_NAME \
    --namespace=$CA_NAMESPACE \
    --user={} \
    --secret=$USER_SECRET \
    --mspid $CA_MSPID \
    --ca-name $CA_TYPE \
    --output $KEYSTORE_DIR/org1.{}.yaml

cat $USER_FILE | xargs -I{} kubectl hlf utils adduser \
    --userPath=$KEYSTORE_DIR/org1.{}.yaml \
    --config=org1.yaml \
    --username={} \
    --mspid=$MSP_ID
