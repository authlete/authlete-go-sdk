# ClientType

The client type, either `CONFIDENTIAL` or `PUBLIC`. See [RFC 6749, 2.1. Client Types](https://datatracker.ietf.org/doc/html/rfc6749#section-2.1)
for details.


## Example Usage

```go
import (
	"github.com/authlete/authlete-go-sdk/models/components"
)

value := components.ClientTypePublic

// Open enum: custom values can be created with a direct type cast
custom := components.ClientType("custom_value")
```


## Values

| Name                     | Value                    |
| ------------------------ | ------------------------ |
| `ClientTypePublic`       | PUBLIC                   |
| `ClientTypeConfidential` | CONFIDENTIAL             |