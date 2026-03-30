# Token

## Overview

### Available Operations

* [Process](#process) - Process Token Request
* [Fail](#fail) - Fail Token Request
* [Issue](#issue) - Issue Token Response

## Process

This API parses request parameters of an authorization request and returns necessary data for the
authorization server implementation to process the authorization request further.


### Example Usage

<!-- UsageSnippet language="go" operationID="auth_token_api" method="post" path="/api/{serviceId}/auth/token" -->
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

    res, err := s.Token.Process(ctx, "<id>", components.TokenRequest{
        Parameters: "grant_type=authorization_code&code=Xv_su944auuBgc5mfUnxXayiiQU9Z4-T_Yae_UfExmo&redirect_uri=https%3A%2F%2Fmy-client.example.com%2Fcb1&code_verifier=dBjftJeZ4CVP-mB92K27uhbUJU1p1r_wW1gFWFOEjXk",
        ClientID: authlete.Pointer("26478243745571"),
        ClientSecret: authlete.Pointer("gXz97ISgLs4HuXwOZWch8GEmgL4YMvUJwu3er_kDVVGcA0UOhA9avLPbEmoeZdagi9yC_-tEiT2BdRyH9dbrQQ"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.TokenResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                          | Type                                                               | Required                                                           | Description                                                        |
| ------------------------------------------------------------------ | ------------------------------------------------------------------ | ------------------------------------------------------------------ | ------------------------------------------------------------------ |
| `ctx`                                                              | [context.Context](https://pkg.go.dev/context#Context)              | :heavy_check_mark:                                                 | The context to use for the request.                                |
| `serviceID`                                                        | `string`                                                           | :heavy_check_mark:                                                 | A service ID.                                                      |
| `body`                                                             | [components.TokenRequest](../../models/components/tokenrequest.md) | :heavy_check_mark:                                                 | N/A                                                                |
| `opts`                                                             | [][operations.Option](../../models/operations/option.md)           | :heavy_minus_sign:                                                 | The options for this request.                                      |

### Response

**[*operations.AuthTokenAPIResponse](../../models/operations/authtokenapiresponse.md), error**

### Errors

| Error Type            | Status Code           | Content Type          |
| --------------------- | --------------------- | --------------------- |
| apierrors.ResultError | 400, 401, 403         | application/json      |
| apierrors.ResultError | 500                   | application/json      |
| apierrors.APIError    | 4XX, 5XX              | \*/\*                 |

## Fail

This API generates a content of an error token response that the authorization server implementation
returns to the client application.


### Example Usage

<!-- UsageSnippet language="go" operationID="auth_token_fail_api" method="post" path="/api/{serviceId}/auth/token/fail" -->
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

    res, err := s.Token.Fail(ctx, "<id>", components.TokenFailRequest{
        Ticket: "83BNqKIhGMyrkvop_7jQjv2Z1612LNdGSQKkvkrf47c",
        Reason: components.TokenFailRequestReasonInvalidResourceOwnerCredentials,
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.TokenFailResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                  | Type                                                                       | Required                                                                   | Description                                                                |
| -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- |
| `ctx`                                                                      | [context.Context](https://pkg.go.dev/context#Context)                      | :heavy_check_mark:                                                         | The context to use for the request.                                        |
| `serviceID`                                                                | `string`                                                                   | :heavy_check_mark:                                                         | A service ID.                                                              |
| `body`                                                                     | [components.TokenFailRequest](../../models/components/tokenfailrequest.md) | :heavy_check_mark:                                                         | N/A                                                                        |
| `opts`                                                                     | [][operations.Option](../../models/operations/option.md)                   | :heavy_minus_sign:                                                         | The options for this request.                                              |

### Response

**[*operations.AuthTokenFailAPIResponse](../../models/operations/authtokenfailapiresponse.md), error**

### Errors

| Error Type            | Status Code           | Content Type          |
| --------------------- | --------------------- | --------------------- |
| apierrors.ResultError | 400, 401, 403         | application/json      |
| apierrors.ResultError | 500                   | application/json      |
| apierrors.APIError    | 4XX, 5XX              | \*/\*                 |

## Issue

This API generates a content of a successful token response that the authorization server implementation
returns to the client application.


### Example Usage

<!-- UsageSnippet language="go" operationID="auth_token_issue_api" method="post" path="/api/{serviceId}/auth/token/issue" -->
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

    res, err := s.Token.Issue(ctx, "<id>", components.TokenIssueRequest{
        Ticket: "p7SXQ9JFjng7KFOZdCMBKcoR3ift7B54l1LGIgQXqEM",
        Subject: "john",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.TokenIssueResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                    | Type                                                                         | Required                                                                     | Description                                                                  |
| ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- |
| `ctx`                                                                        | [context.Context](https://pkg.go.dev/context#Context)                        | :heavy_check_mark:                                                           | The context to use for the request.                                          |
| `serviceID`                                                                  | `string`                                                                     | :heavy_check_mark:                                                           | A service ID.                                                                |
| `body`                                                                       | [components.TokenIssueRequest](../../models/components/tokenissuerequest.md) | :heavy_check_mark:                                                           | N/A                                                                          |
| `opts`                                                                       | [][operations.Option](../../models/operations/option.md)                     | :heavy_minus_sign:                                                           | The options for this request.                                                |

### Response

**[*operations.AuthTokenIssueAPIResponse](../../models/operations/authtokenissueapiresponse.md), error**

### Errors

| Error Type            | Status Code           | Content Type          |
| --------------------- | --------------------- | --------------------- |
| apierrors.ResultError | 400, 401, 403         | application/json      |
| apierrors.ResultError | 500                   | application/json      |
| apierrors.APIError    | 4XX, 5XX              | \*/\*                 |