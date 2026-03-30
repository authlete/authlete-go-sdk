# HskCreateResponseAction

Result of the API call.

## Example Usage

```go
import (
	"github.com/authlete/authlete-go-sdk/models/components"
)

value := components.HskCreateResponseActionSuccess

// Open enum: custom values can be created with a direct type cast
custom := components.HskCreateResponseAction("custom_value")
```


## Values

| Name                                    | Value                                   |
| --------------------------------------- | --------------------------------------- |
| `HskCreateResponseActionSuccess`        | SUCCESS                                 |
| `HskCreateResponseActionInvalidRequest` | INVALID_REQUEST                         |
| `HskCreateResponseActionNotFound`       | NOT_FOUND                               |
| `HskCreateResponseActionServerError`    | SERVER_ERROR                            |