# ClientSource

Source of this client record.


## Example Usage

```go
import (
	"github.com/authlete/authlete-go-sdk/models/components"
)

value := components.ClientSourceDynamicRegistration

// Open enum: custom values can be created with a direct type cast
custom := components.ClientSource("custom_value")
```


## Values

| Name                                | Value                               |
| ----------------------------------- | ----------------------------------- |
| `ClientSourceDynamicRegistration`   | DYNAMIC_REGISTRATION                |
| `ClientSourceAutomaticRegistration` | AUTOMATIC_REGISTRATION              |
| `ClientSourceExplicitRegistration`  | EXPLICIT_REGISTRATION               |
| `ClientSourceMetadataDocument`      | METADATA_DOCUMENT                   |
| `ClientSourceStaticRegistration`    | STATIC_REGISTRATION                 |