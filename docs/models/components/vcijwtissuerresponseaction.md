# VciJwtissuerResponseAction

The next action that the implementation of the JWT issuer metadata
endpoint (`/.well-known/jwt-issuer`) should take after getting
a response from Authlete's `/vci/jwtissuer` API.


## Example Usage

```go
import (
	"github.com/authlete/authlete-go-sdk/models/components"
)

value := components.VciJwtissuerResponseActionOk

// Open enum: custom values can be created with a direct type cast
custom := components.VciJwtissuerResponseAction("custom_value")
```


## Values

| Name                                            | Value                                           |
| ----------------------------------------------- | ----------------------------------------------- |
| `VciJwtissuerResponseActionOk`                  | OK                                              |
| `VciJwtissuerResponseActionNotFound`            | NOT_FOUND                                       |
| `VciJwtissuerResponseActionInternalServerError` | INTERNAL_SERVER_ERROR                           |