# DeviceCompleteRequestResult

The result of the end-user authentication and authorization. One of the following. Details are
described in the description.


## Example Usage

```go
import (
	"github.com/authlete/authlete-go-sdk/models/components"
)

value := components.DeviceCompleteRequestResultTransactionFailed
```


## Values

| Name                                           | Value                                          |
| ---------------------------------------------- | ---------------------------------------------- |
| `DeviceCompleteRequestResultTransactionFailed` | TRANSACTION_FAILED                             |
| `DeviceCompleteRequestResultAccessDenied`      | ACCESS_DENIED                                  |
| `DeviceCompleteRequestResultAuthorized`        | AUTHORIZED                                     |