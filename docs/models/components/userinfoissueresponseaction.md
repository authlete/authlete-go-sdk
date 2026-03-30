# UserinfoIssueResponseAction

The next action that the authorization server implementation should take.

## Example Usage

```go
import (
	"github.com/authlete/authlete-go-sdk/models/components"
)

value := components.UserinfoIssueResponseActionInternalServerError

// Open enum: custom values can be created with a direct type cast
custom := components.UserinfoIssueResponseAction("custom_value")
```


## Values

| Name                                             | Value                                            |
| ------------------------------------------------ | ------------------------------------------------ |
| `UserinfoIssueResponseActionInternalServerError` | INTERNAL_SERVER_ERROR                            |
| `UserinfoIssueResponseActionBadRequest`          | BAD_REQUEST                                      |
| `UserinfoIssueResponseActionUnauthorized`        | UNAUTHORIZED                                     |
| `UserinfoIssueResponseActionForbidden`           | FORBIDDEN                                        |
| `UserinfoIssueResponseActionJSON`                | JSON                                             |
| `UserinfoIssueResponseActionJwt`                 | JWT                                              |