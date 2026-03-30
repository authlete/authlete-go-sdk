# VciBatchParseResponseAction

The next action that the batch credential endpoint should take.

## Example Usage

```go
import (
	"github.com/authlete/authlete-go-sdk/models/components"
)

value := components.VciBatchParseResponseActionOk

// Open enum: custom values can be created with a direct type cast
custom := components.VciBatchParseResponseAction("custom_value")
```


## Values

| Name                                             | Value                                            |
| ------------------------------------------------ | ------------------------------------------------ |
| `VciBatchParseResponseActionOk`                  | OK                                               |
| `VciBatchParseResponseActionBadRequest`          | BAD_REQUEST                                      |
| `VciBatchParseResponseActionUnauthorized`        | UNAUTHORIZED                                     |
| `VciBatchParseResponseActionForbidden`           | FORBIDDEN                                        |
| `VciBatchParseResponseActionInternalServerError` | INTERNAL_SERVER_ERROR                            |