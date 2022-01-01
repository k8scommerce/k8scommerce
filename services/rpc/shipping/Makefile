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
ECR_REPO_SHORT = shipping
ECR_REPO = $(BASE_ECR)/$(ECR_REPO_SHORT)

.PHONY: init
init:
	@go get -u google.golang.org/protobuf/proto
	@go install github.com/golang/protobuf/protoc-gen-go@latest
	@go install github.com/asim/go-micro/cmd/protoc-gen-micro/v4@latest

.PHONY: gen
gen:
	@goctl rpc proto -proto_path=../customer -proto_path=. -src ./shipping.proto -dir .

.PHONY: proto
proto:
	@protoc --proto_path=. --micro_out=. --go_out=:. ../../../protos/shipping.proto

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
	@goctl kube deploy -name shipping -namespace ecomm -image $(ECR_REPO):$(TAG) -o ../../k8s/gitops/infrastructure/rpc-shipping/sources/shipping.yaml -port 8080

# .PHONY: docker
# docker:
# docker login -uadmin -ppassword ${BASE_ECR}

# docker build -t $(ECR_REPO):$(BRANCH)-latest -t $(ECR_REPO):$(TAG) .
# docker push $(ECR_REPO):$(BRANCH)-latest
# docker push $(ECR_REPO):$(TAG)


# NOTE: you must have credentials for AWS -- tries to auth docker here
.PHONY: repo-login
repo-login:
	$(eval TOKEN=$(shell aws ecr get-login-password --region ${AWS_ECR_REGION}))
	@echo ${TOKEN} | docker login --username AWS --password-stdin ${BASE_ECR}

.PHONY: build
build:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o app *.go
	go build -o app-mac *.go

.PHONY: docker
docker:
	docker build -t $(ECR_REPO):$(BRANCH)-latest -t $(ECR_REPO):$(TAG) -t $(ECR_REPO):latest .
	docker build --build-arg arg=2.3 .

.PHONY: push
push: test build docker
	docker push $(ECR_REPO):$(BRANCH)-latest
	docker push $(ECR_REPO):$(TAG)
	docker push $(ECR_REPO):latest
	@rm app
