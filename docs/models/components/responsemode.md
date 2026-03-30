# ResponseMode

## Example Usage

```go
import (
	"github.com/authlete/authlete-go-sdk/models/components"
)

value := components.ResponseModeQuery

// Open enum: custom values can be created with a direct type cast
custom := components.ResponseMode("custom_value")
```


## Values

| Name                      | Value                     |
| ------------------------- | ------------------------- |
| `ResponseModeQuery`       | QUERY                     |
| `ResponseModeFragment`    | FRAGMENT                  |
| `ResponseModeFormPost`    | FORM_POST                 |
| `ResponseModeJwt`         | JWT                       |
| `ResponseModeQueryJwt`    | QUERY_JWT                 |
| `ResponseModeFragmentJwt` | FRAGMENT_JWT              |
| `ResponseModeFormPostJwt` | FORM_POST_JWT             |