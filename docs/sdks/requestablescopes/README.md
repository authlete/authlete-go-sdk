# Clients.RequestableScopes

## Overview

### Available Operations

* [Update](#update) - Update Requestable Scopes

## Update

Update requestable scopes of a client


### Example Usage

<!-- UsageSnippet language="go" operationID="client_extension_requestables_scopes_update_api_post" method="post" path="/api/{serviceId}/client/extension/requestable_scopes/update/{clientId}" -->
```go
package main

import(
	"context"
	"os"
	authlete "github.com/authlete/authlete-go-sdk"
	"github.com/authlete/authlete-go-sdk/models/components"
	"log"
)

func main() {
    ctx := context.Background()

    s := authlete.New(
        authlete.WithSecurity(os.Getenv("AUTHLETE_BEARER")),
    )

    res, err := s.Clients.RequestableScopes.Update(ctx, "<id>", "<id>", components.ClientExtensionRequestableScopesUpdateRequest{})
    if err != nil {
        log.Fatal(err)
    }
    if res.ClientExtensionRequestableScopesUpdateResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                            | Type                                                                                                                                 | Required                                                                                                                             | Description                                                                                                                          |
| ------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                                                | :heavy_check_mark:                                                                                                                   | The context to use for the request.                                                                                                  |
| `serviceID`                                                                                                                          | `string`                                                                                                                             | :heavy_check_mark:                                                                                                                   | A service ID.                                                                                                                        |
| `clientID`                                                                                                                           | `string`                                                                                                                             | :heavy_check_mark:                                                                                                                   | A client ID.<br/>                                                                                                                    |
| `body`                                                                                                                               | [components.ClientExtensionRequestableScopesUpdateRequest](../../models/components/clientextensionrequestablescopesupdaterequest.md) | :heavy_check_mark:                                                                                                                   | N/A                                                                                                                                  |
| `opts`                                                                                                                               | [][operations.Option](../../models/operations/option.md)                                                                             | :heavy_minus_sign:                                                                                                                   | The options for this request.                                                                                                        |

### Response

**[*operations.ClientExtensionRequestablesScopesUpdateAPIPostResponse](../../models/operations/clientextensionrequestablesscopesupdateapipostresponse.md), error**

### Errors

| Error Type            | Status Code           | Content Type          |
| --------------------- | --------------------- | --------------------- |
| apierrors.ResultError | 400, 401, 403         | application/json      |
| apierrors.ResultError | 500                   | application/json      |
| apierrors.APIError    | 4XX, 5XX              | \*/\*                 |