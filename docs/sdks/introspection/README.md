# Introspection

## Overview

### Available Operations

* [Process](#process) - Process Introspection Request
* [StandardProcess](#standardprocess) - Process OAuth 2.0 Introspection Request

## Process

This API gathers information about an access token.


### Example Usage

<!-- UsageSnippet language="go" operationID="auth_introspection_api" method="post" path="/api/{serviceId}/auth/introspection" -->
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

    res, err := s.Introspection.Process(ctx, operations.AuthIntrospectionAPIRequest{
        Body: components.IntrospectionRequest{
            Token: "VFGsNK-5sXiqterdaR7b5QbRX9VTwVCQB87jbr2_xAI",
            Scopes: []string{
                "history.read",
                "timeline.read",
            },
            Subject: authlete.Pointer("john"),
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.IntrospectionResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                            | :heavy_check_mark:                                                                               | The context to use for the request.                                                              |
| `request`                                                                                        | [operations.AuthIntrospectionAPIRequest](../../models/operations/authintrospectionapirequest.md) | :heavy_check_mark:                                                                               | The request object to use for the request.                                                       |
| `opts`                                                                                           | [][operations.Option](../../models/operations/option.md)                                         | :heavy_minus_sign:                                                                               | The options for this request.                                                                    |

### Response

**[*operations.AuthIntrospectionAPIResponse](../../models/operations/authintrospectionapiresponse.md), error**

### Errors

| Error Type            | Status Code           | Content Type          |
| --------------------- | --------------------- | --------------------- |
| apierrors.ResultError | 400, 401, 403         | application/json      |
| apierrors.ResultError | 429                   | application/json      |
| apierrors.ResultError | 500                   | application/json      |
| apierrors.APIError    | 4XX, 5XX              | \*/\*                 |

## StandardProcess

This API exists to help your authorization server provide its own introspection API which complies with [RFC 7662](https://tools.ietf.org/html/rfc7662) (OAuth 2.0 Token Introspection).


### Example Usage

<!-- UsageSnippet language="go" operationID="auth_introspection_standard_api" method="post" path="/api/{serviceId}/auth/introspection/standard" -->
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

    res, err := s.Introspection.StandardProcess(ctx, operations.AuthIntrospectionStandardAPIRequest{
        Body: components.StandardIntrospectionRequest{
            Parameters: "token=VFGsNK-5sXiqterdaR7b5QbRX9VTwVCQB87jbr2_xAI&token_type_hint=access_token",
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.StandardIntrospectionResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                        | Type                                                                                                             | Required                                                                                                         | Description                                                                                                      |
| ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                                            | :heavy_check_mark:                                                                                               | The context to use for the request.                                                                              |
| `request`                                                                                                        | [operations.AuthIntrospectionStandardAPIRequest](../../models/operations/authintrospectionstandardapirequest.md) | :heavy_check_mark:                                                                                               | The request object to use for the request.                                                                       |
| `opts`                                                                                                           | [][operations.Option](../../models/operations/option.md)                                                         | :heavy_minus_sign:                                                                                               | The options for this request.                                                                                    |

### Response

**[*operations.AuthIntrospectionStandardAPIResponse](../../models/operations/authintrospectionstandardapiresponse.md), error**

### Errors

| Error Type            | Status Code           | Content Type          |
| --------------------- | --------------------- | --------------------- |
| apierrors.ResultError | 400, 401, 403         | application/json      |
| apierrors.ResultError | 429                   | application/json      |
| apierrors.ResultError | 500                   | application/json      |
| apierrors.APIError    | 4XX, 5XX              | \*/\*                 |