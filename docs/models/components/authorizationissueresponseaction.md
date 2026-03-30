# AuthorizationIssueResponseAction

The next action that the authorization server implementation should take.

## Example Usage

```go
import (
	"github.com/authlete/authlete-go-sdk/models/components"
)

value := components.AuthorizationIssueResponseActionInternalServerError

// Open enum: custom values can be created with a direct type cast
custom := components.AuthorizationIssueResponseAction("custom_value")
```


## Values

| Name                                                  | Value                                                 |
| ----------------------------------------------------- | ----------------------------------------------------- |
| `AuthorizationIssueResponseActionInternalServerError` | INTERNAL_SERVER_ERROR                                 |
| `AuthorizationIssueResponseActionBadRequest`          | BAD_REQUEST                                           |
| `AuthorizationIssueResponseActionLocation`            | LOCATION                                              |
| `AuthorizationIssueResponseActionForm`                | FORM                                                  |