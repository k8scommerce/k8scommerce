


  

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

###  cart

| Method  | URI     | Name   | Summary |
|---------|---------|--------|---------|
| POST | /v1/cart/{customerId} | [add item to cart](#add-item-to-cart) | Add Item to Cart |
| DELETE | /v1/cart/{customerId} | [clear cart](#clear-cart) | Clear Cart |
| GET | /v1/cart/{customerId} | [get cart](#get-cart) | Get Cart |
| DELETE | /v1/cart/{customerId}/{sku} | [remove cart item](#remove-cart-item) | Remove Item |
| PUT | /v1/cart/{customerId}/{sku} | [update cart item quantity](#update-cart-item-quantity) | Update Item Quantity |
  


## Paths

### <span id="add-item-to-cart"></span> Add Item to Cart (*addItemToCart*)

```
POST /v1/cart/{customerId}
```

adds an item to an existing cart

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| customerId | `path` | string | `string` |  | ✓ |  | a customer's id |
| body | `body` | [AddItemToCartRequest](#add-item-to-cart-request) | `models.AddItemToCartRequest` | | ✓ | |  |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#add-item-to-cart-200) | OK | A successful response. |  | [schema](#add-item-to-cart-200-schema) |

#### Responses


##### <span id="add-item-to-cart-200"></span> 200 - A successful response.
Status: OK

###### <span id="add-item-to-cart-200-schema"></span> Schema
   
  

[AddItemToCartResponse](#add-item-to-cart-response)

### <span id="clear-cart"></span> Clear Cart (*clearCart*)

```
DELETE /v1/cart/{customerId}
```

clear a customer's cart

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| customerId | `path` | string | `string` |  | ✓ |  | a customer's id |
| body | `body` | [ClearCartRequest](#clear-cart-request) | `models.ClearCartRequest` | | ✓ | |  |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#clear-cart-200) | OK | A successful response. |  | [schema](#clear-cart-200-schema) |

#### Responses


##### <span id="clear-cart-200"></span> 200 - A successful response.
Status: OK

###### <span id="clear-cart-200-schema"></span> Schema
   
  

[ClearCartResponse](#clear-cart-response)

### <span id="get-cart"></span> Get Cart (*getCart*)

```
GET /v1/cart/{customerId}
```

returns a shopping cart if one exists

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| customerId | `path` | integer | `int64` |  | ✓ |  | a customer's id |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#get-cart-200) | OK | A successful response. |  | [schema](#get-cart-200-schema) |

#### Responses


##### <span id="get-cart-200"></span> 200 - A successful response.
Status: OK

###### <span id="get-cart-200-schema"></span> Schema
   
  

[GetCartResponse](#get-cart-response)

### <span id="remove-cart-item"></span> Remove Item (*removeCartItem*)

```
DELETE /v1/cart/{customerId}/{sku}
```

removes an item from a customer's cart

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| customerId | `path` | string | `string` |  | ✓ |  | a customer's id |
| sku | `path` | string | `string` |  | ✓ |  | an Item's sku |
| body | `body` | [RemoveCartItemRequest](#remove-cart-item-request) | `models.RemoveCartItemRequest` | | ✓ | |  |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#remove-cart-item-200) | OK | A successful response. |  | [schema](#remove-cart-item-200-schema) |

#### Responses


##### <span id="remove-cart-item-200"></span> 200 - A successful response.
Status: OK

###### <span id="remove-cart-item-200-schema"></span> Schema
   
  

[RemoveCartItemResponse](#remove-cart-item-response)

### <span id="update-cart-item-quantity"></span> Update Item Quantity (*updateCartItemQuantity*)

```
PUT /v1/cart/{customerId}/{sku}
```

updates a cart item's quantity

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| customerId | `path` | string | `string` |  | ✓ |  | a customer's id |
| sku | `path` | string | `string` |  | ✓ |  | an item's sku |
| body | `body` | [UpdateCartItemQuantityRequest](#update-cart-item-quantity-request) | `models.UpdateCartItemQuantityRequest` | | ✓ | |  |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#update-cart-item-quantity-200) | OK | A successful response. |  | [schema](#update-cart-item-quantity-200-schema) |

#### Responses


##### <span id="update-cart-item-quantity-200"></span> 200 - A successful response.
Status: OK

###### <span id="update-cart-item-quantity-200-schema"></span> Schema
   
  

[UpdateCartItemQuantityResponse](#update-cart-item-quantity-response)

## Models

### <span id="add-item-to-cart-request"></span> AddItemToCartRequest


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| customerId | int64 (formatted integer)| `int64` | ✓ | | a customer's id |  |
| item | [Item](#item)| `Item` | ✓ | | an Item object |  |



### <span id="add-item-to-cart-response"></span> AddItemToCartResponse


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| cart | [Cart](#cart)| `Cart` | ✓ | | a Cart object |  |
| status | [ResponseStatus](#response-status)| `ResponseStatus` | ✓ | | a ResponseStatus object |  |



### <span id="cart"></span> Cart


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| items | [][Item](#item)| `[]*Item` | ✓ | | a collection of Item |  |
| totalPrice | double (formatted number)| `float64` | ✓ | | the sum total of the cart |  |



### <span id="clear-cart-request"></span> ClearCartRequest


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| customerId | int64 (formatted integer)| `int64` | ✓ | | a customer's id |  |
| status | [ResponseStatus](#response-status)| `ResponseStatus` | ✓ | | a ResponseStatus object |  |



### <span id="clear-cart-response"></span> ClearCartResponse


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| status | [ResponseStatus](#response-status)| `ResponseStatus` | ✓ | | a ResponseStatus object |  |



### <span id="get-cart-request"></span> GetCartRequest


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| customerId | int64 (formatted integer)| `int64` | ✓ | | a customer's id |  |



### <span id="get-cart-response"></span> GetCartResponse


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| cart | [Cart](#cart)| `Cart` | ✓ | | a Cart object |  |
| status | [ResponseStatus](#response-status)| `ResponseStatus` | ✓ | | a ResponseStatus object |  |



### <span id="item"></span> Item


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| expiresAt | string| `string` | ✓ | | when this item expires in the cart |  |
| price | double (formatted number)| `float64` | ✓ | | the item's price |  |
| quantity | int32 (formatted integer)| `int32` | ✓ | | how many of identical items |  |
| sku | string| `string` | ✓ | | an item's variant sku number |  |



### <span id="jwt-token"></span> JwtToken


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| accessExpire | int64 (formatted integer)| `int64` | ✓ | |  |  |
| accessToken | string| `string` | ✓ | |  |  |
| refreshAfter | int64 (formatted integer)| `int64` | ✓ | |  |  |



### <span id="remove-cart-item-request"></span> RemoveCartItemRequest


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| customerId | int64 (formatted integer)| `int64` | ✓ | | a customer's id |  |
| quanity | int32 (formatted integer)| `int32` | ✓ | | a new quantity |  |
| sku | string| `string` | ✓ | | an item's variant sku number |  |



### <span id="remove-cart-item-response"></span> RemoveCartItemResponse


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| cart | [Cart](#cart)| `Cart` | ✓ | | a Cart object |  |
| status | [ResponseStatus](#response-status)| `ResponseStatus` | ✓ | | a ResponseStatus object |  |



### <span id="response-status"></span> ResponseStatus


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| statusCode | int64 (formatted integer)| `int64` | ✓ | | RFC http status code, ie. 204, etc - https:go.dev/src/net/http/status.go |  |
| statusMessage | string| `string` | ✓ | | status message |  |



### <span id="update-cart-item-quantity-request"></span> UpdateCartItemQuantityRequest


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| customerId | int64 (formatted integer)| `int64` | ✓ | | a customer's id |  |
| quanity | int32 (formatted integer)| `int32` | ✓ | | a new quantity |  |
| sku | string| `string` | ✓ | | an item's variant sku number |  |



### <span id="update-cart-item-quantity-response"></span> UpdateCartItemQuantityResponse


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| cart | [Cart](#cart)| `Cart` | ✓ | | a Cart object |  |
| status | [ResponseStatus](#response-status)| `ResponseStatus` | ✓ | | a ResponseStatus object |  |


