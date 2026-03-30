# Revocation

## Overview

### Available Operations

* [Process](#process) - Process Revocation Request

## Process

This API revokes access tokens and refresh tokens.


### Example Usage

<!-- UsageSnippet language="go" operationID="auth_revocation_api" method="post" path="/api/{serviceId}/auth/revocation" -->
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

    res, err := s.Revocation.Process(ctx, "<id>", components.RevocationRequest{
        Parameters: "VFGsNK-5sXiqterdaR7b5QbRX9VTwVCQB87jbr2_xAI&token_type_hint=access_token",
        ClientID: authlete.Pointer("26478243745571"),
        ClientSecret: authlete.Pointer("gXz97ISgLs4HuXwOZWch8GEmgL4YMvUJwu3er_kDVVGcA0UOhA9avLPbEmoeZdagi9yC_-tEiT2BdRyH9dbrQQ"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.RevocationResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                    | Type                                                                         | Required                                                                     | Description                                                                  |
| ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- |
| `ctx`                                                                        | [context.Context](https://pkg.go.dev/context#Context)                        | :heavy_check_mark:                                                           | The context to use for the request.                                          |
| `serviceID`                                                                  | `string`                                                                     | :heavy_check_mark:                                                           | A service ID.                                                                |
| `body`                                                                       | [components.RevocationRequest](../../models/components/revocationrequest.md) | :heavy_check_mark:                                                           | N/A                                                                          |
| `opts`                                                                       | [][operations.Option](../../models/operations/option.md)                     | :heavy_minus_sign:                                                           | The options for this request.                                                |

### Response

**[*operations.AuthRevocationAPIResponse](../../models/operations/authrevocationapiresponse.md), error**

### Errors

| Error Type            | Status Code           | Content Type          |
| --------------------- | --------------------- | --------------------- |
| apierrors.ResultError | 400, 401, 403         | application/json      |
| apierrors.ResultError | 500                   | application/json      |
| apierrors.APIError    | 4XX, 5XX              | \*/\*                 |