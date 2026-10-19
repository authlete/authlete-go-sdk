# NativeSso

## Overview

### Available Operations

* [Process](#process) - Native SSO Processing
* [Logout](#logout) - Native SSO Logout Processing

## Process

This API should be called by the implementation of a token endpoint to generate the ID token and token response that comply with [OpenID Connect Native SSO for Mobile Apps 1.0](https://openid.net/specs/openid-connect-native-sso-1_0.html) (Native SSO) when Authlete’s `/auth/token` response indicates `action = NATIVE_SSO` (after you validate the session id and verify or generate the device secret as required by the flow). The token endpoint implementation should retrieve the value of `action` from the response and take the following steps according to the value.


### Example Usage

<!-- UsageSnippet language="go" operationID="native_sso_api" method="post" path="/api/{serviceId}/nativesso" -->
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
        authlete.WithServiceID("715948317"),
        authlete.WithSecurity(os.Getenv("AUTHLETE_BEARER")),
    )

    res, err := s.NativeSso.Process(ctx, operations.NativeSsoAPIRequest{
        Body: components.NativeSsoRequest{
            AccessToken: "_kh1aygxZ5NKLYKCJRM8M_AYvDg2wCWoprQDjfO87ZWq",
            RefreshToken: authlete.Pointer("kHUGSt_d3LSgiCQzH7wa5TpwIHWgjAZGw14zZV7hRqw"),
            Claims: authlete.Pointer("{\"given_name\":\"John\",\"family_name\":\"Brown\",\"email\":\"test@example.com\"}"),
            DeviceSecret: "my-ds",
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.NativeSsoResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                        | Type                                                                             | Required                                                                         | Description                                                                      |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `ctx`                                                                            | [context.Context](https://pkg.go.dev/context#Context)                            | :heavy_check_mark:                                                               | The context to use for the request.                                              |
| `request`                                                                        | [operations.NativeSsoAPIRequest](../../models/operations/nativessoapirequest.md) | :heavy_check_mark:                                                               | The request object to use for the request.                                       |
| `opts`                                                                           | [][operations.Option](../../models/operations/option.md)                         | :heavy_minus_sign:                                                               | The options for this request.                                                    |

### Response

**[*operations.NativeSsoAPIResponse](../../models/operations/nativessoapiresponse.md), error**

### Errors

| Error Type            | Status Code           | Content Type          |
| --------------------- | --------------------- | --------------------- |
| apierrors.ResultError | 400, 401, 403         | application/json      |
| apierrors.ResultError | 429                   | application/json      |
| apierrors.ResultError | 500                   | application/json      |
| apierrors.APIError    | 4XX, 5XX              | \*/\*                 |

## Logout

The `/nativesso/logout` API is supposed to be used to support the concept of "logout from all applications" in the context of [OpenID Connect Native SSO for Mobile Apps 1.0](https://openid.net/specs/openid-connect-native-sso-1_0.html) (Native SSO). This is accomplished by deleting access/refresh token records associated with the specified session ID. In Authlete's implementation, access/refresh token records can be associated with a session ID only through the mechanism introduced by Native SSO.


### Example Usage

<!-- UsageSnippet language="go" operationID="native_sso_logout_api" method="post" path="/api/{serviceId}/nativesso/logout" -->
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

    res, err := s.NativeSso.Logout(ctx, operations.NativeSsoLogoutAPIRequest{
        Body: components.NativeSsoLogoutRequest{
            SessionID: "my-sid",
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.NativeSsoLogoutResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [operations.NativeSsoLogoutAPIRequest](../../models/operations/nativessologoutapirequest.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |
| `opts`                                                                                       | [][operations.Option](../../models/operations/option.md)                                     | :heavy_minus_sign:                                                                           | The options for this request.                                                                |

### Response

**[*operations.NativeSsoLogoutAPIResponse](../../models/operations/nativessologoutapiresponse.md), error**

### Errors

| Error Type            | Status Code           | Content Type          |
| --------------------- | --------------------- | --------------------- |
| apierrors.ResultError | 400, 401, 403         | application/json      |
| apierrors.ResultError | 429                   | application/json      |
| apierrors.ResultError | 500                   | application/json      |
| apierrors.APIError    | 4XX, 5XX              | \*/\*                 |