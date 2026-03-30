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
	"log"
)

func main() {
    ctx := context.Background()

    s := authlete.New(
        authlete.WithSecurity(os.Getenv("AUTHLETE_BEARER")),
    )

    res, err := s.Authorization.Management.GetTicketInfo(ctx, "<id>", components.AuthorizationTicketInfoRequest{
        Ticket: "<value>",
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

| Parameter                                                                                              | Type                                                                                                   | Required                                                                                               | Description                                                                                            |
| ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                  | :heavy_check_mark:                                                                                     | The context to use for the request.                                                                    |
| `serviceID`                                                                                            | `string`                                                                                               | :heavy_check_mark:                                                                                     | A service ID.                                                                                          |
| `body`                                                                                                 | [components.AuthorizationTicketInfoRequest](../../models/components/authorizationticketinforequest.md) | :heavy_check_mark:                                                                                     | N/A                                                                                                    |
| `opts`                                                                                                 | [][operations.Option](../../models/operations/option.md)                                               | :heavy_minus_sign:                                                                                     | The options for this request.                                                                          |

### Response

**[*operations.GetAuthorizationTicketInfoResponse](../../models/operations/getauthorizationticketinforesponse.md), error**

### Errors

| Error Type            | Status Code           | Content Type          |
| --------------------- | --------------------- | --------------------- |
| apierrors.ResultError | 400, 401, 403         | application/json      |
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
	"log"
)

func main() {
    ctx := context.Background()

    s := authlete.New(
        authlete.WithSecurity(os.Getenv("AUTHLETE_BEARER")),
    )

    res, err := s.Authorization.Management.UpdateTicket(ctx, "<id>", components.AuthorizationTicketUpdateRequest{
        Ticket: "<value>",
        Info: "<value>",
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
| `serviceID`                                                                                                | `string`                                                                                                   | :heavy_check_mark:                                                                                         | A service ID.                                                                                              |
| `body`                                                                                                     | [components.AuthorizationTicketUpdateRequest](../../models/components/authorizationticketupdaterequest.md) | :heavy_check_mark:                                                                                         | N/A                                                                                                        |
| `opts`                                                                                                     | [][operations.Option](../../models/operations/option.md)                                                   | :heavy_minus_sign:                                                                                         | The options for this request.                                                                              |

### Response

**[*operations.UpdateAuthorizationTicketResponse](../../models/operations/updateauthorizationticketresponse.md), error**

### Errors

| Error Type            | Status Code           | Content Type          |
| --------------------- | --------------------- | --------------------- |
| apierrors.ResultError | 400, 401, 403         | application/json      |
| apierrors.ResultError | 500                   | application/json      |
| apierrors.APIError    | 4XX, 5XX              | \*/\*                 |