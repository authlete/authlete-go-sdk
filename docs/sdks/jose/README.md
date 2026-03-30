# Jose

## Overview

### Available Operations

* [Verify](#verify) - Verify JOSE

## Verify

This API verifies a JOSE object.


### Example Usage

<!-- UsageSnippet language="go" operationID="jose_verify_api" method="post" path="/api/{serviceId}/jose/verify" -->
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

    res, err := s.Jose.Verify(ctx, "<id>", &components.JoseVerifyRequest{
        Jose: "eyJhbGciOiJFUzI1NiJ9.eyJleHAiOjE1NTk4MTE3NTAsImlzcyI6IjU3Mjk3NDA4ODY3In0K.csmdholMVcmjqHe59YWgLGNvm7I5Whp4phQCoGxyrlRGMnTgsfxtwyxBgMXQqEPD5q5k9FaEWNk37K8uAtSwrA",
        ClockSkew: authlete.Pointer[int](100),
        ClientIdentifier: authlete.Pointer("57297408867"),
        SignedByClient: authlete.Pointer(true),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.JoseVerifyResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                     | Type                                                                          | Required                                                                      | Description                                                                   |
| ----------------------------------------------------------------------------- | ----------------------------------------------------------------------------- | ----------------------------------------------------------------------------- | ----------------------------------------------------------------------------- |
| `ctx`                                                                         | [context.Context](https://pkg.go.dev/context#Context)                         | :heavy_check_mark:                                                            | The context to use for the request.                                           |
| `serviceID`                                                                   | `string`                                                                      | :heavy_check_mark:                                                            | A service ID.                                                                 |
| `body`                                                                        | [*components.JoseVerifyRequest](../../models/components/joseverifyrequest.md) | :heavy_minus_sign:                                                            | N/A                                                                           |
| `opts`                                                                        | [][operations.Option](../../models/operations/option.md)                      | :heavy_minus_sign:                                                            | The options for this request.                                                 |

### Response

**[*operations.JoseVerifyAPIResponse](../../models/operations/joseverifyapiresponse.md), error**

### Errors

| Error Type            | Status Code           | Content Type          |
| --------------------- | --------------------- | --------------------- |
| apierrors.ResultError | 400, 401, 403         | application/json      |
| apierrors.ResultError | 500                   | application/json      |
| apierrors.APIError    | 4XX, 5XX              | \*/\*                 |