# ClientClientSource

Source of this client record.


## Example Usage

```go
import (
	"github.com/authlete/authlete-go-sdk/models/components"
)

value := components.ClientClientSourceDynamicRegistration

// Open enum: custom values can be created with a direct type cast
custom := components.ClientClientSource("custom_value")
```


## Values

| Name                                      | Value                                     |
| ----------------------------------------- | ----------------------------------------- |
| `ClientClientSourceDynamicRegistration`   | DYNAMIC_REGISTRATION                      |
| `ClientClientSourceAutomaticRegistration` | AUTOMATIC_REGISTRATION                    |
| `ClientClientSourceExplicitRegistration`  | EXPLICIT_REGISTRATION                     |
| `ClientClientSourceMetadataDocument`      | METADATA_DOCUMENT                         |
| `ClientClientSourceStaticRegistration`    | STATIC_REGISTRATION                       |