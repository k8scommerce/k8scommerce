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

# install the dev-cluster
echo "> Installing dev-cluster"
k3d cluster create dev-cluster --port 8080:80@loadbalancer --port 8443:443@loadbalancer --registry-create dev-cluster-registry

sleep 30

# connect the database to the 
echo "> Connecting db to dev-cluster network..."
docker network connect "k3d-dev-cluster" "k8scommerce-db"

# # install argocd
# echo "> Installing argocd"
# kubectl create namespace argocd
# kubectl create -n argocd -f install.yaml

# sleep 30

# #Create an Ingress to redirect /argocd to the argocd service
# kubectl apply -f ingress.yaml -n argocd

# sleep 30

# # echo "> Getting argocd password"
# # kubectl -n argocd get secret argocd-initial-admin-secret -o jsonpath="{.data.password}" | base64 -d && echo

# ubectl get pods -n argocd -l app.kubernetes.io/name=argocd-server -o name | cut -d'/' -f 2