#!/bin/bash
set -eo pipefail

NAMESPACE="ingress-nginx"

if [ "$1" == "kind" ]; then
    kind delete cluster || true
    kind create cluster --config=./kind.config.yaml --wait=240s
fi

NAMESPACE=$NAMESPACE ./generate_certs.sh
NAMESPACE=$NAMESPACE ./nginx-ingress.sh install
