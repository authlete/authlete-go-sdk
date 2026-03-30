# SubjectType

The subject type that the client application requests. Details about the subject type are described in
[OpenID Connect Core 1.0, 8. Subjct Identifier Types](https://openid.net/specs/openid-connect-core-1_0.html#SubjectIDTypes).

This property corresponds to `subject_type` in
[OpenID Connect Dynamic Client Registration 1.0, 2. Client Metadata](https://openid.net/specs/openid-connect-registration-1_0.html#ClientMetadata).


## Example Usage

```go
import (
	"github.com/authlete/authlete-go-sdk/models/components"
)

value := components.SubjectTypePublic

// Open enum: custom values can be created with a direct type cast
custom := components.SubjectType("custom_value")
```


## Values

| Name                  | Value                 |
| --------------------- | --------------------- |
| `SubjectTypePublic`   | PUBLIC                |
| `SubjectTypePairwise` | PAIRWISE              |