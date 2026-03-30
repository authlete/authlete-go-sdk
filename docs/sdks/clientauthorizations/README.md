# ClientAuthorizations

## Overview

### Available Operations

* [List](#list) - Get Authorized Applications

## List

Get a list of client applications that an end-user has authorized.

The subject parameter is required.


### Example Usage

<!-- UsageSnippet language="go" operationID="client_authorization_get_list_api_post" method="post" path="/api/{serviceId}/client/authorization/get/list" -->
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

    res, err := s.ClientAuthorizations.List(ctx, "<id>", components.ClientAuthorizationGetListRequest{
        Subject: "<value>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ClientAuthorizationGetListResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                    | Type                                                                                                         | Required                                                                                                     | Description                                                                                                  |
| ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                                        | :heavy_check_mark:                                                                                           | The context to use for the request.                                                                          |
| `serviceID`                                                                                                  | `string`                                                                                                     | :heavy_check_mark:                                                                                           | A service ID.                                                                                                |
| `body`                                                                                                       | [components.ClientAuthorizationGetListRequest](../../models/components/clientauthorizationgetlistrequest.md) | :heavy_check_mark:                                                                                           | N/A                                                                                                          |
| `opts`                                                                                                       | [][operations.Option](../../models/operations/option.md)                                                     | :heavy_minus_sign:                                                                                           | The options for this request.                                                                                |

### Response

**[*operations.ClientAuthorizationGetListAPIPostResponse](../../models/operations/clientauthorizationgetlistapipostresponse.md), error**

### Errors

| Error Type            | Status Code           | Content Type          |
| --------------------- | --------------------- | --------------------- |
| apierrors.ResultError | 400, 401, 403         | application/json      |
| apierrors.ResultError | 500                   | application/json      |
| apierrors.APIError    | 4XX, 5XX              | \*/\*                 |