# ClientManagement

## Overview

### Available Operations

* [ListAuthorizedApplications](#listauthorizedapplications) - Get Authorized Applications
* [RevokeClientTokens](#revokeclienttokens) - Delete Client Tokens
* [RevokeGrantedScopes](#revokegrantedscopes) - Delete Granted Scopes

## ListAuthorizedApplications

Get a list of client applications that an end-user has authorized.

The subject parameter is required and can be provided as a query parameter.


### Example Usage

<!-- UsageSnippet language="go" operationID="client_authorization_get_list_api" method="get" path="/api/{serviceId}/client/authorization/get/list" -->
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

    res, err := s.ClientManagement.ListAuthorizedApplications(ctx, operations.ClientAuthorizationGetListAPIRequest{
        ServiceID: "<id>",
        Subject: "<value>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ClientAuthorizationGetListResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                          | Type                                                                                                               | Required                                                                                                           | Description                                                                                                        |
| ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                                              | :heavy_check_mark:                                                                                                 | The context to use for the request.                                                                                |
| `request`                                                                                                          | [operations.ClientAuthorizationGetListAPIRequest](../../models/operations/clientauthorizationgetlistapirequest.md) | :heavy_check_mark:                                                                                                 | The request object to use for the request.                                                                         |
| `opts`                                                                                                             | [][operations.Option](../../models/operations/option.md)                                                           | :heavy_minus_sign:                                                                                                 | The options for this request.                                                                                      |

### Response

**[*operations.ClientAuthorizationGetListAPIResponse](../../models/operations/clientauthorizationgetlistapiresponse.md), error**

### Errors

| Error Type            | Status Code           | Content Type          |
| --------------------- | --------------------- | --------------------- |
| apierrors.ResultError | 400, 401, 403         | application/json      |
| apierrors.ResultError | 500                   | application/json      |
| apierrors.APIError    | 4XX, 5XX              | \*/\*                 |

## RevokeClientTokens

Delete all existing access tokens issued to a client application by an end-user.

The subject parameter is required.


### Example Usage

<!-- UsageSnippet language="go" operationID="client_authorization_delete_api_post" method="post" path="/api/{serviceId}/client/authorization/delete/{clientId}" -->
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

    res, err := s.ClientManagement.RevokeClientTokens(ctx, "<id>", "<id>", operations.ClientAuthorizationDeleteAPIPostRequestBody{
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

| Parameter                                                                                                                        | Type                                                                                                                             | Required                                                                                                                         | Description                                                                                                                      |
| -------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                                                            | :heavy_check_mark:                                                                                                               | The context to use for the request.                                                                                              |
| `serviceID`                                                                                                                      | `string`                                                                                                                         | :heavy_check_mark:                                                                                                               | A service ID.                                                                                                                    |
| `clientID`                                                                                                                       | `string`                                                                                                                         | :heavy_check_mark:                                                                                                               | A client ID.<br/>                                                                                                                |
| `body`                                                                                                                           | [operations.ClientAuthorizationDeleteAPIPostRequestBody](../../models/operations/clientauthorizationdeleteapipostrequestbody.md) | :heavy_check_mark:                                                                                                               | N/A                                                                                                                              |
| `opts`                                                                                                                           | [][operations.Option](../../models/operations/option.md)                                                                         | :heavy_minus_sign:                                                                                                               | The options for this request.                                                                                                    |

### Response

**[*operations.ClientAuthorizationDeleteAPIPostResponse](../../models/operations/clientauthorizationdeleteapipostresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## RevokeGrantedScopes

Delete the set of scopes that an end-user has granted to a client application.

Even if records about granted scopes are deleted by calling this API, existing access tokens are
not deleted and scopes of existing access tokens are not changed.
The subject parameter is required and must be provided as a query parameter.


### Example Usage

<!-- UsageSnippet language="go" operationID="client_granted_scopes_delete_api" method="delete" path="/api/{serviceId}/client/granted_scopes/delete/{clientId}" -->
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

    res, err := s.ClientManagement.RevokeGrantedScopes(ctx, "<id>", "<id>", "<value>")
    if err != nil {
        log.Fatal(err)
    }
    if res.ClientGrantedScopesDeleteResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `serviceID`                                              | `string`                                                 | :heavy_check_mark:                                       | A service ID.                                            |
| `clientID`                                               | `string`                                                 | :heavy_check_mark:                                       | A client ID.<br/>                                        |
| `subject`                                                | `string`                                                 | :heavy_check_mark:                                       | Unique user ID of an end-user.<br/>                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.ClientGrantedScopesDeleteAPIResponse](../../models/operations/clientgrantedscopesdeleteapiresponse.md), error**

### Errors

| Error Type            | Status Code           | Content Type          |
| --------------------- | --------------------- | --------------------- |
| apierrors.ResultError | 400, 401, 403         | application/json      |
| apierrors.ResultError | 500                   | application/json      |
| apierrors.APIError    | 4XX, 5XX              | \*/\*                 |