# ClientManagement

## Overview

API endpoints for managing OAuth clients, including creation, update, and deletion of clients.

### Available Operations

* [ClientCursorListAPI](#clientcursorlistapi) - List Clients (Cursor Pagination)
* [ClientSearchAPI](#clientsearchapi) - Search Clients

## ClientCursorListAPI

Retrieve the clients in the requested service via cursor based iteration.


### Example Usage

<!-- UsageSnippet language="go" operationID="client_cursor_list_api" method="get" path="/api/{serviceId}/client" -->
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

    res, err := s.ClientManagement.ClientCursorListAPI(ctx, 255950, nil, nil, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.OneOf != nil {
        for {
            // handle items

            res, err = res.Next()

            if err != nil {
                // handle error
            }

            if res == nil {
                break
            }
        }
    }
}
```

### Parameters

| Parameter                                                                                                                                                                                                                                                   | Type                                                                                                                                                                                                                                                        | Required                                                                                                                                                                                                                                                    | Description                                                                                                                                                                                                                                                 |
| ----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                                                                                                                       | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                                                                                                       | :heavy_check_mark:                                                                                                                                                                                                                                          | The context to use for the request.                                                                                                                                                                                                                         |
| `serviceID`                                                                                                                                                                                                                                                 | `int64`                                                                                                                                                                                                                                                     | :heavy_check_mark:                                                                                                                                                                                                                                          | The ID of the service to retrieve a list of clients from.                                                                                                                                                                                                   |
| `cursor`                                                                                                                                                                                                                                                    | `*string`                                                                                                                                                                                                                                                   | :heavy_minus_sign:                                                                                                                                                                                                                                          | This should be set to the value of `nextCursor` from a previous call to this endpoint. Which will retrieve the next set of clients available and receive the next available set. If this is not provided the list operation will start from the beginning.<br/> |
| `limit`                                                                                                                                                                                                                                                     | `*int`                                                                                                                                                                                                                                                      | :heavy_minus_sign:                                                                                                                                                                                                                                          | The maximum amount of results to return from this request. This value will be clamped between 1 and the allowed maximum. The default will be used if no value is provided.<br/>                                                                             |
| `includeTotalCount`                                                                                                                                                                                                                                         | `*bool`                                                                                                                                                                                                                                                     | :heavy_minus_sign:                                                                                                                                                                                                                                          | Indicates whether the `totalCount` property should be calculated and included in the response, otherwise it will be unset. For some large datasets this can be expensive and is not always required to calculate.<br/>                                      |
| `opts`                                                                                                                                                                                                                                                      | [][operations.Option](../../models/operations/option.md)                                                                                                                                                                                                    | :heavy_minus_sign:                                                                                                                                                                                                                                          | The options for this request.                                                                                                                                                                                                                               |

### Response

**[*operations.ClientCursorListAPIResponse](../../models/operations/clientcursorlistapiresponse.md), error**

### Errors

| Error Type            | Status Code           | Content Type          |
| --------------------- | --------------------- | --------------------- |
| apierrors.ResultError | 400, 401, 403         | application/json      |
| apierrors.ResultError | 500                   | application/json      |
| apierrors.APIError    | 4XX, 5XX              | \*/\*                 |

## ClientSearchAPI

Search the clients in the requested service that match the condition provided in the `q` query parameter. Results are returned via cursor based iteration. To iterate through all clients without a search condition, use the `/api/&#123;serviceId&#125;/client` API instead.


### Example Usage

<!-- UsageSnippet language="go" operationID="client_search_api" method="get" path="/api/{serviceId}/client/search" -->
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

    res, err := s.ClientManagement.ClientSearchAPI(ctx, operations.ClientSearchAPIRequest{
        ServiceID: 581816,
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.OneOf != nil {
        for {
            // handle items

            res, err = res.Next()

            if err != nil {
                // handle error
            }

            if res == nil {
                break
            }
        }
    }
}
```

### Parameters

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |
| `request`                                                                              | [operations.ClientSearchAPIRequest](../../models/operations/clientsearchapirequest.md) | :heavy_check_mark:                                                                     | The request object to use for the request.                                             |
| `opts`                                                                                 | [][operations.Option](../../models/operations/option.md)                               | :heavy_minus_sign:                                                                     | The options for this request.                                                          |

### Response

**[*operations.ClientSearchAPIResponse](../../models/operations/clientsearchapiresponse.md), error**

### Errors

| Error Type            | Status Code           | Content Type          |
| --------------------- | --------------------- | --------------------- |
| apierrors.ResultError | 400, 401, 403         | application/json      |
| apierrors.ResultError | 429                   | application/json      |
| apierrors.ResultError | 500                   | application/json      |
| apierrors.APIError    | 4XX, 5XX              | \*/\*                 |