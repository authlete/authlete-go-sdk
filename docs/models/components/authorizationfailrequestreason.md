# AuthorizationFailRequestReason

The reason of the failure of the authorization request.
For more details, see [NO_INTERACTION] in the description of `/auth/authorization` API.


## Example Usage

```go
import (
	"github.com/authlete/authlete-go-sdk/models/components"
)

value := components.AuthorizationFailRequestReasonUnknown
```


## Values

| Name                                                     | Value                                                    |
| -------------------------------------------------------- | -------------------------------------------------------- |
| `AuthorizationFailRequestReasonUnknown`                  | UNKNOWN                                                  |
| `AuthorizationFailRequestReasonNotLoggedIn`              | NOT_LOGGED_IN                                            |
| `AuthorizationFailRequestReasonMaxAgeNotSupported`       | MAX_AGE_NOT_SUPPORTED                                    |
| `AuthorizationFailRequestReasonExceedsMaxAge`            | EXCEEDS_MAX_AGE                                          |
| `AuthorizationFailRequestReasonDifferentSubject`         | DIFFERENT_SUBJECT                                        |
| `AuthorizationFailRequestReasonAcrNotSatisfied`          | ACR_NOT_SATISFIED                                        |
| `AuthorizationFailRequestReasonDenied`                   | DENIED                                                   |
| `AuthorizationFailRequestReasonServerError`              | SERVER_ERROR                                             |
| `AuthorizationFailRequestReasonNotAuthenticated`         | NOT_AUTHENTICATED                                        |
| `AuthorizationFailRequestReasonAccountSelectionRequired` | ACCOUNT_SELECTION_REQUIRED                               |
| `AuthorizationFailRequestReasonConsentRequired`          | CONSENT_REQUIRED                                         |
| `AuthorizationFailRequestReasonInteractionRequired`      | INTERACTION_REQUIRED                                     |
| `AuthorizationFailRequestReasonInvalidTarget`            | INVALID_TARGET                                           |