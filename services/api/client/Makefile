# Check that given variables are set and all have non-empty values,
# die with an error otherwise.
#
# Params:
#   1. Variable name(s) to test.
#   2. (optional) Error message to print.
check_defined = \
    $(strip $(foreach 1,$1, \
        $(call __check_defined,$1,$(strip $(value 2)))))
__check_defined = \
    $(if $(value $1),, \
      $(error Undefined $1$(if $2, ($2))))

GOPATH:=$(shell go env GOPATH)
# Use sed here to remove the '/' char from branch names like feature/foobar or issue/foobar
BRANCH = $$(git branch --show-current | sed -e 's/\//-/g')
# Use epoch secs in the tag so flux sorts tags
EPOCH_SECS := $$(date +%s)
# Get a short hash of the git had for building images.
HASH = $$(git rev-parse --short HEAD)
TAG := $(shell echo $(BRANCH)-$(EPOCH_SECS)-$(HASH))
# TS = $$(date -u +'%Y-%m-%dT%H:%M:%SZ')

$(call check_defined, BRANCH HASH TAG)

####### BEGIN DOCKER #######
AWS_REGION = ${AWS_ECR_REGION}
# We don't have production ECR repo yet, so it will error out if you set the env var $PRODUCTION
# BASE_ECR = $(if $(PRODUCTION),ERROR_UNDEFINED,$(AWS_ECR_PATH))
BASE_ECR = 127.0.0.1:5000
ECR_REPO_SHORT = client
ECR_REPO = $(BASE_ECR)/$(ECR_REPO_SHORT)

.PHONY: init
init:

.PHONY: gen
gen:
	@goctl api go -api client.api -dir .

.PHONY: update
update:
	@go get -u

.PHONY: tidy
tidy:
	@go mod tidy

.PHONY: test
test:
	@go test -v ./... -cover

.PHONY: k8s
k8s:
	@goctl kube deploy -name client -namespace ecomm -image $(ECR_REPO):$(TAG) -o ../../k8s/gitops/infrastructure/api-client/sources/client.yaml -port 8080

.PHONY: build
build:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o app *.go
	go build -o app-mac *.go

.PHONY: docker
docker:
	docker build -t $(ECR_REPO):$(BRANCH)-latest -t $(ECR_REPO):$(TAG) -t $(ECR_REPO):latest .

.PHONY: push
push: test build docker
	docker push $(ECR_REPO):$(BRANCH)-latest
	docker push $(ECR_REPO):$(TAG)
	docker push $(ECR_REPO):latest
	@rm app

