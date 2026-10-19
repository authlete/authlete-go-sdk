# Audit

## Overview

### Available Operations

* [Get](#get) - Get Audit Logs
* [GetTypes](#gettypes) - Get Audit Log Event Types

## Get

Retrieve audit logs as a cursor-paginated list. Results can be filtered by time range, event type, and organization.


### Example Usage

<!-- UsageSnippet language="go" operationID="audit_entries_get_idp_api" method="get" path="/api/audit/entries" -->
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

    res, err := s.Audit.Get(ctx, operations.AuditEntriesGetIdpAPIRequest{
        Types: []string{
            "service.create",
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AuditEntriesGetResponse != nil {
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

| Parameter                                                                                          | Type                                                                                               | Required                                                                                           | Description                                                                                        |
| -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                              | :heavy_check_mark:                                                                                 | The context to use for the request.                                                                |
| `request`                                                                                          | [operations.AuditEntriesGetIdpAPIRequest](../../models/operations/auditentriesgetidpapirequest.md) | :heavy_check_mark:                                                                                 | The request object to use for the request.                                                         |
| `opts`                                                                                             | [][operations.Option](../../models/operations/option.md)                                           | :heavy_minus_sign:                                                                                 | The options for this request.                                                                      |

### Response

**[*operations.AuditEntriesGetIdpAPIResponse](../../models/operations/auditentriesgetidpapiresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.IdpError | 400, 401, 403      | application/json   |
| apierrors.IdpError | 500                | application/json   |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## GetTypes

Returns the list of available audit log event types for this environment.

Accessible with a user access token or an organization token that has the `view_audit_log` permission (e.g. a token created with the `admin` or `audit_reader` preset).

This endpoint is hosted on the Authlete IdP server (`https://login.authlete.com`), not on the regional API clusters.


### Example Usage

<!-- UsageSnippet language="go" operationID="audit_types_get_idp_api" method="get" path="/api/audit/types" -->
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

    res, err := s.Audit.GetTypes(ctx)
    if err != nil {
        log.Fatal(err)
    }
    if res.Strings != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.AuditTypesGetIdpAPIResponse](../../models/operations/audittypesgetidpapiresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.IdpError | 400, 401, 403      | application/json   |
| apierrors.IdpError | 500                | application/json   |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |