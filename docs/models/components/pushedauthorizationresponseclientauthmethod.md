# PushedAuthorizationResponseClientAuthMethod

The client authentication method that the client application declares that it uses at the token
endpoint. This property corresponds to `token_endpoint_auth_method` in [OpenID Connect Dynamic
Client Registration 1.0, 2. Client Metadata](https://openid.net/specs/openid-connect-registration-1_0.html#ClientMetadata).


## Example Usage

```go
import (
	"github.com/authlete/authlete-go-sdk/models/components"
)

value := components.PushedAuthorizationResponseClientAuthMethodNone

// Open enum: custom values can be created with a direct type cast
custom := components.PushedAuthorizationResponseClientAuthMethod("custom_value")
```


## Values

| Name                                                                 | Value                                                                |
| -------------------------------------------------------------------- | -------------------------------------------------------------------- |
| `PushedAuthorizationResponseClientAuthMethodNone`                    | NONE                                                                 |
| `PushedAuthorizationResponseClientAuthMethodClientSecretBasic`       | CLIENT_SECRET_BASIC                                                  |
| `PushedAuthorizationResponseClientAuthMethodClientSecretPost`        | CLIENT_SECRET_POST                                                   |
| `PushedAuthorizationResponseClientAuthMethodClientSecretJwt`         | CLIENT_SECRET_JWT                                                    |
| `PushedAuthorizationResponseClientAuthMethodPrivateKeyJwt`           | PRIVATE_KEY_JWT                                                      |
| `PushedAuthorizationResponseClientAuthMethodTLSClientAuth`           | TLS_CLIENT_AUTH                                                      |
| `PushedAuthorizationResponseClientAuthMethodSelfSignedTLSClientAuth` | SELF_SIGNED_TLS_CLIENT_AUTH                                          |
| `PushedAuthorizationResponseClientAuthMethodAttestJwtClientAuth`     | ATTEST_JWT_CLIENT_AUTH                                               |