#!/bin/sh

# bootstrap 

set -o errexit

# remove any exisiting clusters
echo "> Cleaning up existing cluster..."
. ./destroy.sh

# start the database
echo "> Starting local database..."
docker-compose -f ./../../db/docker-compose.yml up -d --force --remove-orphans

# populate the database
echo "> Adding demo data..."
go run ./../../db/tools .

# install kind
echo "> Starting k8s cluster..."
. ./scripts/kind.sh

sleep 30

# connect the database to the 
echo "> Connecting db to kind network..."
docker network connect "kind" "k8scommerce-db"

echo "> Adding ingress nginx..."
kubectl apply -f https://raw.githubusercontent.com/kubernetes/ingress-nginx/master/deploy/static/provider/kind/deploy.yaml

sleep 30

echo "> Adding cert manager..."
kubectl apply --validate=false -f https://github.com/jetstack/cert-manager/releases/download/v1.0.2/cert-manager.yaml

sleep 30

echo "> Adding cert-issuer.yaml..."
cat <<EOF | kubectl apply -f -
apiVersion: cert-manager.io/v1
kind: ClusterIssuer
metadata:
  name: test-selfsigned
spec:
  selfSigned: {}
EOF


. ./scripts/argo.sh


echo "> Adding argocd..."


sleep 60

kubectl get pods -n argocd -l app.kubernetes.io/name=argocd-server -o name | cut -d'/' -f 2






# install & start flux
# echo "> Starting flux..."
# . ./scripts/flux.sh

# echo "> Creating aws-credentials secret..."
# kubectl create secret generic aws-credentials \
#   --namespace=flux-system \
#   --from-literal=AWS_ACCESS_KEY_ID="${ECOMM_DEV_AWS_ACCESS_KEY_ID}" \
#   --from-literal=AWS_SECRET_ACCESS_KEY="${ECOMM_DEV_AWS_SECRET_ACCESS_KEY}"

# sleep 10
# echo "> Manually running ecr-credentials-sync-init..."
# kubectl create job \
#   --from=cronjob/ecr-credentials-sync \
#   --namespace=flux-system \
#   ecr-credentials-sync-init