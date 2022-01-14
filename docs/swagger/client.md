# Client Gateway API
client gateway api

## Version: 1.0.0

### Security
**apiKey**  

|apiKey|*API Key*|
|---|---|
|Description|Enter JWT Bearer token **_only_**|
|Name|Authorization|
|In|header|

### /v1/cart/{customerId}

#### GET
##### Summary

returns a shopping cart if one exists

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| customerId | path |  | Yes | string |

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | A successful response. | [GetCartResponse](#getcartresponse) |

#### POST
##### Summary

creates a shopping cart for the associated customerId. Each customer can only have 1 cart ever.

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| customerId | path |  | Yes | string |
| body | body |  | Yes | [CreateCartRequest](#createcartrequest) |

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | A successful response. | [CreateCartResponse](#createcartresponse) |

### /v1/categories

#### GET
##### Summary

getAllCategories
 returns all categories by storeId

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | A successful response. | [GetAllCategoriesResponse](#getallcategoriesresponse) |

### /v1/category/slug/{slug}

#### GET
##### Summary

getCategoryBySlug
 returns a category by url slug and storeId

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| slug | path |  | Yes | string |
| slug | query |  | Yes | string |

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | A successful response. | [GetCategoryBySlugResponse](#getcategorybyslugresponse) |

### /v1/category/{id}

#### GET
##### Summary

getCategoryById
 returns a category by id

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| id | path |  | Yes | string |
| id | query |  | Yes | long |

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | A successful response. | [GetCategoryByIdResponse](#getcategorybyidresponse) |

### /v1/customer/login

#### POST
##### Summary

customer logic

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| body | body |  | Yes | [CustomerLoginRequest](#customerloginrequest) |

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | A successful response. | [CustomerLoginResponse](#customerloginresponse) |

### /v1/product/sku/{sku}

#### GET
##### Summary

getProductBySku
 returns a product by sku

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| sku | path |  | Yes | string |

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | A successful response. | [Product](#product) |

### /v1/product/slug/{slug}

#### GET
##### Summary

getProductBySlug
 returns a product by url slug

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| slug | path |  | Yes | string |

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | A successful response. | [Product](#product) |

### /v1/product/{id}

#### GET
##### Summary

getProductById
 returns a product by id

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| id | path |  | Yes | string |

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | A successful response. | [Product](#product) |

### /v1/products/category/{categoryId}/{currentPage}/{pageSize}

#### GET
##### Summary

getProductsByCategoryId
 returns all products belonging to a category id

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| categoryId | path |  | Yes | string |
| currentPage | path |  | Yes | string |
| pageSize | path |  | Yes | string |
| sortOn | query |  | No | string |

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | A successful response. | [GetProductsByCategoryIdResponse](#getproductsbycategoryidresponse) |

### /v1/products/{currentPage}/{pageSize}

#### GET
##### Summary

getAllProducts
 returns all products

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| currentPage | path |  | Yes | string |
| pageSize | path |  | Yes | string |
| sortOn | query |  | Yes | string |

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | A successful response. | [GetAllProductsResponse](#getallproductsresponse) |

### Models

#### AddItemToCartRequest

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| customerId | long |  | Yes |
| item | [Item](#item) |  | Yes |

#### AddItemToCartResponse

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| cart | [Cart](#cart) |  | Yes |
| sessionId | string |  | Yes |

#### Cart

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| items | [ [Item](#item) ] |  | Yes |
| totalPrice | double |  | Yes |

#### Category

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| id | long |  | Yes |
| parentId | long |  | Yes |
| slug | string |  | Yes |
| name | string |  | Yes |
| description | string |  | Yes |
| metaTitle | string |  | Yes |
| metaDescription | string |  | Yes |
| metaKeywords | string |  | Yes |
| sortOrder | integer |  | Yes |

#### ClearCartRequest

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| ClearCartRequest | object |  |  |

#### ClearCartResponse

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| statusCode | long |  | Yes |
| statusMessage | string |  | Yes |

#### CreateCartRequest

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| customerId | string |  | Yes |

#### CreateCartResponse

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| cart | [Cart](#cart) |  | Yes |
| sessionId | string |  | Yes |

#### CreateCategoryRequest

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| category | [Category](#category) |  | Yes |

#### CreateCategoryResponse

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| category | [Category](#category) |  | Yes |
| statusCode | long |  | Yes |
| statusMessage | string |  | Yes |

#### CreateProductRequest

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| product:omitempty | [Product](#product) |  | Yes |

#### CreateProductResponse

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| product:omitempty | [Product](#product) |  | Yes |
| statusCode:omitempty | long |  | Yes |
| statusMessage:omitempty | string |  | Yes |

#### Customer

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| id | long |  | Yes |
| firstName | string |  | Yes |
| lastName | string |  | Yes |
| email | string |  | Yes |
| password | string |  | Yes |

#### CustomerLoginRequest

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| email | string |  | Yes |
| password | string |  | Yes |

#### CustomerLoginResponse

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| jwt | [JwtToken](#jwttoken) |  | Yes |
| customer | [Customer](#customer) |  | Yes |
| statusCode | long |  | Yes |
| statusMessage | string |  | Yes |

#### DeleteCategoryRequest

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| id | long |  | Yes |

#### DeleteCategoryResponse

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| statusCode | long |  | Yes |
| statusMessage | string |  | Yes |

#### DeleteProductRequest

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| path:omitempty | long |  | Yes |

#### DeleteProductResponse

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| statusCode:omitempty | long |  | Yes |
| statusMessage:omitempty | string |  | Yes |

#### GetAllCategoriesRequest

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| GetAllCategoriesRequest | object |  |  |

#### GetAllCategoriesResponse

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| categories | [ [Category](#category) ] |  | Yes |
| statusCode | long |  | Yes |
| statusMessage | string |  | Yes |

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

#### GetCartRequest

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| customerId | string |  | Yes |

#### GetCartResponse

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| cart | [Cart](#cart) |  | Yes |
| sessionId | string |  | Yes |

#### GetCategoryByIdRequest

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| id | long |  | Yes |

#### GetCategoryByIdResponse

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| category | [Category](#category) |  | Yes |
| statusCode | long |  | Yes |
| statusMessage | string |  | Yes |

#### GetCategoryBySlugRequest

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| slug | string |  | Yes |

#### GetCategoryBySlugResponse

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| category | [Category](#category) |  | Yes |
| statusCode | long |  | Yes |
| statusMessage | string |  | Yes |

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
| slug | string |  | Yes |

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

#### Item

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| variantId | string |  | Yes |
| quantity | integer |  | Yes |
| price | double |  | Yes |
| expiresAt | string |  | Yes |

#### JwtToken

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| access_token | string |  | Yes |
| access_expire | long |  | Yes |
| refresh_after | long |  | Yes |

#### Price

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| id | long |  | Yes |
| amount | double |  | Yes |
| displayAmount | string |  | Yes |
| compareAtAmount | double |  | Yes |
| displayCompareAtAmount | string |  | Yes |
| currency | string |  | Yes |

#### Product

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| id | long |  | Yes |
| slug | string |  | Yes |
| name | string |  | Yes |
| shortDescription | string |  | Yes |
| description | string |  | Yes |
| metaTitle | string |  | Yes |
| metaDescription | string |  | Yes |
| metaKeywords | string |  | Yes |
| variants | [ [Variant](#variant) ] |  | Yes |

#### RemoveCartItemRequest

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| productId | long |  | Yes |
| variantId | long |  | Yes |
| quanity | integer |  | Yes |

#### RemoveCartItemResponse

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| cart | [Cart](#cart) |  | Yes |
| statusCode | long |  | Yes |
| statusMessage | string |  | Yes |

#### UpdateCartItemRequest

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| productId | long |  | Yes |
| variantId | long |  | Yes |
| quanity | integer |  | Yes |

#### UpdateCartItemResponse

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| cart | [Cart](#cart) |  | Yes |
| statusCode | long |  | Yes |
| statusMessage | string |  | Yes |

#### UpdateCategoryRequest

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| id | long |  | Yes |
| category | [Category](#category) |  | Yes |

#### UpdateCategoryResponse

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| category | [Category](#category) |  | Yes |
| statusCode | long |  | Yes |
| statusMessage | string |  | Yes |

#### UpdateProductRequest

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| path:omitempty | long |  | Yes |
| product:omitempty | [Product](#product) |  | Yes |

#### UpdateProductResponse

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| product:omitempty | [Product](#product) |  | Yes |
| statusCode:omitempty | long |  | Yes |
| statusMessage:omitempty | string |  | Yes |

#### Variant

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| id | long |  | Yes |
| isDefault | boolean (boolean) |  | Yes |
| sku | string |  | Yes |
| weight | double |  | Yes |
| height | double |  | Yes |
| width | double |  | Yes |
| depth | double |  | Yes |
| price | [Price](#price) |  | Yes |
