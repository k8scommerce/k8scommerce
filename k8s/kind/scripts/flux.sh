#!/bin/sh

set -o errexit

. ./vars.sh

# Start gitops
flux bootstrap github \
  --owner=${GITHUB_USER} \
  --branch=master \
  --repository=services \
  --path=/k8s/gitops/clusters/${KIND_CLUSTER_NAME} \
  --personal=false \
  --private=true \
  --components=source-controller,kustomize-controller,notification-controller \
  --components-extra=image-automation-controller,image-reflector-controller 
  # --read-write-key