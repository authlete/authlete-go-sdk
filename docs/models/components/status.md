# Status

Whether the action completed successfully or failed.

## Example Usage

```go
import (
	"github.com/authlete/authlete-go-sdk/models/components"
)

value := components.StatusCompleted

// Open enum: custom values can be created with a direct type cast
custom := components.Status("custom_value")
```


## Values

| Name              | Value             |
| ----------------- | ----------------- |
| `StatusCompleted` | completed         |
| `StatusFailed`    | failed            |