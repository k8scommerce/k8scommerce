SHELL:=bash

envfile=.env
include ${envfile}
export $(shell sed 's/=.*//' ${envfile})

check_defined = \
    $(strip $(foreach 1,$1, \
        $(call __check_defined,$1,$(strip $(value 2)))))
__check_defined = \
    $(if $(value $1),, \
      $(error Undefined $1$(if $2, ($2))))

GOPATH:=$(shell go env GOPATH)
BRANCH:=$(shell git branch --show-current | sed -e 's/\//-/g')
RELEASE_BRANCH:=main
HASH:=$$(git rev-parse --short HEAD)
# TAG := $(shell echo $(BRANCH)-$(HASH))
TAG := $(shell echo $(HASH))

# TS = $$(date -u +'%Y-%m-%dT%H:%M:%SZ')

VERSION=0.0.1


$(call check_defined, BRANCH HASH TAG)

# if the release branch is main push to prod
ifeq ($(BRANCH), $(RELEASE_BRANCH))
IMAGE_REPO=k8scommerce
else
IMAGE_REPO=127.0.0.1:5000
endif


remove=sf_*.go\
sf_*.go\
sp_*.go\
*column.go\
goose*.go\
spatialrefsy*.go\
htally*.go\
geographycolumn*.go\
geometrycolumn*.go

$(value variable)

tablesForModelGen=product\
variant\
store\
category\
price\
property\
option\
option_item\
archetype\
asset

apiServices=client\
admin

# cart depends on inventory, othersbought

# if developing, only enable the services you're not writing code for
# for example if you're working on the customer service:
# 	1. comment out customer below 
# 	2. start the customer microservice independently
# 		- cd services/rpc/cart
#		- go run customer.go -f etc/customer.yaml
# 	3. start the remaining services & api gateways
#		- cd to the project root
# 		- make start
# 	4. to shut down all the services, except the one you're working on
#		- cd to the project root
# 		- make stop
#	5. to shut down the service(s) you're working on
# 		- press CTRL+C
#

rpcServices=customer\
email\
inventory\
othersbought\
payment\
shipping\
similarproducts\
store\
user\
warehouse\
catalog\
cart

consumerServices=imageresizer

# define standard colors
ifneq (,$(findstring xterm,${TERM}))
	BLACK        := $(shell tput -Txterm setaf 0)
	RED          := $(shell tput -Txterm setaf 1)
	GREEN        := $(shell tput -Txterm setaf 2)
	YELLOW       := $(shell tput -Txterm setaf 3)
	LIGHTPURPLE  := $(shell tput -Txterm setaf 4)
	PURPLE       := $(shell tput -Txterm setaf 5)
	BLUE         := $(shell tput -Txterm setaf 6)
	WHITE        := $(shell tput -Txterm setaf 7)
	RESET 		 := $(shell tput -Txterm sgr0)
else
	BLACK        := ""
	RED          := ""
	GREEN        := ""
	YELLOW       := ""
	LIGHTPURPLE  := ""
	PURPLE       := ""
	BLUE         := ""
	WHITE        := ""
	RESET        := ""
endif


# determine the platform
# ifeq ($(OS),Windows_NT)
   
# else
#     UNAME_S := $(shell uname -s)
#     ifeq ($(UNAME_S),Linux)
       
#     endif
#     ifeq ($(UNAME_S),Darwin)
    
#     endif
# endif

.PHONY: prereq
prereq:
	brew install etcd

.PHONY: test
test:
	@echo ${envfile}
	@cd ./services/api/admin; make test-all
	@cd ../
	@cd ./services/api/client; make test-all
	@cd ../
	@cd ./services/rpc/customer; make test-all
	@cd ../
	@cd ./services/rpc/email; make test-all
	@cd ../
	@cd ./services/rpc/inventory; make test-all
	@cd ../
	@cd ./services/rpc/othersbought; make test-all
	@cd ../
	@cd ./services/rpc/cart; make test-all
	@cd ../
	@cd ./services/rpc/payment; make test-all
	@cd ../
	@cd ./services/rpc/shipping; make test-all
	@cd ../
	@cd ./services/rpc/similarproducts; make test-all
	@cd ../
	@cd ./services/rpc/store; make test-all
	@cd ../
	@cd ./services/rpc/user; make test-all
	@cd ../
	@cd ./services/rpc/warehouse; make test-all
	@cd ../

.PHONY: cleanup
cleanup:
	@for pattern in $(remove); do \
		find ./ -type f -name $$pattern -exec rm {} +; \
	done
.PHONY: generate-xo
generate-xo:
	@xo schema '${POSTGRES_DSN}' \
	--go-field-tag='`json:"{{ .SQLName }}" db:"{{ .SQLName }}"`' \
	-o ./internal/models \
	-e *.created_at \
	-e *.updated_at \
	-e *.deleted_at \
	-k smart

.PHONY: xo
xo: generate-xo cleanup

.PHONY: model-gen
model-gen:
	@for table in $(tablesForModelGen); do \
		printf "$(BLUE)Generating code files for table:$(RESET) $(WHITE)$$table$(RESET)\n"; \
		base="$${table%.*}" ; \
		newbase="$${base//[^a-zA-Z0-9 ]/_}" ; \
		filename="$${newbase//_/}.go"; \
		path="./internal/repos"; \
		if [ ! -e $$path/$$filename ]; then \
			echo "package repos" > $$path/$$filename; \
		else \
			printf "$(YELLOW)$$filename already exists, ignored.$(RESET)\n"; \
		fi; \
		goctl model pg datasource r --url='postgres://${POSTGRES_DSN}' --table=$$table --dir=./internal/repos -style=go_zero; \
		echo ""; \
	done

.PHONY: client-ts
client-ts:
	@goctl api ts \
		-dir="./tscode/client" \
	 	-api="./services/api/client/client.api" \
		-caller="this.http"

.PHONY: admin-ts
admin-ts:
	@goctl api ts \
		-dir="./tscode/admin" \
	 	-api="./services/api/admin/admin.api" \
		-caller="this.http"

.PHONY: ts
ts: admin-ts client-ts

.PHONY: update
update:
	@go get -u

.PHONY: tidy
tidy:
	@go mod tidy

.PHONY: start
start:
	@etcd > /dev/null 2>&1 &
	@port=65000; \
	for service in $(rpcServices); do \
		printf "$(BLUE)Starting RPC Service: $(WHITE)$$service::$$port$(RESET)\n"; \
		cd ./services/rpc/$$service; \
		(cp ./etc/$$service.yaml ./etc/local-$$service.yaml &); \
		(sleep 0.1); \
		(sed -i -e "s/:8080/:$$port/g" etc/local-$$service.yaml &); \
		(go run . -f etc/local-$$service.yaml -e ../../../.env &); \
		cd ../../../; \
		echo ""; \
		port=$$((port+1)); \
	done
	@for service in $(consumerServices); do \
		printf "$(BLUE)Starting Consumer Service: $(WHITE)$$service$(RESET)\n"; \
		cd ./services/consumer/$$service; \
		(cp ./etc/$$service.yaml ./etc/local-$$service.yaml &); \
		(sleep 0.1); \
		(go run . -f etc/local-$$service.yaml -e ../../../.env &); \
		cd ../../../; \
		echo ""; \
	done
	@sleep 15
	@port=64000; \
	for service in $(apiServices); do \
		printf "$(BLUE)Starting API Service: $(WHITE)$$service::$$port$(RESET)\n"; \
		cd ./services/api/$$service; \
		(cp ./etc/$$service.yaml ./etc/local-$$service.yaml &); \
		(sleep 0.1); \
		(sed -i -e "s/: 8888/: $$port/g" etc/local-$$service.yaml &); \
		(go run . -f etc/local-$$service.yaml -e ../../../.env &); \
		cd ../../../; \
		echo ""; \
		port=$$((port+1)); \
	done
.PHONY: stop
stop:
	@for service in $(rpcServices); do \
		printf "$(BLUE)Stopping RPC Service: $(WHITE)$$service$(RESET)\n"; \
		pkill -9 -f $$service.yaml; \
		rm ./services/rpc/$$service/etc/local-$$service.yaml; \
		echo ""; \
	done
	@for service in $(consumerServices); do \
		printf "$(BLUE)Stopping Consumer Service: $(WHITE)$$service$(RESET)\n"; \
		pkill -9 -f $$service.yaml; \
		rm ./services/consumer/$$service/etc/local-$$service.yaml; \
		echo ""; \
	done
	@for service in $(apiServices); do \
		printf "$(BLUE)Stopping API Service: $(WHITE)$$service$(RESET)\n"; \
		pkill -9 -f $$service.yaml; \
		rm ./services/api/$$service/etc/local-$$service.yaml; \
		echo ""; \
	done
	@killall etcd


.PHONY: build-all
build-all:
	@port=65000; \
	for service in $(rpcServices); do \
		printf "$(BLUE)Building RPC Service: $(WHITE)$$service::$$port$(RESET)\n"; \
		cd ./services/rpc/$$service; \
		(mkdir -p ../../../bin/build/darwin/$(VERSION)/k8scommerce/etc > /dev/null 2>&1); \
		(mkdir -p ../../../bin/build/linux/$(VERSION)/k8scommerce/etc > /dev/null 2>&1); \
		(mkdir -p ../../../bin/build/windows/$(VERSION)/k8scommerce/etc > /dev/null 2>&1); \
		(cp ./etc/$$service.yaml ../../../bin/build/darwin/$(VERSION)/k8scommerce/etc/ &); \
		(sleep 0.1); \
		(sed -i -e "s/:8080/:$$port/g" ../../../bin/build/darwin/$(VERSION)/k8scommerce/etc/$$service.yaml &); \
		(cp ./etc/$$service.yaml ../../../bin/build/linux/$(VERSION)/k8scommerce/etc/ &); \
		(sleep 0.1); \
		(sed -i -e "s/:8080/:$$port/g" ../../../bin/build/linux/$(VERSION)/k8scommerce/etc/$$service.yaml &); \
		(cp ./etc/$$service.yaml ../../../bin/build/windows/$(VERSION)/k8scommerce/etc/ &); \
		(sleep 0.1); \
		(sed -i -e "s/:8080/:$$port/g" ../../../bin/build/windows/$(VERSION)/k8scommerce/etc/$$service.yaml &); \
		(GOARCH=amd64 GOOS=darwin go build -o ../../../bin/build/darwin/$(VERSION)/k8scommerce/$$service); \
		(GOARCH=amd64 GOOS=linux go build -o ../../../bin/build/linux/$(VERSION)/k8scommerce/$$service); \
		(GOARCH=amd64 GOOS=windows go build -o ../../../bin/build/windows/$(VERSION)/k8scommerce/$$service.exe); \
		cd ../../../; \
		echo ""; \
		port=$$((port+1)); \
	done
	@port=80; \
	for service in $(apiServices); do \
		printf "$(BLUE)Building API Service: $(WHITE)$$service::$$port$(RESET)\n"; \
		cd ./services/api/$$service; \
		(cp ./etc/$$service.yaml ../../../bin/build/darwin/$(VERSION)/k8scommerce/etc/ &); \
		(sleep 0.1); \
		(sed -i -e "s/: 8888/: $$port/g" ../../../bin/build/darwin/$(VERSION)/k8scommerce/etc/$$service.yaml &); \
		(cp ./etc/$$service.yaml ../../../bin/build/linux/$(VERSION)/k8scommerce/etc/ &); \
		(sleep 0.1); \
		(sed -i -e "s/: 8888/: $$port/g" ../../../bin/build/linux/$(VERSION)/k8scommerce/etc/$$service.yaml &); \
		(cp ./etc/$$service.yaml ../../../bin/build/windows/$(VERSION)/k8scommerce/etc/ &); \
		(sleep 0.1); \
		(sed -i -e "s/: 8888/: $$port/g" ../../../bin/build/windows/$(VERSION)/k8scommerce/etc/$$service.yaml &); \
		(GOARCH=amd64 GOOS=darwin go build -o ../../../bin/build/darwin/$(VERSION)/k8scommerce/$$service .); \
		(GOARCH=amd64 GOOS=linux go build -o ../../../bin/build/linux/$(VERSION)/k8scommerce/$$service .); \
		(GOARCH=amd64 GOOS=windows go build -o ../../../bin/build/windows/$(VERSION)/k8scommerce/$$service.exe .); \
		cd ../../../; \
		echo ""; \
		port=$$((port+8000)); \
	done

.PHONY: build-release
build-release:
	(cd ./bin/build/darwin/$(VERSION) && tar -C . -czvf k8scommerce-$(VERSION)-darwin.tar.gz k8scommerce && mv k8scommerce-$(VERSION)-darwin.tar.gz ../../../release/)
	(cd ./bin/build/linux/$(VERSION) && tar -C . -czvf k8scommerce-$(VERSION)-linux.tar.gz k8scommerce && mv k8scommerce-$(VERSION)-linux.tar.gz ../../../release/)
	(cd ./bin/build/windows/$(VERSION) && tar -C . -czvf k8scommerce-$(VERSION)-windows.tar.gz k8scommerce && mv k8scommerce-$(VERSION)-windows.tar.gz ../../../release/)

##
## Admin
##
.PHONY: docker-build-admin
docker-build-admin:
	docker build -f Dockerfile.prod \
		-t $(IMAGE_REPO)/admin:latest \
		-t $(IMAGE_REPO)/admin:$(TAG) \
		--build-arg APP_NAME=admin \
		--build-arg APP_PATH=services/api \
		.

.PHONY: docker-push-admin
docker-push-admin:
	docker push $(IMAGE_REPO)/admin:latest
	docker push $(IMAGE_REPO)/admin:$(TAG)

##
## Client
##
.PHONY: docker-build-client
docker-build-client:
	docker build -f Dockerfile.prod \
		-t $(IMAGE_REPO)/client:latest \
		-t $(IMAGE_REPO)/client:$(TAG) \
		--build-arg APP_NAME=client \
		--build-arg APP_PATH=services/api \
		.

.PHONY: docker-push-client
docker-push-client:
	docker push $(IMAGE_REPO)/client:latest
	docker push $(IMAGE_REPO)/client:$(TAG)

##
## Cart
##
.PHONY: docker-build-cart
docker-build-cart:
	docker build -f Dockerfile.prod \
		-t $(IMAGE_REPO)/cart:latest \
		-t $(IMAGE_REPO)/cart:$(TAG) \
		--build-arg APP_NAME=cart \
		--build-arg APP_PATH=services/rpc \
		.

.PHONY: docker-push-cart
docker-push-cart:
	docker push $(IMAGE_REPO)/cart:latest
	docker push $(IMAGE_REPO)/cart:$(TAG)

##
## Catalog
##
.PHONY: docker-build-catalog
docker-build-catalog:
	docker build -f Dockerfile.prod \
		-t $(IMAGE_REPO)/catalog:latest \
		-t $(IMAGE_REPO)/catalog:$(TAG) \
		--build-arg APP_NAME=catalog \
		--build-arg APP_PATH=services/rpc \
		.

.PHONY: docker-push-catalog
docker-push-catalog:
	docker push $(IMAGE_REPO)/catalog:latest
	docker push $(IMAGE_REPO)/catalog:$(TAG)

##
## Customer
##
.PHONY: docker-build-customer
docker-build-customer:
	docker build -f Dockerfile.prod \
		-t $(IMAGE_REPO)/customer:latest \
		-t $(IMAGE_REPO)/customer:$(TAG) \
		--build-arg APP_NAME=customer \
		--build-arg APP_PATH=services/rpc \
		.

.PHONY: docker-push-customer
docker-push-customer:
	docker push $(IMAGE_REPO)/customer:latest
	docker push $(IMAGE_REPO)/customer:$(TAG)

##
## Email
##
.PHONY: docker-build-email
docker-build-email:
	docker build -f Dockerfile.prod \
		-t $(IMAGE_REPO)/email:latest \
		-t $(IMAGE_REPO)/email:$(TAG) \
		--build-arg APP_NAME=email \
		--build-arg APP_PATH=services/rpc \
		.

.PHONY: docker-push-email
docker-push-email:
	docker push $(IMAGE_REPO)/email:latest
	docker push $(IMAGE_REPO)/email:$(TAG)

##
## Inventory
##
.PHONY: docker-build-inventory
docker-build-inventory:
	docker build -f Dockerfile.prod \
		-t $(IMAGE_REPO)/inventory:latest \
		-t $(IMAGE_REPO)/inventory:$(TAG) \
		--build-arg APP_NAME=inventory \
		--build-arg APP_PATH=services/rpc \
		.

.PHONY: docker-push-inventory
docker-push-inventory:
	docker push $(IMAGE_REPO)/inventory:latest
	docker push $(IMAGE_REPO)/inventory:$(TAG)

##
## othersbought
##
.PHONY: docker-build-othersbought
docker-build-othersbought:
	docker build -f Dockerfile.prod \
		-t $(IMAGE_REPO)/othersbought:latest \
		-t $(IMAGE_REPO)/othersbought:$(TAG) \
		--build-arg APP_NAME=othersbought \
		--build-arg APP_PATH=services/rpc \
		.

.PHONY: docker-push-othersbought
docker-push-othersbought:
	docker push $(IMAGE_REPO)/othersbought:latest
	docker push $(IMAGE_REPO)/othersbought:$(TAG)

##
## Payment
##
.PHONY: docker-build-payment
docker-build-payment:
	docker build -f Dockerfile.prod \
		-t $(IMAGE_REPO)/payment:latest \
		-t $(IMAGE_REPO)/payment:$(TAG) \
		--build-arg APP_NAME=payment \
		--build-arg APP_PATH=services/rpc \
		.

.PHONY: docker-push-payment
docker-push-payment:
	docker push $(IMAGE_REPO)/payment:latest
	docker push $(IMAGE_REPO)/payment:$(TAG)

##
## Shipping
##
.PHONY: docker-build-shipping
docker-build-shipping:
	docker build -f Dockerfile.prod \
		-t $(IMAGE_REPO)/shipping:latest \
		-t $(IMAGE_REPO)/shipping:$(TAG) \
		--build-arg APP_NAME=shipping \
		--build-arg APP_PATH=services/rpc \
		.

.PHONY: docker-push-shipping
docker-push-shipping:
	docker push $(IMAGE_REPO)/shipping:latest
	docker push $(IMAGE_REPO)/shipping:$(TAG)

##
## similarproducts
##
.PHONY: docker-build-similarproducts
docker-build-similarproducts:
	docker build -f Dockerfile.prod \
		-t $(IMAGE_REPO)/similarproducts:latest \
		-t $(IMAGE_REPO)/similarproducts:$(TAG) \
		--build-arg APP_NAME=similarproducts \
		--build-arg APP_PATH=services/rpc \
		.

.PHONY: docker-push-similarproducts
docker-push-similarproducts:
	docker push $(IMAGE_REPO)/similarproducts:latest
	docker push $(IMAGE_REPO)/similarproducts:$(TAG)

##
## Store
##
.PHONY: docker-build-store
docker-build-store:
	docker build -f Dockerfile.prod \
		-t $(IMAGE_REPO)/store:latest \
		-t $(IMAGE_REPO)/store:$(TAG) \
		--build-arg APP_NAME=store \
		--build-arg APP_PATH=services/rpc \
		.

.PHONY: docker-push-store
docker-push-store:
	docker push $(IMAGE_REPO)/store:latest
	docker push $(IMAGE_REPO)/store:$(TAG)

##
## user
##
.PHONY: docker-build-user
docker-build-user:
	docker build -f Dockerfile.prod \
		-t $(IMAGE_REPO)/user:latest \
		-t $(IMAGE_REPO)/user:$(TAG) \
		--build-arg APP_NAME=user \
		--build-arg APP_PATH=services/rpc \
		.

.PHONY: docker-push-user
docker-push-user:
	docker push $(IMAGE_REPO)/user:latest
	docker push $(IMAGE_REPO)/user:$(TAG)

##
## Warehouse
##
.PHONY: docker-build-warehouse
docker-build-warehouse:
	docker build -f Dockerfile.prod \
		-t $(IMAGE_REPO)/warehouse:latest \
		-t $(IMAGE_REPO)/warehouse:$(TAG) \
		--build-arg APP_NAME=warehouse \
		--build-arg APP_PATH=services/rpc \
		.

.PHONY: docker-push-warehouse
docker-push-warehouse:
	docker push $(IMAGE_REPO)/warehouse:latest
	docker push $(IMAGE_REPO)/warehouse:$(TAG)

##
##
##

.PHONY: docker-build
docker-build: docker-build-admin docker-build-client docker-build-cart docker-build-catalog docker-build-customer docker-build-email docker-build-inventory docker-build-othersbought docker-build-payment docker-build-shipping docker-build-similarproducts docker-build-store docker-build-user docker-build-warehouse

.PHONY: docker-push
docker-push: docker-push-admin docker-push-client docker-push-cart docker-push-catalog docker-push-customer docker-push-email docker-push-inventory docker-push-othersbought docker-push-payment docker-push-shipping docker-push-similarproducts docker-push-store docker-push-user docker-push-warehouse
