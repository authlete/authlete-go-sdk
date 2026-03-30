# TokenResponseAction

The next action that the authorization server implementation should take.

## Example Usage

```go
import (
	"github.com/authlete/authlete-go-sdk/models/components"
)

value := components.TokenResponseActionInternalServerError

// Open enum: custom values can be created with a direct type cast
custom := components.TokenResponseAction("custom_value")
```


## Values

| Name                                     | Value                                    |
| ---------------------------------------- | ---------------------------------------- |
| `TokenResponseActionInternalServerError` | INTERNAL_SERVER_ERROR                    |
| `TokenResponseActionInvalidClient`       | INVALID_CLIENT                           |
| `TokenResponseActionBadRequest`          | BAD_REQUEST                              |
| `TokenResponseActionPassword`            | PASSWORD                                 |
| `TokenResponseActionOk`                  | OK                                       |
| `TokenResponseActionTokenExchange`       | TOKEN_EXCHANGE                           |
| `TokenResponseActionJwtBearer`           | JWT_BEARER                               |
| `TokenResponseActionNativeSso`           | NATIVE_SSO                               |
| `TokenResponseActionIDTokenReissuable`   | ID_TOKEN_REISSUABLE                      |