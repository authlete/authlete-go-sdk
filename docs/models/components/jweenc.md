# JweEnc

This is the encryption algorithm to be used when encrypting a JWT on client or server side.
Depending upon the context, this refers to encryption done by the client or by the server. For instance:
  - as `authorizationEncryptionEnc` value, it refers to the encryption algorithm used by server when creating a JARM response
  - as `requestEncryptionEnc` value, it refers to the expected encryption algorithm used by the client when encrypting a Request Object
  - as `idTokenEncryptionEnc` value, it refers to the algorithm used by the server to encrypt id_tokens


## Example Usage

```go
import (
	"github.com/authlete/authlete-go-sdk/models/components"
)

value := components.JweEncA128CbcHs256

// Open enum: custom values can be created with a direct type cast
custom := components.JweEnc("custom_value")
```


## Values

| Name                 | Value                |
| -------------------- | -------------------- |
| `JweEncA128CbcHs256` | A128CBC_HS256        |
| `JweEncA192CbcHs384` | A192CBC_HS384        |
| `JweEncA256CbcHs512` | A256CBC_HS512        |
| `JweEncA128Gcm`      | A128GCM              |
| `JweEncA192Gcm`      | A192GCM              |
| `JweEncA256Gcm`      | A256GCM              |