# GrantManagement

## Overview

### Available Operations

* [ProcessRequest](#processrequest) - Process Grant Management Request

## ProcessRequest

The API is for the implementation of the grant management endpoint which is
defined in "[Grant Management for OAuth 2.0](https://openid.net/specs/fapi-grant-management.html)".


### Example Usage

<!-- UsageSnippet language="go" operationID="grant_m_api" method="post" path="/api/{serviceId}/gm" -->
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

    res, err := s.GrantManagement.ProcessRequest(ctx, "<id>", components.GMRequest{
        AccessToken: authlete.Pointer("eyJhbGciOiJFUzI1NiJ9.eyJleHAiOjE1NTk4MTE3NTAsImlzcyI6IjU3Mjk3NDA4ODY3In0K.csmdholMVcmjqHe59YWgLGNvm7I5Whp4phQCoGxyrlRGMnTgsfxtwyxBgMXQqEPD5q5k9FaEWNk37K8uAtSwrA"),
        GmAction: components.GrantManagementActionRevoke.ToPointer(),
        GrantID: authlete.Pointer("57297408867"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.GMResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                    | Type                                                         | Required                                                     | Description                                                  |
| ------------------------------------------------------------ | ------------------------------------------------------------ | ------------------------------------------------------------ | ------------------------------------------------------------ |
| `ctx`                                                        | [context.Context](https://pkg.go.dev/context#Context)        | :heavy_check_mark:                                           | The context to use for the request.                          |
| `serviceID`                                                  | `string`                                                     | :heavy_check_mark:                                           | A service ID.                                                |
| `body`                                                       | [components.GMRequest](../../models/components/gmrequest.md) | :heavy_check_mark:                                           | N/A                                                          |
| `opts`                                                       | [][operations.Option](../../models/operations/option.md)     | :heavy_minus_sign:                                           | The options for this request.                                |

### Response

**[*operations.GrantMAPIResponse](../../models/operations/grantmapiresponse.md), error**

### Errors

| Error Type            | Status Code           | Content Type          |
| --------------------- | --------------------- | --------------------- |
| apierrors.ResultError | 400, 401, 403         | application/json      |
| apierrors.ResultError | 500                   | application/json      |
| apierrors.APIError    | 4XX, 5XX              | \*/\*                 |