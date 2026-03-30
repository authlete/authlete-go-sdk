# GrantManagementAction

The grant management action of the device authorization request.

The `grant_management_action` request parameter is defined in
[Grant Management for OAuth 2.0](https://openid.net/specs/fapi-grant-management.html).


## Example Usage

```go
import (
	"github.com/authlete/authlete-go-sdk/models/components"
)

value := components.GrantManagementActionCreate

// Open enum: custom values can be created with a direct type cast
custom := components.GrantManagementAction("custom_value")
```


## Values

| Name                           | Value                          |
| ------------------------------ | ------------------------------ |
| `GrantManagementActionCreate`  | CREATE                         |
| `GrantManagementActionQuery`   | QUERY                          |
| `GrantManagementActionReplace` | REPLACE                        |
| `GrantManagementActionRevoke`  | REVOKE                         |
| `GrantManagementActionMerge`   | MERGE                          |