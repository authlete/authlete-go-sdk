# HskGetListResponseAction

Result of the API call

## Example Usage

```go
import (
	"github.com/authlete/authlete-go-sdk/models/components"
)

value := components.HskGetListResponseActionSuccess

// Open enum: custom values can be created with a direct type cast
custom := components.HskGetListResponseAction("custom_value")
```


## Values

| Name                                     | Value                                    |
| ---------------------------------------- | ---------------------------------------- |
| `HskGetListResponseActionSuccess`        | SUCCESS                                  |
| `HskGetListResponseActionInvalidRequest` | INVALID_REQUEST                          |
| `HskGetListResponseActionServerError`    | SERVER_ERROR                             |