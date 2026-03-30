# TokenGetListResponse


## Fields

| Field                                                                 | Type                                                                  | Required                                                              | Description                                                           |
| --------------------------------------------------------------------- | --------------------------------------------------------------------- | --------------------------------------------------------------------- | --------------------------------------------------------------------- |
| `Start`                                                               | `*int`                                                                | :heavy_minus_sign:                                                    | Start index of search results (inclusive).<br/>                       |
| `End`                                                                 | `*int`                                                                | :heavy_minus_sign:                                                    | End index of search results (exclusive).<br/>                         |
| `TotalCount`                                                          | `*int`                                                                | :heavy_minus_sign:                                                    | Unique ID of a client developer.<br/>                                 |
| `Client`                                                              | [*components.ClientLimited](../../models/components/clientlimited.md) | :heavy_minus_sign:                                                    | N/A                                                                   |
| `Subject`                                                             | `*string`                                                             | :heavy_minus_sign:                                                    | Unique user ID of an end-user.<br/>                                   |
| `AccessTokens`                                                        | [][components.AccessToken](../../models/components/accesstoken.md)    | :heavy_minus_sign:                                                    | An array of access tokens.<br/>                                       |