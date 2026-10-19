# Service

## Overview

### Available Operations

* [Get](#get) - Get Service
* [List](#list) - List Services
* [Update](#update) - Update Service
* [Create](#create) - Create Service (IDP)
* [Remove](#remove) - Remove Service (IDP) ⚡
* [GetConfiguration](#getconfiguration) - Get Service Configuration

## Get

Get a service.

If the access token can only view or modify clients underneath this service, but does not have access to view this service directly, a limited view of the service will be returned.


### Example Usage: full

<!-- UsageSnippet language="go" operationID="service_get_api" method="get" path="/api/{serviceId}/service/get" example="full" -->
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
        authlete.WithServiceID("<id>"),
        authlete.WithSecurity(os.Getenv("AUTHLETE_BEARER")),
    )

    res, err := s.Service.Get(ctx, operations.ServiceGetAPIRequest{})
    if err != nil {
        log.Fatal(err)
    }
    if res.Service != nil {
        // handle response
    }
}
```
### Example Usage: limited

<!-- UsageSnippet language="go" operationID="service_get_api" method="get" path="/api/{serviceId}/service/get" example="limited" -->
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
        authlete.WithServiceID("<id>"),
        authlete.WithSecurity(os.Getenv("AUTHLETE_BEARER")),
    )

    res, err := s.Service.Get(ctx, operations.ServiceGetAPIRequest{})
    if err != nil {
        log.Fatal(err)
    }
    if res.Service != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `ctx`                                                                              | [context.Context](https://pkg.go.dev/context#Context)                              | :heavy_check_mark:                                                                 | The context to use for the request.                                                |
| `request`                                                                          | [operations.ServiceGetAPIRequest](../../models/operations/servicegetapirequest.md) | :heavy_check_mark:                                                                 | The request object to use for the request.                                         |
| `opts`                                                                             | [][operations.Option](../../models/operations/option.md)                           | :heavy_minus_sign:                                                                 | The options for this request.                                                      |

### Response

**[*operations.ServiceGetAPIResponse](../../models/operations/servicegetapiresponse.md), error**

### Errors

| Error Type            | Status Code           | Content Type          |
| --------------------- | --------------------- | --------------------- |
| apierrors.ResultError | 400, 401, 403         | application/json      |
| apierrors.ResultError | 429                   | application/json      |
| apierrors.ResultError | 500                   | application/json      |
| apierrors.APIError    | 4XX, 5XX              | \*/\*                 |

## List

Get a list of services.

If the access token can only view or modify clients underneath a service, but does not have access to view that service directly, a limited view of the service will be returned. Otherwise, all properties of the service are returned.

If the access token is an administrative token, this returns a list of all services on the Authlete instance. Otherwise, all services that the access token can view, even in a limited fashion, are returned.


### Example Usage: full

<!-- UsageSnippet language="go" operationID="service_get_list_api" method="get" path="/api/service/get/list" example="full" -->
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

    res, err := s.Service.List(ctx, operations.ServiceGetListAPIRequest{})
    if err != nil {
        log.Fatal(err)
    }
    if res.ServiceGetListResponse != nil {
        // handle response
    }
}
```
### Example Usage: limited

<!-- UsageSnippet language="go" operationID="service_get_list_api" method="get" path="/api/service/get/list" example="limited" -->
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

    res, err := s.Service.List(ctx, operations.ServiceGetListAPIRequest{})
    if err != nil {
        log.Fatal(err)
    }
    if res.ServiceGetListResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `request`                                                                                  | [operations.ServiceGetListAPIRequest](../../models/operations/servicegetlistapirequest.md) | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |
| `opts`                                                                                     | [][operations.Option](../../models/operations/option.md)                                   | :heavy_minus_sign:                                                                         | The options for this request.                                                              |

### Response

**[*operations.ServiceGetListAPIResponse](../../models/operations/servicegetlistapiresponse.md), error**

### Errors

| Error Type            | Status Code           | Content Type          |
| --------------------- | --------------------- | --------------------- |
| apierrors.ResultError | 400, 401, 403         | application/json      |
| apierrors.ResultError | 429                   | application/json      |
| apierrors.ResultError | 500                   | application/json      |
| apierrors.APIError    | 4XX, 5XX              | \*/\*                 |

## Update

Update a service.


### Example Usage

<!-- UsageSnippet language="go" operationID="service_update_api" method="post" path="/api/{serviceId}/service/update" -->
```go
package main

import(
	"context"
	"os"
	authlete "github.com/authlete/authlete-go-sdk"
	"github.com/authlete/authlete-go-sdk/optionalnullable"
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

    res, err := s.Service.Update(ctx, operations.ServiceUpdateAPIRequest{
        Body: &components.ServiceInput{
            ServiceName: optionalnullable.From(authlete.Pointer("My updated service")),
            Issuer: optionalnullable.From(authlete.Pointer("https://my-service.example.com")),
            ClientsPerDeveloper: authlete.Pointer[int](0),
            SupportedGrantTypes: optionalnullable.From(authlete.Pointer([]components.GrantType{
                components.GrantTypeAuthorizationCode,
                components.GrantTypeRefreshToken,
            })),
            SupportedResponseTypes: optionalnullable.From(authlete.Pointer([]components.ResponseType{
                components.ResponseTypeCode,
            })),
            ErrorDescriptionOmitted: authlete.Pointer(false),
            ErrorURIOmitted: authlete.Pointer(false),
            AuthorizationEndpoint: optionalnullable.From(authlete.Pointer("https://my-service.example.com/authz")),
            SupportedDisplays: optionalnullable.From(authlete.Pointer([]components.Display{
                components.DisplayPage,
            })),
            PkceRequired: authlete.Pointer(true),
            PkceS256Required: authlete.Pointer(false),
            AuthorizationResponseDuration: authlete.Pointer[int64](0),
            AuthorizationCodeDuration: authlete.Pointer[int64](0),
            TokenEndpoint: optionalnullable.From(authlete.Pointer("https://my-service.example.com/token")),
            SupportedTokenAuthMethods: optionalnullable.From(authlete.Pointer([]components.ClientAuthMethod{
                components.ClientAuthMethodClientSecretBasic,
            })),
            MissingClientIDAllowed: authlete.Pointer(false),
            RevocationEndpoint: optionalnullable.From(authlete.Pointer("https://my-service.example.com/revocation")),
            SupportedRevocationAuthMethods: optionalnullable.From(authlete.Pointer([]components.ClientAuthMethod{
                components.ClientAuthMethodClientSecretBasic,
            })),
            IntrospectionEndpoint: optionalnullable.From(authlete.Pointer("https://my-service.example.com/introspection")),
            SupportedIntrospectionAuthMethods: optionalnullable.From(authlete.Pointer([]components.ClientAuthMethod{
                components.ClientAuthMethodClientSecretBasic,
            })),
            PushedAuthReqDuration: authlete.Pointer[int64](0),
            ParRequired: authlete.Pointer(false),
            RequestObjectRequired: authlete.Pointer(false),
            TraditionalRequestObjectProcessingApplied: authlete.Pointer(false),
            MutualTLSValidatePkiCertChain: authlete.Pointer(false),
            AccessTokenType: optionalnullable.From(authlete.Pointer("Bearer")),
            TLSClientCertificateBoundAccessTokens: authlete.Pointer(false),
            AccessTokenDuration: authlete.Pointer[int64](3600),
            SingleAccessTokenPerSubject: authlete.Pointer(false),
            RefreshTokenDuration: authlete.Pointer[int64](3600),
            RefreshTokenDurationKept: authlete.Pointer(false),
            RefreshTokenDurationReset: authlete.Pointer(false),
            RefreshTokenKept: authlete.Pointer(false),
            SupportedScopes: optionalnullable.From(authlete.Pointer([]components.Scope{
                components.Scope{
                    Name: authlete.Pointer("history.read"),
                    DefaultEntry: authlete.Pointer(false),
                    Description: authlete.Pointer("A permission to read your history."),
                },
                components.Scope{
                    Name: authlete.Pointer("timeline.read"),
                    DefaultEntry: authlete.Pointer(false),
                    Description: authlete.Pointer("A permission to read your timeline."),
                },
            })),
            ScopeRequired: authlete.Pointer(false),
            IDTokenDuration: authlete.Pointer[int64](0),
            AllowableClockSkew: authlete.Pointer[int](0),
            SupportedClaimTypes: optionalnullable.From(authlete.Pointer([]components.ClaimType{
                components.ClaimTypeNormal,
            })),
            ClaimShortcutRestrictive: authlete.Pointer(false),
            DynamicRegistrationSupported: authlete.Pointer(false),
            BackchannelAuthReqIDDuration: authlete.Pointer[int](0),
            BackchannelPollingInterval: authlete.Pointer[int](0),
            BackchannelUserCodeParameterSupported: authlete.Pointer(false),
            BackchannelBindingMessageRequiredInFapi: authlete.Pointer(false),
            DeviceFlowCodeDuration: authlete.Pointer[int](0),
            DeviceFlowPollingInterval: authlete.Pointer[int](0),
            UserCodeLength: authlete.Pointer[int](0),
            Attributes: optionalnullable.From(authlete.Pointer([]components.Pair{
                components.Pair{
                    Key: authlete.Pointer("attribute1-key"),
                    Value: authlete.Pointer("attribute1-value"),
                },
                components.Pair{
                    Key: authlete.Pointer("attribute2-key"),
                    Value: authlete.Pointer("attribute2-value"),
                },
            })),
            NbfOptional: authlete.Pointer(false),
            IssSuppressed: authlete.Pointer(false),
            TokenExpirationLinked: authlete.Pointer(false),
            FrontChannelRequestObjectEncryptionRequired: authlete.Pointer(false),
            RequestObjectEncryptionAlgMatchRequired: authlete.Pointer(false),
            RequestObjectEncryptionEncMatchRequired: authlete.Pointer(false),
            HsmEnabled: authlete.Pointer(false),
            GrantManagementActionRequired: authlete.Pointer(false),
            UnauthorizedOnClientConfigSupported: authlete.Pointer(false),
            DcrScopeUsedAsRequestable: authlete.Pointer(false),
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.Service != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `request`                                                                                | [operations.ServiceUpdateAPIRequest](../../models/operations/serviceupdateapirequest.md) | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |
| `opts`                                                                                   | [][operations.Option](../../models/operations/option.md)                                 | :heavy_minus_sign:                                                                       | The options for this request.                                                            |

### Response

**[*operations.ServiceUpdateAPIResponse](../../models/operations/serviceupdateapiresponse.md), error**

### Errors

| Error Type            | Status Code           | Content Type          |
| --------------------- | --------------------- | --------------------- |
| apierrors.ResultError | 400, 401, 403         | application/json      |
| apierrors.ResultError | 429                   | application/json      |
| apierrors.ResultError | 500                   | application/json      |
| apierrors.APIError    | 4XX, 5XX              | \*/\*                 |

## Create

Create a new service on the API server, belonging to the specified organization. A service can be created with an organization token or by a user with the CREATE_SERVICE role.

This endpoint is hosted on the Authlete IdP server (`https://login.authlete.com`), not on the regional API clusters.


### Example Usage

<!-- UsageSnippet language="go" operationID="service_create_idp_api" method="post" path="/api/service" -->
```go
package main

import(
	"context"
	"os"
	authlete "github.com/authlete/authlete-go-sdk"
	"github.com/authlete/authlete-go-sdk/optionalnullable"
	"github.com/authlete/authlete-go-sdk/models/components"
	"log"
)

func main() {
    ctx := context.Background()

    s := authlete.New(
        authlete.WithSecurity(os.Getenv("AUTHLETE_BEARER")),
    )

    res, err := s.Service.Create(ctx, components.ServiceCreateIdpRequest{
        APIServerID: authlete.Pointer[int64](76281),
        OrganizationID: 123456789012345,
        Service: &components.ServiceInput{
            ServiceName: optionalnullable.From(authlete.Pointer("My service")),
            Issuer: optionalnullable.From(authlete.Pointer("https://my-service.example.com")),
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.Service != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `request`                                                                                | [components.ServiceCreateIdpRequest](../../models/components/servicecreateidprequest.md) | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |
| `opts`                                                                                   | [][operations.Option](../../models/operations/option.md)                                 | :heavy_minus_sign:                                                                       | The options for this request.                                                            |

### Response

**[*operations.ServiceCreateIdpAPIResponse](../../models/operations/servicecreateidpapiresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.IdpError | 400, 401, 403      | application/json   |
| apierrors.IdpError | 500                | application/json   |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## Remove

Delete a service from the API server. A service can be deleted with an organization token or by a user with the MODIFY_SERVICE role.

This endpoint is hosted on the Authlete IdP server (`https://login.authlete.com`), not on the regional API clusters.


### Example Usage

<!-- UsageSnippet language="go" operationID="service_remove_idp_api" method="post" path="/api/service/remove" -->
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

    res, err := s.Service.Remove(ctx, components.ServiceRemoveIdpRequest{
        APIServerID: authlete.Pointer[int64](76281),
        OrganizationID: 123456789012345,
        ServiceID: 21653835348762,
    })
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `request`                                                                                | [components.ServiceRemoveIdpRequest](../../models/components/serviceremoveidprequest.md) | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |
| `opts`                                                                                   | [][operations.Option](../../models/operations/option.md)                                 | :heavy_minus_sign:                                                                       | The options for this request.                                                            |

### Response

**[*operations.ServiceRemoveIdpAPIResponse](../../models/operations/serviceremoveidpapiresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.IdpError | 400, 401, 403      | application/json   |
| apierrors.IdpError | 500                | application/json   |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## GetConfiguration

This API gathers configuration information about a service.

This API is supposed to be called from within the implementation of the configuration endpoint of the service where the service that supports OpenID Connect and [OpenID Connect Discovery 1.0](https://openid.net/specs/openid-connect-discovery-1_0.html) must expose its configuration information in a JSON format. Details about the format are described in "[3. OpenID Provider Metadata](https://openid.net/specs/openid-connect-discovery-1_0.html#ProviderMetadata)" in OpenID Connect Discovery 1.0.


### Example Usage

<!-- UsageSnippet language="go" operationID="service_configuration_api" method="get" path="/api/{serviceId}/service/configuration" -->
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
        authlete.WithServiceID("<id>"),
        authlete.WithSecurity(os.Getenv("AUTHLETE_BEARER")),
    )

    res, err := s.Service.GetConfiguration(ctx, operations.ServiceConfigurationAPIRequest{})
    if err != nil {
        log.Fatal(err)
    }
    if res.Object != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                              | Type                                                                                                   | Required                                                                                               | Description                                                                                            |
| ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                  | :heavy_check_mark:                                                                                     | The context to use for the request.                                                                    |
| `request`                                                                                              | [operations.ServiceConfigurationAPIRequest](../../models/operations/serviceconfigurationapirequest.md) | :heavy_check_mark:                                                                                     | The request object to use for the request.                                                             |
| `opts`                                                                                                 | [][operations.Option](../../models/operations/option.md)                                               | :heavy_minus_sign:                                                                                     | The options for this request.                                                                          |

### Response

**[*operations.ServiceConfigurationAPIResponse](../../models/operations/serviceconfigurationapiresponse.md), error**

### Errors

| Error Type            | Status Code           | Content Type          |
| --------------------- | --------------------- | --------------------- |
| apierrors.ResultError | 400, 401, 403         | application/json      |
| apierrors.ResultError | 429                   | application/json      |
| apierrors.ResultError | 500                   | application/json      |
| apierrors.APIError    | 4XX, 5XX              | \*/\*                 |