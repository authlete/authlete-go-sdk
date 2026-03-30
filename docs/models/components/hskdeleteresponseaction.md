# HskDeleteResponseAction

Result of the API call

## Example Usage

```go
import (
	"github.com/authlete/authlete-go-sdk/models/components"
)

value := components.HskDeleteResponseActionSuccess

// Open enum: custom values can be created with a direct type cast
custom := components.HskDeleteResponseAction("custom_value")
```


## Values

| Name                                    | Value                                   |
| --------------------------------------- | --------------------------------------- |
| `HskDeleteResponseActionSuccess`        | SUCCESS                                 |
| `HskDeleteResponseActionInvalidRequest` | INVALID_REQUEST                         |
| `HskDeleteResponseActionNotFound`       | NOT_FOUND                               |
| `HskDeleteResponseActionServerError`    | SERVER_ERROR                            |