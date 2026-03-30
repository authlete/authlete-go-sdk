# Federation

## Overview

### Available Operations

* [Configuration](#configuration) - Process Entity Configuration Request
* [Registration](#registration) - Process Federation Registration Request

## Configuration

This API gathers the federation configuration about a service.
The authorization server implementation should
retrieve the value of the `action`
response parameter from the API response and take the following steps
according to the value.


### Example Usage

<!-- UsageSnippet language="go" operationID="federation_configuration_api" method="post" path="/api/{serviceId}/federation/configuration" -->
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

    res, err := s.Federation.Configuration(ctx, "<id>", nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.FederationConfigurationResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                             | Type                                                                                                                  | Required                                                                                                              | Description                                                                                                           |
| --------------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                 | [context.Context](https://pkg.go.dev/context#Context)                                                                 | :heavy_check_mark:                                                                                                    | The context to use for the request.                                                                                   |
| `serviceID`                                                                                                           | `string`                                                                                                              | :heavy_check_mark:                                                                                                    | A service ID.                                                                                                         |
| `body`                                                                                                                | [*operations.FederationConfigurationAPIRequestBody](../../models/operations/federationconfigurationapirequestbody.md) | :heavy_minus_sign:                                                                                                    | N/A                                                                                                                   |
| `opts`                                                                                                                | [][operations.Option](../../models/operations/option.md)                                                              | :heavy_minus_sign:                                                                                                    | The options for this request.                                                                                         |

### Response

**[*operations.FederationConfigurationAPIResponse](../../models/operations/federationconfigurationapiresponse.md), error**

### Errors

| Error Type            | Status Code           | Content Type          |
| --------------------- | --------------------- | --------------------- |
| apierrors.ResultError | 400, 401, 403         | application/json      |
| apierrors.ResultError | 500                   | application/json      |
| apierrors.APIError    | 4XX, 5XX              | \*/\*                 |

## Registration

The Authlete API is for implementations of the **federation registration
endpoint** that accepts "explicit client registration". Its details are
defined in [OpenID Connect Federation 1.0](https://openid.net/specs/openid-connect-federation-1_0.html).
The endpoint accepts `POST` requests whose `Content-Type`
is either of the following.
1. `application/entity-statement+jwt`- `application/trust-chain+json`
When the `Content-Type` of a request is
`application/entity-statement+jwt`, the content of the request is
the entity configuration of a relying party that is to be registered.
In this case, the implementation of the federation registration endpoint
should call Authlete's `/federation/registration` API with the
entity configuration set to the `entityConfiguration` request
parameter.
On the other hand, when the `Content-Type` of a request is
`application/trust-chain+json`, the content of the request is a
JSON array that contains entity statements in JWT format. The sequence
of the entity statements composes the trust chain of a relying party
that is to be registered. In this case, the implementation of the
federation registration endpoint should call Authlete's
`/federation/registration` API with the trust chain set to the
`trustChain` request parameter.


### Example Usage

<!-- UsageSnippet language="go" operationID="federation_registration_api" method="post" path="/api/{serviceId}/federation/registration" -->
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

    res, err := s.Federation.Registration(ctx, "<id>", components.FederationRegistrationRequest{})
    if err != nil {
        log.Fatal(err)
    }
    if res.FederationRegistrationResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                            | Type                                                                                                 | Required                                                                                             | Description                                                                                          |
| ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                | :heavy_check_mark:                                                                                   | The context to use for the request.                                                                  |
| `serviceID`                                                                                          | `string`                                                                                             | :heavy_check_mark:                                                                                   | A service ID.                                                                                        |
| `body`                                                                                               | [components.FederationRegistrationRequest](../../models/components/federationregistrationrequest.md) | :heavy_check_mark:                                                                                   | N/A                                                                                                  |
| `opts`                                                                                               | [][operations.Option](../../models/operations/option.md)                                             | :heavy_minus_sign:                                                                                   | The options for this request.                                                                        |

### Response

**[*operations.FederationRegistrationAPIResponse](../../models/operations/federationregistrationapiresponse.md), error**

### Errors

| Error Type            | Status Code           | Content Type          |
| --------------------- | --------------------- | --------------------- |
| apierrors.ResultError | 400, 401, 403         | application/json      |
| apierrors.ResultError | 500                   | application/json      |
| apierrors.APIError    | 4XX, 5XX              | \*/\*                 |