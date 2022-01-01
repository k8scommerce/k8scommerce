#!/bin/sh

set -o errexit

docker-compose -f ./../../db/docker-compose.yml down 

reg_name='local-registry'
reg_port='5000'

# create registry container unless it already exists
running="$(docker inspect -f '{{.State.Running}}' "${reg_name}" 2>/dev/null || true)"
if [ "${running}" == 'true' ]; then
  cid="$(docker inspect -f '{{.ID}}' "${reg_name}")"
  echo "> Stopping and deleting registry container..."
  docker stop $cid >/dev/null
  docker rm $cid >/dev/null
fi

echo "> Deleting dev-cluster..."
k3d cluster delete dev-cluster