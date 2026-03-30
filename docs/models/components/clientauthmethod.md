# ClientAuthMethod

The client authentication method that the client application declares that it uses at the token
endpoint. This property corresponds to `token_endpoint_auth_method` in [OpenID Connect Dynamic
Client Registration 1.0, 2. Client Metadata](https://openid.net/specs/openid-connect-registration-1_0.html#ClientMetadata).


## Example Usage

```go
import (
	"github.com/authlete/authlete-go-sdk/models/components"
)

value := components.ClientAuthMethodNone

// Open enum: custom values can be created with a direct type cast
custom := components.ClientAuthMethod("custom_value")
```


## Values

| Name                                      | Value                                     |
| ----------------------------------------- | ----------------------------------------- |
| `ClientAuthMethodNone`                    | NONE                                      |
| `ClientAuthMethodClientSecretBasic`       | CLIENT_SECRET_BASIC                       |
| `ClientAuthMethodClientSecretPost`        | CLIENT_SECRET_POST                        |
| `ClientAuthMethodClientSecretJwt`         | CLIENT_SECRET_JWT                         |
| `ClientAuthMethodPrivateKeyJwt`           | PRIVATE_KEY_JWT                           |
| `ClientAuthMethodTLSClientAuth`           | TLS_CLIENT_AUTH                           |
| `ClientAuthMethodSelfSignedTLSClientAuth` | SELF_SIGNED_TLS_CLIENT_AUTH               |
| `ClientAuthMethodAttestJwtClientAuth`     | ATTEST_JWT_CLIENT_AUTH                    |