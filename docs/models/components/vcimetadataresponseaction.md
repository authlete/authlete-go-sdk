# VciMetadataResponseAction

The next action that the implementation of the credential issuer
metadata endpoint (`/.well-known/openid-credential-issuer`)
should take after getting a response from Authlete's
`/vci/metadata` API.


## Example Usage

```go
import (
	"github.com/authlete/authlete-go-sdk/models/components"
)

value := components.VciMetadataResponseActionOk

// Open enum: custom values can be created with a direct type cast
custom := components.VciMetadataResponseAction("custom_value")
```


## Values

| Name                                           | Value                                          |
| ---------------------------------------------- | ---------------------------------------------- |
| `VciMetadataResponseActionOk`                  | OK                                             |
| `VciMetadataResponseActionNotFound`            | NOT_FOUND                                      |
| `VciMetadataResponseActionInternalServerError` | INTERNAL_SERVER_ERROR                          |