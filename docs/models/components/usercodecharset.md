# UserCodeCharset

The character set for end-user verification codes (`user_code`) for Device Flow.


## Example Usage

```go
import (
	"github.com/authlete/authlete-go-sdk/models/components"
)

value := components.UserCodeCharsetBase20

// Open enum: custom values can be created with a direct type cast
custom := components.UserCodeCharset("custom_value")
```


## Values

| Name                     | Value                    |
| ------------------------ | ------------------------ |
| `UserCodeCharsetBase20`  | BASE20                   |
| `UserCodeCharsetNumeric` | NUMERIC                  |