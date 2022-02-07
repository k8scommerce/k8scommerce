


  

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

###  users

| Method  | URI     | Name   | Summary |
|---------|---------|--------|---------|
| GET | /v1/users/{currentPage}/{pageSize} | [get all users](#get-all-users) | Get Users |
| POST | /v1/user/login | [login](#login) | Login |
  


## Paths

### <span id="get-all-users"></span> Get Users (*getAllUsers*)

```
GET /v1/users/{currentPage}/{pageSize}
```

returns all users

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| currentPage | `path` | integer | `int64` |  | ✓ |  | current page number |
| pageSize | `path` | integer | `int64` |  | ✓ |  | number of records per page |
| sortOn | `query` | string | `string` |  |  |  |  |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#get-all-users-200) | OK | A successful response. |  | [schema](#get-all-users-200-schema) |

#### Responses


##### <span id="get-all-users-200"></span> 200 - A successful response.
Status: OK

###### <span id="get-all-users-200-schema"></span> Schema
   
  

[GetAllUsersResponse](#get-all-users-response)

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

### <span id="get-all-users-request"></span> GetAllUsersRequest


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| currentPage | int64 (formatted integer)| `int64` | ✓ | |  |  |
| pageSize | int64 (formatted integer)| `int64` | ✓ | |  |  |
| sortOn | string| `string` |  | |  |  |



### <span id="get-all-users-response"></span> GetAllUsersResponse


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| totalPages | int64 (formatted integer)| `int64` | ✓ | |  |  |
| totalRecords | int64 (formatted integer)| `int64` | ✓ | |  |  |
| users | [][User](#user)| `[]*User` | ✓ | |  |  |



### <span id="jwt-token"></span> JwtToken


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| accessExpire | int64 (formatted integer)| `int64` | ✓ | |  |  |
| accessToken | string| `string` | ✓ | |  |  |
| refreshAfter | int64 (formatted integer)| `int64` | ✓ | |  |  |



### <span id="permission-group"></span> PermissionGroup


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| groupName | string| `string` | ✓ | | groupName |  |
| id | int64 (formatted integer)| `int64` | ✓ | | permission group id |  |



### <span id="response-status"></span> ResponseStatus


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| statusCode | int64 (formatted integer)| `int64` | ✓ | | RFC http status code, ie. 204, etc - https:go.dev/src/net/http/status.go |  |
| statusMessage | string| `string` | ✓ | | status message |  |



### <span id="user"></span> User


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| email | string| `string` | ✓ | | email address |  |
| firstName | string| `string` | ✓ | | first name |  |
| id | int64 (formatted integer)| `int64` | ✓ | | user id |  |
| lastName | string| `string` | ✓ | | last name |  |
| password | string| `string` |  | | password |  |



### <span id="user-login-request"></span> UserLoginRequest


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| email | string| `string` | ✓ | | email address |  |
| password | string| `string` | ✓ | | password |  |



### <span id="user-login-response"></span> UserLoginResponse


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| jwt | [JwtToken](#jwt-token)| `JwtToken` | ✓ | | JwtToken object |  |
| user | [User](#user)| `User` | ✓ | | User object |  |



### <span id="users-permission-groups"></span> UsersPermissionGroups


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| permissionGroupId | int64 (formatted integer)| `int64` | ✓ | | permission group id |  |
| userId | int64 (formatted integer)| `int64` | ✓ | | user id |  |


