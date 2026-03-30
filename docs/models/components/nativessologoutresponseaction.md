# NativeSsoLogoutResponseAction

The next action that the API caller should take.


## Example Usage

```go
import (
	"github.com/authlete/authlete-go-sdk/models/components"
)

value := components.NativeSsoLogoutResponseActionOk

// Open enum: custom values can be created with a direct type cast
custom := components.NativeSsoLogoutResponseAction("custom_value")
```


## Values

| Name                                       | Value                                      |
| ------------------------------------------ | ------------------------------------------ |
| `NativeSsoLogoutResponseActionOk`          | OK                                         |
| `NativeSsoLogoutResponseActionServerError` | SERVER_ERROR                               |
| `NativeSsoLogoutResponseActionCallerError` | CALLER_ERROR                               |