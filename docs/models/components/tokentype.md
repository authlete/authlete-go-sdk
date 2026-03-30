# TokenType

The token type identifier used in OAuth 2.0 Token Exchange (RFC 8693).
The API returns short codes (enum constant names) in response fields.


## Example Usage

```go
import (
	"github.com/authlete/authlete-go-sdk/models/components"
)

value := components.TokenTypeJwt

// Open enum: custom values can be created with a direct type cast
custom := components.TokenType("custom_value")
```


## Values

| Name                     | Value                    |
| ------------------------ | ------------------------ |
| `TokenTypeJwt`           | JWT                      |
| `TokenTypeAccessToken`   | ACCESS_TOKEN             |
| `TokenTypeRefreshToken`  | REFRESH_TOKEN            |
| `TokenTypeIDToken`       | ID_TOKEN                 |
| `TokenTypeSaml1`         | SAML1                    |
| `TokenTypeSaml2`         | SAML2                    |
| `TokenTypeDeviceSecret`  | DEVICE_SECRET            |
| `TokenTypeDeviceCode`    | DEVICE_CODE              |
| `TokenTypeTokenExchange` | TOKEN_EXCHANGE           |
| `TokenTypeJwtBearer`     | JWT_BEARER               |