# JweAlg

this is the 'alg' header value for encrypted JWT tokens.
Depending upon the context, this refers to key transport scheme to be used by the client and by the server. For instance:
- as `authorizationEncryptionAlg` value, it refers to the encoding algorithm used by server for transporting they keys on JARM objects
- as `requestEncryptionAlg` value, it refers to the expected key transport encoding algorithm that server expect from client when encrypting a Request Object
- as `idTokenEncryptionAlg` value, it refers to the algorithm used by the server to key transport of id_tokens

**Please note that some of the algorithms are more secure than others, some are not supported very well cross platforms and some (like RSA1_5) is known to be weak**.


## Example Usage

```go
import (
	"github.com/authlete/authlete-go-sdk/models/components"
)

value := components.JweAlgRsa15

// Open enum: custom values can be created with a direct type cast
custom := components.JweAlg("custom_value")
```


## Values

| Name                     | Value                    |
| ------------------------ | ------------------------ |
| `JweAlgRsa15`            | RSA1_5                   |
| `JweAlgRsaOaep`          | RSA_OAEP                 |
| `JweAlgRsaOaep256`       | RSA_OAEP_256             |
| `JweAlgA128Kw`           | A128KW                   |
| `JweAlgA192Kw`           | A192KW                   |
| `JweAlgA256Kw`           | A256KW                   |
| `JweAlgDir`              | DIR                      |
| `JweAlgEcdhEs`           | ECDH_ES                  |
| `JweAlgEcdhEsA128Kw`     | ECDH_ES_A128KW           |
| `JweAlgEcdhEsA192Kw`     | ECDH_ES_A192KW           |
| `JweAlgEcdhEsA256Kw`     | ECDH_ES_A256KW           |
| `JweAlgA128Gcmkw`        | A128GCMKW                |
| `JweAlgA192Gcmkw`        | A192GCMKW                |
| `JweAlgA256Gcmkw`        | A256GCMKW                |
| `JweAlgPbes2Hs256A128Kw` | PBES2_HS256_A128KW       |
| `JweAlgPbes2Hs384A192Kw` | PBES2_HS384_A192KW       |
| `JweAlgPbes2Hs512A256Kw` | PBES2_HS512_A256KW       |