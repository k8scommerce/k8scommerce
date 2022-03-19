# http -f POST http://api.local.k8sly.com:64001/v1/asset/1/2/image asset@../internal/storage/asset/testfiles/logo.png Store-Key:yxvzrvla
# http -f POST http://api.local.k8sly.com:64001/v1/asset/1/2/image asset@../internal/storage/asset/testfiles/Pizigani_1367_Chart_10MB.jpeg Store-Key:yxvzrvla
# http -f POST http://api.local.k8sly.com:64001/v1/asset/1/2/image asset@../internal/storage/asset/testfiles/galaxy-gbf877d301_1920.jpeg Store-Key:yxvzrvla

# http GET http://localhost:8888/v1/categories Store-Key:yxvzrvla
# http GET http://localhost:8888/v1/products/1/0/12?sortOn= Store-Key:yxvzrvla

echo '{ "email": "test@k8scommerce.com" }' | http -v POST http://api.local.k8sly.com:64000/v1/customer/resend-confirm-email Store-Key:eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdG9yZUtleSI6MSwidXJsIjoiaHR0cDovL2xvY2FsaG9zdDo0MjAwLyIsImlzcyI6Iks4c0NvbW1lcmNlIiwiZXhwIjo0ODAxMDQ3OTM3LCJuYmYiOjE2NDc0NDc5MzcsImlhdCI6MTY0NzQ0NzkzN30.wEqu_TnSNmSjfFWlpA0YzTsMjzMTkiwXcxc1j3IgMno

# http -v http://api.local.k8sly.com:64001/v1/store/generate-token/1
