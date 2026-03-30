# Service

## Overview

### Available Operations

* [Get](#get) - Get Service
* [List](#list) - List Services
* [Update](#update) - Update Service
* [Delete](#delete) - Delete Service ⚡
* [GetConfiguration](#getconfiguration) - Get Service Configuration

## Get

Get a service.

If the access token can only view or modify clients underneath this service, but does not
have access to view this service directly, a limited view of the service will be returned.


### Example Usage: full

<!-- UsageSnippet language="go" operationID="service_get_api" method="get" path="/api/{serviceId}/service/get" example="full" -->
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

    res, err := s.Service.Get(ctx, "<id>")
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
	"log"
)

func main() {
    ctx := context.Background()

    s := authlete.New(
        authlete.WithSecurity(os.Getenv("AUTHLETE_BEARER")),
    )

    res, err := s.Service.Get(ctx, "<id>")
    if err != nil {
        log.Fatal(err)
    }
    if res.Service != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `serviceID`                                              | `string`                                                 | :heavy_check_mark:                                       | A service ID.                                            |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.ServiceGetAPIResponse](../../models/operations/servicegetapiresponse.md), error**

### Errors

| Error Type            | Status Code           | Content Type          |
| --------------------- | --------------------- | --------------------- |
| apierrors.ResultError | 400, 401, 403         | application/json      |
| apierrors.ResultError | 500                   | application/json      |
| apierrors.APIError    | 4XX, 5XX              | \*/\*                 |

## List

Get a list of services.

If the access token can only view or modify clients underneath a service, but does not
have access to view that service directly, a limited view of the service will be returned.
Otherwise, all properties of the service are returned.

If the access token is an administrative token, this returns a list of all services on the Authlete instance.
Otherwise, all services that the access token can view, even in a limited fashion, are returned.


### Example Usage: full

<!-- UsageSnippet language="go" operationID="service_get_list_api" method="get" path="/api/service/get/list" example="full" -->
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

    res, err := s.Service.List(ctx, nil, nil)
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
	"log"
)

func main() {
    ctx := context.Background()

    s := authlete.New(
        authlete.WithSecurity(os.Getenv("AUTHLETE_BEARER")),
    )

    res, err := s.Service.List(ctx, nil, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.ServiceGetListResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                         | Type                                                                                              | Required                                                                                          | Description                                                                                       |
| ------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                             | [context.Context](https://pkg.go.dev/context#Context)                                             | :heavy_check_mark:                                                                                | The context to use for the request.                                                               |
| `start`                                                                                           | `*int`                                                                                            | :heavy_minus_sign:                                                                                | Start index (inclusive) of the result set. The default value is 0. Must not be a negative number. |
| `end`                                                                                             | `*int`                                                                                            | :heavy_minus_sign:                                                                                | End index (exclusive) of the result set. The default value is 5. Must not be a negative number.   |
| `opts`                                                                                            | [][operations.Option](../../models/operations/option.md)                                          | :heavy_minus_sign:                                                                                | The options for this request.                                                                     |

### Response

**[*operations.ServiceGetListAPIResponse](../../models/operations/servicegetlistapiresponse.md), error**

### Errors

| Error Type            | Status Code           | Content Type          |
| --------------------- | --------------------- | --------------------- |
| apierrors.ResultError | 400, 401, 403         | application/json      |
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
	"github.com/authlete/authlete-go-sdk/models/components"
	"log"
)

func main() {
    ctx := context.Background()

    s := authlete.New(
        authlete.WithSecurity(os.Getenv("AUTHLETE_BEARER")),
    )

    res, err := s.Service.Update(ctx, "<id>", &components.ServiceInput{
        ServiceName: authlete.Pointer("My updated service"),
        Issuer: authlete.Pointer("https://my-service.example.com"),
        ClientsPerDeveloper: authlete.Pointer[int](0),
        ClientIDAliasEnabled: authlete.Pointer(true),
        SupportedGrantTypes: []components.GrantType{
            components.GrantTypeAuthorizationCode,
            components.GrantTypeRefreshToken,
        },
        SupportedResponseTypes: []components.ResponseType{
            components.ResponseTypeCode,
        },
        ErrorDescriptionOmitted: authlete.Pointer(false),
        ErrorURIOmitted: authlete.Pointer(false),
        AuthorizationEndpoint: authlete.Pointer("https://my-service.example.com/authz"),
        DirectAuthorizationEndpointEnabled: authlete.Pointer(false),
        SupportedDisplays: []components.Display{
            components.DisplayPage,
        },
        PkceRequired: authlete.Pointer(true),
        PkceS256Required: authlete.Pointer(false),
        AuthorizationResponseDuration: authlete.Pointer[int64](0),
        TokenEndpoint: authlete.Pointer("https://my-service.example.com/token"),
        DirectTokenEndpointEnabled: authlete.Pointer(false),
        SupportedTokenAuthMethods: []components.ClientAuthMethod{
            components.ClientAuthMethodClientSecretBasic,
        },
        MissingClientIDAllowed: authlete.Pointer(false),
        RevocationEndpoint: authlete.Pointer("https://my-service.example.com/revocation"),
        DirectRevocationEndpointEnabled: authlete.Pointer(false),
        SupportedRevocationAuthMethods: []components.ClientAuthMethod{
            components.ClientAuthMethodClientSecretBasic,
        },
        IntrospectionEndpoint: authlete.Pointer("https://my-service.example.com/introspection"),
        DirectIntrospectionEndpointEnabled: authlete.Pointer(false),
        SupportedIntrospectionAuthMethods: []components.ClientAuthMethod{
            components.ClientAuthMethodClientSecretBasic,
        },
        PushedAuthReqDuration: authlete.Pointer[int64](0),
        ParRequired: authlete.Pointer(false),
        RequestObjectRequired: authlete.Pointer(false),
        TraditionalRequestObjectProcessingApplied: authlete.Pointer(false),
        MutualTLSValidatePkiCertChain: authlete.Pointer(false),
        AccessTokenType: authlete.Pointer("Bearer"),
        TLSClientCertificateBoundAccessTokens: authlete.Pointer(false),
        AccessTokenDuration: authlete.Pointer[int64](3600),
        SingleAccessTokenPerSubject: authlete.Pointer(false),
        RefreshTokenDuration: authlete.Pointer[int64](3600),
        RefreshTokenDurationKept: authlete.Pointer(false),
        RefreshTokenDurationReset: authlete.Pointer(false),
        RefreshTokenKept: authlete.Pointer(false),
        SupportedScopes: []components.Scope{
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
        },
        ScopeRequired: authlete.Pointer(false),
        IDTokenDuration: authlete.Pointer[int64](0),
        AllowableClockSkew: authlete.Pointer[int](0),
        SupportedClaimTypes: []components.ClaimType{
            components.ClaimTypeNormal,
        },
        ClaimShortcutRestrictive: authlete.Pointer(false),
        DirectJwksEndpointEnabled: authlete.Pointer(false),
        DirectUserInfoEndpointEnabled: authlete.Pointer(false),
        DynamicRegistrationSupported: authlete.Pointer(false),
        BackchannelAuthReqIDDuration: authlete.Pointer[int](0),
        BackchannelPollingInterval: authlete.Pointer[int](0),
        BackchannelUserCodeParameterSupported: authlete.Pointer(false),
        BackchannelBindingMessageRequiredInFapi: authlete.Pointer(false),
        DeviceFlowCodeDuration: authlete.Pointer[int](0),
        DeviceFlowPollingInterval: authlete.Pointer[int](0),
        UserCodeLength: authlete.Pointer[int](0),
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

| Parameter                                                                                                                                                                                                                                               | Type                                                                                                                                                                                                                                                    | Required                                                                                                                                                                                                                                                | Description                                                                                                                                                                                                                                             | Example                                                                                                                                                                                                                                                 |
| ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                                                                                                                   | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                                                                                                   | :heavy_check_mark:                                                                                                                                                                                                                                      | The context to use for the request.                                                                                                                                                                                                                     |                                                                                                                                                                                                                                                         |
| `serviceID`                                                                                                                                                                                                                                             | `string`                                                                                                                                                                                                                                                | :heavy_check_mark:                                                                                                                                                                                                                                      | A service ID.                                                                                                                                                                                                                                           |                                                                                                                                                                                                                                                         |
| `body`                                                                                                                                                                                                                                                  | [*components.ServiceInput](../../models/components/serviceinput.md)                                                                                                                                                                                     | :heavy_minus_sign:                                                                                                                                                                                                                                      | N/A                                                                                                                                                                                                                                                     | {<br/>"number": 715948317,<br/>"serviceName": "My Test Service",<br/>"issuer": "https://example.com",<br/>"supportedScopes": [<br/>"profile",<br/>"email",<br/>"openid"<br/>],<br/>"supportedResponseTypes": [<br/>"CODE"<br/>],<br/>"supportedGrantTypes": [<br/>"AUTHORIZATION_CODE",<br/>"REFRESH_TOKEN"<br/>]<br/>} |
| `opts`                                                                                                                                                                                                                                                  | [][operations.Option](../../models/operations/option.md)                                                                                                                                                                                                | :heavy_minus_sign:                                                                                                                                                                                                                                      | The options for this request.                                                                                                                                                                                                                           |                                                                                                                                                                                                                                                         |

### Response

**[*operations.ServiceUpdateAPIResponse](../../models/operations/serviceupdateapiresponse.md), error**

### Errors

| Error Type            | Status Code           | Content Type          |
| --------------------- | --------------------- | --------------------- |
| apierrors.ResultError | 400, 401, 403         | application/json      |
| apierrors.ResultError | 500                   | application/json      |
| apierrors.APIError    | 4XX, 5XX              | \*/\*                 |

## Delete

Delete a service.


### Example Usage

<!-- UsageSnippet language="go" operationID="service_delete_api" method="delete" path="/api/{serviceId}/service/delete" -->
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

    res, err := s.Service.Delete(ctx, "<id>")
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
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.ServiceDeleteAPIResponse](../../models/operations/servicedeleteapiresponse.md), error**

### Errors

| Error Type            | Status Code           | Content Type          |
| --------------------- | --------------------- | --------------------- |
| apierrors.ResultError | 400, 401, 403         | application/json      |
| apierrors.ResultError | 500                   | application/json      |
| apierrors.APIError    | 4XX, 5XX              | \*/\*                 |

## GetConfiguration

This API gathers configuration information about a service.

This API is supposed to be called from within the implementation of the configuration endpoint of
the service where the service that supports OpenID Connect and [OpenID Connect Discovery 1.0](https://openid.net/specs/openid-connect-discovery-1_0.html)
must expose its configuration information in a JSON format. Details about the format are described
in "[3. OpenID Provider Metadata](https://openid.net/specs/openid-connect-discovery-1_0.html#ProviderMetadata)"
in OpenID Connect Discovery 1.0.


### Example Usage

<!-- UsageSnippet language="go" operationID="service_configuration_api" method="get" path="/api/{serviceId}/service/configuration" -->
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

    res, err := s.Service.GetConfiguration(ctx, "<id>", nil, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.Object != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                                        | Type                                                                                                                                                                             | Required                                                                                                                                                                         | Description                                                                                                                                                                      |
| -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                            | :heavy_check_mark:                                                                                                                                                               | The context to use for the request.                                                                                                                                              |
| `serviceID`                                                                                                                                                                      | `string`                                                                                                                                                                         | :heavy_check_mark:                                                                                                                                                               | A service ID.                                                                                                                                                                    |
| `pretty`                                                                                                                                                                         | `*bool`                                                                                                                                                                          | :heavy_minus_sign:                                                                                                                                                               | This boolean value indicates whether the JSON in the response should be formatted or not. If `true`, the JSON in the response is pretty-formatted. The default value is `false`. |
| `patch`                                                                                                                                                                          | `*string`                                                                                                                                                                        | :heavy_minus_sign:                                                                                                                                                               | Get the JSON Patch [RFC 6902 JavaScript Object Notation (JSON) Patch](https://www.rfc-editor.org/rfc/rfc6902) to be applied.                                                     |
| `opts`                                                                                                                                                                           | [][operations.Option](../../models/operations/option.md)                                                                                                                         | :heavy_minus_sign:                                                                                                                                                               | The options for this request.                                                                                                                                                    |

### Response

**[*operations.ServiceConfigurationAPIResponse](../../models/operations/serviceconfigurationapiresponse.md), error**

### Errors

| Error Type            | Status Code           | Content Type          |
| --------------------- | --------------------- | --------------------- |
| apierrors.ResultError | 400, 401, 403         | application/json      |
| apierrors.ResultError | 500                   | application/json      |
| apierrors.APIError    | 4XX, 5XX              | \*/\*                 |