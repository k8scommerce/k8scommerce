
# KUBERNETES_REGISTRY=aws-ecr
# DOCKER_ECR=00000000000.dkr.ecr.us-east-1.amazonaws.com
# DOCKER_USERNAME=AWS
# DOCKER_EMAIL=your@email.com
# aws ecr get-login-password --region us-west-1 --profil ecomm-readonly | sed -e 's/.*-p //' -e 's/ .*$//' > ecr-token
# echo $(<./ecr-token)
# kubectl delete secrets ${KUBERNETES_REGISTRY} 2> /dev/null
# kubectl create secret docker-registry ${KUBERNETES_REGISTRY} \
# --docker-server=${DOCKER_ECR} \
# --docker-username=${DOCKER_USERNAME} \
# --docker-password=${DOCKER_SECRET} \
# --docker-email=${DOCKER_EMAIL}


# SECRET_NAME=ecr-credentials
# ECR_REGISTRY=260264107230.dkr.ecr.us-west-1.amazonaws.com

# kubectl delete secrets ${SECRET_NAME} --namespace=flux-system 2> /dev/null
# kubectl create secret docker-registry $SECRET_NAME \
#                     --docker-server="$ECR_REGISTRY" \
#                     --docker-username=AWS \
#                     --docker-password="$(<./ecr-token)" \
#                     --namespace=flux-system \
#                     --docker-email="alma.tuck@gmail.com"

# --output json | base64 --decode | jq '.payload' | tr -d '"'
# aws ecr get-login-password --region us-west-1 --profil ecomm-readonly > ./ecr-token
echo $(<./ecr-token) | docker login --username AWS --password-stdin 260264107230.dkr.ecr.us-west-1.amazonaws.com


# kubectl delete secret ecr-credentials
# kubectl create secret docker-registry ecr-credentials \
#   --docker-server=260264107230.dkr.ecr.us-west-1.amazonaws.com \
#   --docker-username=AWS \
#   --docker-password=$SECRET \
#   --docker-email=alma.tuck@gmail.com


# KUBECTL='kubectl --dry-run=client'
# KUBECTL='kubectl'

# ENVIRONMENT=260264107230
# AWS_DEFAULT_REGION=us-west-1

# # ecomm-readonly

# EXISTS=$($KUBECTL get secret "ecr-credentials" --namespace ecomm-dev | tail -n 1 | cut -d ' ' -f 1)
# if [ "$EXISTS" = "ecr-credentials" ]; then
#   echo "Secret exists, deleting"
#   $KUBECTL delete secrets "ecr-credentials" --namespace ecomm-dev
# fi

# PASS=$(aws ecr get-login-password --region $AWS_DEFAULT_REGION --profile ecomm-readonly)
# echo $PASS
# $KUBECTL create secret docker-registry ecr-credentials \
#     --docker-server=$AWS_ACCOUNT_ID.dkr.ecr.$AWS_DEFAULT_REGION.amazonaws.com \
#     --docker-username=AWS \
#     --docker-password=$PASS \
#     --docker-email=alma.tuck@gmail.com \
#     --namespace ecomm-dev

