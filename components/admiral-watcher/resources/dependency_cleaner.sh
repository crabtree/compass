#!/bin/bash

set -e

function cleanup() {
  export KUBECONFIG=
}

trap cleanup EXIT

ROOT_PATH=$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )

echo $ROOT_PATH

if [ "$#" -gt "0" ]; then
  dep=$1
fi

if [ "$#" -gt "1" ]; then
  remote=$2
fi

admiral_cluster=${ROOT_PATH}/kubeconfigs/admiral.yaml
remote_cluster=${ROOT_PATH}/kubeconfigs/${remote}

export KUBECONFIG=$admiral_cluster
kubectl delete dependency ${dep}

export KUBECONFIG=$remote_cluster
kubectl delete se --all -n admiral-sync
kubectl delete dr --all -n admiral-sync

export KUBECONFIG=$admiral_cluster
kubectl rollout restart deployment/admiral -n admiral

export KUBECONFIG=$remote_cluster
kubectl delete se --all -n admiral-sync
kubectl delete dr --all -n admiral-sync