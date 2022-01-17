# Admin Gateway API
admin gateway api

## Version: 1

### Security
**apiKey**  

|apiKey|*API Key*|
|---|---|
|Description|Enter JWT Bearer token **_only_**|
|Name|Authorization|
|In|header|

### /v1/categories

#### GET
##### Summary

Get All Categories

##### Description

returns all categories belonging to a store

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | A successful response. | [GetAllCategoriesResponse](#getallcategoriesresponse) |

### /v1/category

#### POST
##### Summary

Create Category

##### Description

creates a category

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| body | body |  | Yes | [CreateCategoryRequest](#createcategoryrequest) |

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | A successful response. | [CreateCategoryResponse](#createcategoryresponse) |

### /v1/category/slug/{slug}

#### GET
##### Summary

Get Category By Slug

##### Description

returns all categories by slug belonging to a store

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| slug | path | category slug | Yes | string |

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | A successful response. | [GetCategoryBySlugResponse](#getcategorybyslugresponse) |

### /v1/category/{id}

#### GET
##### Summary

Get Category By Id

##### Description

returns all categories by id belonging to a store

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| id | path | category id | Yes | string |
| id | query |  | Yes | long |

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | A successful response. | [GetCategoryByIdResponse](#getcategorybyidresponse) |

#### DELETE
##### Summary

Delete Category

##### Description

deletes a category

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| id | path | category id | Yes | string |
| body | body |  | Yes | [DeleteCategoryRequest](#deletecategoryrequest) |

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | A successful response. | [DeleteCategoryResponse](#deletecategoryresponse) |

#### PUT
##### Summary

Update Category

##### Description

updates a category

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| id | path | category id | Yes | string |
| body | body |  | Yes | [UpdateCategoryRequest](#updatecategoryrequest) |

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | A successful response. | [UpdateCategoryResponse](#updatecategoryresponse) |

### /v1/customer

#### POST
##### Summary

Create Customer

##### Description

creates a new customer

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| body | body |  | Yes | [CreateCustomerRequest](#createcustomerrequest) |

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | A successful response. | [CreateCustomerResponse](#createcustomerresponse) |

### /v1/customer/login

#### POST
##### Summary

Login

##### Description

login for customers

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| body | body |  | Yes | [CustomerLoginRequest](#customerloginrequest) |

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | A successful response. | [CustomerLoginResponse](#customerloginresponse) |

### /v1/product

#### POST
##### Summary

Create Product

##### Description

creates a product

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| body | body |  | Yes | [CreateProductRequest](#createproductrequest) |

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | A successful response. | [CreateProductResponse](#createproductresponse) |

### /v1/product/sku/{sku}

#### GET
##### Summary

Get Product By Sku

##### Description

returns all products by sku belonging to a store

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| sku | path | product sku | Yes | string |

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | A successful response. | [GetProductResponse](#getproductresponse) |

### /v1/product/slug/{slug}

#### GET
##### Summary

Get Product By Slug

##### Description

returns matching product by slug

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| slug | path | product slug | Yes | string |

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | A successful response. | [GetProductResponse](#getproductresponse) |

### /v1/product/{id}

#### GET
##### Summary

Get Product By Id

##### Description

returns matching product by id

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| id | path | product id | Yes | string |

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | A successful response. | [GetProductResponse](#getproductresponse) |

#### DELETE
##### Summary

Delete Product

##### Description

delete a product

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| id | path | product id | Yes | string |
| body | body |  | Yes | [DeleteProductRequest](#deleteproductrequest) |

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | A successful response. | [DeleteProductResponse](#deleteproductresponse) |

#### PUT
##### Summary

Update Product

##### Description

updates a product

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| id | path | product id | Yes | string |
| body | body |  | Yes | [UpdateProductRequest](#updateproductrequest) |

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | A successful response. | [UpdateProductResponse](#updateproductresponse) |

### /v1/products/{categoryId}/{currentPage}/{pageSize}

#### GET
##### Summary

Get Products By Category Id

##### Description

returns all products by category id belonging to a store

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| categoryId | path | category id | Yes | string |
| currentPage | path | current page number | Yes | string |
| pageSize | path | number of records per page | Yes | string |
| sortOn | query |  | No | string |

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | A successful response. | [GetProductsByCategoryIdResponse](#getproductsbycategoryidresponse) |

### /v1/products/{currentPage}/{pageSize}

#### GET
##### Summary

Get All Products

##### Description

returns all products belonging to a store

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| currentPage | path | current page number | Yes | string |
| pageSize | path | number of records per page | Yes | string |
| sortOn | query |  | Yes | string |

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | A successful response. | [GetAllProductsResponse](#getallproductsresponse) |

### /v1/user/login

#### POST
##### Summary

Login

##### Description

login for administration users

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| body | body |  | Yes | [UserLoginRequest](#userloginrequest) |

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | A successful response. | [UserLoginResponse](#userloginresponse) |

### Models

#### Address

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| street | string |  street name, ie: 1723 NW 23rd Ave. | Yes |
| city | string |  city name | Yes |
| stateProvince | string |  state or province name | Yes |
| country | string |  IISO 3166-1 alpha-2 country code. https:en.wikipedia.org/wiki/List_of_ISO_3166_country_codes | Yes |
| postalCode | string |  postal or zip code | Yes |
| isDefault | boolean (boolean) |  indicates if this is a default address | Yes |

#### Category

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| id | long |  category id | Yes |
| parentId | long |  parent category id. references Category.Id | Yes |
| slug | string |  slug name of the category | Yes |
| name | string |  name of category | Yes |
| description | string |  description of category | Yes |
| metaTitle | string |  metatag title for SEO | Yes |
| metaDescription | string |  metatag description for SEO | Yes |
| metaKeywords | string |  metatag keywords for SEO | Yes |
| sortOrder | integer |  sort order of menu items on the same level and same parent id | Yes |

#### CreateCategoryRequest

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| category | [Category](#category) |  | Yes |

#### CreateCategoryResponse

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| category | [Category](#category) |  | Yes |
| status | [ResponseStatus](#responsestatus) |  a ResponseStatus object | Yes |

#### CreateCustomerRequest

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| customer | [NewCustomer](#newcustomer) |  NewCustomer object | Yes |

#### CreateCustomerResponse

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| customer | [Customer](#customer) |  Customer object | Yes |
| status | [ResponseStatus](#responsestatus) |  a ResponseStatus object | Yes |

#### CreateProductRequest

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| product:omitempty | [Product](#product) |  | Yes |

#### CreateProductResponse

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| product:omitempty | [Product](#product) |  | Yes |
| status | [ResponseStatus](#responsestatus) |  a ResponseStatus object | Yes |

#### Customer

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| id | long |  customer id | Yes |
| firstName | string |  first name | Yes |
| lastName | string |  last or given name | Yes |
| email | string |  email address | Yes |
| password | string |  password | Yes |

#### CustomerAccount

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| id | long |  customer id | Yes |
| billingAddress | [Address](#address) |  Address object | Yes |
| shippingAddresses | [ [Address](#address) ] |  collection of Address objects | Yes |

#### CustomerLoginRequest

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| email | string |  email address, unique to each store id | Yes |
| password | string |  password | Yes |

#### CustomerLoginResponse

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| jwt | [JwtToken](#jwttoken) |  jwt token | Yes |
| customer | [Customer](#customer) |  Customer object | Yes |
| status | [ResponseStatus](#responsestatus) |  a ResponseStatus object | Yes |

#### DeleteCategoryRequest

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| id | long |  | Yes |

#### DeleteCategoryResponse

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| status | [ResponseStatus](#responsestatus) |  a ResponseStatus object | Yes |

#### DeleteProductRequest

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| path:omitempty | long |  | Yes |

#### DeleteProductResponse

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| status | [ResponseStatus](#responsestatus) |  a ResponseStatus object | Yes |

#### GetAllCategoriesResponse

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| categories | [ [Category](#category) ] |  a collection of Category | Yes |
| status | [ResponseStatus](#responsestatus) |  a ResponseStatus object | Yes |

#### GetAllProductsRequest

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| currentPage | long |  | Yes |
| pageSize | long |  | Yes |
| sortOn | string |  | Yes |

#### GetAllProductsResponse

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| products | [ [Product](#product) ] |  | Yes |
| totalRecords | long |  | Yes |
| totalPages | long |  | Yes |
| status | [ResponseStatus](#responsestatus) |  a ResponseStatus object | Yes |

#### GetCategoryByIdRequest

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| id | long |  | Yes |

#### GetCategoryByIdResponse

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| category | [Category](#category) |  | Yes |
| status | [ResponseStatus](#responsestatus) |  a ResponseStatus object | Yes |

#### GetCategoryBySlugRequest

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| slug | string |  slug name of the category | Yes |

#### GetCategoryBySlugResponse

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| category | [Category](#category) |  | Yes |
| status | [ResponseStatus](#responsestatus) |  a ResponseStatus object | Yes |

#### GetProductByIdRequest

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| id | long |  | Yes |

#### GetProductBySkuRequest

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| sku | string |  | Yes |

#### GetProductBySlugRequest

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| slug | string |  slug name of the category | Yes |

#### GetProductResponse

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| product | [Product](#product) |  slug name of the category | Yes |
| status | [ResponseStatus](#responsestatus) |  a ResponseStatus object | Yes |

#### GetProductsByCategoryIdRequest

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| categoryId | long |  | Yes |
| currentPage | long |  | Yes |
| pageSize | long |  | Yes |
| sortOn | string |  | No |

#### GetProductsByCategoryIdResponse

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| products | [ [Product](#product) ] |  | Yes |
| totalRecords | long |  | Yes |
| totalPages | long |  | Yes |
| status | [ResponseStatus](#responsestatus) |  a ResponseStatus object | Yes |

#### JwtToken

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| access_token | string |  | Yes |
| access_expire | long |  | Yes |
| refresh_after | long |  | Yes |

#### NewCustomer

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| firstName | string |  first name | Yes |
| lastName | string |  last or given name | Yes |
| email | string |  email address, unique per store id | Yes |
| password | string |  password | Yes |
| billingAddress | [Address](#address) |  Address object | No |
| shippingAddresses | [Address](#address) |  Address object | No |

#### Price

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| id | long |  price id | Yes |
| amount | double |  price amount | Yes |
| displayAmount | string |  price display amount | Yes |
| compareAtAmount | double |  price compare amount | Yes |
| displayCompareAtAmount | string |  price display compare amount | Yes |
| currency | string |  price currency. example: USD, CAN, etc. | Yes |

#### Product

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| id | long |  product id | Yes |
| slug | string |  product slug | Yes |
| name | string |  product name | Yes |
| shortDescription | string |  product short description. used in category pages | Yes |
| description | string |  category description | Yes |
| metaTitle | string |  metatag title for SEO | Yes |
| metaDescription | string |  metatag description for SEO | Yes |
| metaKeywords | string |  metatag keywords for SEO | Yes |
| variants | [ [Variant](#variant) ] |  collection of Variant objects | Yes |

#### ResponseStatus

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| statusCode | long |  RFC http status code, ie. 204, etc - https:go.dev/src/net/http/status.go | Yes |
| statusMessage | string |  status message | Yes |

#### UpdateCategoryRequest

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| id | long |  | Yes |
| category | [Category](#category) |  | Yes |

#### UpdateCategoryResponse

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| category | [Category](#category) |  | Yes |
| status | [ResponseStatus](#responsestatus) |  a ResponseStatus object | Yes |

#### UpdateProductRequest

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| path:omitempty | long |  | Yes |
| product:omitempty | [Product](#product) |  | Yes |

#### UpdateProductResponse

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| product:omitempty | [Product](#product) |  | Yes |
| status | [ResponseStatus](#responsestatus) |  a ResponseStatus object | Yes |

#### User

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| id | long |  | Yes |
| firstName | string |  | Yes |
| lastName | string |  | Yes |
| email | string |  | Yes |
| password | string |  | Yes |

#### UserLoginRequest

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| email | string |  | Yes |
| password | string |  | Yes |

#### UserLoginResponse

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| jwt | [JwtToken](#jwttoken) |  | Yes |
| user | [User](#user) |  | Yes |
| statusCode | long |  | Yes |
| statusMessage | string |  | Yes |

#### Variant

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| id | long |  variant id | Yes |
| isDefault | boolean (boolean) |  is default variant. each product must have exactly 1 default variant | Yes |
| sku | string |  variant sku. usually the product sku with appended identification tags | Yes |
| weight | double |  variant weight. used in calculating shipping | Yes |
| height | double |  variant height. used in calculating shipping | Yes |
| width | double |  variant width. used in calculating shipping | Yes |
| depth | double |  variant depth. used in calculating shipping | Yes |
| price | [Price](#price) |  variant Price object | Yes |
