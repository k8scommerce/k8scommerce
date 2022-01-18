SHELL := bash

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

$(call check_defined, BRANCH HASH TAG)

# IMAGE_REPO=127.0.0.1:5000
IMAGE_REPO=k8scommerce


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
archetype

apiServices=client admin

# cart depends on inventory, othersbought

rpcServices=customer\
email\
inventory\
othersbought\
cart\
payment\
shipping\
similarproducts\
store\
user\
warehouse \
catalog

# rpcServices=inventory\
# othersbought\
# cart


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
	@xo schema 'pgsql://${DB_CONN_STR}' \
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
		goctl model pg datasource r --url='postgres://${DB_CONN_STR}' --table=$$table --dir=./internal/repos -style=go_zero; \
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
		(sed -i '' -e "s/:8080/:$$port/g" etc/local-$$service.yaml &); \
		(go run . -f etc/local-$$service.yaml &); \
		cd ../../../; \
		echo ""; \
		port=$$((port+1)); \
	done
	@sleep 15
	@port=8888; \
	for service in $(apiServices); do \
		printf "$(BLUE)Starting API Service: $(WHITE)$$service::$$port$(RESET)\n"; \
		cd ./services/api/$$service; \
		(cp ./etc/$$service.yaml ./etc/local-$$service.yaml &); \
		(sleep 0.1); \
		(sed -i '' -e "s/: 8888/: $$port/g" etc/local-$$service.yaml &); \
		(go run . -f etc/local-$$service.yaml &); \
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
	@for service in $(apiServices); do \
		printf "$(BLUE)Stopping API Service: $(WHITE)$$service$(RESET)\n"; \
		pkill -9 -f $$service.yaml; \
		rm ./services/api/$$service/etc/local-$$service.yaml; \
		echo ""; \
	done
	@killall etcd

##
## Admin
##
.PHONY: docker-build-admin
docker-build-admin:
ifeq ($(BRANCH), $(RELEASE_BRANCH))
	docker build -f Dockerfile.prod \
		-t $(IMAGE_REPO)/admin:latest \
		-t $(IMAGE_REPO)/admin:$(TAG) \
		--build-arg APP_NAME=admin \
		--build-arg APP_PATH=services/api \
		.
endif

.PHONY: docker-push-admin
docker-push-admin:
ifeq ($(BRANCH), $(RELEASE_BRANCH))
	docker push $(IMAGE_REPO)/admin:latest
	docker push $(IMAGE_REPO)/admin:$(TAG)
endif

##
## Client
##
.PHONY: docker-build-client
docker-build-client:
ifeq ($(BRANCH), $(RELEASE_BRANCH))
	docker build -f Dockerfile.prod \
		-t $(IMAGE_REPO)/client:latest \
		-t $(IMAGE_REPO)/client:$(TAG) \
		--build-arg APP_NAME=client \
		--build-arg APP_PATH=services/api \
		.
endif

.PHONY: docker-push-client
docker-push-client:
ifeq ($(BRANCH), $(RELEASE_BRANCH))
	docker push $(IMAGE_REPO)/client:latest
	docker push $(IMAGE_REPO)/client:$(TAG)
endif

##
## Cart
##
.PHONY: docker-build-cart
docker-build-cart:
ifeq ($(BRANCH), $(RELEASE_BRANCH))
	docker build -f Dockerfile.prod \
		-t $(IMAGE_REPO)/cart:latest \
		-t $(IMAGE_REPO)/cart:$(TAG) \
		--build-arg APP_NAME=cart \
		--build-arg APP_PATH=services/rpc \
		.
endif

.PHONY: docker-push-cart
docker-push-cart:
ifeq ($(BRANCH), $(RELEASE_BRANCH))
	docker push $(IMAGE_REPO)/cart:latest
	docker push $(IMAGE_REPO)/cart:$(TAG)
endif

##
## Catalog
##
.PHONY: docker-build-catalog
docker-build-catalog:
ifeq ($(BRANCH), $(RELEASE_BRANCH))
	docker build -f Dockerfile.prod \
		-t $(IMAGE_REPO)/catalog:latest \
		-t $(IMAGE_REPO)/catalog:$(TAG) \
		--build-arg APP_NAME=catalog \
		--build-arg APP_PATH=services/rpc \
		.
endif

.PHONY: docker-push-catalog
docker-push-catalog:
ifeq ($(BRANCH), $(RELEASE_BRANCH))
	docker push $(IMAGE_REPO)/catalog:latest
	docker push $(IMAGE_REPO)/catalog:$(TAG)
endif

##
## Customer
##
.PHONY: docker-build-customer
docker-build-customer:
ifeq ($(BRANCH), $(RELEASE_BRANCH))
	docker build -f Dockerfile.prod \
		-t $(IMAGE_REPO)/customer:latest \
		-t $(IMAGE_REPO)/customer:$(TAG) \
		--build-arg APP_NAME=customer \
		--build-arg APP_PATH=services/rpc \
		.
endif

.PHONY: docker-push-customer
docker-push-customer:
ifeq ($(BRANCH), $(RELEASE_BRANCH))
	docker push $(IMAGE_REPO)/customer:latest
	docker push $(IMAGE_REPO)/customer:$(TAG)
endif

##
## Email
##
.PHONY: docker-build-email
docker-build-email:
ifeq ($(BRANCH), $(RELEASE_BRANCH))
	docker build -f Dockerfile.prod \
		-t $(IMAGE_REPO)/email:latest \
		-t $(IMAGE_REPO)/email:$(TAG) \
		--build-arg APP_NAME=email \
		--build-arg APP_PATH=services/rpc \
		.
endif

.PHONY: docker-push-email
docker-push-email:
ifeq ($(BRANCH), $(RELEASE_BRANCH))
	docker push $(IMAGE_REPO)/email:latest
	docker push $(IMAGE_REPO)/email:$(TAG)
endif

##
## Inventory
##
.PHONY: docker-build-inventory
docker-build-inventory:
ifeq ($(BRANCH), $(RELEASE_BRANCH))
	docker build -f Dockerfile.prod \
		-t $(IMAGE_REPO)/inventory:latest \
		-t $(IMAGE_REPO)/inventory:$(TAG) \
		--build-arg APP_NAME=inventory \
		--build-arg APP_PATH=services/rpc \
		.
endif

.PHONY: docker-push-inventory
docker-push-inventory:
ifeq ($(BRANCH), $(RELEASE_BRANCH))
	docker push $(IMAGE_REPO)/inventory:latest
	docker push $(IMAGE_REPO)/inventory:$(TAG)
endif

##
## othersbought
##
.PHONY: docker-build-othersbought
docker-build-othersbought:
ifeq ($(BRANCH), $(RELEASE_BRANCH))
	docker build -f Dockerfile.prod \
		-t $(IMAGE_REPO)/othersbought:latest \
		-t $(IMAGE_REPO)/othersbought:$(TAG) \
		--build-arg APP_NAME=othersbought \
		--build-arg APP_PATH=services/rpc \
		.
endif

.PHONY: docker-push-othersbought
docker-push-othersbought:
ifeq ($(BRANCH), $(RELEASE_BRANCH))
	docker push $(IMAGE_REPO)/othersbought:latest
	docker push $(IMAGE_REPO)/othersbought:$(TAG)
endif

##
## Payment
##
.PHONY: docker-build-payment
docker-build-payment:
ifeq ($(BRANCH), $(RELEASE_BRANCH))
	docker build -f Dockerfile.prod \
		-t $(IMAGE_REPO)/payment:latest \
		-t $(IMAGE_REPO)/payment:$(TAG) \
		--build-arg APP_NAME=payment \
		--build-arg APP_PATH=services/rpc \
		.
endif

.PHONY: docker-push-payment
docker-push-payment:
ifeq ($(BRANCH), $(RELEASE_BRANCH))
	docker push $(IMAGE_REPO)/payment:latest
	docker push $(IMAGE_REPO)/payment:$(TAG)
endif

##
## Shipping
##
.PHONY: docker-build-shipping
docker-build-shipping:
ifeq ($(BRANCH), $(RELEASE_BRANCH))
	docker build -f Dockerfile.prod \
		-t $(IMAGE_REPO)/shipping:latest \
		-t $(IMAGE_REPO)/shipping:$(TAG) \
		--build-arg APP_NAME=shipping \
		--build-arg APP_PATH=services/rpc \
		.
endif

.PHONY: docker-push-shipping
docker-push-shipping:
ifeq ($(BRANCH), $(RELEASE_BRANCH))
	docker push $(IMAGE_REPO)/shipping:latest
	docker push $(IMAGE_REPO)/shipping:$(TAG)
endif

##
## similarproducts
##
.PHONY: docker-build-similarproducts
docker-build-similarproducts:
ifeq ($(BRANCH), $(RELEASE_BRANCH))
	docker build -f Dockerfile.prod \
		-t $(IMAGE_REPO)/similarproducts:latest \
		-t $(IMAGE_REPO)/similarproducts:$(TAG) \
		--build-arg APP_NAME=similarproducts \
		--build-arg APP_PATH=services/rpc \
		.
endif

.PHONY: docker-push-similarproducts
docker-push-similarproducts:
ifeq ($(BRANCH), $(RELEASE_BRANCH))
	docker push $(IMAGE_REPO)/similarproducts:latest
	docker push $(IMAGE_REPO)/similarproducts:$(TAG)
endif

##
## Store
##
.PHONY: docker-build-store
docker-build-store:
ifeq ($(BRANCH), $(RELEASE_BRANCH))
	docker build -f Dockerfile.prod \
		-t $(IMAGE_REPO)/store:latest \
		-t $(IMAGE_REPO)/store:$(TAG) \
		--build-arg APP_NAME=store \
		--build-arg APP_PATH=services/rpc \
		.
endif

.PHONY: docker-push-store
docker-push-store:
ifeq ($(BRANCH), $(RELEASE_BRANCH))
	docker push $(IMAGE_REPO)/store:latest
	docker push $(IMAGE_REPO)/store:$(TAG)
endif

##
## user
##
.PHONY: docker-build-user
docker-build-user:
ifeq ($(BRANCH), $(RELEASE_BRANCH))
	docker build -f Dockerfile.prod \
		-t $(IMAGE_REPO)/user:latest \
		-t $(IMAGE_REPO)/user:$(TAG) \
		--build-arg APP_NAME=user \
		--build-arg APP_PATH=services/rpc \
		.
endif

.PHONY: docker-push-user
docker-push-user:
ifeq ($(BRANCH), $(RELEASE_BRANCH))
	docker push $(IMAGE_REPO)/user:latest
	docker push $(IMAGE_REPO)/user:$(TAG)
endif

##
## Warehouse
##
.PHONY: docker-build-warehouse
docker-build-warehouse:
ifeq ($(BRANCH), $(RELEASE_BRANCH))
	docker build -f Dockerfile.prod \
		-t $(IMAGE_REPO)/warehouse:latest \
		-t $(IMAGE_REPO)/warehouse:$(TAG) \
		--build-arg APP_NAME=warehouse \
		--build-arg APP_PATH=services/rpc \
		.
endif

.PHONY: docker-push-warehouse
docker-push-warehouse:
ifeq ($(BRANCH), $(RELEASE_BRANCH))
	docker push $(IMAGE_REPO)/warehouse:latest
	docker push $(IMAGE_REPO)/warehouse:$(TAG)
endif

##
##
##

.PHONY: docker-build
docker-build: docker-build-admin docker-build-client docker-build-cart docker-build-catalog docker-build-customer docker-build-email docker-build-inventory docker-build-othersbought docker-build-payment docker-build-shipping docker-build-similarproducts docker-build-store docker-build-user docker-build-warehouse

.PHONY: docker-push
docker-push: docker-push-admin docker-push-client docker-push-cart docker-push-catalog docker-push-customer docker-push-email docker-push-inventory docker-push-othersbought docker-push-payment docker-push-shipping docker-push-similarproducts docker-push-store docker-push-user docker-push-warehouse
