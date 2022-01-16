


# Client Gateway API
client gateway api
  

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
| POST | /v1/customer/login | [login](#login) | Login |
  


## Paths

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

### <span id="customer"></span> Customer


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| email | string| `string` | ✓ | |  |  |
| firstName | string| `string` | ✓ | |  |  |
| id | int64 (formatted integer)| `int64` | ✓ | |  |  |
| lastName | string| `string` | ✓ | |  |  |
| password | string| `string` | ✓ | |  |  |



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
| statusCode | int64 (formatted integer)| `int64` | ✓ | |  |  |
| statusMessage | string| `string` | ✓ | |  |  |



### <span id="jwt-token"></span> JwtToken


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| access_expire | int64 (formatted integer)| `int64` | ✓ | |  |  |
| access_token | string| `string` | ✓ | |  |  |
| refresh_after | int64 (formatted integer)| `int64` | ✓ | |  |  |


