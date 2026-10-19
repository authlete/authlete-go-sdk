# IdpError

Error response returned by the Authlete IdP server. Unlike the main API's `resultCode`/`resultMessage` format, IdP errors carry a human-readable `error` message, optionally accompanied by contextual fields (such as `organizationId` or `apiServerId`). Request validation failures instead return an `errors` array of per-field messages.



## Fields

| Field                                                                         | Type                                                                          | Required                                                                      | Description                                                                   |
| ----------------------------------------------------------------------------- | ----------------------------------------------------------------------------- | ----------------------------------------------------------------------------- | ----------------------------------------------------------------------------- |
| `Error`                                                                       | `*string`                                                                     | :heavy_minus_sign:                                                            | A human-readable error message.                                               |
| `Errors`                                                                      | []`string`                                                                    | :heavy_minus_sign:                                                            | Per-field validation error messages, present for request validation failures. |
| `AdditionalProperties`                                                        | map[string]`any`                                                              | :heavy_minus_sign:                                                            | N/A                                                                           |
| `HTTPMeta`                                                                    | [components.HTTPMetadata](../../models/components/httpmetadata.md)            | :heavy_check_mark:                                                            | N/A                                                                           |