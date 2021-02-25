#!/bin/bash
set -euxo pipefail

NAMESPACE="ingress-nginx"

if [ "$#" -eq 1 ]; then
    if [ "$1" == "delete" ]; then
        k3d cluster delete || true
    fi
fi

k3d cluster create --wait --no-lb -p "32080:32080@server[*]" -p "32443:32443@server[*]"

NAMESPACE=$NAMESPACE ./generate_certs.sh
NAMESPACE=$NAMESPACE ./nginx-ingress.sh install

make run_wait_for_status_code
