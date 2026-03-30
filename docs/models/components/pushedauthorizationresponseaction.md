# PushedAuthorizationResponseAction

The next action that the authorization server implementation should take. Any other value other than "CREATED" should be handled as unsuccessful result.

## Example Usage

```go
import (
	"github.com/authlete/authlete-go-sdk/models/components"
)

value := components.PushedAuthorizationResponseActionCreated

// Open enum: custom values can be created with a direct type cast
custom := components.PushedAuthorizationResponseAction("custom_value")
```


## Values

| Name                                                   | Value                                                  |
| ------------------------------------------------------ | ------------------------------------------------------ |
| `PushedAuthorizationResponseActionCreated`             | CREATED                                                |
| `PushedAuthorizationResponseActionBadRequest`          | BAD_REQUEST                                            |
| `PushedAuthorizationResponseActionUnauthorized`        | UNAUTHORIZED                                           |
| `PushedAuthorizationResponseActionForbidden`           | FORBIDDEN                                              |
| `PushedAuthorizationResponseActionPayloadTooLarge`     | PAYLOAD_TOO_LARGE                                      |
| `PushedAuthorizationResponseActionInternalServerError` | INTERNAL_SERVER_ERROR                                  |