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
export IMAGE_NAME="nova38/saacs:latest"

### Configure Internal DNS
CLUSTER_IP=$(kubectl -n istio-system get svc istio-ingressgateway -o json | jq -r .spec.clusterIP)
kubectl apply -f - <<EOF
kind: ConfigMap
apiVersion: v1
metadata:
  name: coredns
  namespace: kube-system
data:
  Corefile: |
    .:53 {
        errors
        health {
           lameduck 5s
        }
        rewrite name regex (.*)\.nip\.io host.ingress.internal
        hosts {
          ${CLUSTER_IP} host.ingress.internal
          fallthrough
        }
        ready
        kubernetes cluster.local in-addr.arpa ip6.arpa {
           pods insecure
           fallthrough in-addr.arpa ip6.arpa
           ttl 30
        }
        prometheus :9153
        forward . /etc/resolv.conf {
           max_concurrent 1000
        }
        cache 30
        loop
        reload
        loadbalance
    }
EOF

echo $SEP
echo "Installing CA"
echo $SEP

### Deploy a certificate authority
kubectl hlf ca create \
    --image=$CA_IMAGE \
    --version=$CA_VERSION \
    --storage-class=standard \
    --capacity=1Gi \
    --name=org1-ca \
    --enroll-id=enroll \
    --enroll-pw=enrollpw \
    --hosts=org1-ca.$DOMAIN \
    --istio-port=443

kubectl wait --timeout=180s --for=condition=Running fabriccas.hlf.kungfusoftware.es --all

### Varify the certificate authority is working
curl -k https://org1-ca.$DOMAIN:443/cainfo

echo $SEP
echo "Installing Peers"
echo $SEP

# register user in CA for peers
kubectl hlf ca register \
    --name=org1-ca \
    --user=peer \
    --secret=peerpw \
    --type=peer \
    --enroll-id enroll \
    --enroll-secret=enrollpw \
    --mspid Org1MSP

## Deploy Peers
kubectl hlf peer create \
    --statedb=couchdb \
    --image=$PEER_IMAGE \
    --version=$PEER_VERSION \
    --storage-class=standard \
    --enroll-id=peer \
    --mspid=Org1MSP \
    --enroll-pw=peerpw \
    --capacity=10Gi \
    --name=org1-peer0 \
    --ca-name=org1-ca.default \
    --hosts=peer0-org1.$DOMAIN \
    --istio-port=443

kubectl hlf peer create \
    --statedb=couchdb \
    --image=$PEER_IMAGE \
    --version=$PEER_VERSION \
    --storage-class=standard \
    --enroll-id=peer \
    --mspid=Org1MSP \
    --enroll-pw=peerpw \
    --capacity=10Gi \
    --name=org1-peer1 \
    --ca-name=org1-ca.default \
    --hosts=peer1-org1.$DOMAIN \
    --istio-port=443

kubectl wait \
    --timeout=180s \
    --for=condition=Running \
    fabricpeers.hlf.kungfusoftware.es \
    --all

## Check the peers are working
openssl s_client -connect peer0-org1.$DOMAIN:443
openssl s_client -connect peer1-org1.$DOMAIN:443

echo $SEP
echo "Installing Oderer"
echo $SEP

# Deploy an `Orderer` organization
kubectl hlf ca create \
    --image=$CA_IMAGE \
    --version=$CA_VERSION \
    --storage-class=standard \
    --capacity=1Gi \
    --name=ord-ca \
    --enroll-id=enroll \
    --enroll-pw=enrollpw \
    --hosts=ord-ca.$DOMAIN \
    --istio-port=443

kubectl wait \
    --timeout=180s \
    --for=condition=Running \
    fabriccas.hlf.kungfusoftware.es \
    --all

## Verify the certificate authority is working
curl -vik https://ord-ca.$DOMAIN:443/cainfo

### Register user `orderer`
kubectl hlf ca register \
    --name=ord-ca \
    --user=orderer \
    --secret=ordererpw \
    --type=orderer \
    --enroll-id enroll \
    --enroll-secret=enrollpw \
    --mspid=OrdererMSP \
    --ca-url="https://ord-ca.$DOMAIN:443"

### Deploy an orderer node
kubectl hlf ordnode create \
    --image=$ORDERER_IMAGE \
    --version=$ORDERER_VERSION \
    --storage-class=standard \
    --enroll-id=orderer \
    --mspid=OrdererMSP \
    --enroll-pw=ordererpw \
    --capacity=2Gi \
    --name=ord-node1 \
    --ca-name=ord-ca.default \
    --hosts=orderer0-ord.$DOMAIN \
    --istio-port=443

#### Wait for the orderer node to be ready
kubectl wait \
    --timeout=180s \
    --for=condition=Running \
    fabricorderernodes.hlf.kungfusoftware.es \
    --all

#### Check the orderer node is working
kubectl get pods
openssl s_client -connect orderer0-ord.$DOMAIN:443
