# AuthTokenGetListAPIRequest


## Fields

| Field                                                              | Type                                                               | Required                                                           | Description                                                        |
| ------------------------------------------------------------------ | ------------------------------------------------------------------ | ------------------------------------------------------------------ | ------------------------------------------------------------------ |
| `ServiceID`                                                        | `string`                                                           | :heavy_check_mark:                                                 | A service ID.                                                      |
| `ClientIdentifier`                                                 | `*string`                                                          | :heavy_minus_sign:                                                 | Client Identifier (client ID or client ID alias).<br/>             |
| `Subject`                                                          | `*string`                                                          | :heavy_minus_sign:                                                 | Unique user ID.<br/>                                               |
| `Start`                                                            | `*int`                                                             | :heavy_minus_sign:                                                 | Start index of search results (inclusive). The default value is 0. |
| `End`                                                              | `*int`                                                             | :heavy_minus_sign:                                                 | End index of search results (exclusive). The default value is 5.<br/> |