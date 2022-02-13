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
# plow http://api.local.k8sly.com:8889/v1/product/slug/ergonomic-steel-hat -c 50 -d 30s -T 'application/json' -H 'Store-Key: yxvzrvla' -m GET


# http -vvv GET http://api.local.k8sly.com:8888/v1/products/0/20?sortOn=-id \
#     Store-Key:yxvzrvla
#     Content-Type:application/json


# http -vvv GET "http://api.local.k8sly.com:8888/v1/products/0/20?filter=p-name-like-Practica&filter=p-name-nlike-Cotton&filter=v-sku-eq-practical-cotton-gloves&sortOn=p-name&sortOn=-pr-amount" Store-Key:yxvzrvla
 


http -vvv GET "http://api.local.k8sly.com:8888/v1/products/0/20?filter=p-name-like-practical&sortOn=p-id" Store-Key:yxvzrvla




#  ?cars[]=Saab&cars[]=Audi