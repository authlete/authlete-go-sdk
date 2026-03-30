# ClaimType

## Example Usage

```go
import (
	"github.com/authlete/authlete-go-sdk/models/components"
)

value := components.ClaimTypeNormal

// Open enum: custom values can be created with a direct type cast
custom := components.ClaimType("custom_value")
```


## Values

| Name                   | Value                  |
| ---------------------- | ---------------------- |
| `ClaimTypeNormal`      | NORMAL                 |
| `ClaimTypeAggregated`  | AGGREGATED             |
| `ClaimTypeDistributed` | DISTRIBUTED            |