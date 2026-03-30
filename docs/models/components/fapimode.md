# FapiMode

## Example Usage

```go
import (
	"github.com/authlete/authlete-go-sdk/models/components"
)

value := components.FapiModeFapi1Advanced

// Open enum: custom values can be created with a direct type cast
custom := components.FapiMode("custom_value")
```


## Values

| Name                                          | Value                                         |
| --------------------------------------------- | --------------------------------------------- |
| `FapiModeFapi1Advanced`                       | FAPI1_ADVANCED                                |
| `FapiModeFapi1Baseline`                       | FAPI1_BASELINE                                |
| `FapiModeFapi2MessageSigningAuthReq`          | FAPI2_MESSAGE_SIGNING_AUTH_REQ                |
| `FapiModeFapi2MessageSigningAuthRes`          | FAPI2_MESSAGE_SIGNING_AUTH_RES                |
| `FapiModeFapi2MessageSigningIntrospectionRes` | FAPI2_MESSAGE_SIGNING_INTROSPECTION_RES       |
| `FapiModeFapi2Security`                       | FAPI2_SECURITY                                |