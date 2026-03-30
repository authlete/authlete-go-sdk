# GrantType

The grant type of the access token when the access token was created.


## Example Usage

```go
import (
	"github.com/authlete/authlete-go-sdk/models/components"
)

value := components.GrantTypeAuthorizationCode

// Open enum: custom values can be created with a direct type cast
custom := components.GrantType("custom_value")
```


## Values

| Name                         | Value                        |
| ---------------------------- | ---------------------------- |
| `GrantTypeAuthorizationCode` | AUTHORIZATION_CODE           |
| `GrantTypeImplicit`          | IMPLICIT                     |
| `GrantTypePassword`          | PASSWORD                     |
| `GrantTypeClientCredentials` | CLIENT_CREDENTIALS           |
| `GrantTypeRefreshToken`      | REFRESH_TOKEN                |
| `GrantTypeCiba`              | CIBA                         |
| `GrantTypeDeviceCode`        | DEVICE_CODE                  |
| `GrantTypeTokenExchange`     | TOKEN_EXCHANGE               |
| `GrantTypeJwtBearer`         | JWT_BEARER                   |
| `GrantTypePreAuthorizedCode` | PRE_AUTHORIZED_CODE          |