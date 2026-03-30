# UserinfoResponseAction

The next action that the authorization server implementation should take.

## Example Usage

```go
import (
	"github.com/authlete/authlete-go-sdk/models/components"
)

value := components.UserinfoResponseActionInternalServerError

// Open enum: custom values can be created with a direct type cast
custom := components.UserinfoResponseAction("custom_value")
```


## Values

| Name                                        | Value                                       |
| ------------------------------------------- | ------------------------------------------- |
| `UserinfoResponseActionInternalServerError` | INTERNAL_SERVER_ERROR                       |
| `UserinfoResponseActionBadRequest`          | BAD_REQUEST                                 |
| `UserinfoResponseActionUnauthorized`        | UNAUTHORIZED                                |
| `UserinfoResponseActionForbidden`           | FORBIDDEN                                   |
| `UserinfoResponseActionOk`                  | OK                                          |