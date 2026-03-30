# DeviceAuthorizationResponseAction

The next action that the authorization server implementation should take.

## Example Usage

```go
import (
	"github.com/authlete/authlete-go-sdk/models/components"
)

value := components.DeviceAuthorizationResponseActionInternalServerError

// Open enum: custom values can be created with a direct type cast
custom := components.DeviceAuthorizationResponseAction("custom_value")
```


## Values

| Name                                                   | Value                                                  |
| ------------------------------------------------------ | ------------------------------------------------------ |
| `DeviceAuthorizationResponseActionInternalServerError` | INTERNAL_SERVER_ERROR                                  |
| `DeviceAuthorizationResponseActionBadRequest`          | BAD_REQUEST                                            |
| `DeviceAuthorizationResponseActionUnauthorized`        | UNAUTHORIZED                                           |
| `DeviceAuthorizationResponseActionOk`                  | OK                                                     |