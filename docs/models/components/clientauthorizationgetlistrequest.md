# ClientAuthorizationGetListRequest


## Fields

| Field                                      | Type                                       | Required                                   | Description                                |
| ------------------------------------------ | ------------------------------------------ | ------------------------------------------ | ------------------------------------------ |
| `Subject`                                  | `string`                                   | :heavy_check_mark:                         | Unique user ID of an end-user.             |
| `Developer`                                | `*string`                                  | :heavy_minus_sign:                         | Unique ID of a client developer.           |
| `Start`                                    | `*int`                                     | :heavy_minus_sign:                         | Start index of search results (inclusive). |
| `End`                                      | `*int`                                     | :heavy_minus_sign:                         | End index of search results (exclusive).   |