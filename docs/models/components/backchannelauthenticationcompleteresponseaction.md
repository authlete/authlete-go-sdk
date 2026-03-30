# BackchannelAuthenticationCompleteResponseAction

The next action that the authorization server implementation should take.


## Example Usage

```go
import (
	"github.com/authlete/authlete-go-sdk/models/components"
)

value := components.BackchannelAuthenticationCompleteResponseActionServerError

// Open enum: custom values can be created with a direct type cast
custom := components.BackchannelAuthenticationCompleteResponseAction("custom_value")
```


## Values

| Name                                                          | Value                                                         |
| ------------------------------------------------------------- | ------------------------------------------------------------- |
| `BackchannelAuthenticationCompleteResponseActionServerError`  | SERVER_ERROR                                                  |
| `BackchannelAuthenticationCompleteResponseActionNoAction`     | NO_ACTION                                                     |
| `BackchannelAuthenticationCompleteResponseActionNotification` | NOTIFICATION                                                  |