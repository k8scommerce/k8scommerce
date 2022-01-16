


# User API Endpoints
admin gateway user api endpoints
  

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

###  admin

| Method  | URI     | Name   | Summary |
|---------|---------|--------|---------|
| POST | /v1/user/login | [login](#login) | Login |
  


## Paths

### <span id="login"></span> Login (*login*)

```
POST /v1/user/login
```

login for administration users

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| body | `body` | [UserLoginRequest](#user-login-request) | `models.UserLoginRequest` | | ✓ | |  |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#login-200) | OK | A successful response. |  | [schema](#login-200-schema) |

#### Responses


##### <span id="login-200"></span> 200 - A successful response.
Status: OK

###### <span id="login-200-schema"></span> Schema
   
  

[UserLoginResponse](#user-login-response)

## Models

### <span id="jwt-token"></span> JwtToken


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| access_expire | int64 (formatted integer)| `int64` | ✓ | |  |  |
| access_token | string| `string` | ✓ | |  |  |
| refresh_after | int64 (formatted integer)| `int64` | ✓ | |  |  |



### <span id="user"></span> User


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| email | string| `string` | ✓ | |  |  |
| firstName | string| `string` | ✓ | |  |  |
| id | int64 (formatted integer)| `int64` | ✓ | |  |  |
| lastName | string| `string` | ✓ | |  |  |
| password | string| `string` | ✓ | |  |  |



### <span id="user-login-request"></span> UserLoginRequest


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| email | string| `string` | ✓ | |  |  |
| password | string| `string` | ✓ | |  |  |



### <span id="user-login-response"></span> UserLoginResponse


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| jwt | [JwtToken](#jwt-token)| `JwtToken` | ✓ | |  |  |
| statusCode | int64 (formatted integer)| `int64` | ✓ | |  |  |
| statusMessage | string| `string` | ✓ | |  |  |
| user | [User](#user)| `User` | ✓ | |  |  |


