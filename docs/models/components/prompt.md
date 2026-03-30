# Prompt

The prompt that the UI displayed to the end-user must satisfy as the minimum level. This value comes from `prompt` request parameter.

When the authorization request does not contain `prompt` request parameter, `CONSENT` is used as the default value.

See "[OpenID Connect Core 1.0, 3.1.2.1. Authentication Request](https://openid.net/specs/openid-connect-core-1_0.html#AuthRequest), prompt" for `prompt` request parameter.


## Example Usage

```go
import (
	"github.com/authlete/authlete-go-sdk/models/components"
)

value := components.PromptNone

// Open enum: custom values can be created with a direct type cast
custom := components.Prompt("custom_value")
```


## Values

| Name                  | Value                 |
| --------------------- | --------------------- |
| `PromptNone`          | NONE                  |
| `PromptLogin`         | LOGIN                 |
| `PromptConsent`       | CONSENT               |
| `PromptSelectAccount` | SELECT_ACCOUNT        |
| `PromptCreate`        | CREATE                |