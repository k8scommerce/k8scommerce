.PHONY: test-all
test-all: # Run full barrage of unit tests
test-all:
	$(eval PACKAGES=$(shell cd ./internal && go list ./... | grep -Ev "handler|svc" | tr '\n' ' '))
	@go test ${PACKAGES} -ginkgo.failFast -test.parallel 10
