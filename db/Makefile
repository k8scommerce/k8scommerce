
check_defined = \
    $(strip $(foreach 1,$1, \
        $(call __check_defined,$1,$(strip $(value 2)))))
__check_defined = \
    $(if $(value $1),, \
      $(error Undefined $1$(if $2, ($2))))

BRANCH = $$(git branch --show-current | sed -e 's/\//-/g')


HASH = $$(git rev-parse --short HEAD)
TAG := $(shell echo $(BRANCH)-$(HASH))
TS = $$(date -u +'%Y-%m-%dT%H:%M:%SZ')

# Fail this build if any of these are undefined so we don't push images without proper tags
$(call check_defined, BRANCH HASH TAG TS)

####### BEGIN DOCKER #######
BASE_ECR = $(if $(PRODUCTION),ERROR_UNDEFINED,127.0.0.1:5000)
ECR_REPO = $(BASE_ECR)/migrations

.PHONY: docker

# NOTE: you must have credentials for AWS -- tries to auth docker here
repo-login:
	$(eval TOKEN=$(shell aws ecr get-login-password --region ${AWS_REGION}))
	@echo ${TOKEN} | docker login --username AWS --password-stdin ${BASE_ECR}

docker:
	docker build -t $(ECR_REPO):$(BRANCH)-latest -t $(ECR_REPO):$(TAG) -t $(ECR_REPO):latest .

.PHONY: push
push: docker
	docker push $(ECR_REPO):$(BRANCH)-latest
	docker push $(ECR_REPO):$(TAG)
	docker push $(ECR_REPO):latest

