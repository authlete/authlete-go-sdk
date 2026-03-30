# ResponseType

## Example Usage

```go
import (
	"github.com/authlete/authlete-go-sdk/models/components"
)

value := components.ResponseTypeNone

// Open enum: custom values can be created with a direct type cast
custom := components.ResponseType("custom_value")
```


## Values

| Name                           | Value                          |
| ------------------------------ | ------------------------------ |
| `ResponseTypeNone`             | NONE                           |
| `ResponseTypeCode`             | CODE                           |
| `ResponseTypeToken`            | TOKEN                          |
| `ResponseTypeIDToken`          | ID_TOKEN                       |
| `ResponseTypeCodeToken`        | CODE_TOKEN                     |
| `ResponseTypeCodeIDToken`      | CODE_ID_TOKEN                  |
| `ResponseTypeIDTokenToken`     | ID_TOKEN_TOKEN                 |
| `ResponseTypeCodeIDTokenToken` | CODE_ID_TOKEN_TOKEN            |