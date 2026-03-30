# ApplicationType

The application type. The value of this property affects the validation steps for a redirect URI.
See the description about `redirectUris` property for more details.


## Example Usage

```go
import (
	"github.com/authlete/authlete-go-sdk/models/components"
)

value := components.ApplicationTypeWeb

// Open enum: custom values can be created with a direct type cast
custom := components.ApplicationType("custom_value")
```


## Values

| Name                    | Value                   |
| ----------------------- | ----------------------- |
| `ApplicationTypeWeb`    | WEB                     |
| `ApplicationTypeNative` | NATIVE                  |