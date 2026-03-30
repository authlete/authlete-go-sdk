# IdtokenReissueResponseAction

The next action that the implementation of the token endpoint should take.

## Example Usage

```go
import (
	"github.com/authlete/authlete-go-sdk/models/components"
)

value := components.IdtokenReissueResponseActionOk

// Open enum: custom values can be created with a direct type cast
custom := components.IdtokenReissueResponseAction("custom_value")
```


## Values

| Name                                              | Value                                             |
| ------------------------------------------------- | ------------------------------------------------- |
| `IdtokenReissueResponseActionOk`                  | OK                                                |
| `IdtokenReissueResponseActionInternalServerError` | INTERNAL_SERVER_ERROR                             |
| `IdtokenReissueResponseActionCallerError`         | CALLER_ERROR                                      |