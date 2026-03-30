# VerifiedClaimsValidationSchema

The verified claims validation schema set.


## Example Usage

```go
import (
	"github.com/authlete/authlete-go-sdk/models/components"
)

value := components.VerifiedClaimsValidationSchemaStandard

// Open enum: custom values can be created with a direct type cast
custom := components.VerifiedClaimsValidationSchema("custom_value")
```


## Values

| Name                                                   | Value                                                  |
| ------------------------------------------------------ | ------------------------------------------------------ |
| `VerifiedClaimsValidationSchemaStandard`               | standard                                               |
| `VerifiedClaimsValidationSchemaStandardPlusIDDocument` | standard+id_document                                   |