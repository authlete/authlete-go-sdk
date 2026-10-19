# BackChannelLogout

## Overview

### Available Operations

* [BackchannelLogoutTokenAPI](#backchannellogouttokenapi) - Backchannel Logout Token Issuing

## BackchannelLogoutTokenAPI

The `/backchannel/logout/token` API issues a logout token for a client application in the context of [OpenID Connect Back-Channel Logout 1.0](https://openid.net/specs/openid-connect-backchannel-1_0.html).


### Example Usage

<!-- UsageSnippet language="go" operationID="backchannel_logout_token_api" method="post" path="/api/{serviceId}/backchannel/logout/token" -->
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

    res, err := s.BackChannelLogout.BackchannelLogoutTokenAPI(ctx, operations.BackchannelLogoutTokenAPIRequest{
        Body: components.BackchannelLogoutTokenRequest{
            ClientIdentifier: "1140735077",
            Subject: authlete.Pointer("user123"),
            SessionID: authlete.Pointer("my-sid"),
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.BackchannelLogoutTokenResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                  | Type                                                                                                       | Required                                                                                                   | Description                                                                                                |
| ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                      | :heavy_check_mark:                                                                                         | The context to use for the request.                                                                        |
| `request`                                                                                                  | [operations.BackchannelLogoutTokenAPIRequest](../../models/operations/backchannellogouttokenapirequest.md) | :heavy_check_mark:                                                                                         | The request object to use for the request.                                                                 |
| `opts`                                                                                                     | [][operations.Option](../../models/operations/option.md)                                                   | :heavy_minus_sign:                                                                                         | The options for this request.                                                                              |

### Response

**[*operations.BackchannelLogoutTokenAPIResponse](../../models/operations/backchannellogouttokenapiresponse.md), error**

### Errors

| Error Type            | Status Code           | Content Type          |
| --------------------- | --------------------- | --------------------- |
| apierrors.ResultError | 400, 401, 403         | application/json      |
| apierrors.ResultError | 429                   | application/json      |
| apierrors.ResultError | 500                   | application/json      |
| apierrors.APIError    | 4XX, 5XX              | \*/\*                 |