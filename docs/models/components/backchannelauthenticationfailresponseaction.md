# BackchannelAuthenticationFailResponseAction

The next action that the authorization server implementation should take.

## Example Usage

```go
import (
	"github.com/authlete/authlete-go-sdk/models/components"
)

value := components.BackchannelAuthenticationFailResponseActionInternalServerError

// Open enum: custom values can be created with a direct type cast
custom := components.BackchannelAuthenticationFailResponseAction("custom_value")
```


## Values

| Name                                                             | Value                                                            |
| ---------------------------------------------------------------- | ---------------------------------------------------------------- |
| `BackchannelAuthenticationFailResponseActionInternalServerError` | INTERNAL_SERVER_ERROR                                            |
| `BackchannelAuthenticationFailResponseActionBadRequest`          | BAD_REQUEST                                                      |
| `BackchannelAuthenticationFailResponseActionForbidden`           | FORBIDDEN                                                        |