# Userinfo

## Overview

### Available Operations

* [Process](#process) - Process UserInfo Request
* [Issue](#issue) - Issue UserInfo Response

## Process

This API gathers information about a user.


### Example Usage

<!-- UsageSnippet language="go" operationID="auth_userinfo_api" method="post" path="/api/{serviceId}/auth/userinfo" -->
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

    res, err := s.Userinfo.Process(ctx, "<id>", components.UserinfoRequest{
        Token: "Ntm9MDb8WXQAevqrBkd84KTTHbYHVQrTjgUZCOWqEUI",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.UserinfoResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                | Type                                                                     | Required                                                                 | Description                                                              |
| ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ |
| `ctx`                                                                    | [context.Context](https://pkg.go.dev/context#Context)                    | :heavy_check_mark:                                                       | The context to use for the request.                                      |
| `serviceID`                                                              | `string`                                                                 | :heavy_check_mark:                                                       | A service ID.                                                            |
| `body`                                                                   | [components.UserinfoRequest](../../models/components/userinforequest.md) | :heavy_check_mark:                                                       | N/A                                                                      |
| `opts`                                                                   | [][operations.Option](../../models/operations/option.md)                 | :heavy_minus_sign:                                                       | The options for this request.                                            |

### Response

**[*operations.AuthUserinfoAPIResponse](../../models/operations/authuserinfoapiresponse.md), error**

### Errors

| Error Type            | Status Code           | Content Type          |
| --------------------- | --------------------- | --------------------- |
| apierrors.ResultError | 400, 401, 403         | application/json      |
| apierrors.ResultError | 500                   | application/json      |
| apierrors.APIError    | 4XX, 5XX              | \*/\*                 |

## Issue

This API generates an ID token.


### Example Usage

<!-- UsageSnippet language="go" operationID="auth_userinfo_issue_api" method="post" path="/api/{serviceId}/auth/userinfo/issue" -->
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

    res, err := s.Userinfo.Issue(ctx, "<id>", components.UserinfoIssueRequest{
        Token: "Ntm9MDb8WXQAevqrBkd84KTTHbYHVQrTjgUZCOWqEUI",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.UserinfoIssueResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `ctx`                                                                              | [context.Context](https://pkg.go.dev/context#Context)                              | :heavy_check_mark:                                                                 | The context to use for the request.                                                |
| `serviceID`                                                                        | `string`                                                                           | :heavy_check_mark:                                                                 | A service ID.                                                                      |
| `body`                                                                             | [components.UserinfoIssueRequest](../../models/components/userinfoissuerequest.md) | :heavy_check_mark:                                                                 | N/A                                                                                |
| `opts`                                                                             | [][operations.Option](../../models/operations/option.md)                           | :heavy_minus_sign:                                                                 | The options for this request.                                                      |

### Response

**[*operations.AuthUserinfoIssueAPIResponse](../../models/operations/authuserinfoissueapiresponse.md), error**

### Errors

| Error Type            | Status Code           | Content Type          |
| --------------------- | --------------------- | --------------------- |
| apierrors.ResultError | 400, 401, 403         | application/json      |
| apierrors.ResultError | 500                   | application/json      |
| apierrors.APIError    | 4XX, 5XX              | \*/\*                 |