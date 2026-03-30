# VciOfferCreateResponseAction

The result of the `/vci/offer/create` API call.

## Example Usage

```go
import (
	"github.com/authlete/authlete-go-sdk/models/components"
)

value := components.VciOfferCreateResponseActionCreated

// Open enum: custom values can be created with a direct type cast
custom := components.VciOfferCreateResponseAction("custom_value")
```


## Values

| Name                                        | Value                                       |
| ------------------------------------------- | ------------------------------------------- |
| `VciOfferCreateResponseActionCreated`       | CREATED                                     |
| `VciOfferCreateResponseActionForbidden`     | FORBIDDEN                                   |
| `VciOfferCreateResponseActionCallerError`   | CALLER_ERROR                                |
| `VciOfferCreateResponseActionAuthleteError` | AUTHLETE_ERROR                              |