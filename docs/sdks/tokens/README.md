# Clients.Tokens

## Overview

### Available Operations

* [Delete](#delete) - Delete Client Tokens

## Delete

Delete all existing access tokens issued to a client application by an end-user.

The subject parameter is required and must be provided as a query parameter.


### Example Usage

<!-- UsageSnippet language="go" operationID="client_authorization_delete_api" method="delete" path="/api/{serviceId}/client/authorization/delete/{clientId}" -->
```go
package main

import(
	"context"
	"os"
	authlete "github.com/authlete/authlete-go-sdk"
	"log"
)

func main() {
    ctx := context.Background()

    s := authlete.New(
        authlete.WithSecurity(os.Getenv("AUTHLETE_BEARER")),
    )

    res, err := s.Clients.Tokens.Delete(ctx, "<id>", "<id>", "<value>")
    if err != nil {
        log.Fatal(err)
    }
    if res.ClientAuthorizationDeleteResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `serviceID`                                              | `string`                                                 | :heavy_check_mark:                                       | A service ID.                                            |
| `clientID`                                               | `string`                                                 | :heavy_check_mark:                                       | A client ID.<br/>                                        |
| `subject`                                                | `string`                                                 | :heavy_check_mark:                                       | Unique user ID of an end-user.<br/>                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.ClientAuthorizationDeleteAPIResponse](../../models/operations/clientauthorizationdeleteapiresponse.md), error**

### Errors

| Error Type            | Status Code           | Content Type          |
| --------------------- | --------------------- | --------------------- |
| apierrors.ResultError | 400, 401, 403         | application/json      |
| apierrors.ResultError | 500                   | application/json      |
| apierrors.APIError    | 4XX, 5XX              | \*/\*                 |