# VciDeferredIssueResponseAction

The next action that the implementation of the deferred credential
endpoint should take.


## Example Usage

```go
import (
	"github.com/authlete/authlete-go-sdk/models/components"
)

value := components.VciDeferredIssueResponseActionOk

// Open enum: custom values can be created with a direct type cast
custom := components.VciDeferredIssueResponseAction("custom_value")
```


## Values

| Name                                                | Value                                               |
| --------------------------------------------------- | --------------------------------------------------- |
| `VciDeferredIssueResponseActionOk`                  | OK                                                  |
| `VciDeferredIssueResponseActionForbidden`           | FORBIDDEN                                           |
| `VciDeferredIssueResponseActionInternalServerError` | INTERNAL_SERVER_ERROR                               |
| `VciDeferredIssueResponseActionCallerError`         | CALLER_ERROR                                        |