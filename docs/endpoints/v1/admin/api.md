


  

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

###  apiops

| Method  | URI     | Name   | Summary |
|---------|---------|--------|---------|
| GET | /v1/api/ping | [ping](#ping) | Ping |
  


## Paths

### <span id="ping"></span> Ping (*ping*)

```
GET /v1/api/ping
```

Ping for API up validation. On valid returns {ping: PONG}

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#ping-200) | OK | A successful response. |  | [schema](#ping-200-schema) |

#### Responses


##### <span id="ping-200"></span> 200 - A successful response.
Status: OK

###### <span id="ping-200-schema"></span> Schema
   
  

[PingResponse](#ping-response)

## Models

### <span id="jwt-token"></span> JwtToken


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| accessExpire | int64 (formatted integer)| `int64` | ✓ | |  |  |
| accessToken | string| `string` | ✓ | |  |  |
| refreshAfter | int64 (formatted integer)| `int64` | ✓ | |  |  |



### <span id="ping-response"></span> PingResponse


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| ping | string| `string` | ✓ | |  |  |



### <span id="response-status"></span> ResponseStatus


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| statusCode | int64 (formatted integer)| `int64` | ✓ | | RFC http status code, ie. 204, etc - https:go.dev/src/net/http/status.go |  |
| statusMessage | string| `string` | ✓ | | status message |  |


