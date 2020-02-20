#!/usr/bin/env bash

export WORK_DIR=$(cd `dirname $0`; pwd)

# delete pulsar cluster operator
kubectl delete -f ${WORK_DIR}/release.yaml

# delete pulsar cluster operator account and role
kubectl delete -f ${WORK_DIR}/rbac/all_namespace_rbac.yaml

# delete pulsar cluster crd
kubectl delete -f ${WORK_DIR}/crds/crds.yaml