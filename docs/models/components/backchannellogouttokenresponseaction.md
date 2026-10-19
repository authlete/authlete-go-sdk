# BackchannelLogoutTokenResponseAction

The next action that the API caller should take.


## Example Usage

```go
import (
	"github.com/authlete/authlete-go-sdk/models/components"
)

value := components.BackchannelLogoutTokenResponseActionOk

// Open enum: custom values can be created with a direct type cast
custom := components.BackchannelLogoutTokenResponseAction("custom_value")
```


## Values

| Name                                              | Value                                             |
| ------------------------------------------------- | ------------------------------------------------- |
| `BackchannelLogoutTokenResponseActionOk`          | OK                                                |
| `BackchannelLogoutTokenResponseActionServerError` | SERVER_ERROR                                      |
| `BackchannelLogoutTokenResponseActionCallerError` | CALLER_ERROR                                      |