# TokenCreateResponseAction

The next action that the authorization server implementation should take.

## Example Usage

```go
import (
	"github.com/authlete/authlete-go-sdk/models/components"
)

value := components.TokenCreateResponseActionInternalServerError

// Open enum: custom values can be created with a direct type cast
custom := components.TokenCreateResponseAction("custom_value")
```


## Values

| Name                                           | Value                                          |
| ---------------------------------------------- | ---------------------------------------------- |
| `TokenCreateResponseActionInternalServerError` | INTERNAL_SERVER_ERROR                          |
| `TokenCreateResponseActionBadRequest`          | BAD_REQUEST                                    |
| `TokenCreateResponseActionForbidden`           | FORBIDDEN                                      |
| `TokenCreateResponseActionOk`                  | OK                                             |