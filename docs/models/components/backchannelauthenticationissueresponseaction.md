# BackchannelAuthenticationIssueResponseAction

The next action that the authorization server implementation should take.

## Example Usage

```go
import (
	"github.com/authlete/authlete-go-sdk/models/components"
)

value := components.BackchannelAuthenticationIssueResponseActionInternalServerError

// Open enum: custom values can be created with a direct type cast
custom := components.BackchannelAuthenticationIssueResponseAction("custom_value")
```


## Values

| Name                                                              | Value                                                             |
| ----------------------------------------------------------------- | ----------------------------------------------------------------- |
| `BackchannelAuthenticationIssueResponseActionInternalServerError` | INTERNAL_SERVER_ERROR                                             |
| `BackchannelAuthenticationIssueResponseActionInvalidTicket`       | INVALID_TICKET                                                    |
| `BackchannelAuthenticationIssueResponseActionOk`                  | OK                                                                |