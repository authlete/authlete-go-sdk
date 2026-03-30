# VciJwksResponseAction

The next action that the implementation of the JWK Set document
endpoint of the credential issuer should take after getting a
response from Authlete's `/vci/jwks` API.


## Example Usage

```go
import (
	"github.com/authlete/authlete-go-sdk/models/components"
)

value := components.VciJwksResponseActionOk

// Open enum: custom values can be created with a direct type cast
custom := components.VciJwksResponseAction("custom_value")
```


## Values

| Name                                       | Value                                      |
| ------------------------------------------ | ------------------------------------------ |
| `VciJwksResponseActionOk`                  | OK                                         |
| `VciJwksResponseActionNotFound`            | NOT_FOUND                                  |
| `VciJwksResponseActionInternalServerError` | INTERNAL_SERVER_ERROR                      |