# BackchannelAuthenticationCompleteRequestResult

The result of the end-user authentication and authorization. One of the following. Details are
described in the description.


## Example Usage

```go
import (
	"github.com/authlete/authlete-go-sdk/models/components"
)

value := components.BackchannelAuthenticationCompleteRequestResultTransactionFailed
```


## Values

| Name                                                              | Value                                                             |
| ----------------------------------------------------------------- | ----------------------------------------------------------------- |
| `BackchannelAuthenticationCompleteRequestResultTransactionFailed` | TRANSACTION_FAILED                                                |
| `BackchannelAuthenticationCompleteRequestResultAccessDenied`      | ACCESS_DENIED                                                     |
| `BackchannelAuthenticationCompleteRequestResultAuthorized`        | AUTHORIZED                                                        |