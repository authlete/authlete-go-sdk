# Authorization

## Overview

### Available Operations

* [ProcessRequest](#processrequest) - Process Authorization Request
* [Fail](#fail) - Fail Authorization Request
* [Issue](#issue) - Issue Authorization Response

## ProcessRequest

This API parses request parameters of an authorization request and returns necessary data for the authorization server
implementation to process the authorization request further.


### Example Usage

<!-- UsageSnippet language="go" operationID="auth_authorization_api" method="post" path="/api/{serviceId}/auth/authorization" -->
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

    res, err := s.Authorization.ProcessRequest(ctx, "<id>", components.AuthorizationRequest{
        Parameters: "response_type=code&client_id=26478243745571&redirect_uri=https%3A%2F%2Fmy-client.example.com%2Fcb1&scope=timeline.read+history.read&code_challenge=E9Melhoa2OwvFrEMTJguCHaoeK1t8URWbuGJSstw-cM&code_challenge_method=S256",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AuthorizationResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `ctx`                                                                              | [context.Context](https://pkg.go.dev/context#Context)                              | :heavy_check_mark:                                                                 | The context to use for the request.                                                |
| `serviceID`                                                                        | `string`                                                                           | :heavy_check_mark:                                                                 | A service ID.                                                                      |
| `body`                                                                             | [components.AuthorizationRequest](../../models/components/authorizationrequest.md) | :heavy_check_mark:                                                                 | N/A                                                                                |
| `opts`                                                                             | [][operations.Option](../../models/operations/option.md)                           | :heavy_minus_sign:                                                                 | The options for this request.                                                      |

### Response

**[*operations.AuthAuthorizationAPIResponse](../../models/operations/authauthorizationapiresponse.md), error**

### Errors

| Error Type            | Status Code           | Content Type          |
| --------------------- | --------------------- | --------------------- |
| apierrors.ResultError | 400, 401, 403         | application/json      |
| apierrors.ResultError | 500                   | application/json      |
| apierrors.APIError    | 4XX, 5XX              | \*/\*                 |

## Fail

This API generates a content of an error authorization response that the authorization server implementation
returns to the client application.


### Example Usage

<!-- UsageSnippet language="go" operationID="auth_authorization_fail_api" method="post" path="/api/{serviceId}/auth/authorization/fail" -->
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

    res, err := s.Authorization.Fail(ctx, "<id>", components.AuthorizationFailRequest{
        Ticket: "qA7wGybwArICpbUSutrf5Xc9-i1fHE0ySOHxR1eBoBQ",
        Reason: components.AuthorizationFailRequestReasonNotAuthenticated,
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AuthorizationFailResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `serviceID`                                                                                | `string`                                                                                   | :heavy_check_mark:                                                                         | A service ID.                                                                              |
| `body`                                                                                     | [components.AuthorizationFailRequest](../../models/components/authorizationfailrequest.md) | :heavy_check_mark:                                                                         | N/A                                                                                        |
| `opts`                                                                                     | [][operations.Option](../../models/operations/option.md)                                   | :heavy_minus_sign:                                                                         | The options for this request.                                                              |

### Response

**[*operations.AuthAuthorizationFailAPIResponse](../../models/operations/authauthorizationfailapiresponse.md), error**

### Errors

| Error Type            | Status Code           | Content Type          |
| --------------------- | --------------------- | --------------------- |
| apierrors.ResultError | 400, 401, 403         | application/json      |
| apierrors.ResultError | 500                   | application/json      |
| apierrors.APIError    | 4XX, 5XX              | \*/\*                 |

## Issue

This API parses request parameters of an authorization request and returns necessary data for the
authorization server implementation to process the authorization request further.


### Example Usage

<!-- UsageSnippet language="go" operationID="auth_authorization_issue_api" method="post" path="/api/{serviceId}/auth/authorization/issue" -->
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

    res, err := s.Authorization.Issue(ctx, "<id>", components.AuthorizationIssueRequest{
        Ticket: "FFgB9gwb_WXh6g1u-UQ8ZI-d_k4B-o-cm7RkVzI8Vnc",
        Subject: "john",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AuthorizationIssueResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `serviceID`                                                                                  | `string`                                                                                     | :heavy_check_mark:                                                                           | A service ID.                                                                                |
| `body`                                                                                       | [components.AuthorizationIssueRequest](../../models/components/authorizationissuerequest.md) | :heavy_check_mark:                                                                           | N/A                                                                                          |
| `opts`                                                                                       | [][operations.Option](../../models/operations/option.md)                                     | :heavy_minus_sign:                                                                           | The options for this request.                                                                |

### Response

**[*operations.AuthAuthorizationIssueAPIResponse](../../models/operations/authauthorizationissueapiresponse.md), error**

### Errors

| Error Type            | Status Code           | Content Type          |
| --------------------- | --------------------- | --------------------- |
| apierrors.ResultError | 400, 401, 403         | application/json      |
| apierrors.ResultError | 500                   | application/json      |
| apierrors.APIError    | 4XX, 5XX              | \*/\*                 |