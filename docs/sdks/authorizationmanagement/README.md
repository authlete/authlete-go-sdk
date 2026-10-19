# Authorization.Management

## Overview

### Available Operations

* [GetTicketInfo](#getticketinfo) - Get Ticket Information
* [UpdateTicket](#updateticket) - Update Ticket Information

## GetTicketInfo

Get Ticket Information

### Example Usage

<!-- UsageSnippet language="go" operationID="getAuthorizationTicketInfo" method="post" path="/api/{serviceId}/auth/authorization/ticket/info" -->
```go
package main

import(
	"context"
	"os"
	authlete "github.com/authlete/authlete-go-sdk"
	"github.com/authlete/authlete-go-sdk/models/components"
	"github.com/authlete/authlete-go-sdk/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := authlete.New(
        authlete.WithServiceID("<id>"),
        authlete.WithSecurity(os.Getenv("AUTHLETE_BEARER")),
    )

    res, err := s.Authorization.Management.GetTicketInfo(ctx, operations.GetAuthorizationTicketInfoRequest{
        Body: components.AuthorizationTicketInfoRequest{
            Ticket: "<value>",
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AuthorizationTicketInfoResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                    | Type                                                                                                         | Required                                                                                                     | Description                                                                                                  |
| ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                                        | :heavy_check_mark:                                                                                           | The context to use for the request.                                                                          |
| `request`                                                                                                    | [operations.GetAuthorizationTicketInfoRequest](../../models/operations/getauthorizationticketinforequest.md) | :heavy_check_mark:                                                                                           | The request object to use for the request.                                                                   |
| `opts`                                                                                                       | [][operations.Option](../../models/operations/option.md)                                                     | :heavy_minus_sign:                                                                                           | The options for this request.                                                                                |

### Response

**[*operations.GetAuthorizationTicketInfoResponse](../../models/operations/getauthorizationticketinforesponse.md), error**

### Errors

| Error Type            | Status Code           | Content Type          |
| --------------------- | --------------------- | --------------------- |
| apierrors.ResultError | 400, 401, 403         | application/json      |
| apierrors.ResultError | 429                   | application/json      |
| apierrors.ResultError | 500                   | application/json      |
| apierrors.APIError    | 4XX, 5XX              | \*/\*                 |

## UpdateTicket

Update Ticket Information

### Example Usage

<!-- UsageSnippet language="go" operationID="updateAuthorizationTicket" method="post" path="/api/{serviceId}/auth/authorization/ticket/update" -->
```go
package main

import(
	"context"
	"os"
	authlete "github.com/authlete/authlete-go-sdk"
	"github.com/authlete/authlete-go-sdk/models/components"
	"github.com/authlete/authlete-go-sdk/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := authlete.New(
        authlete.WithServiceID("<id>"),
        authlete.WithSecurity(os.Getenv("AUTHLETE_BEARER")),
    )

    res, err := s.Authorization.Management.UpdateTicket(ctx, operations.UpdateAuthorizationTicketRequest{
        Body: components.AuthorizationTicketUpdateRequest{
            Ticket: "<value>",
            Info: components.AuthorizationTicketInfo{},
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AuthorizationTicketUpdateResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                  | Type                                                                                                       | Required                                                                                                   | Description                                                                                                |
| ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                      | :heavy_check_mark:                                                                                         | The context to use for the request.                                                                        |
| `request`                                                                                                  | [operations.UpdateAuthorizationTicketRequest](../../models/operations/updateauthorizationticketrequest.md) | :heavy_check_mark:                                                                                         | The request object to use for the request.                                                                 |
| `opts`                                                                                                     | [][operations.Option](../../models/operations/option.md)                                                   | :heavy_minus_sign:                                                                                         | The options for this request.                                                                              |

### Response

**[*operations.UpdateAuthorizationTicketResponse](../../models/operations/updateauthorizationticketresponse.md), error**

### Errors

| Error Type            | Status Code           | Content Type          |
| --------------------- | --------------------- | --------------------- |
| apierrors.ResultError | 400, 401, 403         | application/json      |
| apierrors.ResultError | 429                   | application/json      |
| apierrors.ResultError | 500                   | application/json      |
| apierrors.APIError    | 4XX, 5XX              | \*/\*                 |