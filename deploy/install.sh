#!/usr/bin/env bash

export WORK_DIR=$(cd `dirname $0`; pwd)

# create pulsar cluster crd
kubectl create -f ${WORK_DIR}/crds/crds.yaml

# create pulsar cluster operator account and role
kubectl create -f ${WORK_DIR}/rbac/all_namespace_rbac.yaml

# install pulsar cluster operator
kubectl create -f ${WORK_DIR}/release.yaml