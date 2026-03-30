# DeliveryMode

## Example Usage

```go
import (
	"github.com/authlete/authlete-go-sdk/models/components"
)

value := components.DeliveryModePing

// Open enum: custom values can be created with a direct type cast
custom := components.DeliveryMode("custom_value")
```


## Values

| Name               | Value              |
| ------------------ | ------------------ |
| `DeliveryModePing` | PING               |
| `DeliveryModePoll` | POLL               |
| `DeliveryModePush` | PUSH               |