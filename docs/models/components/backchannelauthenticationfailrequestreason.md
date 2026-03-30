# BackchannelAuthenticationFailRequestReason

The reason of the failure of the backchannel authentication request. This request parameter is
not mandatory but optional. However, giving this parameter is recommended. If omitted, `SERVER_ERROR`
is used as a reason.


## Example Usage

```go
import (
	"github.com/authlete/authlete-go-sdk/models/components"
)

value := components.BackchannelAuthenticationFailRequestReasonAccessDenied
```


## Values

| Name                                                              | Value                                                             |
| ----------------------------------------------------------------- | ----------------------------------------------------------------- |
| `BackchannelAuthenticationFailRequestReasonAccessDenied`          | ACCESS_DENIED                                                     |
| `BackchannelAuthenticationFailRequestReasonExpiredLoginHintToken` | EXPIRED_LOGIN_HINT_TOKEN                                          |
| `BackchannelAuthenticationFailRequestReasonInvalidBindingMessage` | INVALID_BINDING_MESSAGE                                           |
| `BackchannelAuthenticationFailRequestReasonInvalidTarget`         | INVALID_TARGET                                                    |
| `BackchannelAuthenticationFailRequestReasonInvalidUserCode`       | INVALID_USER_CODE                                                 |
| `BackchannelAuthenticationFailRequestReasonMissingUserCode`       | MISSING_USER_CODE                                                 |
| `BackchannelAuthenticationFailRequestReasonServerError`           | SERVER_ERROR                                                      |
| `BackchannelAuthenticationFailRequestReasonUnauthorizedClient`    | UNAUTHORIZED_CLIENT                                               |
| `BackchannelAuthenticationFailRequestReasonUnknownUserID`         | UNKNOWN_USER_ID                                                   |