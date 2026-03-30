# PushedAuthorization

## Overview

### Available Operations

* [Create](#create) - Process Pushed Authorization Request

## Create

This API creates a pushed request authorization. It authenticates the client and creates a authorization_uri to be returned by the authorization server.


### Example Usage

<!-- UsageSnippet language="go" operationID="pushed_auth_req_api" method="post" path="/api/{serviceId}/pushed_auth_req" -->
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

    res, err := s.PushedAuthorization.Create(ctx, "<id>", components.PushedAuthorizationRequest{
        Parameters: "response_type=code%20id_token&client_id=5921531358155430&redirect_uri=https%3A%2F%2Fserver.example.com%2Fcb&state=SOME_VALUE_ABLE_TO_PREVENT_CSRF&scope=openid&nonce=SOME_VALUE_ABLE_TO_PREVENT_REPLAY_ATTACK&code_challenge=5ZWDQJiryK3eaLtSeFV8y1XySMCWtyITxICLaTwvK8g&code_challenge_method=S256",
        ClientID: authlete.Pointer("5921531358155430"),
        ClientSecret: authlete.Pointer("P_FouxWlI7zcOep_9vBwR9qMAVJQiCiUiK1HrAP4GziOyezHQpqY0f5dHXK4JT4tnvI51OkbWVoEM9GnOyJViA"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.PushedAuthorizationResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `serviceID`                                                                                    | `string`                                                                                       | :heavy_check_mark:                                                                             | A service ID.                                                                                  |
| `body`                                                                                         | [components.PushedAuthorizationRequest](../../models/components/pushedauthorizationrequest.md) | :heavy_check_mark:                                                                             | N/A                                                                                            |
| `opts`                                                                                         | [][operations.Option](../../models/operations/option.md)                                       | :heavy_minus_sign:                                                                             | The options for this request.                                                                  |

### Response

**[*operations.PushedAuthReqAPIResponse](../../models/operations/pushedauthreqapiresponse.md), error**

### Errors

| Error Type            | Status Code           | Content Type          |
| --------------------- | --------------------- | --------------------- |
| apierrors.ResultError | 400, 401, 403         | application/json      |
| apierrors.ResultError | 500                   | application/json      |
| apierrors.APIError    | 4XX, 5XX              | \*/\*                 |