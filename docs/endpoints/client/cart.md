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

#### GetCartRequest

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| customerId | string |  | Yes |

#### GetCartResponse

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| cart | [Cart](#cart) |  | Yes |
| sessionId | string |  | Yes |

#### Item

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| variantId | string |  | Yes |
| quantity | integer |  | Yes |
| price | double |  | Yes |
| expiresAt | string |  | Yes |

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
