


  

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

###  categories

| Method  | URI     | Name   | Summary |
|---------|---------|--------|---------|
| POST | /v1/category | [create category](#create-category) | Create Category |
| DELETE | /v1/category/{id} | [delete category](#delete-category) | Delete Category |
| GET | /v1/categories/{storeId}/{currentPage}/{pageSize} | [get all categories](#get-all-categories) | Get All Categories |
| GET | /v1/category/{id} | [get category by Id](#get-category-by-id) | Get Category By Id |
| GET | /v1/category/slug/{slug} | [get category by slug](#get-category-by-slug) | Get Category By Slug |
| PUT | /v1/category/{id} | [update category](#update-category) | Update Category |
  


###  products

| Method  | URI     | Name   | Summary |
|---------|---------|--------|---------|
| POST | /v1/product | [create product](#create-product) | Create Product |
| DELETE | /v1/product/{id} | [delete product](#delete-product) | Delete Product |
| GET | /v1/products/{currentPage}/{pageSize} | [get all products](#get-all-products) | Get All Products |
| GET | /v1/product/{id} | [get product by Id](#get-product-by-id) | Get Product By Id |
| GET | /v1/product/sku/{sku} | [get product by sku](#get-product-by-sku) | Get Product By Sku |
| GET | /v1/product/slug/{slug} | [get product by slug](#get-product-by-slug) | Get Product By Slug |
| GET | /v1/products/{categoryId}/{currentPage}/{pageSize} | [get products by category Id](#get-products-by-category-id) | Get Products By Category Id |
| PUT | /v1/product/{id} | [update product](#update-product) | Update Product |
  


## Paths

### <span id="create-category"></span> Create Category (*createCategory*)

```
POST /v1/category
```

creates a category

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| body | `body` | [CreateCategoryRequest](#create-category-request) | `models.CreateCategoryRequest` | | ✓ | |  |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#create-category-200) | OK | A successful response. |  | [schema](#create-category-200-schema) |

#### Responses


##### <span id="create-category-200"></span> 200 - A successful response.
Status: OK

###### <span id="create-category-200-schema"></span> Schema
   
  

[Category](#category)

### <span id="create-product"></span> Create Product (*createProduct*)

```
POST /v1/product
```

creates a product

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| body | `body` | [CreateProductRequest](#create-product-request) | `models.CreateProductRequest` | | ✓ | |  |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#create-product-200) | OK | A successful response. |  | [schema](#create-product-200-schema) |

#### Responses


##### <span id="create-product-200"></span> 200 - A successful response.
Status: OK

###### <span id="create-product-200-schema"></span> Schema
   
  

[Product](#product)

### <span id="delete-category"></span> Delete Category (*deleteCategory*)

```
DELETE /v1/category/{id}
```

deletes a category

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| id | `path` | string | `string` |  | ✓ |  | category id |
| body | `body` | [DeleteCategoryRequest](#delete-category-request) | `models.DeleteCategoryRequest` | | ✓ | |  |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#delete-category-200) | OK | A successful response. |  | [schema](#delete-category-200-schema) |

#### Responses


##### <span id="delete-category-200"></span> 200 - A successful response.
Status: OK

###### <span id="delete-category-200-schema"></span> Schema
   
  

[DeleteCategoryResponse](#delete-category-response)

### <span id="delete-product"></span> Delete Product (*deleteProduct*)

```
DELETE /v1/product/{id}
```

delete a product

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| id | `path` | string | `string` |  | ✓ |  | product id |
| body | `body` | [DeleteProductRequest](#delete-product-request) | `models.DeleteProductRequest` | | ✓ | |  |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#delete-product-200) | OK | A successful response. |  | [schema](#delete-product-200-schema) |

#### Responses


##### <span id="delete-product-200"></span> 200 - A successful response.
Status: OK

###### <span id="delete-product-200-schema"></span> Schema
   
  

[DeleteProductResponse](#delete-product-response)

### <span id="get-all-categories"></span> Get All Categories (*getAllCategories*)

```
GET /v1/categories/{storeId}/{currentPage}/{pageSize}
```

returns all categories belonging to a store

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| currentPage | `path` | integer | `int64` |  | ✓ |  | current page number |
| pageSize | `path` | integer | `int64` |  | ✓ |  | number of records per page |
| storeId | `path` | string | `string` |  | ✓ |  | store id |
| sortOn | `query` | string | `string` |  |  |  |  |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#get-all-categories-200) | OK | A successful response. |  | [schema](#get-all-categories-200-schema) |

#### Responses


##### <span id="get-all-categories-200"></span> 200 - A successful response.
Status: OK

###### <span id="get-all-categories-200-schema"></span> Schema
   
  

[GetAllCategoriesResponse](#get-all-categories-response)

### <span id="get-all-products"></span> Get All Products (*getAllProducts*)

```
GET /v1/products/{currentPage}/{pageSize}
```

returns all products belonging to a store

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| currentPage | `path` | integer | `int64` |  | ✓ |  | current page number |
| pageSize | `path` | integer | `int64` |  | ✓ |  | number of records per page |
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

### <span id="get-category-by-id"></span> Get Category By Id (*getCategoryById*)

```
GET /v1/category/{id}
```

returns all categories by id belonging to a store

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| id | `path` | integer | `int64` |  | ✓ |  | category id |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#get-category-by-id-200) | OK | A successful response. |  | [schema](#get-category-by-id-200-schema) |

#### Responses


##### <span id="get-category-by-id-200"></span> 200 - A successful response.
Status: OK

###### <span id="get-category-by-id-200-schema"></span> Schema
   
  

[Category](#category)

### <span id="get-category-by-slug"></span> Get Category By Slug (*getCategoryBySlug*)

```
GET /v1/category/slug/{slug}
```

returns all categories by slug belonging to a store

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| slug | `path` | string | `string` |  | ✓ |  | category slug |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#get-category-by-slug-200) | OK | A successful response. |  | [schema](#get-category-by-slug-200-schema) |

#### Responses


##### <span id="get-category-by-slug-200"></span> 200 - A successful response.
Status: OK

###### <span id="get-category-by-slug-200-schema"></span> Schema
   
  

[Category](#category)

### <span id="get-product-by-id"></span> Get Product By Id (*getProductById*)

```
GET /v1/product/{id}
```

returns matching product by id

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| id | `path` | integer | `int64` |  | ✓ |  | product id |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#get-product-by-id-200) | OK | A successful response. |  | [schema](#get-product-by-id-200-schema) |

#### Responses


##### <span id="get-product-by-id-200"></span> 200 - A successful response.
Status: OK

###### <span id="get-product-by-id-200-schema"></span> Schema
   
  

[Product](#product)

### <span id="get-product-by-sku"></span> Get Product By Sku (*getProductBySku*)

```
GET /v1/product/sku/{sku}
```

returns all products by sku belonging to a store

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| sku | `path` | string | `string` |  | ✓ |  | product sku |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#get-product-by-sku-200) | OK | A successful response. |  | [schema](#get-product-by-sku-200-schema) |

#### Responses


##### <span id="get-product-by-sku-200"></span> 200 - A successful response.
Status: OK

###### <span id="get-product-by-sku-200-schema"></span> Schema
   
  

[Product](#product)

### <span id="get-product-by-slug"></span> Get Product By Slug (*getProductBySlug*)

```
GET /v1/product/slug/{slug}
```

returns matching product by slug

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| slug | `path` | string | `string` |  | ✓ |  | product slug |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#get-product-by-slug-200) | OK | A successful response. |  | [schema](#get-product-by-slug-200-schema) |

#### Responses


##### <span id="get-product-by-slug-200"></span> 200 - A successful response.
Status: OK

###### <span id="get-product-by-slug-200-schema"></span> Schema
   
  

[Product](#product)

### <span id="get-products-by-category-id"></span> Get Products By Category Id (*getProductsByCategoryId*)

```
GET /v1/products/{categoryId}/{currentPage}/{pageSize}
```

returns all products by category id belonging to a store

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| categoryId | `path` | integer | `int64` |  | ✓ |  | category id |
| currentPage | `path` | integer | `int64` |  | ✓ |  | current page number |
| pageSize | `path` | integer | `int64` |  | ✓ |  | number of records per page |
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

### <span id="update-category"></span> Update Category (*updateCategory*)

```
PUT /v1/category/{id}
```

updates a category

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| id | `path` | string | `string` |  | ✓ |  | category id |
| body | `body` | [UpdateCategoryRequest](#update-category-request) | `models.UpdateCategoryRequest` | | ✓ | |  |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#update-category-200) | OK | A successful response. |  | [schema](#update-category-200-schema) |

#### Responses


##### <span id="update-category-200"></span> 200 - A successful response.
Status: OK

###### <span id="update-category-200-schema"></span> Schema
   
  

[Category](#category)

### <span id="update-product"></span> Update Product (*updateProduct*)

```
PUT /v1/product/{id}
```

updates a product

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| id | `path` | string | `string` |  | ✓ |  | product id |
| body | `body` | [UpdateProductRequest](#update-product-request) | `models.UpdateProductRequest` | | ✓ | |  |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#update-product-200) | OK | A successful response. |  | [schema](#update-product-200-schema) |

#### Responses


##### <span id="update-product-200"></span> 200 - A successful response.
Status: OK

###### <span id="update-product-200-schema"></span> Schema
   
  

[Product](#product)

## Models

### <span id="category"></span> Category


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| depth | int32 (formatted integer)| `int32` |  | | category level depth |  |
| description | string| `string` |  | | description of category |  |
| id | int64 (formatted integer)| `int64` | ✓ | | category id |  |
| metaDescription | string| `string` |  | | metatag description for SEO |  |
| metaKeywords | string| `string` |  | | metatag keywords for SEO |  |
| metaTitle | string| `string` |  | | metatag title for SEO |  |
| name | string| `string` | ✓ | | name of category |  |
| parentId | int64 (formatted integer)| `int64` |  | | parent category id. references Category.Id |  |
| slug | string| `string` | ✓ | | slug name of the category |  |
| sortOrder | int32 (formatted integer)| `int32` |  | | sort order of menu items on the same level and same parent id |  |



### <span id="create-category-request"></span> CreateCategoryRequest


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| category | [Category](#category)| `Category` | ✓ | |  |  |



### <span id="create-product-request"></span> CreateProductRequest


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| product | [Product](#product)| `Product` | ✓ | |  |  |



### <span id="delete-category-request"></span> DeleteCategoryRequest


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| id | int64 (formatted integer)| `int64` | ✓ | |  |  |



### <span id="delete-category-response"></span> DeleteCategoryResponse


  

[interface{}](#interface)

### <span id="delete-product-request"></span> DeleteProductRequest


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| id | int64 (formatted integer)| `int64` | ✓ | |  |  |



### <span id="delete-product-response"></span> DeleteProductResponse


  

[interface{}](#interface)

### <span id="get-all-categories-request"></span> GetAllCategoriesRequest


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| currentPage | int64 (formatted integer)| `int64` | ✓ | |  |  |
| pageSize | int64 (formatted integer)| `int64` | ✓ | |  |  |
| sortOn | string| `string` |  | |  |  |



### <span id="get-all-categories-response"></span> GetAllCategoriesResponse


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| categories | [][Category](#category)| `[]*Category` | ✓ | | a collection of Category |  |
| totalPages | int64 (formatted integer)| `int64` | ✓ | |  |  |
| totalRecords | int64 (formatted integer)| `int64` | ✓ | |  |  |



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



### <span id="get-category-by-slug-request"></span> GetCategoryBySlugRequest


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| slug | string| `string` | ✓ | | slug name of the category |  |



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
| slug | string| `string` | ✓ | | slug name of the category |  |



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



### <span id="jwt-token"></span> JwtToken


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| accessExpire | int64 (formatted integer)| `int64` | ✓ | |  |  |
| accessToken | string| `string` | ✓ | |  |  |
| refreshAfter | int64 (formatted integer)| `int64` | ✓ | |  |  |



### <span id="price"></span> Price


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| amount | double (formatted number)| `float64` | ✓ | | price amount |  |
| compareAtAmount | double (formatted number)| `float64` |  | | price compare amount |  |
| currency | string| `string` |  | | price currency. example: USD, CAN, etc. |  |
| displayAmount | string| `string` | ✓ | | price display amount |  |
| displayCompareAtAmount | string| `string` |  | | price display compare amount |  |
| id | int64 (formatted integer)| `int64` |  | | price id |  |



### <span id="product"></span> Product


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| description | string| `string` |  | | category description |  |
| id | int64 (formatted integer)| `int64` | ✓ | | product id |  |
| metaDescription | string| `string` |  | | metatag description for SEO |  |
| metaKeywords | string| `string` |  | | metatag keywords for SEO |  |
| metaTitle | string| `string` |  | | metatag title for SEO |  |
| name | string| `string` | ✓ | | product name |  |
| shortDescription | string| `string` |  | | product short description. used in category pages |  |
| slug | string| `string` | ✓ | | product slug |  |
| variants | [][Variant](#variant)| `[]*Variant` |  | | collection of Variant objects |  |



### <span id="response-status"></span> ResponseStatus


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| statusCode | int64 (formatted integer)| `int64` | ✓ | | RFC http status code, ie. 204, etc - https:go.dev/src/net/http/status.go |  |
| statusMessage | string| `string` | ✓ | | status message |  |



### <span id="update-category-request"></span> UpdateCategoryRequest


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| category | [Category](#category)| `Category` | ✓ | |  |  |
| id | int64 (formatted integer)| `int64` | ✓ | |  |  |



### <span id="update-product-request"></span> UpdateProductRequest


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| path | int64 (formatted integer)| `int64` | ✓ | |  |  |
| product | [Product](#product)| `Product` | ✓ | |  |  |



### <span id="variant"></span> Variant


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| depth | double (formatted number)| `float64` |  | | variant depth. used in calculating shipping |  |
| height | double (formatted number)| `float64` |  | | variant height. used in calculating shipping |  |
| id | int64 (formatted integer)| `int64` | ✓ | | variant id |  |
| isDefault | boolean (formatted boolean)| `bool` | ✓ | | is default variant. each product must have exactly 1 default variant |  |
| price | [Price](#price)| `Price` |  | | variant Price object |  |
| sku | string| `string` | ✓ | | variant sku. usually the product sku with appended identification tags |  |
| weight | double (formatted number)| `float64` |  | | variant weight. used in calculating shipping |  |
| width | double (formatted number)| `float64` |  | | variant width. used in calculating shipping |  |


