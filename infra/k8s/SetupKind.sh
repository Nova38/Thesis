#!/usr/bin/env bash

# Create a kind cluster
kind create cluster --config=./kind-config.yaml


# Install HLF Operator
helm repo add kfs https://kfsoftware.github.io/hlf-helm-charts --force-update
helm install hlf-operator --version=1.9.0 -- kfs/hlf-operator

kubectl create namespace istio-system
istioctl operator init


kubectl wait --timeout=180s --for=jsonpath='{.status.status}'=HEALTHY istiooperator istio-gateway --namespace=istio-system

# Configure DNS
  CLUSTER_IP=$(kubectl -n istio-system get svc istio-ingressgateway -o json | jq -r .spec.clusterIP)
  echo "CLUSTER_IP=${CLUSTER_IP}"
  kubectl apply -f - <<EOF

  EOF

  kubectl get configmap coredns -n kube-system -o yaml
