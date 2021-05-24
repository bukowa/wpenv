#!/bin/bash
set -euxo pipefail

NAMESPACE="ingress-nginx"

kubectl --context=k3d-wpenv create namespace $NAMESPACE || true

KEY_FILE=cert.key
CERT_FILE=cert.crt
HOST=localhost
CERT_NAME=default-tls

rm cert.key cert.cert || true
kubectl --context=k3d-wpenv --namespace=$NAMESPACE delete secret $CERT_NAME || true

openssl req -x509 -nodes -days 365 -newkey rsa:2048 -keyout ${KEY_FILE} -out ${CERT_FILE} -subj "/CN=${HOST}/O=${HOST}"
kubectl --context=k3d-wpenv --namespace=$NAMESPACE create secret tls ${CERT_NAME} --key ${KEY_FILE} --cert ${CERT_FILE}
