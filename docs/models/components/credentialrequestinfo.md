# CredentialRequestInfo


## Fields

| Field                                                               | Type                                                                | Required                                                            | Description                                                         |
| ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- |
| `Identifier`                                                        | `*string`                                                           | :heavy_minus_sign:                                                  | The identifier of the credential offer.                             |
| `Format`                                                            | `*string`                                                           | :heavy_minus_sign:                                                  | The value of the format parameter in the credential request.        |
| `BindingKey`                                                        | `*string`                                                           | :heavy_minus_sign:                                                  | The binding key specified by the proof in the credential request.   |
| `BindingKeys`                                                       | []`string`                                                          | :heavy_minus_sign:                                                  | The binding keys specified by the proofs in the credential request. |
| `Details`                                                           | `*string`                                                           | :heavy_minus_sign:                                                  | The details about the credential request.                           |