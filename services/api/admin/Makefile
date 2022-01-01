GOPATH:=$(shell go env GOPATH)
VERSION:=1.0.0

.PHONY: init
init:

.PHONY: gen
gen:
	@goctl api go -api gateway.api -dir .

.PHONY: update
update:
	@go get -u

.PHONY: tidy
tidy:
	@go mod tidy

.PHONY: build
build:
	@CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o gateway *.go

.PHONY: test
test:
	@go test -v ./... -cover

.PHONY: k8s
k8s:
	@goctl kube deploy -name gateway -namespace micro-services -image localhost:5000/gateway:${VERSION} -o k8s/gateway.yaml -port 8080

.PHONY: docker
docker:
	@docker build -t gateway:${VERSION} -t localhost:5000/gateway:${VERSION} -t localhost:5000/gateway:latest .
	@docker push localhost:5000/gateway:${VERSION}
	@docker push localhost:5000/gateway:latest

