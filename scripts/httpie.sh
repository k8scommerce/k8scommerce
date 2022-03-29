storeKey=60f44r

# http -f POST http://api.local.k8sly.com:64001/v1/asset/1/2/image asset@../internal/storage/asset/testfiles/logo.png Store-Key:yxvzrvla
# http -f POST http://api.local.k8sly.com:64001/v1/asset/1/2/image asset@../internal/storage/asset/testfiles/Pizigani_1367_Chart_10MB.jpeg Store-Key:yxvzrvla
# http -f POST http://api.local.k8sly.com:64001/v1/asset/1/2/image asset@../internal/storage/asset/testfiles/galaxy-gbf877d301_1920.jpeg Store-Key:yxvzrvla

# http GET http://localhost:8888/v1/categories Store-Key:yxvzrvla
# http GET http://localhost:8888/v1/products/1/0/12?sortOn= Store-Key:yxvzrvla

# echo '{ "email": "alma.tuck@gmail.com" }' | http -v POST http://api.local.k8sly.com:8888/v1/customer/resend-confirm-email Store-Key:$storeKey

# http -v http://api.local.k8sly.com:64001/v1/store/generate-token/1

# http -v http://api.local.k8sly.com:64000/v1/products/0/1000 Store-Key:$storeKey

# http -v http://api.local.k8sly.com:64000/v1/product/slug/antapple-kickeragate-doomtwilight Store-Key:$storeKey

# http -v http://api.local.k8sly.com:64000/v1/categories Store-Key:$storeKey

# echo '{ "email": "alma.tuck@gmail.com", "password": "!Dingo12" }' | http -v POST http://api.local.k8sly.com:8888/v1/customer/login Store-Key:$storeKey

# http -vvv GET "http://api.local.k8sly.com:8888/v1/products/0/20?sortOn=p-name&sortOn=-pr-amount" Store-Key:$storeKey
# http -vvv GET http://api.local.k8sly.com:8888/v1/product/slug/arrowrazor-snoutthorn-palmleather Store-Key:$storeKey

# echo $(rawurlencode "home/rugs/all-rugs")

# echo '{ "slug": "home/rugs/all-rugs" }' | http -vvv POST http://api.local.k8sly.com:8888/v1/category/slug Store-Key:$storeKey
# http -vvv GET http://api.local.k8sly.com:8888/v1/product/1 Store-Key:$storeKey
# http -vvv GET http://api.local.k8sly.com:8888/v1/product/sku/carvercandle-oxemerald-queenblaze Store-Key:$storeKey
http -vvv GET http://api.local.k8sly.com:8888/v1/product/slug/carvercandle-oxemerald-queenblaze Store-Key:$storeKey






# echo '{ "email": "alma.tuck@gmail.com" }' | http -v POST http://api.local.k8sly.com:8888/v1/customer/email Store-Key:$storeKey
