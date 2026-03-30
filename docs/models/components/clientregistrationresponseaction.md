# ClientRegistrationResponseAction

The next action that the authorization server implementation should take.


## Example Usage

```go
import (
	"github.com/authlete/authlete-go-sdk/models/components"
)

value := components.ClientRegistrationResponseActionBadRequest

// Open enum: custom values can be created with a direct type cast
custom := components.ClientRegistrationResponseAction("custom_value")
```


## Values

| Name                                                  | Value                                                 |
| ----------------------------------------------------- | ----------------------------------------------------- |
| `ClientRegistrationResponseActionBadRequest`          | BAD_REQUEST                                           |
| `ClientRegistrationResponseActionCreated`             | CREATED                                               |
| `ClientRegistrationResponseActionDeleted`             | DELETED                                               |
| `ClientRegistrationResponseActionInternalServerError` | INTERNAL_SERVER_ERROR                                 |
| `ClientRegistrationResponseActionOk`                  | OK                                                    |
| `ClientRegistrationResponseActionUnauthorized`        | UNAUTHORIZED                                          |
| `ClientRegistrationResponseActionUpdated`             | UPDATED                                               |