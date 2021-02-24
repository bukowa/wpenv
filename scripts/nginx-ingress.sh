#!/bin/bash
set -eo pipefail

NAMESPACE="ingress-nginx"

if [ "$1" == "" ]; then
  printf "Usage:\n 'nginx-ingress.sh install'\n 'nginx-ingress.sh upgrade'"
  exit 1
fi

helm repo add ingress-nginx https://kubernetes.github.io/ingress-nginx
helm repo update
helm repo ls

kubectl create namespace $NAMESPACE || true

if [ "$1" == "install" ]; then
  echo "Installing..."
  helm install \
    nginx-ingress ingress-nginx/ingress-nginx \
    --namespace $NAMESPACE \
    -f nginx-ingress.values.yaml
  exit 0
fi

if [ "$1" == "upgrade" ]; then
  echo "Upgrading..."
  helm upgrade \
    nginx-ingress ingress-nginx/ingress-nginx \
    --namespace $NAMESPACE \
    -f nginx-ingress.values.yaml
  exit 0
fi


echo "Wrong argument provided: $1"
exit 1