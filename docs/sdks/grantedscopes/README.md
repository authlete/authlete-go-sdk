# Clients.GrantedScopes

## Overview

### Available Operations

* [Get](#get) - Get Granted Scopes
* [Create](#create) - Get Granted Scopes

## Get

Get the set of scopes that a user has granted to a client application.


### Example Usage

<!-- UsageSnippet language="go" operationID="client_granted_scopes_get_api" method="get" path="/api/{serviceId}/client/granted_scopes/get/{clientId}" -->
```go
package main

import(
	"context"
	"os"
	authlete "github.com/authlete/authlete-go-sdk"
	"log"
)

func main() {
    ctx := context.Background()

    s := authlete.New(
        authlete.WithSecurity(os.Getenv("AUTHLETE_BEARER")),
    )

    res, err := s.Clients.GrantedScopes.Get(ctx, "715948317", "1140735077", "<value>")
    if err != nil {
        log.Fatal(err)
    }
    if res.ClientAuthorizationDeleteResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              | Example                                                  |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |                                                          |
| `serviceID`                                              | `string`                                                 | :heavy_check_mark:                                       | A service ID.                                            | 715948317                                                |
| `clientID`                                               | `string`                                                 | :heavy_check_mark:                                       | A client ID.<br/>                                        | 1140735077                                               |
| `subject`                                                | `string`                                                 | :heavy_check_mark:                                       | Unique user ID of an end-user.<br/>                      |                                                          |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.ClientGrantedScopesGetAPIResponse](../../models/operations/clientgrantedscopesgetapiresponse.md), error**

### Errors

| Error Type            | Status Code           | Content Type          |
| --------------------- | --------------------- | --------------------- |
| apierrors.ResultError | 400, 401, 403         | application/json      |
| apierrors.ResultError | 500                   | application/json      |
| apierrors.APIError    | 4XX, 5XX              | \*/\*                 |

## Create

Get the set of scopes that a user has granted to a client application.

The subject parameter is required.


### Example Usage

<!-- UsageSnippet language="go" operationID="client_granted_scopes_get_api_post" method="post" path="/api/{serviceId}/client/granted_scopes/get/{clientId}" -->
```go
package main

import(
	"context"
	"os"
	authlete "github.com/authlete/authlete-go-sdk"
	"github.com/authlete/authlete-go-sdk/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := authlete.New(
        authlete.WithSecurity(os.Getenv("AUTHLETE_BEARER")),
    )

    res, err := s.Clients.GrantedScopes.Create(ctx, "<id>", "<id>", operations.ClientGrantedScopesGetAPIPostRequestBody{
        Subject: "<value>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ClientAuthorizationDeleteResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                  | Type                                                                                                                       | Required                                                                                                                   | Description                                                                                                                |
| -------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                                      | :heavy_check_mark:                                                                                                         | The context to use for the request.                                                                                        |
| `serviceID`                                                                                                                | `string`                                                                                                                   | :heavy_check_mark:                                                                                                         | A service ID.                                                                                                              |
| `clientID`                                                                                                                 | `string`                                                                                                                   | :heavy_check_mark:                                                                                                         | A client ID.<br/>                                                                                                          |
| `body`                                                                                                                     | [operations.ClientGrantedScopesGetAPIPostRequestBody](../../models/operations/clientgrantedscopesgetapipostrequestbody.md) | :heavy_check_mark:                                                                                                         | N/A                                                                                                                        |
| `opts`                                                                                                                     | [][operations.Option](../../models/operations/option.md)                                                                   | :heavy_minus_sign:                                                                                                         | The options for this request.                                                                                              |

### Response

**[*operations.ClientGrantedScopesGetAPIPostResponse](../../models/operations/clientgrantedscopesgetapipostresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |