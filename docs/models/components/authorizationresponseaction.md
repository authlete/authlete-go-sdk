# AuthorizationResponseAction

The next action that the authorization server implementation should take.

## Example Usage

```go
import (
	"github.com/authlete/authlete-go-sdk/models/components"
)

value := components.AuthorizationResponseActionInternalServerError

// Open enum: custom values can be created with a direct type cast
custom := components.AuthorizationResponseAction("custom_value")
```


## Values

| Name                                             | Value                                            |
| ------------------------------------------------ | ------------------------------------------------ |
| `AuthorizationResponseActionInternalServerError` | INTERNAL_SERVER_ERROR                            |
| `AuthorizationResponseActionBadRequest`          | BAD_REQUEST                                      |
| `AuthorizationResponseActionLocation`            | LOCATION                                         |
| `AuthorizationResponseActionForm`                | FORM                                             |
| `AuthorizationResponseActionNoInteraction`       | NO_INTERACTION                                   |
| `AuthorizationResponseActionInteraction`         | INTERACTION                                      |