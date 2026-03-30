# RevocationResponseAction

The next action that the authorization server implementation should take.

## Example Usage

```go
import (
	"github.com/authlete/authlete-go-sdk/models/components"
)

value := components.RevocationResponseActionInternalServerError

// Open enum: custom values can be created with a direct type cast
custom := components.RevocationResponseAction("custom_value")
```


## Values

| Name                                          | Value                                         |
| --------------------------------------------- | --------------------------------------------- |
| `RevocationResponseActionInternalServerError` | INTERNAL_SERVER_ERROR                         |
| `RevocationResponseActionInvalidClient`       | INVALID_CLIENT                                |
| `RevocationResponseActionBadRequest`          | BAD_REQUEST                                   |
| `RevocationResponseActionOk`                  | OK                                            |