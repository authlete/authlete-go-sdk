# ClientAuthorizationGetListBySubjectAPIRequest


## Fields

| Field                                                              | Type                                                               | Required                                                           | Description                                                        |
| ------------------------------------------------------------------ | ------------------------------------------------------------------ | ------------------------------------------------------------------ | ------------------------------------------------------------------ |
| `ServiceID`                                                        | `string`                                                           | :heavy_check_mark:                                                 | A service ID.                                                      |
| `Subject`                                                          | `string`                                                           | :heavy_check_mark:                                                 | Unique user ID of an end-user.<br/>                                |
| `Developer`                                                        | `*string`                                                          | :heavy_minus_sign:                                                 | Unique ID of a client developer.<br/>                              |
| `Start`                                                            | `*int`                                                             | :heavy_minus_sign:                                                 | Start index of search results (inclusive). The default value is 0. |
| `End`                                                              | `*int`                                                             | :heavy_minus_sign:                                                 | End index of search results (exclusive). The default value is 5.<br/> |