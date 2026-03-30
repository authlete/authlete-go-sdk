# Client.Management

## Overview

### Available Operations

* [UpdateLockFlag](#updatelockflag) - Update Client Lock
* [RefreshSecret](#refreshsecret) - Rotate Client Secret
* [UpdateSecret](#updatesecret) - Update Client Secret
* [ListAuthorizedApplications](#listauthorizedapplications) - Get Authorized Applications
* [ListAuthorizedApplicationsWithBody](#listauthorizedapplicationswithbody) - Get Authorized Applications
* [ListAuthorizations](#listauthorizations) - Get Authorized Applications (by Subject)
* [UpdateAuthorizations](#updateauthorizations) - Update Client Tokens
* [DeleteClientTokens](#deleteclienttokens) - Delete Client Tokens
* [RevokeClientTokens](#revokeclienttokens) - Delete Client Tokens
* [DeleteAuthorizations](#deleteauthorizations) - Delete Client Tokens (by Subject)
* [GetGrantedScopesForClient](#getgrantedscopesforclient) - Get Granted Scopes
* [CreateGrantedScopes](#creategrantedscopes) - Get Granted Scopes
* [GetGrantedScopes](#getgrantedscopes) - Get Granted Scopes (by Subject)
* [DeleteGrantedScopesForClient](#deletegrantedscopesforclient) - Delete Granted Scopes
* [DeleteGrantedScopes](#deletegrantedscopes) - Delete Granted Scopes (by Subject)
* [GetRequestableScopes](#getrequestablescopes) - Get Requestable Scopes
* [UpdateRequestableScopesWithBody](#updaterequestablescopeswithbody) - Update Requestable Scopes
* [UpdateRequestableScopes](#updaterequestablescopes) - Update Requestable Scopes
* [DeleteRequestableScopes](#deleterequestablescopes) - Delete Requestable Scopes

## UpdateLockFlag

Lock and unlock a client


### Example Usage

<!-- UsageSnippet language="go" operationID="client_flag_update_api" method="post" path="/api/{serviceId}/client/lock_flag/update/{clientIdentifier}" -->
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

    res, err := s.Client.Management.UpdateLockFlag(ctx, "<id>", "<value>", &components.ClientFlagUpdateRequest{
        ClientLocked: true,
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ClientFlagUpdateResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                 | Type                                                                                      | Required                                                                                  | Description                                                                               |
| ----------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------- |
| `ctx`                                                                                     | [context.Context](https://pkg.go.dev/context#Context)                                     | :heavy_check_mark:                                                                        | The context to use for the request.                                                       |
| `serviceID`                                                                               | `string`                                                                                  | :heavy_check_mark:                                                                        | A service ID.                                                                             |
| `clientIdentifier`                                                                        | `string`                                                                                  | :heavy_check_mark:                                                                        | A client ID.                                                                              |
| `body`                                                                                    | [*components.ClientFlagUpdateRequest](../../models/components/clientflagupdaterequest.md) | :heavy_minus_sign:                                                                        | N/A                                                                                       |
| `opts`                                                                                    | [][operations.Option](../../models/operations/option.md)                                  | :heavy_minus_sign:                                                                        | The options for this request.                                                             |

### Response

**[*operations.ClientFlagUpdateAPIResponse](../../models/operations/clientflagupdateapiresponse.md), error**

### Errors

| Error Type            | Status Code           | Content Type          |
| --------------------- | --------------------- | --------------------- |
| apierrors.ResultError | 400, 401, 403         | application/json      |
| apierrors.ResultError | 500                   | application/json      |
| apierrors.APIError    | 4XX, 5XX              | \*/\*                 |

## RefreshSecret

Refresh the client secret of a client. A new value of the client secret will be generated by the
Authlete server.

If you want to specify a new value, use `/api/client/secret/update` API.


### Example Usage

<!-- UsageSnippet language="go" operationID="client_secret_refresh_api" method="get" path="/api/{serviceId}/client/secret/refresh/{clientIdentifier}" -->
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

    res, err := s.Client.Management.RefreshSecret(ctx, "<id>", "<value>")
    if err != nil {
        log.Fatal(err)
    }
    if res.ClientSecretRefreshResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `serviceID`                                              | `string`                                                 | :heavy_check_mark:                                       | A service ID.                                            |
| `clientIdentifier`                                       | `string`                                                 | :heavy_check_mark:                                       | The client ID or the client ID alias of a client.<br/>   |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.ClientSecretRefreshAPIResponse](../../models/operations/clientsecretrefreshapiresponse.md), error**

### Errors

| Error Type            | Status Code           | Content Type          |
| --------------------- | --------------------- | --------------------- |
| apierrors.ResultError | 400, 401, 403         | application/json      |
| apierrors.ResultError | 500                   | application/json      |
| apierrors.APIError    | 4XX, 5XX              | \*/\*                 |

## UpdateSecret

Update the client secret of a client.

If you want to have the Authlete server generate a new value of the client secret, use `/api/client/secret/refresh`
API.


### Example Usage

<!-- UsageSnippet language="go" operationID="client_secret_update_api" method="post" path="/api/{serviceId}/client/secret/update/{clientIdentifier}" -->
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

    res, err := s.Client.Management.UpdateSecret(ctx, "<id>", "<value>", components.ClientSecretUpdateRequest{
        ClientSecret: "my_updated_client_secret",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ClientSecretUpdateResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `serviceID`                                                                                  | `string`                                                                                     | :heavy_check_mark:                                                                           | A service ID.                                                                                |
| `clientIdentifier`                                                                           | `string`                                                                                     | :heavy_check_mark:                                                                           | The client ID or the client ID alias of a client.<br/>                                       |
| `body`                                                                                       | [components.ClientSecretUpdateRequest](../../models/components/clientsecretupdaterequest.md) | :heavy_check_mark:                                                                           | N/A                                                                                          |
| `opts`                                                                                       | [][operations.Option](../../models/operations/option.md)                                     | :heavy_minus_sign:                                                                           | The options for this request.                                                                |

### Response

**[*operations.ClientSecretUpdateAPIResponse](../../models/operations/clientsecretupdateapiresponse.md), error**

### Errors

| Error Type            | Status Code           | Content Type          |
| --------------------- | --------------------- | --------------------- |
| apierrors.ResultError | 400, 401, 403         | application/json      |
| apierrors.ResultError | 500                   | application/json      |
| apierrors.APIError    | 4XX, 5XX              | \*/\*                 |

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

    res, err := s.Client.Management.ListAuthorizedApplications(ctx, operations.ClientAuthorizationGetListAPIRequest{
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

## ListAuthorizedApplicationsWithBody

Get a list of client applications that an end-user has authorized.

The subject parameter is required.


### Example Usage

<!-- UsageSnippet language="go" operationID="client_authorization_get_list_api_post" method="post" path="/api/{serviceId}/client/authorization/get/list" -->
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

    res, err := s.Client.Management.ListAuthorizedApplicationsWithBody(ctx, "<id>", components.ClientAuthorizationGetListRequest{
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

| Parameter                                                                                                    | Type                                                                                                         | Required                                                                                                     | Description                                                                                                  |
| ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                                        | :heavy_check_mark:                                                                                           | The context to use for the request.                                                                          |
| `serviceID`                                                                                                  | `string`                                                                                                     | :heavy_check_mark:                                                                                           | A service ID.                                                                                                |
| `body`                                                                                                       | [components.ClientAuthorizationGetListRequest](../../models/components/clientauthorizationgetlistrequest.md) | :heavy_check_mark:                                                                                           | N/A                                                                                                          |
| `opts`                                                                                                       | [][operations.Option](../../models/operations/option.md)                                                     | :heavy_minus_sign:                                                                                           | The options for this request.                                                                                |

### Response

**[*operations.ClientAuthorizationGetListAPIPostResponse](../../models/operations/clientauthorizationgetlistapipostresponse.md), error**

### Errors

| Error Type            | Status Code           | Content Type          |
| --------------------- | --------------------- | --------------------- |
| apierrors.ResultError | 400, 401, 403         | application/json      |
| apierrors.ResultError | 500                   | application/json      |
| apierrors.APIError    | 4XX, 5XX              | \*/\*                 |

## ListAuthorizations

Get a list of client applications that an end-user has authorized.
In this variant, the subject is provided in the path.


### Example Usage

<!-- UsageSnippet language="go" operationID="client_authorization_get_list_by_subject_api" method="get" path="/api/{serviceId}/client/authorization/get/list/{subject}" -->
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

    res, err := s.Client.Management.ListAuthorizations(ctx, operations.ClientAuthorizationGetListBySubjectAPIRequest{
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

| Parameter                                                                                                                            | Type                                                                                                                                 | Required                                                                                                                             | Description                                                                                                                          |
| ------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                                                | :heavy_check_mark:                                                                                                                   | The context to use for the request.                                                                                                  |
| `request`                                                                                                                            | [operations.ClientAuthorizationGetListBySubjectAPIRequest](../../models/operations/clientauthorizationgetlistbysubjectapirequest.md) | :heavy_check_mark:                                                                                                                   | The request object to use for the request.                                                                                           |
| `opts`                                                                                                                               | [][operations.Option](../../models/operations/option.md)                                                                             | :heavy_minus_sign:                                                                                                                   | The options for this request.                                                                                                        |

### Response

**[*operations.ClientAuthorizationGetListBySubjectAPIResponse](../../models/operations/clientauthorizationgetlistbysubjectapiresponse.md), error**

### Errors

| Error Type            | Status Code           | Content Type          |
| --------------------- | --------------------- | --------------------- |
| apierrors.ResultError | 400, 401, 403         | application/json      |
| apierrors.ResultError | 500                   | application/json      |
| apierrors.APIError    | 4XX, 5XX              | \*/\*                 |

## UpdateAuthorizations

Update attributes of all existing access tokens given to a client application.


### Example Usage

<!-- UsageSnippet language="go" operationID="client_authorization_update_api" method="post" path="/api/{serviceId}/client/authorization/update/{clientId}" -->
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

    res, err := s.Client.Management.UpdateAuthorizations(ctx, "<id>", "<id>", &components.ClientAuthorizationUpdateRequest{
        Subject: "john",
        Scopes: []string{
            "history.read",
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ClientAuthorizationUpdateResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                   | Type                                                                                                        | Required                                                                                                    | Description                                                                                                 |
| ----------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                       | [context.Context](https://pkg.go.dev/context#Context)                                                       | :heavy_check_mark:                                                                                          | The context to use for the request.                                                                         |
| `serviceID`                                                                                                 | `string`                                                                                                    | :heavy_check_mark:                                                                                          | A service ID.                                                                                               |
| `clientID`                                                                                                  | `string`                                                                                                    | :heavy_check_mark:                                                                                          | A client ID.<br/>                                                                                           |
| `body`                                                                                                      | [*components.ClientAuthorizationUpdateRequest](../../models/components/clientauthorizationupdaterequest.md) | :heavy_minus_sign:                                                                                          | N/A                                                                                                         |
| `opts`                                                                                                      | [][operations.Option](../../models/operations/option.md)                                                    | :heavy_minus_sign:                                                                                          | The options for this request.                                                                               |

### Response

**[*operations.ClientAuthorizationUpdateAPIResponse](../../models/operations/clientauthorizationupdateapiresponse.md), error**

### Errors

| Error Type            | Status Code           | Content Type          |
| --------------------- | --------------------- | --------------------- |
| apierrors.ResultError | 400, 401, 403         | application/json      |
| apierrors.ResultError | 500                   | application/json      |
| apierrors.APIError    | 4XX, 5XX              | \*/\*                 |

## DeleteClientTokens

Delete all existing access tokens issued to a client application by an end-user.

The subject parameter is required and must be provided as a query parameter.


### Example Usage

<!-- UsageSnippet language="go" operationID="client_authorization_delete_api" method="delete" path="/api/{serviceId}/client/authorization/delete/{clientId}" -->
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

    res, err := s.Client.Management.DeleteClientTokens(ctx, "<id>", "<id>", "<value>")
    if err != nil {
        log.Fatal(err)
    }
    if res.ClientAuthorizationDeleteResponse != nil {
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

**[*operations.ClientAuthorizationDeleteAPIResponse](../../models/operations/clientauthorizationdeleteapiresponse.md), error**

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

    res, err := s.Client.Management.RevokeClientTokens(ctx, "<id>", "<id>", operations.ClientAuthorizationDeleteAPIPostRequestBody{
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

## DeleteAuthorizations

Delete all existing access tokens issued to a client application by an end-user.
In this variant, the subject is provided in the path.


### Example Usage

<!-- UsageSnippet language="go" operationID="client_authorization_delete_by_subject_api" method="delete" path="/api/{serviceId}/client/authorization/delete/{clientId}/{subject}" -->
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

    res, err := s.Client.Management.DeleteAuthorizations(ctx, "<id>", "<id>", "<value>")
    if err != nil {
        log.Fatal(err)
    }
    if res.ClientAuthorizationDeleteResponse != nil {
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

**[*operations.ClientAuthorizationDeleteBySubjectAPIResponse](../../models/operations/clientauthorizationdeletebysubjectapiresponse.md), error**

### Errors

| Error Type            | Status Code           | Content Type          |
| --------------------- | --------------------- | --------------------- |
| apierrors.ResultError | 400, 401, 403         | application/json      |
| apierrors.ResultError | 500                   | application/json      |
| apierrors.APIError    | 4XX, 5XX              | \*/\*                 |

## GetGrantedScopesForClient

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

    res, err := s.Client.Management.GetGrantedScopesForClient(ctx, "715948317", "1140735077", "<value>")
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

## CreateGrantedScopes

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

    res, err := s.Client.Management.CreateGrantedScopes(ctx, "<id>", "<id>", operations.ClientGrantedScopesGetAPIPostRequestBody{
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

## GetGrantedScopes

Get the set of scopes that a user has granted to a client application.
In this variant, the subject is provided in the path.


### Example Usage

<!-- UsageSnippet language="go" operationID="client_granted_scopes_get_by_subject_api" method="get" path="/api/{serviceId}/client/granted_scopes/get/{clientId}/{subject}" -->
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

    res, err := s.Client.Management.GetGrantedScopes(ctx, "<id>", "<id>", "<value>")
    if err != nil {
        log.Fatal(err)
    }
    if res.ClientAuthorizationDeleteResponse != nil {
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

**[*operations.ClientGrantedScopesGetBySubjectAPIResponse](../../models/operations/clientgrantedscopesgetbysubjectapiresponse.md), error**

### Errors

| Error Type            | Status Code           | Content Type          |
| --------------------- | --------------------- | --------------------- |
| apierrors.ResultError | 400, 401, 403         | application/json      |
| apierrors.ResultError | 500                   | application/json      |
| apierrors.APIError    | 4XX, 5XX              | \*/\*                 |

## DeleteGrantedScopesForClient

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

    res, err := s.Client.Management.DeleteGrantedScopesForClient(ctx, "<id>", "<id>", "<value>")
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

## DeleteGrantedScopes

Delete the set of scopes that an end-user has granted to a client application.
In this variant, the subject is provided in the path.


### Example Usage

<!-- UsageSnippet language="go" operationID="client_granted_scopes_delete_by_subject_api" method="delete" path="/api/{serviceId}/client/granted_scopes/delete/{clientId}/{subject}" -->
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

    res, err := s.Client.Management.DeleteGrantedScopes(ctx, "<id>", "<id>", "<value>")
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

**[*operations.ClientGrantedScopesDeleteBySubjectAPIResponse](../../models/operations/clientgrantedscopesdeletebysubjectapiresponse.md), error**

### Errors

| Error Type            | Status Code           | Content Type          |
| --------------------- | --------------------- | --------------------- |
| apierrors.ResultError | 400, 401, 403         | application/json      |
| apierrors.ResultError | 500                   | application/json      |
| apierrors.APIError    | 4XX, 5XX              | \*/\*                 |

## GetRequestableScopes

Get the requestable scopes per client


### Example Usage

<!-- UsageSnippet language="go" operationID="client_extension_requestables_scopes_get_api" method="get" path="/api/{serviceId}/client/extension/requestable_scopes/get/{clientId}" -->
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

    res, err := s.Client.Management.GetRequestableScopes(ctx, "<id>", "<id>")
    if err != nil {
        log.Fatal(err)
    }
    if res.ClientExtensionRequestableScopesGetResponse != nil {
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
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.ClientExtensionRequestablesScopesGetAPIResponse](../../models/operations/clientextensionrequestablesscopesgetapiresponse.md), error**

### Errors

| Error Type            | Status Code           | Content Type          |
| --------------------- | --------------------- | --------------------- |
| apierrors.ResultError | 400, 401, 403         | application/json      |
| apierrors.ResultError | 500                   | application/json      |
| apierrors.APIError    | 4XX, 5XX              | \*/\*                 |

## UpdateRequestableScopesWithBody

Update requestable scopes of a client


### Example Usage

<!-- UsageSnippet language="go" operationID="client_extension_requestables_scopes_update_api_post" method="post" path="/api/{serviceId}/client/extension/requestable_scopes/update/{clientId}" -->
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

    res, err := s.Client.Management.UpdateRequestableScopesWithBody(ctx, "<id>", "<id>", components.ClientExtensionRequestableScopesUpdateRequest{})
    if err != nil {
        log.Fatal(err)
    }
    if res.ClientExtensionRequestableScopesUpdateResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                            | Type                                                                                                                                 | Required                                                                                                                             | Description                                                                                                                          |
| ------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                                                | :heavy_check_mark:                                                                                                                   | The context to use for the request.                                                                                                  |
| `serviceID`                                                                                                                          | `string`                                                                                                                             | :heavy_check_mark:                                                                                                                   | A service ID.                                                                                                                        |
| `clientID`                                                                                                                           | `string`                                                                                                                             | :heavy_check_mark:                                                                                                                   | A client ID.<br/>                                                                                                                    |
| `body`                                                                                                                               | [components.ClientExtensionRequestableScopesUpdateRequest](../../models/components/clientextensionrequestablescopesupdaterequest.md) | :heavy_check_mark:                                                                                                                   | N/A                                                                                                                                  |
| `opts`                                                                                                                               | [][operations.Option](../../models/operations/option.md)                                                                             | :heavy_minus_sign:                                                                                                                   | The options for this request.                                                                                                        |

### Response

**[*operations.ClientExtensionRequestablesScopesUpdateAPIPostResponse](../../models/operations/clientextensionrequestablesscopesupdateapipostresponse.md), error**

### Errors

| Error Type            | Status Code           | Content Type          |
| --------------------- | --------------------- | --------------------- |
| apierrors.ResultError | 400, 401, 403         | application/json      |
| apierrors.ResultError | 500                   | application/json      |
| apierrors.APIError    | 4XX, 5XX              | \*/\*                 |

## UpdateRequestableScopes

Update requestable scopes of a client


### Example Usage

<!-- UsageSnippet language="go" operationID="client_extension_requestables_scopes_update_api" method="put" path="/api/{serviceId}/client/extension/requestable_scopes/update/{clientId}" -->
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

    res, err := s.Client.Management.UpdateRequestableScopes(ctx, "<id>", "<id>", components.ClientExtensionRequestableScopesUpdateRequest{})
    if err != nil {
        log.Fatal(err)
    }
    if res.ClientExtensionRequestableScopesUpdateResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                            | Type                                                                                                                                 | Required                                                                                                                             | Description                                                                                                                          |
| ------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                                                | :heavy_check_mark:                                                                                                                   | The context to use for the request.                                                                                                  |
| `serviceID`                                                                                                                          | `string`                                                                                                                             | :heavy_check_mark:                                                                                                                   | A service ID.                                                                                                                        |
| `clientID`                                                                                                                           | `string`                                                                                                                             | :heavy_check_mark:                                                                                                                   | A client ID.<br/>                                                                                                                    |
| `body`                                                                                                                               | [components.ClientExtensionRequestableScopesUpdateRequest](../../models/components/clientextensionrequestablescopesupdaterequest.md) | :heavy_check_mark:                                                                                                                   | N/A                                                                                                                                  |
| `opts`                                                                                                                               | [][operations.Option](../../models/operations/option.md)                                                                             | :heavy_minus_sign:                                                                                                                   | The options for this request.                                                                                                        |

### Response

**[*operations.ClientExtensionRequestablesScopesUpdateAPIResponse](../../models/operations/clientextensionrequestablesscopesupdateapiresponse.md), error**

### Errors

| Error Type            | Status Code           | Content Type          |
| --------------------- | --------------------- | --------------------- |
| apierrors.ResultError | 400, 401, 403         | application/json      |
| apierrors.ResultError | 500                   | application/json      |
| apierrors.APIError    | 4XX, 5XX              | \*/\*                 |

## DeleteRequestableScopes

Delete requestable scopes of a client


### Example Usage

<!-- UsageSnippet language="go" operationID="client_extension_requestables_scopes_delete_api" method="delete" path="/api/{serviceId}/client/extension/requestable_scopes/delete/{clientId}" -->
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

    res, err := s.Client.Management.DeleteRequestableScopes(ctx, "<id>", "<id>")
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
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
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.ClientExtensionRequestablesScopesDeleteAPIResponse](../../models/operations/clientextensionrequestablesscopesdeleteapiresponse.md), error**

### Errors

| Error Type            | Status Code           | Content Type          |
| --------------------- | --------------------- | --------------------- |
| apierrors.ResultError | 400, 401, 403         | application/json      |
| apierrors.ResultError | 500                   | application/json      |
| apierrors.APIError    | 4XX, 5XX              | \*/\*                 |