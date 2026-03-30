# ClientLimitedAuthorizationClientSource

Source of this client record.


## Example Usage

```go
import (
	"github.com/authlete/authlete-go-sdk/models/components"
)

value := components.ClientLimitedAuthorizationClientSourceDynamicRegistration

// Open enum: custom values can be created with a direct type cast
custom := components.ClientLimitedAuthorizationClientSource("custom_value")
```


## Values

| Name                                                          | Value                                                         |
| ------------------------------------------------------------- | ------------------------------------------------------------- |
| `ClientLimitedAuthorizationClientSourceDynamicRegistration`   | DYNAMIC_REGISTRATION                                          |
| `ClientLimitedAuthorizationClientSourceAutomaticRegistration` | AUTOMATIC_REGISTRATION                                        |
| `ClientLimitedAuthorizationClientSourceExplicitRegistration`  | EXPLICIT_REGISTRATION                                         |
| `ClientLimitedAuthorizationClientSourceMetadataDocument`      | METADATA_DOCUMENT                                             |
| `ClientLimitedAuthorizationClientSourceStaticRegistration`    | STATIC_REGISTRATION                                           |