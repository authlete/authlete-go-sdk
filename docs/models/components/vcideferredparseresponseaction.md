# VciDeferredParseResponseAction

The next action that the deferred credential endpoint should take.

## Example Usage

```go
import (
	"github.com/authlete/authlete-go-sdk/models/components"
)

value := components.VciDeferredParseResponseActionOk

// Open enum: custom values can be created with a direct type cast
custom := components.VciDeferredParseResponseAction("custom_value")
```


## Values

| Name                                                | Value                                               |
| --------------------------------------------------- | --------------------------------------------------- |
| `VciDeferredParseResponseActionOk`                  | OK                                                  |
| `VciDeferredParseResponseActionBadRequest`          | BAD_REQUEST                                         |
| `VciDeferredParseResponseActionUnauthorized`        | UNAUTHORIZED                                        |
| `VciDeferredParseResponseActionForbidden`           | FORBIDDEN                                           |
| `VciDeferredParseResponseActionInternalServerError` | INTERNAL_SERVER_ERROR                               |