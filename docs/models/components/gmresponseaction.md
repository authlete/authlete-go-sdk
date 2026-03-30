# GMResponseAction

The next action that the authorization server implementation should take.

## Example Usage

```go
import (
	"github.com/authlete/authlete-go-sdk/models/components"
)

value := components.GMResponseActionOk

// Open enum: custom values can be created with a direct type cast
custom := components.GMResponseAction("custom_value")
```


## Values

| Name                            | Value                           |
| ------------------------------- | ------------------------------- |
| `GMResponseActionOk`            | OK                              |
| `GMResponseActionNoContent`     | NO_CONTENT                      |
| `GMResponseActionUnauthorized`  | UNAUTHORIZED                    |
| `GMResponseActionForbidden`     | FORBIDDEN                       |
| `GMResponseActionNotFound`      | NOT_FOUND                       |
| `GMResponseActionCallerError`   | CALLER_ERROR                    |
| `GMResponseActionAuthleteError` | AUTHLETE_ERROR                  |