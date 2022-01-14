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

### /v1/customer/login

#### POST
##### Summary

customer logic

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| body | body |  | Yes | [CustomerLoginRequest](#customerloginrequest) |

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | A successful response. | [CustomerLoginResponse](#customerloginresponse) |

### Models

#### Customer

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| id | long |  | Yes |
| firstName | string |  | Yes |
| lastName | string |  | Yes |
| email | string |  | Yes |
| password | string |  | Yes |

#### CustomerLoginRequest

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| email | string |  | Yes |
| password | string |  | Yes |

#### CustomerLoginResponse

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| jwt | [JwtToken](#jwttoken) |  | Yes |
| customer | [Customer](#customer) |  | Yes |
| statusCode | long |  | Yes |
| statusMessage | string |  | Yes |

#### JwtToken

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| access_token | string |  | Yes |
| access_expire | long |  | Yes |
| refresh_after | long |  | Yes |
