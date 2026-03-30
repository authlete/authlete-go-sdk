# FederationConfigurationResponseAction

The next action that the authorization server implementation should take.

## Example Usage

```go
import (
	"github.com/authlete/authlete-go-sdk/models/components"
)

value := components.FederationConfigurationResponseActionOk

// Open enum: custom values can be created with a direct type cast
custom := components.FederationConfigurationResponseAction("custom_value")
```


## Values

| Name                                                       | Value                                                      |
| ---------------------------------------------------------- | ---------------------------------------------------------- |
| `FederationConfigurationResponseActionOk`                  | OK                                                         |
| `FederationConfigurationResponseActionNotFound`            | NOT_FOUND                                                  |
| `FederationConfigurationResponseActionInternalServerError` | INTERNAL_SERVER_ERROR                                      |