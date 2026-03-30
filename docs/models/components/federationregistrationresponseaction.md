# FederationRegistrationResponseAction

The next action that the authorization server implementation should take.

## Example Usage

```go
import (
	"github.com/authlete/authlete-go-sdk/models/components"
)

value := components.FederationRegistrationResponseActionOk

// Open enum: custom values can be created with a direct type cast
custom := components.FederationRegistrationResponseAction("custom_value")
```


## Values

| Name                                                      | Value                                                     |
| --------------------------------------------------------- | --------------------------------------------------------- |
| `FederationRegistrationResponseActionOk`                  | OK                                                        |
| `FederationRegistrationResponseActionBadRequest`          | BAD_REQUEST                                               |
| `FederationRegistrationResponseActionNotFound`            | NOT_FOUND                                                 |
| `FederationRegistrationResponseActionInternalServerError` | INTERNAL_SERVER_ERROR                                     |