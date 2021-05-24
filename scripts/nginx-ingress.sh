#!/bin/bash
set -euxo pipefail

NAMESPACE="ingress-nginx"

if [ "$1" == "" ]; then
  printf "Usage:\n 'nginx-ingress.sh install'\n 'nginx-ingress.sh upgrade'"
  exit 1
fi

helm --kube-context=k3d-wpenv repo add ingress-nginx https://kubernetes.github.io/ingress-nginx
helm --kube-context=k3d-wpenv repo update
helm --kube-context=k3d-wpenv repo ls

kubectl --context=k3d-wpenv create namespace $NAMESPACE || true

if [ "$1" == "install" ]; then
  echo "Installing..."
  helm --kube-context=k3d-wpenv install \
    nginx-ingress ingress-nginx/ingress-nginx \
    --namespace $NAMESPACE \
    -f nginx-ingress.values.yaml
  exit 0
fi

if [ "$1" == "upgrade" ]; then
  echo "Upgrading..."
  helm --kube-context=k3d-wpenv upgrade \
    nginx-ingress ingress-nginx/ingress-nginx \
    --namespace $NAMESPACE \
    -f nginx-ingress.values.yaml
  exit 0
fi

if [ "$1" == "uninstall" ]; then
  echo "Uninstalling..."
  helm --kube-context=k3d-wpenv uninstall -n $NAMESPACE \
    nginx-ingress
  exit 0
fi


echo "Wrong argument provided: $1"
exit 1


# kubectl get -n $NAMESPACE service nginx-ingress-ingress-nginx-controller -o jsonpath='{.spec.ports[?(@.name=="http")].nodePort}'
# kubectl -n $NAMESPACE port-forward service/nginx-ingress-ingress-nginx-controller 443
