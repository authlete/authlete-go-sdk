# AuthorizationTicketInfoResponseAction

The result of the `/auth/authorization/ticket/info` API call.

## Example Usage

```go
import (
	"github.com/authlete/authlete-go-sdk/models/components"
)

value := components.AuthorizationTicketInfoResponseActionOk

// Open enum: custom values can be created with a direct type cast
custom := components.AuthorizationTicketInfoResponseAction("custom_value")
```


## Values

| Name                                                 | Value                                                |
| ---------------------------------------------------- | ---------------------------------------------------- |
| `AuthorizationTicketInfoResponseActionOk`            | OK                                                   |
| `AuthorizationTicketInfoResponseActionNotFound`      | NOT_FOUND                                            |
| `AuthorizationTicketInfoResponseActionCallerError`   | CALLER_ERROR                                         |
| `AuthorizationTicketInfoResponseActionAuthleteError` | AUTHLETE_ERROR                                       |