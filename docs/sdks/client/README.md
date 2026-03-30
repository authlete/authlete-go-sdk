# Client

## Overview

### Available Operations

* [Get](#get) - Get Client
* [List](#list) - List Clients
* [Create](#create) - Create Client
* [Update](#update) - Update Client
* [UpdateForm](#updateform) - Update Client
* [Delete](#delete) - Delete Client ⚡

## Get

Get a client.


### Example Usage

<!-- UsageSnippet language="go" operationID="client_get_api" method="get" path="/api/{serviceId}/client/get/{clientId}" -->
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

    res, err := s.Client.Get(ctx, "<id>", "<id>")
    if err != nil {
        log.Fatal(err)
    }
    if res.Client != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `serviceID`                                              | `string`                                                 | :heavy_check_mark:                                       | A service ID.                                            |
| `clientID`                                               | `string`                                                 | :heavy_check_mark:                                       | A client ID.                                             |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.ClientGetAPIResponse](../../models/operations/clientgetapiresponse.md), error**

### Errors

| Error Type            | Status Code           | Content Type          |
| --------------------- | --------------------- | --------------------- |
| apierrors.ResultError | 400, 401, 403         | application/json      |
| apierrors.ResultError | 500                   | application/json      |
| apierrors.APIError    | 4XX, 5XX              | \*/\*                 |

## List

Get a list of clients on a service.

If the access token can view a full service (including an admin), all clients within the
service are returned. Otherwise, only clients that the access token can view within the
service are returned.
- ViewClient: []


### Example Usage: full

<!-- UsageSnippet language="go" operationID="client_get_list_api" method="get" path="/api/{serviceId}/client/get/list" example="full" -->
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

    res, err := s.Client.List(ctx, "<id>", nil, nil, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.ClientGetListResponse != nil {
        // handle response
    }
}
```
### Example Usage: limited

<!-- UsageSnippet language="go" operationID="client_get_list_api" method="get" path="/api/{serviceId}/client/get/list" example="limited" -->
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

    res, err := s.Client.List(ctx, "<id>", nil, nil, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.ClientGetListResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                                                                                                         | Type                                                                                                                                                                                                                                              | Required                                                                                                                                                                                                                                          | Description                                                                                                                                                                                                                                       |
| ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                                                                                                             | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                                                                                             | :heavy_check_mark:                                                                                                                                                                                                                                | The context to use for the request.                                                                                                                                                                                                               |
| `serviceID`                                                                                                                                                                                                                                       | `string`                                                                                                                                                                                                                                          | :heavy_check_mark:                                                                                                                                                                                                                                | A service ID.                                                                                                                                                                                                                                     |
| `developer`                                                                                                                                                                                                                                       | `*string`                                                                                                                                                                                                                                         | :heavy_minus_sign:                                                                                                                                                                                                                                | The developer of client applications. The default value is null. If this parameter is not set<br/>to `null`, client application of the specified developer are returned. Otherwise, all client<br/>applications that belong to the service are returned.<br/> |
| `start`                                                                                                                                                                                                                                           | `*int`                                                                                                                                                                                                                                            | :heavy_minus_sign:                                                                                                                                                                                                                                | Start index (inclusive) of the result set. The default value is 0. Must not be a negative number.                                                                                                                                                 |
| `end`                                                                                                                                                                                                                                             | `*int`                                                                                                                                                                                                                                            | :heavy_minus_sign:                                                                                                                                                                                                                                | End index (exclusive) of the result set. The default value is 5. Must not be a negative number.                                                                                                                                                   |
| `opts`                                                                                                                                                                                                                                            | [][operations.Option](../../models/operations/option.md)                                                                                                                                                                                          | :heavy_minus_sign:                                                                                                                                                                                                                                | The options for this request.                                                                                                                                                                                                                     |

### Response

**[*operations.ClientGetListAPIResponse](../../models/operations/clientgetlistapiresponse.md), error**

### Errors

| Error Type            | Status Code           | Content Type          |
| --------------------- | --------------------- | --------------------- |
| apierrors.ResultError | 400, 401, 403         | application/json      |
| apierrors.ResultError | 500                   | application/json      |
| apierrors.APIError    | 4XX, 5XX              | \*/\*                 |

## Create

Create a new client.


### Example Usage

<!-- UsageSnippet language="go" operationID="client_create_api" method="post" path="/api/{serviceId}/client/create" -->
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

    res, err := s.Client.Create(ctx, "<id>", &components.ClientInput{
        ClientName: authlete.Pointer("My Client"),
        ClientIDAlias: authlete.Pointer("my-client"),
        ClientIDAliasEnabled: authlete.Pointer(true),
        ClientType: components.ClientTypeConfidential.ToPointer(),
        ApplicationType: components.ApplicationTypeWeb.ToPointer(),
        Developer: authlete.Pointer("john"),
        GrantTypes: []components.GrantType{
            components.GrantTypeAuthorizationCode,
            components.GrantTypeRefreshToken,
        },
        ResponseTypes: []components.ResponseType{
            components.ResponseTypeCode,
            components.ResponseTypeToken,
        },
        RedirectUris: []string{
            "https://my-client.example.com/cb1",
            "https://my-client.example.com/cb2",
        },
        TokenAuthMethod: components.ClientAuthMethodClientSecretBasic.ToPointer(),
        Attributes: []components.Pair{
            components.Pair{
                Key: authlete.Pointer("attribute1-key"),
                Value: authlete.Pointer("attribute1-value"),
            },
            components.Pair{
                Key: authlete.Pointer("attribute2-key"),
                Value: authlete.Pointer("attribute2-value"),
            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.Client != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                                                                                                                                                                                                                                 | Type                                                                                                                                                                                                                                                                                                                                                                      | Required                                                                                                                                                                                                                                                                                                                                                                  | Description                                                                                                                                                                                                                                                                                                                                                               | Example                                                                                                                                                                                                                                                                                                                                                                   |
| ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                                                                                                                                                                                                                                     | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                                                                                                                                                                                                                     | :heavy_check_mark:                                                                                                                                                                                                                                                                                                                                                        | The context to use for the request.                                                                                                                                                                                                                                                                                                                                       |                                                                                                                                                                                                                                                                                                                                                                           |
| `serviceID`                                                                                                                                                                                                                                                                                                                                                               | `string`                                                                                                                                                                                                                                                                                                                                                                  | :heavy_check_mark:                                                                                                                                                                                                                                                                                                                                                        | A service ID.                                                                                                                                                                                                                                                                                                                                                             |                                                                                                                                                                                                                                                                                                                                                                           |
| `body`                                                                                                                                                                                                                                                                                                                                                                    | [*components.ClientInput](../../models/components/clientinput.md)                                                                                                                                                                                                                                                                                                         | :heavy_minus_sign:                                                                                                                                                                                                                                                                                                                                                        | N/A                                                                                                                                                                                                                                                                                                                                                                       | {<br/>"number": 1140735077,<br/>"serviceNumber": 715948317,<br/>"clientName": "My Test Client",<br/>"clientId": "1140735077",<br/>"clientSecret": "gXz97ISgLs4HuXwOZWch8GEmgL4YMvUJwu3er_kDVVGcA0UOhA9avLPbEmoeZdagi9yC_-tEiT2BdRyH9dbrQQ",<br/>"clientType": "PUBLIC",<br/>"redirectUris": [<br/>"https://example.com/callback"<br/>],<br/>"responseTypes": [<br/>"CODE"<br/>],<br/>"grantTypes": [<br/>"AUTHORIZATION_CODE"<br/>]<br/>} |
| `opts`                                                                                                                                                                                                                                                                                                                                                                    | [][operations.Option](../../models/operations/option.md)                                                                                                                                                                                                                                                                                                                  | :heavy_minus_sign:                                                                                                                                                                                                                                                                                                                                                        | The options for this request.                                                                                                                                                                                                                                                                                                                                             |                                                                                                                                                                                                                                                                                                                                                                           |

### Response

**[*operations.ClientCreateAPIResponse](../../models/operations/clientcreateapiresponse.md), error**

### Errors

| Error Type            | Status Code           | Content Type          |
| --------------------- | --------------------- | --------------------- |
| apierrors.ResultError | 400, 401, 403         | application/json      |
| apierrors.ResultError | 500                   | application/json      |
| apierrors.APIError    | 4XX, 5XX              | \*/\*                 |

## Update

Update a client.


### Example Usage

<!-- UsageSnippet language="go" operationID="client_update_api" method="post" path="/api/{serviceId}/client/update/{clientId}" -->
```go
package main

import(
	"context"
	"os"
	authlete "github.com/authlete/authlete-go-sdk"
	"github.com/authlete/authlete-go-sdk/models/components"
	"github.com/authlete/authlete-go-sdk/optionalnullable"
	"log"
)

func main() {
    ctx := context.Background()

    s := authlete.New(
        authlete.WithSecurity(os.Getenv("AUTHLETE_BEARER")),
    )

    res, err := s.Client.Update(ctx, "<id>", "<id>", &components.ClientInput{
        ClientName: authlete.Pointer("My updated client"),
        ClientIDAlias: authlete.Pointer("my-client"),
        ClientIDAliasEnabled: authlete.Pointer(true),
        ClientType: components.ClientTypeConfidential.ToPointer(),
        ApplicationType: components.ApplicationTypeWeb.ToPointer(),
        TLSClientCertificateBoundAccessTokens: authlete.Pointer(false),
        Developer: authlete.Pointer("john"),
        GrantTypes: []components.GrantType{
            components.GrantTypeAuthorizationCode,
            components.GrantTypeRefreshToken,
        },
        ResponseTypes: []components.ResponseType{
            components.ResponseTypeCode,
            components.ResponseTypeToken,
        },
        RedirectUris: []string{
            "https://my-client.example.com/cb1",
            "https://my-client.example.com/cb2",
        },
        TokenAuthMethod: components.ClientAuthMethodClientSecretBasic.ToPointer(),
        ParRequired: authlete.Pointer(false),
        RequestObjectRequired: authlete.Pointer(false),
        DefaultMaxAge: authlete.Pointer[int](0),
        IDTokenSignAlg: optionalnullable.From(authlete.Pointer(components.JwsAlgRs256)),
        AuthTimeRequired: authlete.Pointer(false),
        SubjectType: components.SubjectTypePublic.ToPointer(),
        BcUserCodeRequired: authlete.Pointer(false),
        Attributes: []components.Pair{
            components.Pair{
                Key: authlete.Pointer("attribute1-key"),
                Value: authlete.Pointer("attribute1-value"),
            },
            components.Pair{
                Key: authlete.Pointer("attribute2-key"),
                Value: authlete.Pointer("attribute2-value"),
            },
        },
        FrontChannelRequestObjectEncryptionRequired: authlete.Pointer(false),
        RequestObjectEncryptionAlgMatchRequired: authlete.Pointer(false),
        RequestObjectEncryptionEncMatchRequired: authlete.Pointer(false),
        AdditionalProperties: map[string]any{
            "derivedSectorIdentifier": "my-client.example.com",
            "dynamicallyRegistered": false,
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.Client != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                                                                                                                                                                                                                                 | Type                                                                                                                                                                                                                                                                                                                                                                      | Required                                                                                                                                                                                                                                                                                                                                                                  | Description                                                                                                                                                                                                                                                                                                                                                               | Example                                                                                                                                                                                                                                                                                                                                                                   |
| ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                                                                                                                                                                                                                                     | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                                                                                                                                                                                                                     | :heavy_check_mark:                                                                                                                                                                                                                                                                                                                                                        | The context to use for the request.                                                                                                                                                                                                                                                                                                                                       |                                                                                                                                                                                                                                                                                                                                                                           |
| `serviceID`                                                                                                                                                                                                                                                                                                                                                               | `string`                                                                                                                                                                                                                                                                                                                                                                  | :heavy_check_mark:                                                                                                                                                                                                                                                                                                                                                        | A service ID.                                                                                                                                                                                                                                                                                                                                                             |                                                                                                                                                                                                                                                                                                                                                                           |
| `clientID`                                                                                                                                                                                                                                                                                                                                                                | `string`                                                                                                                                                                                                                                                                                                                                                                  | :heavy_check_mark:                                                                                                                                                                                                                                                                                                                                                        | A client ID.                                                                                                                                                                                                                                                                                                                                                              |                                                                                                                                                                                                                                                                                                                                                                           |
| `body`                                                                                                                                                                                                                                                                                                                                                                    | [*components.ClientInput](../../models/components/clientinput.md)                                                                                                                                                                                                                                                                                                         | :heavy_minus_sign:                                                                                                                                                                                                                                                                                                                                                        | N/A                                                                                                                                                                                                                                                                                                                                                                       | {<br/>"number": 1140735077,<br/>"serviceNumber": 715948317,<br/>"clientName": "My Test Client",<br/>"clientId": "1140735077",<br/>"clientSecret": "gXz97ISgLs4HuXwOZWch8GEmgL4YMvUJwu3er_kDVVGcA0UOhA9avLPbEmoeZdagi9yC_-tEiT2BdRyH9dbrQQ",<br/>"clientType": "PUBLIC",<br/>"redirectUris": [<br/>"https://example.com/callback"<br/>],<br/>"responseTypes": [<br/>"CODE"<br/>],<br/>"grantTypes": [<br/>"AUTHORIZATION_CODE"<br/>]<br/>} |
| `opts`                                                                                                                                                                                                                                                                                                                                                                    | [][operations.Option](../../models/operations/option.md)                                                                                                                                                                                                                                                                                                                  | :heavy_minus_sign:                                                                                                                                                                                                                                                                                                                                                        | The options for this request.                                                                                                                                                                                                                                                                                                                                             |                                                                                                                                                                                                                                                                                                                                                                           |

### Response

**[*operations.ClientUpdateAPIResponse](../../models/operations/clientupdateapiresponse.md), error**

### Errors

| Error Type            | Status Code           | Content Type          |
| --------------------- | --------------------- | --------------------- |
| apierrors.ResultError | 400, 401, 403         | application/json      |
| apierrors.ResultError | 500                   | application/json      |
| apierrors.APIError    | 4XX, 5XX              | \*/\*                 |

## UpdateForm

Update a client.


### Example Usage

<!-- UsageSnippet language="go" operationID="client_update_api_form" method="post" path="/api/{serviceId}/client/update/{clientId}" -->
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

    res, err := s.Client.UpdateForm(ctx, "<id>", "<id>", map[string]any{
        "clientName": "My Test Client",
        "clientType": "PUBLIC",
        "grantTypes": []any{
            "AUTHORIZATION_CODE",
        },
        "responseTypes": []any{
            "CODE",
        },
        "redirectUris": []any{
            "https://example.com/callback",
        },
        "number": 1140735077,
        "serviceNumber": 715948317,
        "clientId": "1140735077",
        "clientSecret": "gXz97ISgLs4HuXwOZWch8GEmgL4YMvUJwu3er_kDVVGcA0UOhA9avLPbEmoeZdagi9yC_-tEiT2BdRyH9dbrQQ",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.Client != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                                                                                                                                                                                                                                 | Type                                                                                                                                                                                                                                                                                                                                                                      | Required                                                                                                                                                                                                                                                                                                                                                                  | Description                                                                                                                                                                                                                                                                                                                                                               | Example                                                                                                                                                                                                                                                                                                                                                                   |
| ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                                                                                                                                                                                                                                     | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                                                                                                                                                                                                                     | :heavy_check_mark:                                                                                                                                                                                                                                                                                                                                                        | The context to use for the request.                                                                                                                                                                                                                                                                                                                                       |                                                                                                                                                                                                                                                                                                                                                                           |
| `serviceID`                                                                                                                                                                                                                                                                                                                                                               | `string`                                                                                                                                                                                                                                                                                                                                                                  | :heavy_check_mark:                                                                                                                                                                                                                                                                                                                                                        | A service ID.                                                                                                                                                                                                                                                                                                                                                             |                                                                                                                                                                                                                                                                                                                                                                           |
| `clientID`                                                                                                                                                                                                                                                                                                                                                                | `string`                                                                                                                                                                                                                                                                                                                                                                  | :heavy_check_mark:                                                                                                                                                                                                                                                                                                                                                        | A client ID.                                                                                                                                                                                                                                                                                                                                                              |                                                                                                                                                                                                                                                                                                                                                                           |
| `body`                                                                                                                                                                                                                                                                                                                                                                    | map[string]`any`                                                                                                                                                                                                                                                                                                                                                          | :heavy_minus_sign:                                                                                                                                                                                                                                                                                                                                                        | N/A                                                                                                                                                                                                                                                                                                                                                                       | {<br/>"number": 1140735077,<br/>"serviceNumber": 715948317,<br/>"clientName": "My Test Client",<br/>"clientId": "1140735077",<br/>"clientSecret": "gXz97ISgLs4HuXwOZWch8GEmgL4YMvUJwu3er_kDVVGcA0UOhA9avLPbEmoeZdagi9yC_-tEiT2BdRyH9dbrQQ",<br/>"clientType": "PUBLIC",<br/>"redirectUris": [<br/>"https://example.com/callback"<br/>],<br/>"responseTypes": [<br/>"CODE"<br/>],<br/>"grantTypes": [<br/>"AUTHORIZATION_CODE"<br/>]<br/>} |
| `opts`                                                                                                                                                                                                                                                                                                                                                                    | [][operations.Option](../../models/operations/option.md)                                                                                                                                                                                                                                                                                                                  | :heavy_minus_sign:                                                                                                                                                                                                                                                                                                                                                        | The options for this request.                                                                                                                                                                                                                                                                                                                                             |                                                                                                                                                                                                                                                                                                                                                                           |

### Response

**[*operations.ClientUpdateAPIFormResponse](../../models/operations/clientupdateapiformresponse.md), error**

### Errors

| Error Type            | Status Code           | Content Type          |
| --------------------- | --------------------- | --------------------- |
| apierrors.ResultError | 400, 401, 403         | application/json      |
| apierrors.ResultError | 500                   | application/json      |
| apierrors.APIError    | 4XX, 5XX              | \*/\*                 |

## Delete

Delete a client.


### Example Usage

<!-- UsageSnippet language="go" operationID="client_delete_api" method="delete" path="/api/{serviceId}/client/delete/{clientId}" -->
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

    res, err := s.Client.Delete(ctx, "<id>", "<id>")
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
| `clientID`                                               | `string`                                                 | :heavy_check_mark:                                       | The client ID.                                           |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.ClientDeleteAPIResponse](../../models/operations/clientdeleteapiresponse.md), error**

### Errors

| Error Type            | Status Code           | Content Type          |
| --------------------- | --------------------- | --------------------- |
| apierrors.ResultError | 400, 401, 403         | application/json      |
| apierrors.ResultError | 500                   | application/json      |
| apierrors.APIError    | 4XX, 5XX              | \*/\*                 |