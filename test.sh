# http -v get http://localhost:8888/api/v1/products/category/1/1/12?sort_on=name,-price


# http http://localhost:8888/api/v1/product/slug/awesome-granite-shirt


# http -v post http://localhost:8888/api/v1/user/login username=alma.tuck@gmail.com password=!Dingo12


# http -v localhost:8888/api/v1/cart/add Accept:application/json Authorization:"Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjMyNTI4ODYzODQsImlhdCI6MTY0MDAxOTMxMCwidXNlcklkIjoyfQ.wG2PXQBRQ0_whyN1KiEyPb4t516xxFVf91TFdQJYyr4"


go-wrk -c 1 -d 1 -M POST http://localhost:8888/api/v1/cart/add -H 'Content-Type: application/json' -H 'Authorization:"Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjMyNTI4ODYzODQsImlhdCI6MTY0MDAxOTMxMCwidXNlcklkIjoyfQ.wG2PXQBRQ0_whyN1KiEyPb4t516xxFVf91TFdQJYyr4"'