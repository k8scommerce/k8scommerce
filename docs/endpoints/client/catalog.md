


# Catalog API Endpoints
admin gateway catalog api endpoints
  

## Informations

### Version

1

## Content negotiation

### URI Schemes
  * http
  * https

### Consumes
  * application/json

### Produces
  * application/json

## Access control

### Security Schemes

#### apiKey (header: Authorization)

Enter JWT Bearer token **_only_**

> **Type**: apikey

### Security Requirements
  * apiKey

## All endpoints

###  client

| Method  | URI     | Name   | Summary |
|---------|---------|--------|---------|
| GET | /v1/categories | [get all categories](#get-all-categories) | Get All Categories |
| GET | /v1/products/{currentPage}/{pageSize} | [get all products](#get-all-products) | getAllProducts
	returns all products |
| GET | /v1/category/{id} | [get category by Id](#get-category-by-id) | getCategoryById
	returns a category by id |
| GET | /v1/category/slug/{slug} | [get category by slug](#get-category-by-slug) | getCategoryBySlug
	returns a category by url slug and storeId |
| GET | /v1/product/{id} | [get product by Id](#get-product-by-id) | getProductById
	returns a product by id |
| GET | /v1/product/sku/{sku} | [get product by sku](#get-product-by-sku) | getProductBySku
	returns a product by sku |
| GET | /v1/product/slug/{slug} | [get product by slug](#get-product-by-slug) | getProductBySlug
	returns a product by url slug |
| GET | /v1/products/category/{categoryId}/{currentPage}/{pageSize} | [get products by category Id](#get-products-by-category-id) | getProductsByCategoryId
	returns all products belonging to a category id |
  


## Paths

### <span id="get-all-categories"></span> Get All Categories (*getAllCategories*)

```
GET /v1/categories
```

returns all categories by belonging to a store

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#get-all-categories-200) | OK | A successful response. |  | [schema](#get-all-categories-200-schema) |

#### Responses


##### <span id="get-all-categories-200"></span> 200 - A successful response.
Status: OK

###### <span id="get-all-categories-200-schema"></span> Schema
   
  

[GetAllCategoriesResponse](#get-all-categories-response)

### <span id="get-all-products"></span> getAllProducts
	returns all products (*getAllProducts*)

```
GET /v1/products/{currentPage}/{pageSize}
```

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| currentPage | `path` | string | `string` |  | ✓ |  |  |
| pageSize | `path` | string | `string` |  | ✓ |  |  |
| sortOn | `query` | string | `string` |  |  |  |  |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#get-all-products-200) | OK | A successful response. |  | [schema](#get-all-products-200-schema) |

#### Responses


##### <span id="get-all-products-200"></span> 200 - A successful response.
Status: OK

###### <span id="get-all-products-200-schema"></span> Schema
   
  

[GetAllProductsResponse](#get-all-products-response)

### <span id="get-category-by-id"></span> getCategoryById
	returns a category by id (*getCategoryById*)

```
GET /v1/category/{id}
```

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| id | `path` | string | `string` |  | ✓ |  |  |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#get-category-by-id-200) | OK | A successful response. |  | [schema](#get-category-by-id-200-schema) |

#### Responses


##### <span id="get-category-by-id-200"></span> 200 - A successful response.
Status: OK

###### <span id="get-category-by-id-200-schema"></span> Schema
   
  

[GetCategoryByIDResponse](#get-category-by-id-response)

### <span id="get-category-by-slug"></span> getCategoryBySlug
	returns a category by url slug and storeId (*getCategoryBySlug*)

```
GET /v1/category/slug/{slug}
```

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| slug | `path` | string | `string` |  | ✓ |  |  |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#get-category-by-slug-200) | OK | A successful response. |  | [schema](#get-category-by-slug-200-schema) |

#### Responses


##### <span id="get-category-by-slug-200"></span> 200 - A successful response.
Status: OK

###### <span id="get-category-by-slug-200-schema"></span> Schema
   
  

[GetCategoryBySlugResponse](#get-category-by-slug-response)

### <span id="get-product-by-id"></span> getProductById
	returns a product by id (*getProductById*)

```
GET /v1/product/{id}
```

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| id | `path` | string | `string` |  | ✓ |  |  |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#get-product-by-id-200) | OK | A successful response. |  | [schema](#get-product-by-id-200-schema) |

#### Responses


##### <span id="get-product-by-id-200"></span> 200 - A successful response.
Status: OK

###### <span id="get-product-by-id-200-schema"></span> Schema
   
  

[Product](#product)

### <span id="get-product-by-sku"></span> getProductBySku
	returns a product by sku (*getProductBySku*)

```
GET /v1/product/sku/{sku}
```

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| sku | `path` | string | `string` |  | ✓ |  |  |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#get-product-by-sku-200) | OK | A successful response. |  | [schema](#get-product-by-sku-200-schema) |

#### Responses


##### <span id="get-product-by-sku-200"></span> 200 - A successful response.
Status: OK

###### <span id="get-product-by-sku-200-schema"></span> Schema
   
  

[Product](#product)

### <span id="get-product-by-slug"></span> getProductBySlug
	returns a product by url slug (*getProductBySlug*)

```
GET /v1/product/slug/{slug}
```

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| slug | `path` | string | `string` |  | ✓ |  |  |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#get-product-by-slug-200) | OK | A successful response. |  | [schema](#get-product-by-slug-200-schema) |

#### Responses


##### <span id="get-product-by-slug-200"></span> 200 - A successful response.
Status: OK

###### <span id="get-product-by-slug-200-schema"></span> Schema
   
  

[Product](#product)

### <span id="get-products-by-category-id"></span> getProductsByCategoryId
	returns all products belonging to a category id (*getProductsByCategoryId*)

```
GET /v1/products/category/{categoryId}/{currentPage}/{pageSize}
```

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| categoryId | `path` | string | `string` |  | ✓ |  |  |
| currentPage | `path` | string | `string` |  | ✓ |  |  |
| pageSize | `path` | string | `string` |  | ✓ |  |  |
| sortOn | `query` | string | `string` |  |  |  |  |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#get-products-by-category-id-200) | OK | A successful response. |  | [schema](#get-products-by-category-id-200-schema) |

#### Responses


##### <span id="get-products-by-category-id-200"></span> 200 - A successful response.
Status: OK

###### <span id="get-products-by-category-id-200-schema"></span> Schema
   
  

[GetProductsByCategoryIDResponse](#get-products-by-category-id-response)

## Models

### <span id="category"></span> Category


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| description | string| `string` | ✓ | | description of category |  |
| id | int64 (formatted integer)| `int64` | ✓ | | category id |  |
| metaDescription | string| `string` | ✓ | | metatag description for SEO |  |
| metaKeywords | string| `string` | ✓ | | metatag keywords for SEO |  |
| metaTitle | string| `string` | ✓ | | metatag title for SEO |  |
| name | string| `string` | ✓ | | name of category |  |
| parentId | int64 (formatted integer)| `int64` | ✓ | | parent category id. references Category.Id |  |
| slug | string| `string` | ✓ | | slug name of the category |  |
| sortOrder | int32 (formatted integer)| `int32` | ✓ | | sort order of menu items on the same level and same parent id |  |



### <span id="create-category-request"></span> CreateCategoryRequest


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| category | [Category](#category)| `Category` | ✓ | |  |  |



### <span id="create-category-response"></span> CreateCategoryResponse


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| category | [Category](#category)| `Category` | ✓ | |  |  |
| statusCode | int64 (formatted integer)| `int64` | ✓ | |  |  |
| statusMessage | string| `string` | ✓ | |  |  |



### <span id="create-product-request"></span> CreateProductRequest


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| product:omitempty | [Product](#product)| `Product` | ✓ | |  |  |



### <span id="create-product-response"></span> CreateProductResponse


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| product:omitempty | [Product](#product)| `Product` | ✓ | |  |  |
| statusCode:omitempty | int64 (formatted integer)| `int64` | ✓ | |  |  |
| statusMessage:omitempty | string| `string` | ✓ | |  |  |



### <span id="delete-category-request"></span> DeleteCategoryRequest


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| id | int64 (formatted integer)| `int64` | ✓ | |  |  |



### <span id="delete-category-response"></span> DeleteCategoryResponse


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| statusCode | int64 (formatted integer)| `int64` | ✓ | |  |  |
| statusMessage | string| `string` | ✓ | |  |  |



### <span id="delete-product-request"></span> DeleteProductRequest


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| path:omitempty | int64 (formatted integer)| `int64` | ✓ | |  |  |



### <span id="delete-product-response"></span> DeleteProductResponse


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| statusCode:omitempty | int64 (formatted integer)| `int64` | ✓ | |  |  |
| statusMessage:omitempty | string| `string` | ✓ | |  |  |



### <span id="get-all-categories-request"></span> GetAllCategoriesRequest


  

[interface{}](#interface)

### <span id="get-all-categories-response"></span> GetAllCategoriesResponse


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| categories | [][Category](#category)| `[]*Category` | ✓ | |  |  |
| statusCode | int64 (formatted integer)| `int64` | ✓ | |  |  |
| statusMessage | string| `string` | ✓ | |  |  |



### <span id="get-all-products-request"></span> GetAllProductsRequest


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| currentPage | int64 (formatted integer)| `int64` | ✓ | |  |  |
| pageSize | int64 (formatted integer)| `int64` | ✓ | |  |  |
| sortOn | string| `string` |  | |  |  |



### <span id="get-all-products-response"></span> GetAllProductsResponse


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| products | [][Product](#product)| `[]*Product` | ✓ | |  |  |
| totalPages | int64 (formatted integer)| `int64` | ✓ | |  |  |
| totalRecords | int64 (formatted integer)| `int64` | ✓ | |  |  |



### <span id="get-category-by-id-request"></span> GetCategoryByIdRequest


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| id | int64 (formatted integer)| `int64` | ✓ | |  |  |



### <span id="get-category-by-id-response"></span> GetCategoryByIdResponse


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| category | [Category](#category)| `Category` | ✓ | |  |  |
| statusCode | int64 (formatted integer)| `int64` | ✓ | |  |  |
| statusMessage | string| `string` | ✓ | |  |  |



### <span id="get-category-by-slug-request"></span> GetCategoryBySlugRequest


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| slug | string| `string` | ✓ | |  |  |



### <span id="get-category-by-slug-response"></span> GetCategoryBySlugResponse


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| category | [Category](#category)| `Category` | ✓ | |  |  |
| statusCode | int64 (formatted integer)| `int64` | ✓ | |  |  |
| statusMessage | string| `string` | ✓ | |  |  |



### <span id="get-product-by-id-request"></span> GetProductByIdRequest


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| id | int64 (formatted integer)| `int64` | ✓ | |  |  |



### <span id="get-product-by-sku-request"></span> GetProductBySkuRequest


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| sku | string| `string` | ✓ | |  |  |



### <span id="get-product-by-slug-request"></span> GetProductBySlugRequest


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| slug | string| `string` | ✓ | |  |  |



### <span id="get-products-by-category-id-request"></span> GetProductsByCategoryIdRequest


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| categoryId | int64 (formatted integer)| `int64` | ✓ | |  |  |
| currentPage | int64 (formatted integer)| `int64` | ✓ | |  |  |
| pageSize | int64 (formatted integer)| `int64` | ✓ | |  |  |
| sortOn | string| `string` |  | |  |  |



### <span id="get-products-by-category-id-response"></span> GetProductsByCategoryIdResponse


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| products | [][Product](#product)| `[]*Product` | ✓ | |  |  |
| totalPages | int64 (formatted integer)| `int64` | ✓ | |  |  |
| totalRecords | int64 (formatted integer)| `int64` | ✓ | |  |  |



### <span id="price"></span> Price


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| amount | double (formatted number)| `float64` | ✓ | | price amount |  |
| compareAtAmount | double (formatted number)| `float64` | ✓ | | price compare amount |  |
| currency | string| `string` | ✓ | | price currency. example: USD, CAN, etc. |  |
| displayAmount | string| `string` | ✓ | | price display amount |  |
| displayCompareAtAmount | string| `string` | ✓ | | price display compare amount |  |
| id | int64 (formatted integer)| `int64` | ✓ | | price id |  |



### <span id="product"></span> Product


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| description | string| `string` | ✓ | | category description |  |
| id | int64 (formatted integer)| `int64` | ✓ | | product id |  |
| metaDescription | string| `string` | ✓ | | metatag description for SEO |  |
| metaKeywords | string| `string` | ✓ | | metatag keywords for SEO |  |
| metaTitle | string| `string` | ✓ | | metatag title for SEO |  |
| name | string| `string` | ✓ | | product name |  |
| shortDescription | string| `string` | ✓ | | product short description. used in category pages |  |
| slug | string| `string` | ✓ | | product slug |  |
| variants | [][Variant](#variant)| `[]*Variant` | ✓ | | collection of Variant objects |  |



### <span id="update-category-request"></span> UpdateCategoryRequest


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| category | [Category](#category)| `Category` | ✓ | |  |  |
| id | int64 (formatted integer)| `int64` | ✓ | |  |  |



### <span id="update-category-response"></span> UpdateCategoryResponse


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| category | [Category](#category)| `Category` | ✓ | |  |  |
| statusCode | int64 (formatted integer)| `int64` | ✓ | |  |  |
| statusMessage | string| `string` | ✓ | |  |  |



### <span id="update-product-request"></span> UpdateProductRequest


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| path:omitempty | int64 (formatted integer)| `int64` | ✓ | |  |  |
| product:omitempty | [Product](#product)| `Product` | ✓ | |  |  |



### <span id="update-product-response"></span> UpdateProductResponse


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| product:omitempty | [Product](#product)| `Product` | ✓ | |  |  |
| statusCode:omitempty | int64 (formatted integer)| `int64` | ✓ | |  |  |
| statusMessage:omitempty | string| `string` | ✓ | |  |  |



### <span id="variant"></span> Variant


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| depth | double (formatted number)| `float64` | ✓ | | variant depth. used in calculating shipping |  |
| height | double (formatted number)| `float64` | ✓ | | variant height. used in calculating shipping |  |
| id | int64 (formatted integer)| `int64` | ✓ | | variant id |  |
| isDefault | boolean (formatted boolean)| `bool` | ✓ | | is default variant. each product must have exactly 1 default variant |  |
| price | [Price](#price)| `Price` | ✓ | | variant Price object |  |
| sku | string| `string` | ✓ | | variant sku. usually the product sku with appended identification tags |  |
| weight | double (formatted number)| `float64` | ✓ | | variant weight. used in calculating shipping |  |
| width | double (formatted number)| `float64` | ✓ | | variant width. used in calculating shipping |  |


