


  

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
| POST | /v1/customer | [create customer](#create-customer) | Create Customer |
| POST | /v1/customer/login | [login](#login) | Login |
  


## Paths

### <span id="create-customer"></span> Create Customer (*createCustomer*)

```
POST /v1/customer
```

creates a new customer

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| body | `body` | [CreateCustomerRequest](#create-customer-request) | `models.CreateCustomerRequest` | | ✓ | |  |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#create-customer-200) | OK | A successful response. |  | [schema](#create-customer-200-schema) |

#### Responses


##### <span id="create-customer-200"></span> 200 - A successful response.
Status: OK

###### <span id="create-customer-200-schema"></span> Schema
   
  

[CreateCustomerResponse](#create-customer-response)

### <span id="login"></span> Login (*login*)

```
POST /v1/customer/login
```

login for customers

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| body | `body` | [CustomerLoginRequest](#customer-login-request) | `models.CustomerLoginRequest` | | ✓ | |  |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#login-200) | OK | A successful response. |  | [schema](#login-200-schema) |

#### Responses


##### <span id="login-200"></span> 200 - A successful response.
Status: OK

###### <span id="login-200-schema"></span> Schema
   
  

[CustomerLoginResponse](#customer-login-response)

## Models

### <span id="address"></span> Address


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| city | string| `string` | ✓ | |  |  |
| country | string| `string` | ✓ | |  |  |
| isDefault | boolean (formatted boolean)| `bool` | ✓ | |  |  |
| postalCode | string| `string` | ✓ | |  |  |
| state | string| `string` | ✓ | |  |  |
| street | string| `string` | ✓ | |  |  |



### <span id="create-customer-request"></span> CreateCustomerRequest


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| customer | [NewCustomer](#new-customer)| `NewCustomer` | ✓ | |  |  |



### <span id="create-customer-response"></span> CreateCustomerResponse


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| customer | [Customer](#customer)| `Customer` | ✓ | |  |  |
| status | [ResponseStatus](#response-status)| `ResponseStatus` | ✓ | | a ResponseStatus object |  |



### <span id="customer"></span> Customer


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| email | string| `string` | ✓ | |  |  |
| firstName | string| `string` | ✓ | |  |  |
| id | int64 (formatted integer)| `int64` | ✓ | |  |  |
| lastName | string| `string` | ✓ | |  |  |
| password | string| `string` | ✓ | |  |  |



### <span id="customer-account"></span> CustomerAccount


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| billingAddress | [Address](#address)| `Address` | ✓ | |  |  |
| id | int64 (formatted integer)| `int64` | ✓ | |  |  |
| shippingAddresses | [][Address](#address)| `[]*Address` | ✓ | |  |  |



### <span id="customer-login-request"></span> CustomerLoginRequest


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| email | string| `string` | ✓ | |  |  |
| password | string| `string` | ✓ | |  |  |



### <span id="customer-login-response"></span> CustomerLoginResponse


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| customer | [Customer](#customer)| `Customer` | ✓ | |  |  |
| jwt | [JwtToken](#jwt-token)| `JwtToken` | ✓ | |  |  |
| status | [ResponseStatus](#response-status)| `ResponseStatus` | ✓ | | a ResponseStatus object |  |



### <span id="jwt-token"></span> JwtToken


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| access_expire | int64 (formatted integer)| `int64` | ✓ | |  |  |
| access_token | string| `string` | ✓ | |  |  |
| refresh_after | int64 (formatted integer)| `int64` | ✓ | |  |  |



### <span id="new-customer"></span> NewCustomer


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| billingAddress | [Address](#address)| `Address` |  | |  |  |
| email | string| `string` | ✓ | |  |  |
| firstName | string| `string` | ✓ | |  |  |
| lastName | string| `string` | ✓ | |  |  |
| password | string| `string` | ✓ | |  |  |
| shippingAddresses | [Address](#address)| `Address` |  | |  |  |



### <span id="response-status"></span> ResponseStatus


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| statusCode | int64 (formatted integer)| `int64` | ✓ | | RFC http status code, ie. 204, etc - https:go.dev/src/net/http/status.go |  |
| statusMessage | string| `string` | ✓ | | status message |  |


