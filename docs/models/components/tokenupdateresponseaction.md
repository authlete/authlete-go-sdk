# TokenUpdateResponseAction

The next action that the authorization server implementation should take.

## Example Usage

```go
import (
	"github.com/authlete/authlete-go-sdk/models/components"
)

value := components.TokenUpdateResponseActionInternalServerError

// Open enum: custom values can be created with a direct type cast
custom := components.TokenUpdateResponseAction("custom_value")
```


## Values

| Name                                           | Value                                          |
| ---------------------------------------------- | ---------------------------------------------- |
| `TokenUpdateResponseActionInternalServerError` | INTERNAL_SERVER_ERROR                          |
| `TokenUpdateResponseActionBadRequest`          | BAD_REQUEST                                    |
| `TokenUpdateResponseActionForbidden`           | FORBIDDEN                                      |
| `TokenUpdateResponseActionNotFound`            | NOT_FOUND                                      |
| `TokenUpdateResponseActionOk`                  | OK                                             |