


  

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

###  customers

| Method  | URI     | Name   | Summary |
|---------|---------|--------|---------|
| POST | /v1/customer | [create customer](#create-customer) | Create Customer |
| POST | /v1/customer/login | [customer login](#customer-login) | Login |
  


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

### <span id="customer-login"></span> Login (*customerLogin*)

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
| [200](#customer-login-200) | OK | A successful response. |  | [schema](#customer-login-200-schema) |

#### Responses


##### <span id="customer-login-200"></span> 200 - A successful response.
Status: OK

###### <span id="customer-login-200-schema"></span> Schema
   
  

[CustomerLoginResponse](#customer-login-response)

## Models

### <span id="address"></span> Address


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| city | string| `string` | ✓ | | city name |  |
| country | string| `string` | ✓ | | IISO 3166-1 alpha-2 country code. https:en.wikipedia.org/wiki/List_of_ISO_3166_country_codes |  |
| isDefault | boolean (formatted boolean)| `bool` | ✓ | | indicates if this is a default address |  |
| postalCode | string| `string` | ✓ | | postal or zip code |  |
| stateProvince | string| `string` | ✓ | | state or province name |  |
| street | string| `string` | ✓ | | street name, ie: 1723 NW 23rd Ave. |  |



### <span id="create-customer-request"></span> CreateCustomerRequest


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| customer | [NewCustomer](#new-customer)| `NewCustomer` | ✓ | | NewCustomer object |  |



### <span id="create-customer-response"></span> CreateCustomerResponse


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| customer | [Customer](#customer)| `Customer` | ✓ | | Customer object |  |
| status | [ResponseStatus](#response-status)| `ResponseStatus` | ✓ | | a ResponseStatus object |  |



### <span id="customer"></span> Customer


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| email | string| `string` | ✓ | | email address |  |
| firstName | string| `string` | ✓ | | first name |  |
| id | int64 (formatted integer)| `int64` | ✓ | | customer id |  |
| lastName | string| `string` | ✓ | | last or given name |  |
| password | string| `string` |  | | password |  |



### <span id="customer-account"></span> CustomerAccount


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| billingAddress | [Address](#address)| `Address` | ✓ | | Address object |  |
| id | int64 (formatted integer)| `int64` | ✓ | | customer id |  |
| shippingAddresses | [][Address](#address)| `[]*Address` | ✓ | | collection of Address objects |  |



### <span id="customer-login-request"></span> CustomerLoginRequest


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| email | string| `string` | ✓ | | email address, unique to each store id |  |
| password | string| `string` | ✓ | | password |  |



### <span id="customer-login-response"></span> CustomerLoginResponse


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| customer | [Customer](#customer)| `Customer` | ✓ | | Customer object |  |
| jwt | [JwtToken](#jwt-token)| `JwtToken` | ✓ | | jwt token |  |
| status | [ResponseStatus](#response-status)| `ResponseStatus` | ✓ | | a ResponseStatus object |  |



### <span id="jwt-token"></span> JwtToken


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| accessExpire | int64 (formatted integer)| `int64` | ✓ | |  |  |
| accessToken | string| `string` | ✓ | |  |  |
| refreshAfter | int64 (formatted integer)| `int64` | ✓ | |  |  |



### <span id="new-customer"></span> NewCustomer


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| billingAddress | [Address](#address)| `Address` |  | | Address object |  |
| email | string| `string` | ✓ | | email address, unique per store id |  |
| firstName | string| `string` | ✓ | | first name |  |
| lastName | string| `string` | ✓ | | last or given name |  |
| password | string| `string` | ✓ | | password |  |
| shippingAddresses | [Address](#address)| `Address` |  | | Address object |  |



### <span id="response-status"></span> ResponseStatus


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| statusCode | int64 (formatted integer)| `int64` | ✓ | | RFC http status code, ie. 204, etc - https:go.dev/src/net/http/status.go |  |
| statusMessage | string| `string` | ✓ | | status message |  |


