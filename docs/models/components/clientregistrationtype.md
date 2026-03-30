# ClientRegistrationType

Values for the `client_registration_types` RP metadata and the
 `client_registration_types_supported` OP metadata that are defined in
 [OpenID Connect Federation 1.0](https://openid.net/specs/openid-connect-federation-1_0.html).


## Example Usage

```go
import (
	"github.com/authlete/authlete-go-sdk/models/components"
)

value := components.ClientRegistrationTypeAutomatic

// Open enum: custom values can be created with a direct type cast
custom := components.ClientRegistrationType("custom_value")
```


## Values

| Name                              | Value                             |
| --------------------------------- | --------------------------------- |
| `ClientRegistrationTypeAutomatic` | AUTOMATIC                         |
| `ClientRegistrationTypeExplicit`  | EXPLICIT                          |