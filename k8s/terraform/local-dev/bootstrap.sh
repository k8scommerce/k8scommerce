#!/bin/sh

set -e

. tfvars.sh

terraform init

terraform apply -auto-approve

# printf "\nWaiting for the echo web server service... \n"
# kubectl apply -f https://kind.sigs.k8s.io/examples/ingress/usage.yaml
# sleep 10