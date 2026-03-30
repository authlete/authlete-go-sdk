# VciBatchIssueResponseAction

The next action that the implementation of the batch credential
endpoint should take.


## Example Usage

```go
import (
	"github.com/authlete/authlete-go-sdk/models/components"
)

value := components.VciBatchIssueResponseActionOk

// Open enum: custom values can be created with a direct type cast
custom := components.VciBatchIssueResponseAction("custom_value")
```


## Values

| Name                                             | Value                                            |
| ------------------------------------------------ | ------------------------------------------------ |
| `VciBatchIssueResponseActionOk`                  | OK                                               |
| `VciBatchIssueResponseActionUnauthorized`        | UNAUTHORIZED                                     |
| `VciBatchIssueResponseActionForbidden`           | FORBIDDEN                                        |
| `VciBatchIssueResponseActionInternalServerError` | INTERNAL_SERVER_ERROR                            |
| `VciBatchIssueResponseActionCallerError`         | CALLER_ERROR                                     |