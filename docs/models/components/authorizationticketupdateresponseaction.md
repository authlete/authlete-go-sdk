# AuthorizationTicketUpdateResponseAction

The result of the /auth/authorization/ticket/info API call.

## Example Usage

```go
import (
	"github.com/authlete/authlete-go-sdk/models/components"
)

value := components.AuthorizationTicketUpdateResponseActionOk

// Open enum: custom values can be created with a direct type cast
custom := components.AuthorizationTicketUpdateResponseAction("custom_value")
```


## Values

| Name                                                   | Value                                                  |
| ------------------------------------------------------ | ------------------------------------------------------ |
| `AuthorizationTicketUpdateResponseActionOk`            | OK                                                     |
| `AuthorizationTicketUpdateResponseActionNotFound`      | NOT_FOUND                                              |
| `AuthorizationTicketUpdateResponseActionCallerError`   | CALLER_ERROR                                           |
| `AuthorizationTicketUpdateResponseActionAuthleteError` | AUTHLETE_ERROR                                         |