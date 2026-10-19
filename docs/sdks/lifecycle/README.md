# Lifecycle

## Overview

### Available Operations

* [GetAPILifecycleHealthcheck](#getapilifecyclehealthcheck) - Health Check

## GetAPILifecycleHealthcheck

Perform a health check of the server.


### Example Usage

<!-- UsageSnippet language="go" operationID="get_/api/lifecycle/healthcheck" method="get" path="/api/lifecycle/healthcheck" -->
```go
package main

import(
	"context"
	authlete "github.com/authlete/authlete-go-sdk"
	"github.com/authlete/authlete-go-sdk/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := authlete.New()

    res, err := s.Lifecycle.GetAPILifecycleHealthcheck(ctx, operations.GetAPILifecycleHealthcheckRequest{})
    if err != nil {
        log.Fatal(err)
    }
    if res.Res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                    | Type                                                                                                         | Required                                                                                                     | Description                                                                                                  |
| ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                                        | :heavy_check_mark:                                                                                           | The context to use for the request.                                                                          |
| `request`                                                                                                    | [operations.GetAPILifecycleHealthcheckRequest](../../models/operations/getapilifecyclehealthcheckrequest.md) | :heavy_check_mark:                                                                                           | The request object to use for the request.                                                                   |
| `opts`                                                                                                       | [][operations.Option](../../models/operations/option.md)                                                     | :heavy_minus_sign:                                                                                           | The options for this request.                                                                                |

### Response

**[*operations.GetAPILifecycleHealthcheckResponse](../../models/operations/getapilifecyclehealthcheckresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |