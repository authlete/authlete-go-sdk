# JwsAlg

The signature algorithm for JWT. This value is represented on 'alg' attribute
of the header of JWT.

it's semantics depends upon where is this defined, for instance:
  - as service accessTokenSignAlg value, it defines that access token are JWT and the algorithm used to sign it. Check your [KB article](https://kb.authlete.com/en/s/oauth-and-openid-connect/a/jwt-based-access-token).
  - as client authorizationSignAlg value, it represents the signature algorithm used when [creating a JARM response](https://kb.authlete.com/en/s/oauth-and-openid-connect/a/enabling-jarm).
  - or as client requestSignAlg value, it specifies which is the expected signature used by [client on a Request Object](https://kb.authlete.com/en/s/oauth-and-openid-connect/a/request-objects).


## Example Usage

```go
import (
	"github.com/authlete/authlete-go-sdk/models/components"
)

value := components.JwsAlgNone

// Open enum: custom values can be created with a direct type cast
custom := components.JwsAlg("custom_value")
```


## Values

| Name           | Value          |
| -------------- | -------------- |
| `JwsAlgNone`   | NONE           |
| `JwsAlgHs256`  | HS256          |
| `JwsAlgHs384`  | HS384          |
| `JwsAlgHs512`  | HS512          |
| `JwsAlgRs256`  | RS256          |
| `JwsAlgRs384`  | RS384          |
| `JwsAlgRs512`  | RS512          |
| `JwsAlgEs256`  | ES256          |
| `JwsAlgEs384`  | ES384          |
| `JwsAlgEs512`  | ES512          |
| `JwsAlgPs256`  | PS256          |
| `JwsAlgPs384`  | PS384          |
| `JwsAlgPs512`  | PS512          |
| `JwsAlgEs256K` | ES256K         |
| `JwsAlgEdDsa`  | EdDSA          |