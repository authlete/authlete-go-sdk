# Lifecycle

## Overview

### Available Operations

* [Healthcheck](#healthcheck) - Health Check

## Healthcheck

Perform a health check of the server.


### Example Usage

<!-- UsageSnippet language="go" operationID="get_/api/lifecycle/healthcheck" method="get" path="/api/lifecycle/healthcheck" -->
```go
package main

import(
	"context"
	authlete "github.com/authlete/authlete-go-sdk"
	"log"
)

func main() {
    ctx := context.Background()

    s := authlete.New()

    res, err := s.Lifecycle.Healthcheck(ctx, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.Res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                | Type                                                                     | Required                                                                 | Description                                                              |
| ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ |
| `ctx`                                                                    | [context.Context](https://pkg.go.dev/context#Context)                    | :heavy_check_mark:                                                       | The context to use for the request.                                      |
| `extended`                                                               | `*bool`                                                                  | :heavy_minus_sign:                                                       | If `true`, perform extended health checks (e.g. database connectivity).<br/> |
| `opts`                                                                   | [][operations.Option](../../models/operations/option.md)                 | :heavy_minus_sign:                                                       | The options for this request.                                            |

### Response

**[*operations.GetAPILifecycleHealthcheckResponse](../../models/operations/getapilifecyclehealthcheckresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |