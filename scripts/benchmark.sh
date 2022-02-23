#!/bin/bash

# brew install httpie
# http -vvv GET http://api.local.k8sly.com:8889/v1/product/slug/ergonomic-steel-hat \
#     Store-Key:yxvzrvla
#     Content-Type:application/json

# go install github.com/tsliwowicz/go-wrk@latest
# go-wrk -c 5 -d 10 -H 'Store-Key:yxvzrvla' http://52.9.133.84:8888/v1/product/slug/incredible-steel-hat

# brew install plow
# https://golangexample.com/a-http-s-benchmarking-tool-written-in-golang/
# when running open http://127.0.0.1:18888/ to see live stats
# plow http://52.9.133.84:8080/v1/product/slug/awesome-cotton-pants -c 1000 -d 30s -T 'application/json' -H 'Store-Key:yxvzrvla' -m GET


# https://woocommerce.com/products/woocommerce-api-manager/
# ab -n 500 -c 1 -H 'Store-Key:yxvzrvla' http://52.9.133.84:8080/v1/product/slug/awesome-cotton-pants
# ab -n 500 -c 10 -H 'Store-Key:yxvzrvla' http://52.9.133.84:8080/v1/product/slug/awesome-cotton-pants
# ab -n 500 -c 100 -H 'Store-Key:yxvzrvla' http://52.9.133.84:8080/v1/product/slug/awesome-cotton-pants
# ab -n 94212 -c 500 -H 'Store-Key:yxvzrvla' http://52.9.133.84:8080/v1/product/slug/awesome-cotton-pants
# ab -n 500 -c 300 -H 'Store-Key:yxvzrvla' http://52.9.133.84:8080/v1/product/slug/awesome-cotton-pants
# ab -n 500 -c 400 -H 'Store-Key:yxvzrvla' http://52.9.133.84:8080/v1/product/slug/awesome-cotton-pants
# ab -n 100000 -c 2500 -H 'Store-Key:yxvzrvla' http://52.9.133.84:8080/v1/product/slug/awesome-cotton-pants
# ab -n 1000 -c 1000 -H 'Store-Key:yxvzrvla' http://52.9.133.84:8080/v1/product/slug/awesome-cotton-pants
# ab -n 2000 -c 2000 -H 'Store-Key:yxvzrvla' http://52.9.133.84:8080/v1/product/slug/awesome-cotton-pants


# http GET http://52.9.133.84:8080/v1/product/slug/awesome-cotton-pants Store-Key:yxvzrvla
# http GET http://localhost:64001/v1/product/slug/awesome-cotton-pants Store-Key:yxvzrvla


http -vvv GET "http://localhost:64001/v1/products/0/20?filter=p-name-like-Practica&sortOn=p-name&sortOn=-pr-amount" Store-Key:yxvzrvla

# http -vvv GET "http://52.9.133.84:8080/v1/products/0/20?filter=p-name-like-Practica&filter=p-name-nlike-Cotton&filter=v-sku-eq-practical-cotton-gloves&sortOn=p-name&sortOn=-pr-amount" Store-Key:yxvzrvla
 



#   http://52.9.133.84:8888/

# http -vvv GET "http://api.local.k8sly.com:8888/v1/products/0/20?filter=p-name-like-practical&sortOn=p-id" Store-Key:yxvzrvla


# http -vvv GET "http://52.9.133.84:8888/v1/product/slug/incredible-steel-hat" Store-Key:yxvzrvla


#  ?cars[]=Saab&cars[]=Audi