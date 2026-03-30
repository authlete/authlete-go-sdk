# TokenFailRequest


## Fields

| Field                                                                                  | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `Ticket`                                                                               | `string`                                                                               | :heavy_check_mark:                                                                     | The ticket issued from Authlete `/auth/token` API.<br/>                                |
| `Reason`                                                                               | [components.TokenFailRequestReason](../../models/components/tokenfailrequestreason.md) | :heavy_check_mark:                                                                     | The reason of the failure of the token request.<br/>                                   |