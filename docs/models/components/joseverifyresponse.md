# JoseVerifyResponse


## Fields

| Field                                                      | Type                                                       | Required                                                   | Description                                                |
| ---------------------------------------------------------- | ---------------------------------------------------------- | ---------------------------------------------------------- | ---------------------------------------------------------- |
| `ResultCode`                                               | `*string`                                                  | :heavy_minus_sign:                                         | The code which represents the result of the API call.      |
| `ResultMessage`                                            | `*string`                                                  | :heavy_minus_sign:                                         | A short message which explains the result of the API call. |
| `Valid`                                                    | `*bool`                                                    | :heavy_minus_sign:                                         | The result of the verification on the JOSE object.<br/>    |
| `SignatureValid`                                           | `*bool`                                                    | :heavy_minus_sign:                                         | The result of the signature verification.<br/>             |
| `MissingClaims`                                            | []`string`                                                 | :heavy_minus_sign:                                         | The list of missing claims.<br/>                           |
| `InvalidClaims`                                            | []`string`                                                 | :heavy_minus_sign:                                         | The list of invalid claims.<br/>                           |
| `ErrorDescriptions`                                        | []`string`                                                 | :heavy_minus_sign:                                         | The list of error messages.<br/>                           |