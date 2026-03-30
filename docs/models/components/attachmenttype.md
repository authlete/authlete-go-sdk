# AttachmentType

Supported attachment types. This property corresponds to the `attachments_supported`
 server metadata which was added by the third implementer's draft of OpenID Connect
 for Identity Assurance 1.0.


## Example Usage

```go
import (
	"github.com/authlete/authlete-go-sdk/models/components"
)

value := components.AttachmentTypeEmbedded

// Open enum: custom values can be created with a direct type cast
custom := components.AttachmentType("custom_value")
```


## Values

| Name                     | Value                    |
| ------------------------ | ------------------------ |
| `AttachmentTypeEmbedded` | EMBEDDED                 |
| `AttachmentTypeExternal` | EXTERNAL                 |