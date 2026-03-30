# TokenFailResponseAction

The next action that the authorization server implementation should take.

## Example Usage

```go
import (
	"github.com/authlete/authlete-go-sdk/models/components"
)

value := components.TokenFailResponseActionInternalServerError

// Open enum: custom values can be created with a direct type cast
custom := components.TokenFailResponseAction("custom_value")
```


## Values

| Name                                         | Value                                        |
| -------------------------------------------- | -------------------------------------------- |
| `TokenFailResponseActionInternalServerError` | INTERNAL_SERVER_ERROR                        |
| `TokenFailResponseActionBadRequest`          | BAD_REQUEST                                  |