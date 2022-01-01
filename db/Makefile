
# Makefile for building Goose + Migration SQL scripts Docker image

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

# Use sed here to remove the '/' char from branch names like feature/foobar or issue/foobar
BRANCH = $$(git branch --show-current | sed -e 's/\//-/g')
# Use epoch secs in the tag so flux sorts tags
EPOCH_SECS := $$(date +%s)
# Get a short hash of the git had for building images.
HASH = $$(git rev-parse --short HEAD)
TAG := $(shell echo $(BRANCH)-$(EPOCH_SECS)-$(HASH))
TS = $$(date -u +'%Y-%m-%dT%H:%M:%SZ')

# Fail this build if any of these are undefined so we don't push images without proper tags and timestamps
$(call check_defined, BRANCH HASH TAG TS)

####### BEGIN DOCKER #######
AWS_REGION = us-east-1
# We don't have production ECR repo yet, so it will error out if you set the env var $PRODUCTION
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

.PHONY: restart
restart:
	rm -rf ~/docker/volumes/localrivet && docker-compose up --force
