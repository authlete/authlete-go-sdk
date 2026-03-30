# IntrospectionResponseAction

The next action that the authorization server implementation should take.

## Example Usage

```go
import (
	"github.com/authlete/authlete-go-sdk/models/components"
)

value := components.IntrospectionResponseActionInternalServerError

// Open enum: custom values can be created with a direct type cast
custom := components.IntrospectionResponseAction("custom_value")
```


## Values

| Name                                             | Value                                            |
| ------------------------------------------------ | ------------------------------------------------ |
| `IntrospectionResponseActionInternalServerError` | INTERNAL_SERVER_ERROR                            |
| `IntrospectionResponseActionBadRequest`          | BAD_REQUEST                                      |
| `IntrospectionResponseActionUnauthorized`        | UNAUTHORIZED                                     |
| `IntrospectionResponseActionForbidden`           | FORBIDDEN                                        |
| `IntrospectionResponseActionOk`                  | OK                                               |