# TokenIssueResponseAction

The next action that the authorization server implementation should take.

## Example Usage

```go
import (
	"github.com/authlete/authlete-go-sdk/models/components"
)

value := components.TokenIssueResponseActionInternalServerError

// Open enum: custom values can be created with a direct type cast
custom := components.TokenIssueResponseAction("custom_value")
```


## Values

| Name                                          | Value                                         |
| --------------------------------------------- | --------------------------------------------- |
| `TokenIssueResponseActionInternalServerError` | INTERNAL_SERVER_ERROR                         |
| `TokenIssueResponseActionOk`                  | OK                                            |