# VciBatchIssueRequest


## Fields

| Field                                                                                      | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `AccessToken`                                                                              | `*string`                                                                                  | :heavy_minus_sign:                                                                         | The access token that came along with the credential request.                              |
| `Orders`                                                                                   | [][components.CredentialIssuanceOrder](../../models/components/credentialissuanceorder.md) | :heavy_minus_sign:                                                                         | The instructions for issuance of credentials and/or transaction IDs.                       |