# VciSingleParseResponseAction

The next action that the credential endpoint should take.

## Example Usage

```go
import (
	"github.com/authlete/authlete-go-sdk/models/components"
)

value := components.VciSingleParseResponseActionOk

// Open enum: custom values can be created with a direct type cast
custom := components.VciSingleParseResponseAction("custom_value")
```


## Values

| Name                                              | Value                                             |
| ------------------------------------------------- | ------------------------------------------------- |
| `VciSingleParseResponseActionOk`                  | OK                                                |
| `VciSingleParseResponseActionBadRequest`          | BAD_REQUEST                                       |
| `VciSingleParseResponseActionUnauthorized`        | UNAUTHORIZED                                      |
| `VciSingleParseResponseActionForbidden`           | FORBIDDEN                                         |
| `VciSingleParseResponseActionInternalServerError` | INTERNAL_SERVER_ERROR                             |