# http -f POST http://api.local.k8sly.com:64001/v1/asset/1/2/image asset@../internal/storage/asset/testfiles/logo.png Store-Key:yxvzrvla
# http -f POST http://api.local.k8sly.com:64001/v1/asset/1/2/image asset@../internal/storage/asset/testfiles/Pizigani_1367_Chart_10MB.jpeg Store-Key:yxvzrvla
# http -f POST http://api.local.k8sly.com:64001/v1/asset/1/2/image asset@../internal/storage/asset/testfiles/galaxy-gbf877d301_1920.jpeg Store-Key:yxvzrvla

# http GET http://localhost:8888/v1/categories Store-Key:yxvzrvla
http GET http://localhost:8888/v1/products/1/0/12?sortOn= Store-Key:yxvzrvla
