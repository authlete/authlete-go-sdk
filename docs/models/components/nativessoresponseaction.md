# NativeSsoResponseAction

The next action that the implementation of the token endpoint should take.


## Example Usage

```go
import (
	"github.com/authlete/authlete-go-sdk/models/components"
)

value := components.NativeSsoResponseActionOk

// Open enum: custom values can be created with a direct type cast
custom := components.NativeSsoResponseAction("custom_value")
```


## Values

| Name                                         | Value                                        |
| -------------------------------------------- | -------------------------------------------- |
| `NativeSsoResponseActionOk`                  | OK                                           |
| `NativeSsoResponseActionInternalServerError` | INTERNAL_SERVER_ERROR                        |
| `NativeSsoResponseActionCallerError`         | CALLER_ERROR                                 |