# Admin Gateway API
admin gateway api

## Version: 1.0.0

### Security
**apiKey**  

|apiKey|*API Key*|
|---|---|
|Description|Enter JWT Bearer token **_only_**|
|Name|Authorization|
|In|header|

### /v1/user/login

#### POST
##### Summary

manages user logins

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| body | body |  | Yes | [UserLoginRequest](#userloginrequest) |

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | A successful response. | [UserLoginResponse](#userloginresponse) |

### Models

#### JwtToken

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| access_token | string |  | Yes |
| access_expire | long |  | Yes |
| refresh_after | long |  | Yes |

#### User

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| id | long |  | Yes |
| firstName | string |  | Yes |
| lastName | string |  | Yes |
| email | string |  | Yes |
| password | string |  | Yes |

#### UserLoginRequest

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| email | string |  | Yes |
| password | string |  | Yes |

#### UserLoginResponse

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| jwt | [JwtToken](#jwttoken) |  | Yes |
| user | [User](#user) |  | Yes |
| statusCode | long |  | Yes |
| statusMessage | string |  | Yes |
