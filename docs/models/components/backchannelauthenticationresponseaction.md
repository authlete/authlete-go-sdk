# BackchannelAuthenticationResponseAction

The next action that the authorization server implementation should take.

## Example Usage

```go
import (
	"github.com/authlete/authlete-go-sdk/models/components"
)

value := components.BackchannelAuthenticationResponseActionInternalServerError

// Open enum: custom values can be created with a direct type cast
custom := components.BackchannelAuthenticationResponseAction("custom_value")
```


## Values

| Name                                                         | Value                                                        |
| ------------------------------------------------------------ | ------------------------------------------------------------ |
| `BackchannelAuthenticationResponseActionInternalServerError` | INTERNAL_SERVER_ERROR                                        |
| `BackchannelAuthenticationResponseActionBadRequest`          | BAD_REQUEST                                                  |
| `BackchannelAuthenticationResponseActionUnauthorized`        | UNAUTHORIZED                                                 |
| `BackchannelAuthenticationResponseActionUserIdentification`  | USER_IDENTIFICATION                                          |