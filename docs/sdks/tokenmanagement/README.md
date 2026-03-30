# Token.Management

## Overview

### Available Operations

* [ReissueIDToken](#reissueidtoken) - Reissue ID Token
* [List](#list) - List Issued Tokens
* [Create](#create) - Create Access Token
* [Update](#update) - Update Access Token
* [Delete](#delete) - Delete Access Token
* [Revoke](#revoke) - Revoke Access Token

## ReissueIDToken

The API is expected to be called only when the value of the `action`
parameter in a response from the `/auth/token` API is [ID_TOKEN_REISSUABLE](https://authlete.github.io/authlete-java-common/com/authlete/common/dto/TokenResponse.Action.html#ID_TOKEN_REISSUABLE). The purpose
of the `/idtoken/reissue` API is to generate a token response that
includes a new ID token together with a new access token and a refresh
token.


### Example Usage

<!-- UsageSnippet language="go" operationID="idtoken_reissue_api" method="post" path="/api/{serviceId}/idtoken/reissue" -->
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

    res, err := s.Token.Management.ReissueIDToken(ctx, "<id>", nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.IdtokenReissueResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                             | Type                                                                                  | Required                                                                              | Description                                                                           |
| ------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------- |
| `ctx`                                                                                 | [context.Context](https://pkg.go.dev/context#Context)                                 | :heavy_check_mark:                                                                    | The context to use for the request.                                                   |
| `serviceID`                                                                           | `string`                                                                              | :heavy_check_mark:                                                                    | A service ID.                                                                         |
| `body`                                                                                | [*components.IdtokenReissueRequest](../../models/components/idtokenreissuerequest.md) | :heavy_minus_sign:                                                                    | N/A                                                                                   |
| `opts`                                                                                | [][operations.Option](../../models/operations/option.md)                              | :heavy_minus_sign:                                                                    | The options for this request.                                                         |

### Response

**[*operations.IdtokenReissueAPIResponse](../../models/operations/idtokenreissueapiresponse.md), error**

### Errors

| Error Type            | Status Code           | Content Type          |
| --------------------- | --------------------- | --------------------- |
| apierrors.ResultError | 400, 401, 403         | application/json      |
| apierrors.ResultError | 500                   | application/json      |
| apierrors.APIError    | 4XX, 5XX              | \*/\*                 |

## List

Get the list of access tokens that are associated with the service.


### Example Usage

<!-- UsageSnippet language="go" operationID="auth_token_get_list_api" method="get" path="/api/{serviceId}/auth/token/get/list" -->
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

    res, err := s.Token.Management.List(ctx, operations.AuthTokenGetListAPIRequest{
        ServiceID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.TokenGetListResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [operations.AuthTokenGetListAPIRequest](../../models/operations/authtokengetlistapirequest.md) | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |
| `opts`                                                                                         | [][operations.Option](../../models/operations/option.md)                                       | :heavy_minus_sign:                                                                             | The options for this request.                                                                  |

### Response

**[*operations.AuthTokenGetListAPIResponse](../../models/operations/authtokengetlistapiresponse.md), error**

### Errors

| Error Type            | Status Code           | Content Type          |
| --------------------- | --------------------- | --------------------- |
| apierrors.ResultError | 400, 401, 403         | application/json      |
| apierrors.ResultError | 500                   | application/json      |
| apierrors.APIError    | 4XX, 5XX              | \*/\*                 |

## Create

Create an access token.


### Example Usage

<!-- UsageSnippet language="go" operationID="auth_token_create_api" method="post" path="/api/{serviceId}/auth/token/create" -->
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

    res, err := s.Token.Management.Create(ctx, "<id>", components.TokenCreateRequest{
        GrantType: components.GrantTypeAuthorizationCode,
        ClientID: authlete.Pointer[int64](26888344961664),
        Subject: authlete.Pointer("john"),
        Scopes: []string{
            "history.read",
            "timeline.read",
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.TokenCreateResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                      | Type                                                                           | Required                                                                       | Description                                                                    |
| ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ |
| `ctx`                                                                          | [context.Context](https://pkg.go.dev/context#Context)                          | :heavy_check_mark:                                                             | The context to use for the request.                                            |
| `serviceID`                                                                    | `string`                                                                       | :heavy_check_mark:                                                             | A service ID.                                                                  |
| `body`                                                                         | [components.TokenCreateRequest](../../models/components/tokencreaterequest.md) | :heavy_check_mark:                                                             | N/A                                                                            |
| `opts`                                                                         | [][operations.Option](../../models/operations/option.md)                       | :heavy_minus_sign:                                                             | The options for this request.                                                  |

### Response

**[*operations.AuthTokenCreateAPIResponse](../../models/operations/authtokencreateapiresponse.md), error**

### Errors

| Error Type            | Status Code           | Content Type          |
| --------------------- | --------------------- | --------------------- |
| apierrors.ResultError | 400, 401, 403         | application/json      |
| apierrors.ResultError | 500                   | application/json      |
| apierrors.APIError    | 4XX, 5XX              | \*/\*                 |

## Update

Update an access token.


### Example Usage

<!-- UsageSnippet language="go" operationID="auth_token_update_api" method="post" path="/api/{serviceId}/auth/token/update" -->
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

    res, err := s.Token.Management.Update(ctx, "<id>", components.TokenUpdateRequest{
        AccessToken: authlete.Pointer("Z5a40U6dWvw2gMoCOAFbZcM85q4HC0Z--0YKD9-Nf6Q"),
        Scopes: []string{
            "history.read",
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.TokenUpdateResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                      | Type                                                                           | Required                                                                       | Description                                                                    |
| ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ |
| `ctx`                                                                          | [context.Context](https://pkg.go.dev/context#Context)                          | :heavy_check_mark:                                                             | The context to use for the request.                                            |
| `serviceID`                                                                    | `string`                                                                       | :heavy_check_mark:                                                             | A service ID.                                                                  |
| `body`                                                                         | [components.TokenUpdateRequest](../../models/components/tokenupdaterequest.md) | :heavy_check_mark:                                                             | N/A                                                                            |
| `opts`                                                                         | [][operations.Option](../../models/operations/option.md)                       | :heavy_minus_sign:                                                             | The options for this request.                                                  |

### Response

**[*operations.AuthTokenUpdateAPIResponse](../../models/operations/authtokenupdateapiresponse.md), error**

### Errors

| Error Type            | Status Code           | Content Type          |
| --------------------- | --------------------- | --------------------- |
| apierrors.ResultError | 400, 401, 403         | application/json      |
| apierrors.ResultError | 500                   | application/json      |
| apierrors.APIError    | 4XX, 5XX              | \*/\*                 |

## Delete

Delete an access token.


### Example Usage

<!-- UsageSnippet language="go" operationID="auth_token_delete_api" method="delete" path="/api/{serviceId}/auth/token/delete/{accessTokenIdentifier}" -->
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

    res, err := s.Token.Management.Delete(ctx, "<id>", "<value>")
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                  | Type                                                                                                                                       | Required                                                                                                                                   | Description                                                                                                                                |
| ------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                                                      | :heavy_check_mark:                                                                                                                         | The context to use for the request.                                                                                                        |
| `serviceID`                                                                                                                                | `string`                                                                                                                                   | :heavy_check_mark:                                                                                                                         | A service ID.                                                                                                                              |
| `accessTokenIdentifier`                                                                                                                    | `string`                                                                                                                                   | :heavy_check_mark:                                                                                                                         | The identifier of an existing access token. The identifier is the value of the access token<br/>or the value of the hash of the access token.<br/> |
| `opts`                                                                                                                                     | [][operations.Option](../../models/operations/option.md)                                                                                   | :heavy_minus_sign:                                                                                                                         | The options for this request.                                                                                                              |

### Response

**[*operations.AuthTokenDeleteAPIResponse](../../models/operations/authtokendeleteapiresponse.md), error**

### Errors

| Error Type            | Status Code           | Content Type          |
| --------------------- | --------------------- | --------------------- |
| apierrors.ResultError | 400, 401, 403         | application/json      |
| apierrors.ResultError | 500                   | application/json      |
| apierrors.APIError    | 4XX, 5XX              | \*/\*                 |

## Revoke

Revoke an access token.


### Example Usage

<!-- UsageSnippet language="go" operationID="auth_token_revoke_api" method="post" path="/api/{serviceId}/auth/token/revoke" -->
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

    res, err := s.Token.Management.Revoke(ctx, "<id>", components.TokenRevokeRequest{
        AccessTokenIdentifier: authlete.Pointer("Z5a40U6dWvw2gMoCOAFbZcM85q4HC0Z--0YKD9-Nf6Q"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.TokenRevokeResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                      | Type                                                                           | Required                                                                       | Description                                                                    |
| ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ |
| `ctx`                                                                          | [context.Context](https://pkg.go.dev/context#Context)                          | :heavy_check_mark:                                                             | The context to use for the request.                                            |
| `serviceID`                                                                    | `string`                                                                       | :heavy_check_mark:                                                             | A service ID.                                                                  |
| `body`                                                                         | [components.TokenRevokeRequest](../../models/components/tokenrevokerequest.md) | :heavy_check_mark:                                                             | N/A                                                                            |
| `opts`                                                                         | [][operations.Option](../../models/operations/option.md)                       | :heavy_minus_sign:                                                             | The options for this request.                                                  |

### Response

**[*operations.AuthTokenRevokeAPIResponse](../../models/operations/authtokenrevokeapiresponse.md), error**

### Errors

| Error Type            | Status Code           | Content Type          |
| --------------------- | --------------------- | --------------------- |
| apierrors.ResultError | 400, 401, 403, 404    | application/json      |
| apierrors.ResultError | 500                   | application/json      |
| apierrors.APIError    | 4XX, 5XX              | \*/\*                 |