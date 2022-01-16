# Client Gateway API
client gateway api

## Version: 1

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

Get Cart

##### Description

returns a shopping cart if one exists

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| customerId | path | a customer's id | Yes | string |

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | A successful response. | [GetCartResponse](#getcartresponse) |

#### DELETE
##### Summary

Clear Cart

##### Description

clear a customer's cart

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| customerId | path | a customer's id | Yes | string |
| body | body |  | Yes | [ClearCartRequest](#clearcartrequest) |

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | A successful response. | [ClearCartResponse](#clearcartresponse) |

#### POST
##### Summary

Add Item to Cart

##### Description

adds an item to an existing cart

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| customerId | path | a customer's id | Yes | string |
| body | body |  | Yes | [AddItemToCartRequest](#additemtocartrequest) |

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | A successful response. | [AddItemToCartResponse](#additemtocartresponse) |

### /v1/cart/{customerId}/{sku}

#### DELETE
##### Summary

Remove Item

##### Description

removes an item from a customer's cart

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| customerId | path | a customer's id | Yes | string |
| sku | path | an Item's sku | Yes | string |
| body | body |  | Yes | [RemoveCartItemRequest](#removecartitemrequest) |

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | A successful response. | [RemoveCartItemResponse](#removecartitemresponse) |

#### PUT
##### Summary

Update Item Quantity

##### Description

updates a cart item's quantity

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| customerId | path | a customer's id | Yes | string |
| sku | path | an item's sku | Yes | string |
| body | body |  | Yes | [UpdateCartItemQuantityRequest](#updatecartitemquantityrequest) |

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | A successful response. | [UpdateCartItemQuantityResponse](#updatecartitemquantityresponse) |

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
| slug | query |  slug name of the category | Yes | string |

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
| 200 | A successful response. | [Product](#product) |

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
| status | query |  a ResponseStatus object | Yes | invalid (UNKNOWN) |

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | A successful response. | [Product](#product) |

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
| 200 | A successful response. | [Product](#product) |

### /v1/products/category/{categoryId}/{currentPage}/{pageSize}

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

### Models

#### AddItemToCartRequest

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| customerId | long |  a customer's id | Yes |
| item | [Item](#item) |  an Item object | Yes |

#### AddItemToCartResponse

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| cart | [Cart](#cart) |  a Cart object | Yes |
| status | [ResponseStatus](#responsestatus) |  a ResponseStatus object | Yes |

#### Cart

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| items | [ [Item](#item) ] |  a collection of Item | Yes |
| totalPrice | double |  the sum total of the cart | Yes |

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

#### ClearCartRequest

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| customerId | long |  a customer's id | Yes |
| status | [ResponseStatus](#responsestatus) |  a ResponseStatus object | Yes |

#### ClearCartResponse

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| status | [ResponseStatus](#responsestatus) |  a ResponseStatus object | Yes |

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

#### GetCartRequest

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| customerId | long |  a customer's id | Yes |

#### GetCartResponse

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| cart | [Cart](#cart) |  a Cart object | Yes |
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

#### Item

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| sku | string |  an item's variant sku number | Yes |
| quantity | integer |  how many of identical items | Yes |
| price | double |  the item's price | Yes |
| expiresAt | string |  when this item expires in the cart | Yes |

#### JwtToken

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| access_token | string |  | Yes |
| access_expire | long |  | Yes |
| refresh_after | long |  | Yes |

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

#### RemoveCartItemRequest

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| customerId | long |  a customer's id | Yes |
| sku | string |  an item's variant sku number | Yes |
| quanity | integer |  a new quantity | Yes |

#### RemoveCartItemResponse

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| cart | [Cart](#cart) |  a Cart object | Yes |
| status | [ResponseStatus](#responsestatus) |  a ResponseStatus object | Yes |

#### ResponseStatus

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| statusCode | long |  RFC http status code, ie. 204, etc - https:go.dev/src/net/http/status.go | Yes |
| statusMessage | string |  status message | Yes |

#### UpdateCartItemQuantityRequest

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| customerId | long |  a customer's id | Yes |
| sku | string |  an item's variant sku number | Yes |
| quanity | integer |  a new quantity | Yes |

#### UpdateCartItemQuantityResponse

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| cart | [Cart](#cart) |  a Cart object | Yes |
| status | [ResponseStatus](#responsestatus) |  a ResponseStatus object | Yes |

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
