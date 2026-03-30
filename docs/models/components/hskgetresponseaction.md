# HskGetResponseAction

Result of the API call

## Example Usage

```go
import (
	"github.com/authlete/authlete-go-sdk/models/components"
)

value := components.HskGetResponseActionSuccess

// Open enum: custom values can be created with a direct type cast
custom := components.HskGetResponseAction("custom_value")
```


## Values

| Name                                 | Value                                |
| ------------------------------------ | ------------------------------------ |
| `HskGetResponseActionSuccess`        | SUCCESS                              |
| `HskGetResponseActionInvalidRequest` | INVALID_REQUEST                      |
| `HskGetResponseActionNotFound`       | NOT_FOUND                            |
| `HskGetResponseActionServerError`    | SERVER_ERROR                         |