# Jwks

## Overview

### Available Operations

* [Get](#get) - Get JWK Set

## Get

This API gathers JWK Set information for a service so that its client applications can verify
signatures by the service and encrypt their requests to the service.

This API is supposed to be called from within the implementation of the jwk set endpoint of the
service where the service that supports OpenID Connect must expose its JWK Set information so that
client applications can verify signatures by the service and encrypt their requests to the service.
The URI of the endpoint can be found as the value of `jwks_uri` in [OpenID Provider Metadata](https://openid.net/specs/openid-connect-discovery-1_0.html#ProviderMetadata)
if the service supports [OpenID Connect Discovery 1.0](https://openid.net/specs/openid-connect-discovery-1_0.html).


### Example Usage

<!-- UsageSnippet language="go" operationID="service_jwks_get_api" method="get" path="/api/{serviceId}/service/jwks/get" -->
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

    res, err := s.Jwks.Get(ctx, "<id>", nil, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.ServiceJwksGetResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                                                                         | Type                                                                                                                                                                                                              | Required                                                                                                                                                                                                          | Description                                                                                                                                                                                                       |
| ----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                                                                             | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                                                             | :heavy_check_mark:                                                                                                                                                                                                | The context to use for the request.                                                                                                                                                                               |
| `serviceID`                                                                                                                                                                                                       | `string`                                                                                                                                                                                                          | :heavy_check_mark:                                                                                                                                                                                                | A service ID.                                                                                                                                                                                                     |
| `includePrivateKeys`                                                                                                                                                                                              | `*bool`                                                                                                                                                                                                           | :heavy_minus_sign:                                                                                                                                                                                                | The boolean value that indicates whether the response should include the private keys associated with the service or not. If `true`, the private keys are included in the response. The default value is `false`. |
| `pretty`                                                                                                                                                                                                          | `*bool`                                                                                                                                                                                                           | :heavy_minus_sign:                                                                                                                                                                                                | This boolean value indicates whether the JSON in the response should be formatted or not. If `true`, the JSON in the response is pretty-formatted. The default value is `false`.                                  |
| `opts`                                                                                                                                                                                                            | [][operations.Option](../../models/operations/option.md)                                                                                                                                                          | :heavy_minus_sign:                                                                                                                                                                                                | The options for this request.                                                                                                                                                                                     |

### Response

**[*operations.ServiceJwksGetAPIResponse](../../models/operations/servicejwksgetapiresponse.md), error**

### Errors

| Error Type            | Status Code           | Content Type          |
| --------------------- | --------------------- | --------------------- |
| apierrors.ResultError | 400, 401, 403         | application/json      |
| apierrors.ResultError | 500                   | application/json      |
| apierrors.APIError    | 4XX, 5XX              | \*/\*                 |