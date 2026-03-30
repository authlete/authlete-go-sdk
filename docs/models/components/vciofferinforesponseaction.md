# VciOfferInfoResponseAction

The result of the `/vci/offer/info` API call.

## Example Usage

```go
import (
	"github.com/authlete/authlete-go-sdk/models/components"
)

value := components.VciOfferInfoResponseActionOk

// Open enum: custom values can be created with a direct type cast
custom := components.VciOfferInfoResponseAction("custom_value")
```


## Values

| Name                                      | Value                                     |
| ----------------------------------------- | ----------------------------------------- |
| `VciOfferInfoResponseActionOk`            | OK                                        |
| `VciOfferInfoResponseActionForbidden`     | FORBIDDEN                                 |
| `VciOfferInfoResponseActionNotFound`      | NOT_FOUND                                 |
| `VciOfferInfoResponseActionCallerError`   | CALLER_ERROR                              |
| `VciOfferInfoResponseActionAuthleteError` | AUTHLETE_ERROR                            |