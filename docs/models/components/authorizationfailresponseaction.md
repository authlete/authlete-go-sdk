# AuthorizationFailResponseAction

The next action that the authorization server implementation should take.

## Example Usage

```go
import (
	"github.com/authlete/authlete-go-sdk/models/components"
)

value := components.AuthorizationFailResponseActionInternalServerError

// Open enum: custom values can be created with a direct type cast
custom := components.AuthorizationFailResponseAction("custom_value")
```


## Values

| Name                                                 | Value                                                |
| ---------------------------------------------------- | ---------------------------------------------------- |
| `AuthorizationFailResponseActionInternalServerError` | INTERNAL_SERVER_ERROR                                |
| `AuthorizationFailResponseActionBadRequest`          | BAD_REQUEST                                          |
| `AuthorizationFailResponseActionLocation`            | LOCATION                                             |
| `AuthorizationFailResponseActionForm`                | FORM                                                 |