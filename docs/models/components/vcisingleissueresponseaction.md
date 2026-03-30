# VciSingleIssueResponseAction

The next action that the implementation of the credential endpoint
should take.


## Example Usage

```go
import (
	"github.com/authlete/authlete-go-sdk/models/components"
)

value := components.VciSingleIssueResponseActionOk

// Open enum: custom values can be created with a direct type cast
custom := components.VciSingleIssueResponseAction("custom_value")
```


## Values

| Name                                              | Value                                             |
| ------------------------------------------------- | ------------------------------------------------- |
| `VciSingleIssueResponseActionOk`                  | OK                                                |
| `VciSingleIssueResponseActionAccepted`            | ACCEPTED                                          |
| `VciSingleIssueResponseActionUnauthorized`        | UNAUTHORIZED                                      |
| `VciSingleIssueResponseActionForbidden`           | FORBIDDEN                                         |
| `VciSingleIssueResponseActionInternalServerError` | INTERNAL_SERVER_ERROR                             |
| `VciSingleIssueResponseActionCallerError`         | CALLER_ERROR                                      |