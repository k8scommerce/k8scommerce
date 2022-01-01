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
BRANCH = $(shell git branch --show-current | sed -e 's/\//-/g')
HASH = $$(git rev-parse --short HEAD)
TAG := $(shell echo $(BRANCH)-$(HASH))

# TS = $$(date -u +'%Y-%m-%dT%H:%M:%SZ')

$(call check_defined, BRANCH HASH TAG)

IMAGE_REPO = 127.0.0.1:5000

remove=sf_*.go\
sf_*.go\
sp_*.go\
*column.go\
goose*.go\
spatialrefsy*.go\
htally*.go

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


apiServices=client
rpcServices=inventory\
othersbought\
product\
similarproducts\
user\
cart

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

.PHONY: test
test:
	@echo ${envfile}
	@echo $(value COMMERCE_ADAPTER)

.PHONY: cleanup
cleanup:
	@for pattern in $(remove); do \
		find ./ -type f -name $$pattern -exec rm {} +; \
	done
.PHONY: generate-xo
generate-xo:
	@xo schema 'pgsql://${DB_CONN_STR}' \
	--go-field-tag='`json:"{{ .SQLName }}" db:"{{ .SQLName }}"`' \
	-o ./shared/models \
	-e *.created_at \
	-e *.updated_at \
	-e *.deleted_at \
	-k smart

	@go get -v -u  github.com/k8scommerce/k8scommerce/pkg 

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

.PHONY: ts
ts:
	@goctl api ts \
		-dir="$(HOME)/workspaces/localrivet/shop/src/app/data/services" \
	 	-api="./services/api/client/client.api" \
		-caller="this.http"
# import { HttpClient } from '@angular/common/http';

.PHONY: update
update:
	@go get -u

.PHONY: tidy
tidy:
	@go mod tidy

.PHONY: start
start:
	@etcd > /dev/null 2>&1 &
	@for service in $(rpcServices); do \
		printf "$(BLUE)Starting RPC Service: $(WHITE)$$service$(RESET)\n"; \
		cd ./services/rpc/$$service; \
		(go run . -f etc/$$service.yaml &); \
		cd ../../../; \
		echo ""; \
	done
	@sleep 15
	@for service in $(apiServices); do \
		printf "$(BLUE)Starting RPC Service: $(WHITE)$$service$(RESET)\n"; \
		cd ./services/api/$$service; \
		open -a Terminal "`go run . -f etc/$$service.yaml`"; \
		cd ../../../; \
		echo ""; \
	done
# xterm -hold go run . -f etc/$$service.yaml; \
.PHONY: stop
stop:
	@for service in $(rpcServices); do \
		printf "$(BLUE)Stopping RPC Service: $(WHITE)$$service$(RESET)\n"; \
		pkill -9 -f $$service.yaml; \
		echo ""; \
	done
	@for service in $(apiServices); do \
		printf "$(BLUE)Stopping RPC Service: $(WHITE)$$service$(RESET)\n"; \
		pkill -9 -f $$service.yaml; \
		echo ""; \
	done
	@killall etcd

##
## Client
##
.PHONY: docker-build-client
docker-build-client:
	docker build -f Dockerfile.prod \
		-t $(IMAGE_REPO)/client:latest \
		-t $(IMAGE_REPO)/client:$(BRANCH)-latest \
		-t $(IMAGE_REPO)/client:$(TAG) \
		--build-arg APP_NAME=client \
		--build-arg APP_PATH=services/api \
		--no-cache .

.PHONY: docker-push-client
docker-push-client:
	docker push $(IMAGE_REPO)/client:latest
	docker push $(IMAGE_REPO)/client:$(BRANCH)-latest
	docker push $(IMAGE_REPO)/client:$(TAG)

##
## Cart
##
.PHONY: docker-build-cart
docker-build-cart:
	docker build -f Dockerfile.prod \
		-t $(IMAGE_REPO)/cart:latest \
		-t $(IMAGE_REPO)/cart:$(BRANCH)-latest \
		-t $(IMAGE_REPO)/cart:$(TAG) \
		--build-arg APP_NAME=cart \
		--build-arg APP_PATH=services/rpc \
		--no-cache .

.PHONY: docker-push-cart
docker-push-cart:
	docker push $(IMAGE_REPO)/cart:latest
	docker push $(IMAGE_REPO)/cart:$(BRANCH)-latest
	docker push $(IMAGE_REPO)/cart:$(TAG)

##
## Inventory
##
.PHONY: docker-build-inventory
docker-build-inventory:
	docker build -f Dockerfile.prod \
		-t $(IMAGE_REPO)/inventory:latest \
		-t $(IMAGE_REPO)/inventory:$(BRANCH)-latest \
		-t $(IMAGE_REPO)/inventory:$(TAG) \
		--build-arg APP_NAME=inventory \
		--build-arg APP_PATH=services/rpc \
		--no-cache .

.PHONY: docker-push-inventory
docker-push-inventory:
	docker push $(IMAGE_REPO)/inventory:latest
	docker push $(IMAGE_REPO)/inventory:$(BRANCH)-latest
	docker push $(IMAGE_REPO)/inventory:$(TAG)

##
## othersbought
##
.PHONY: docker-build-othersbought
docker-build-othersbought:
	docker build -f Dockerfile.prod \
		-t $(IMAGE_REPO)/othersbought:latest \
		-t $(IMAGE_REPO)/othersbought:$(BRANCH)-latest \
		-t $(IMAGE_REPO)/othersbought:$(TAG) \
		--build-arg APP_NAME=othersbought \
		--build-arg APP_PATH=services/rpc \
		--no-cache .

.PHONY: docker-push-othersbought
docker-push-othersbought:
	docker push $(IMAGE_REPO)/othersbought:latest
	docker push $(IMAGE_REPO)/othersbought:$(BRANCH)-latest
	docker push $(IMAGE_REPO)/othersbought:$(TAG)


##
## Product
##
.PHONY: docker-build-product
docker-build-product:
	docker build -f Dockerfile.prod \
		-t $(IMAGE_REPO)/product:latest \
		-t $(IMAGE_REPO)/product:$(BRANCH)-latest \
		-t $(IMAGE_REPO)/product:$(TAG) \
		--build-arg APP_NAME=product \
		--build-arg APP_PATH=services/rpc \
		--no-cache .

.PHONY: docker-push-product
docker-push-product:
	docker push $(IMAGE_REPO)/product:latest
	docker push $(IMAGE_REPO)/product:$(BRANCH)-latest
	docker push $(IMAGE_REPO)/product:$(TAG)

##
## similarproducts
##
.PHONY: docker-build-similarproducts
docker-build-similarproducts:
	docker build -f Dockerfile.prod \
		-t $(IMAGE_REPO)/similarproducts:latest \
		-t $(IMAGE_REPO)/similarproducts:$(BRANCH)-latest \
		-t $(IMAGE_REPO)/similarproducts:$(TAG) \
		--build-arg APP_NAME=similarproducts \
		--build-arg APP_PATH=services/rpc \
		--no-cache .

.PHONY: docker-push-similarproducts
docker-push-similarproducts:
	docker push $(IMAGE_REPO)/similarproducts:latest
	docker push $(IMAGE_REPO)/similarproducts:$(BRANCH)-latest
	docker push $(IMAGE_REPO)/similarproducts:$(TAG)

##
## user
##
.PHONY: docker-build-user
docker-build-user:
	docker build -f Dockerfile.prod \
		-t $(IMAGE_REPO)/user:latest \
		-t $(IMAGE_REPO)/user:$(BRANCH)-latest \
		-t $(IMAGE_REPO)/user:$(TAG) \
		--build-arg APP_NAME=user \
		--build-arg APP_PATH=services/rpc \
		--no-cache .

.PHONY: docker-push-user
docker-push-user:
	docker push $(IMAGE_REPO)/user:latest
	docker push $(IMAGE_REPO)/user:$(BRANCH)-latest
	docker push $(IMAGE_REPO)/user:$(TAG)

##
##
##

.PHONY: docker-build
docker-build: docker-build-client docker-build-cart docker-build-inventory docker-build-othersbought docker-build-product docker-build-similarproducts docker-build-user 
	
.PHONY: docker-push
docker-push: docker-push-client docker-push-cart docker-push-inventory docker-push-othersbought docker-push-product docker-push-similarproducts docker-push-user