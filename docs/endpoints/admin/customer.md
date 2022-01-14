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

### Models

#### Customer

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| id | long |  | Yes |
| firstName | string |  | Yes |
| lastName | string |  | Yes |
| email | string |  | Yes |
| password | string |  | Yes |
