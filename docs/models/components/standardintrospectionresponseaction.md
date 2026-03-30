# StandardIntrospectionResponseAction

The next action that the authorization server implementation should take.

## Example Usage

```go
import (
	"github.com/authlete/authlete-go-sdk/models/components"
)

value := components.StandardIntrospectionResponseActionInternalServerError

// Open enum: custom values can be created with a direct type cast
custom := components.StandardIntrospectionResponseAction("custom_value")
```


## Values

| Name                                                     | Value                                                    |
| -------------------------------------------------------- | -------------------------------------------------------- |
| `StandardIntrospectionResponseActionInternalServerError` | INTERNAL_SERVER_ERROR                                    |
| `StandardIntrospectionResponseActionBadRequest`          | BAD_REQUEST                                              |
| `StandardIntrospectionResponseActionOk`                  | OK                                                       |
| `StandardIntrospectionResponseActionJwt`                 | JWT                                                      |