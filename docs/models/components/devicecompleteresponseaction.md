# DeviceCompleteResponseAction

The next action that the authorization server implementation should take.


## Example Usage

```go
import (
	"github.com/authlete/authlete-go-sdk/models/components"
)

value := components.DeviceCompleteResponseActionServerError

// Open enum: custom values can be created with a direct type cast
custom := components.DeviceCompleteResponseAction("custom_value")
```


## Values

| Name                                           | Value                                          |
| ---------------------------------------------- | ---------------------------------------------- |
| `DeviceCompleteResponseActionServerError`      | SERVER_ERROR                                   |
| `DeviceCompleteResponseActionUserCodeNotExist` | USER_CODE_NOT_EXIST                            |
| `DeviceCompleteResponseActionUserCodeExpired`  | USER_CODE_EXPIRED                              |
| `DeviceCompleteResponseActionInvalidRequest`   | INVALID_REQUEST                                |
| `DeviceCompleteResponseActionSuccess`          | SUCCESS                                        |