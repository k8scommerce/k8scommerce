#!/bin/bash

# brew install httpie
# http -vvv GET http://api.local.k8sly.com:8889/v1/product/slug/ergonomic-steel-hat \
#     Store-Key:yxvzrvla
#     Content-Type:application/json

# go install github.com/tsliwowicz/go-wrk@latest
# go-wrk -c 125 -d 10 -H 'Store-Key: yxvzrvla' http://api.local.k8sly.com:8889/v1/product/slug/ergonomic-steel-hat

# brew install plow
# https://golangexample.com/a-http-s-benchmarking-tool-written-in-golang/
# when running open http://127.0.0.1:18888/ to see live stats
plow http://api.local.k8sly.com:8889/v1/product/slug/ergonomic-steel-hat -c 50 -d 30s -T 'application/json' -H 'Store-Key: yxvzrvla' -m GET
