# DeviceFlow

## Overview

### Available Operations

* [Authorization](#authorization) - Process Device Authorization Request
* [Verification](#verification) - Process Device Verification Request
* [Complete](#complete) - Complete Device Authorization

## Authorization

This API parses request parameters of a [device authorization request](https://datatracker.ietf.org/doc/html/rfc8628#section-3.1)
and returns necessary data for the authorization server implementation to process the device authorization
request further.


### Example Usage

<!-- UsageSnippet language="go" operationID="device_authorization_api" method="post" path="/api/{serviceId}/device/authorization" -->
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

    res, err := s.DeviceFlow.Authorization(ctx, "<id>", components.DeviceAuthorizationRequest{
        Parameters: "client_id=26888344961664&scope=history.read",
        ClientID: authlete.Pointer("26888344961664"),
        ClientSecret: authlete.Pointer("SfnYOLkJdofrb_66mTd6q03_SDoDEUnpXtvqFaE4k6L6UcpZzbdVJi2GpBj48AvGeDDllwsTruC62WYqQ_LGog"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.DeviceAuthorizationResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `serviceID`                                                                                    | `string`                                                                                       | :heavy_check_mark:                                                                             | A service ID.                                                                                  |
| `body`                                                                                         | [components.DeviceAuthorizationRequest](../../models/components/deviceauthorizationrequest.md) | :heavy_check_mark:                                                                             | N/A                                                                                            |
| `opts`                                                                                         | [][operations.Option](../../models/operations/option.md)                                       | :heavy_minus_sign:                                                                             | The options for this request.                                                                  |

### Response

**[*operations.DeviceAuthorizationAPIResponse](../../models/operations/deviceauthorizationapiresponse.md), error**

### Errors

| Error Type            | Status Code           | Content Type          |
| --------------------- | --------------------- | --------------------- |
| apierrors.ResultError | 400, 401, 403         | application/json      |
| apierrors.ResultError | 500                   | application/json      |
| apierrors.APIError    | 4XX, 5XX              | \*/\*                 |

## Verification

The API returns information associated with a user code.


### Example Usage

<!-- UsageSnippet language="go" operationID="device_verification_api" method="post" path="/api/{serviceId}/device/verification" -->
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

    res, err := s.DeviceFlow.Verification(ctx, "<id>", components.DeviceVerificationRequest{
        UserCode: "XWWKPBWVXQ",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.DeviceVerificationResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `serviceID`                                                                                  | `string`                                                                                     | :heavy_check_mark:                                                                           | A service ID.                                                                                |
| `body`                                                                                       | [components.DeviceVerificationRequest](../../models/components/deviceverificationrequest.md) | :heavy_check_mark:                                                                           | N/A                                                                                          |
| `opts`                                                                                       | [][operations.Option](../../models/operations/option.md)                                     | :heavy_minus_sign:                                                                           | The options for this request.                                                                |

### Response

**[*operations.DeviceVerificationAPIResponse](../../models/operations/deviceverificationapiresponse.md), error**

### Errors

| Error Type            | Status Code           | Content Type          |
| --------------------- | --------------------- | --------------------- |
| apierrors.ResultError | 400, 401, 403         | application/json      |
| apierrors.ResultError | 500                   | application/json      |
| apierrors.APIError    | 4XX, 5XX              | \*/\*                 |

## Complete

This API returns information about what action the authorization server should take after it receives
the result of end-user's decision about whether the end-user has approved or rejected a client
application's request.


### Example Usage

<!-- UsageSnippet language="go" operationID="device_complete_api" method="post" path="/api/{serviceId}/device/complete" -->
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

    res, err := s.DeviceFlow.Complete(ctx, "<id>", components.DeviceCompleteRequest{
        UserCode: "XWWKPBWVXQ",
        Result: components.DeviceCompleteRequestResultAuthorized,
        Subject: "john",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.DeviceCompleteResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `serviceID`                                                                          | `string`                                                                             | :heavy_check_mark:                                                                   | A service ID.                                                                        |
| `body`                                                                               | [components.DeviceCompleteRequest](../../models/components/devicecompleterequest.md) | :heavy_check_mark:                                                                   | N/A                                                                                  |
| `opts`                                                                               | [][operations.Option](../../models/operations/option.md)                             | :heavy_minus_sign:                                                                   | The options for this request.                                                        |

### Response

**[*operations.DeviceCompleteAPIResponse](../../models/operations/devicecompleteapiresponse.md), error**

### Errors

| Error Type            | Status Code           | Content Type          |
| --------------------- | --------------------- | --------------------- |
| apierrors.ResultError | 400, 401, 403         | application/json      |
| apierrors.ResultError | 500                   | application/json      |
| apierrors.APIError    | 4XX, 5XX              | \*/\*                 |