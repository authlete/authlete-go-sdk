# VerifiableCredentials

## Overview

### Available Operations

* [GetMetadata](#getmetadata) - Get Verifiable Credential Issuer Metadata
* [GetJwtIssuer](#getjwtissuer) - Get JWT Issuer Information
* [GetJwks](#getjwks) - Get JSON Web Key Set
* [CreateOffer](#createoffer) - Create Credential Offer
* [GetOfferInfo](#getofferinfo) - Get Credential Offer Information
* [Parse](#parse) - Parse Single Credential
* [Issue](#issue) - Issue Single Credential
* [BatchParse](#batchparse) - Parse Batch Credentials
* [BatchIssue](#batchissue) - Issue Batch Credentials
* [DeferredParse](#deferredparse) - Parse Deferred Credential
* [DeferredIssue](#deferredissue) - Issue Deferred Credential

## GetMetadata

Get verifiable credential issuer metadata

### Example Usage

<!-- UsageSnippet language="go" operationID="vci_metadata_api" method="post" path="/api/{serviceId}/vci/metadata" -->
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

    res, err := s.VerifiableCredentials.GetMetadata(ctx, "<id>", components.VciMetadataRequest{
        Pretty: true,
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.VciMetadataResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                      | Type                                                                           | Required                                                                       | Description                                                                    |
| ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ |
| `ctx`                                                                          | [context.Context](https://pkg.go.dev/context#Context)                          | :heavy_check_mark:                                                             | The context to use for the request.                                            |
| `serviceID`                                                                    | `string`                                                                       | :heavy_check_mark:                                                             | A service ID.                                                                  |
| `body`                                                                         | [components.VciMetadataRequest](../../models/components/vcimetadatarequest.md) | :heavy_check_mark:                                                             | N/A                                                                            |
| `opts`                                                                         | [][operations.Option](../../models/operations/option.md)                       | :heavy_minus_sign:                                                             | The options for this request.                                                  |

### Response

**[*operations.VciMetadataAPIResponse](../../models/operations/vcimetadataapiresponse.md), error**

### Errors

| Error Type            | Status Code           | Content Type          |
| --------------------- | --------------------- | --------------------- |
| apierrors.ResultError | 400, 401, 403         | application/json      |
| apierrors.ResultError | 500                   | application/json      |
| apierrors.APIError    | 4XX, 5XX              | \*/\*                 |

## GetJwtIssuer

Get JWT issuer information for VCI

### Example Usage

<!-- UsageSnippet language="go" operationID="vci_jwtissuer_api" method="post" path="/api/{serviceId}/vci/jwtissuer" -->
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

    res, err := s.VerifiableCredentials.GetJwtIssuer(ctx, "<id>", components.VciJwtissuerRequest{
        Pretty: true,
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.VciJwtissuerResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                        | Type                                                                             | Required                                                                         | Description                                                                      |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `ctx`                                                                            | [context.Context](https://pkg.go.dev/context#Context)                            | :heavy_check_mark:                                                               | The context to use for the request.                                              |
| `serviceID`                                                                      | `string`                                                                         | :heavy_check_mark:                                                               | A service ID.                                                                    |
| `body`                                                                           | [components.VciJwtissuerRequest](../../models/components/vcijwtissuerrequest.md) | :heavy_check_mark:                                                               | N/A                                                                              |
| `opts`                                                                           | [][operations.Option](../../models/operations/option.md)                         | :heavy_minus_sign:                                                               | The options for this request.                                                    |

### Response

**[*operations.VciJwtissuerAPIResponse](../../models/operations/vcijwtissuerapiresponse.md), error**

### Errors

| Error Type            | Status Code           | Content Type          |
| --------------------- | --------------------- | --------------------- |
| apierrors.ResultError | 400, 401, 403         | application/json      |
| apierrors.ResultError | 500                   | application/json      |
| apierrors.APIError    | 4XX, 5XX              | \*/\*                 |

## GetJwks

Get JSON Web Key Set for VCI

### Example Usage

<!-- UsageSnippet language="go" operationID="vci_jwks_api" method="post" path="/api/{serviceId}/vci/jwks" -->
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

    res, err := s.VerifiableCredentials.GetJwks(ctx, "<id>", components.VciJwksRequest{
        Pretty: false,
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.VciJwksResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                              | Type                                                                   | Required                                                               | Description                                                            |
| ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- |
| `ctx`                                                                  | [context.Context](https://pkg.go.dev/context#Context)                  | :heavy_check_mark:                                                     | The context to use for the request.                                    |
| `serviceID`                                                            | `string`                                                               | :heavy_check_mark:                                                     | A service ID.                                                          |
| `body`                                                                 | [components.VciJwksRequest](../../models/components/vcijwksrequest.md) | :heavy_check_mark:                                                     | N/A                                                                    |
| `opts`                                                                 | [][operations.Option](../../models/operations/option.md)               | :heavy_minus_sign:                                                     | The options for this request.                                          |

### Response

**[*operations.VciJwksAPIResponse](../../models/operations/vcijwksapiresponse.md), error**

### Errors

| Error Type            | Status Code           | Content Type          |
| --------------------- | --------------------- | --------------------- |
| apierrors.ResultError | 400, 401, 403         | application/json      |
| apierrors.ResultError | 500                   | application/json      |
| apierrors.APIError    | 4XX, 5XX              | \*/\*                 |

## CreateOffer

Create a verifiable credential offer

### Example Usage

<!-- UsageSnippet language="go" operationID="vci_offer_create_api" method="post" path="/api/{serviceId}/vci/offer/create" -->
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

    res, err := s.VerifiableCredentials.CreateOffer(ctx, "<id>", components.VciOfferCreateRequest{})
    if err != nil {
        log.Fatal(err)
    }
    if res.VciOfferCreateResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `serviceID`                                                                          | `string`                                                                             | :heavy_check_mark:                                                                   | A service ID.                                                                        |
| `body`                                                                               | [components.VciOfferCreateRequest](../../models/components/vcioffercreaterequest.md) | :heavy_check_mark:                                                                   | N/A                                                                                  |
| `opts`                                                                               | [][operations.Option](../../models/operations/option.md)                             | :heavy_minus_sign:                                                                   | The options for this request.                                                        |

### Response

**[*operations.VciOfferCreateAPIResponse](../../models/operations/vcioffercreateapiresponse.md), error**

### Errors

| Error Type            | Status Code           | Content Type          |
| --------------------- | --------------------- | --------------------- |
| apierrors.ResultError | 400, 401, 403         | application/json      |
| apierrors.ResultError | 500                   | application/json      |
| apierrors.APIError    | 4XX, 5XX              | \*/\*                 |

## GetOfferInfo

Get information about a verifiable credential offer

### Example Usage

<!-- UsageSnippet language="go" operationID="vci_offer_info_api" method="post" path="/api/{serviceId}/vci/offer/info" -->
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

    res, err := s.VerifiableCredentials.GetOfferInfo(ctx, "<id>", components.VciOfferInfoRequest{})
    if err != nil {
        log.Fatal(err)
    }
    if res.VciOfferInfoResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                        | Type                                                                             | Required                                                                         | Description                                                                      |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `ctx`                                                                            | [context.Context](https://pkg.go.dev/context#Context)                            | :heavy_check_mark:                                                               | The context to use for the request.                                              |
| `serviceID`                                                                      | `string`                                                                         | :heavy_check_mark:                                                               | A service ID.                                                                    |
| `body`                                                                           | [components.VciOfferInfoRequest](../../models/components/vciofferinforequest.md) | :heavy_check_mark:                                                               | N/A                                                                              |
| `opts`                                                                           | [][operations.Option](../../models/operations/option.md)                         | :heavy_minus_sign:                                                               | The options for this request.                                                    |

### Response

**[*operations.VciOfferInfoAPIResponse](../../models/operations/vciofferinfoapiresponse.md), error**

### Errors

| Error Type            | Status Code           | Content Type          |
| --------------------- | --------------------- | --------------------- |
| apierrors.ResultError | 400, 401, 403         | application/json      |
| apierrors.ResultError | 500                   | application/json      |
| apierrors.APIError    | 4XX, 5XX              | \*/\*                 |

## Parse

Parse a single verifiable credential

### Example Usage

<!-- UsageSnippet language="go" operationID="vci_single_parse_api" method="post" path="/api/{serviceId}/vci/single/parse" -->
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

    res, err := s.VerifiableCredentials.Parse(ctx, "<id>", components.VciSingleParseRequest{})
    if err != nil {
        log.Fatal(err)
    }
    if res.VciSingleParseResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `serviceID`                                                                          | `string`                                                                             | :heavy_check_mark:                                                                   | A service ID.                                                                        |
| `body`                                                                               | [components.VciSingleParseRequest](../../models/components/vcisingleparserequest.md) | :heavy_check_mark:                                                                   | N/A                                                                                  |
| `opts`                                                                               | [][operations.Option](../../models/operations/option.md)                             | :heavy_minus_sign:                                                                   | The options for this request.                                                        |

### Response

**[*operations.VciSingleParseAPIResponse](../../models/operations/vcisingleparseapiresponse.md), error**

### Errors

| Error Type            | Status Code           | Content Type          |
| --------------------- | --------------------- | --------------------- |
| apierrors.ResultError | 400, 401, 403         | application/json      |
| apierrors.ResultError | 500                   | application/json      |
| apierrors.APIError    | 4XX, 5XX              | \*/\*                 |

## Issue

Issue a single verifiable credential

### Example Usage

<!-- UsageSnippet language="go" operationID="vci_single_issue_api" method="post" path="/api/{serviceId}/vci/single/issue" -->
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

    res, err := s.VerifiableCredentials.Issue(ctx, "<id>", components.VciSingleIssueRequest{})
    if err != nil {
        log.Fatal(err)
    }
    if res.VciSingleIssueResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `serviceID`                                                                          | `string`                                                                             | :heavy_check_mark:                                                                   | A service ID.                                                                        |
| `body`                                                                               | [components.VciSingleIssueRequest](../../models/components/vcisingleissuerequest.md) | :heavy_check_mark:                                                                   | N/A                                                                                  |
| `opts`                                                                               | [][operations.Option](../../models/operations/option.md)                             | :heavy_minus_sign:                                                                   | The options for this request.                                                        |

### Response

**[*operations.VciSingleIssueAPIResponse](../../models/operations/vcisingleissueapiresponse.md), error**

### Errors

| Error Type            | Status Code           | Content Type          |
| --------------------- | --------------------- | --------------------- |
| apierrors.ResultError | 400, 401, 403         | application/json      |
| apierrors.ResultError | 500                   | application/json      |
| apierrors.APIError    | 4XX, 5XX              | \*/\*                 |

## BatchParse

Parse multiple verifiable credentials in batch

### Example Usage

<!-- UsageSnippet language="go" operationID="vci_batch_parse_api" method="post" path="/api/{serviceId}/vci/batch/parse" -->
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

    res, err := s.VerifiableCredentials.BatchParse(ctx, "<id>", components.VciBatchParseRequest{})
    if err != nil {
        log.Fatal(err)
    }
    if res.VciBatchParseResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `ctx`                                                                              | [context.Context](https://pkg.go.dev/context#Context)                              | :heavy_check_mark:                                                                 | The context to use for the request.                                                |
| `serviceID`                                                                        | `string`                                                                           | :heavy_check_mark:                                                                 | A service ID.                                                                      |
| `body`                                                                             | [components.VciBatchParseRequest](../../models/components/vcibatchparserequest.md) | :heavy_check_mark:                                                                 | N/A                                                                                |
| `opts`                                                                             | [][operations.Option](../../models/operations/option.md)                           | :heavy_minus_sign:                                                                 | The options for this request.                                                      |

### Response

**[*operations.VciBatchParseAPIResponse](../../models/operations/vcibatchparseapiresponse.md), error**

### Errors

| Error Type            | Status Code           | Content Type          |
| --------------------- | --------------------- | --------------------- |
| apierrors.ResultError | 400, 401, 403         | application/json      |
| apierrors.ResultError | 500                   | application/json      |
| apierrors.APIError    | 4XX, 5XX              | \*/\*                 |

## BatchIssue

Issue multiple verifiable credentials in batch

### Example Usage

<!-- UsageSnippet language="go" operationID="vci_batch_issue_api" method="post" path="/api/{serviceId}/vci/batch/issue" -->
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

    res, err := s.VerifiableCredentials.BatchIssue(ctx, "<id>", components.VciBatchIssueRequest{})
    if err != nil {
        log.Fatal(err)
    }
    if res.VciBatchIssueResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `ctx`                                                                              | [context.Context](https://pkg.go.dev/context#Context)                              | :heavy_check_mark:                                                                 | The context to use for the request.                                                |
| `serviceID`                                                                        | `string`                                                                           | :heavy_check_mark:                                                                 | A service ID.                                                                      |
| `body`                                                                             | [components.VciBatchIssueRequest](../../models/components/vcibatchissuerequest.md) | :heavy_check_mark:                                                                 | N/A                                                                                |
| `opts`                                                                             | [][operations.Option](../../models/operations/option.md)                           | :heavy_minus_sign:                                                                 | The options for this request.                                                      |

### Response

**[*operations.VciBatchIssueAPIResponse](../../models/operations/vcibatchissueapiresponse.md), error**

### Errors

| Error Type            | Status Code           | Content Type          |
| --------------------- | --------------------- | --------------------- |
| apierrors.ResultError | 400, 401, 403         | application/json      |
| apierrors.ResultError | 500                   | application/json      |
| apierrors.APIError    | 4XX, 5XX              | \*/\*                 |

## DeferredParse

Parse a deferred verifiable credential

### Example Usage

<!-- UsageSnippet language="go" operationID="vci_deferred_parse_api" method="post" path="/api/{serviceId}/vci/deferred/parse" -->
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

    res, err := s.VerifiableCredentials.DeferredParse(ctx, "<id>", components.VciDeferredParseRequest{})
    if err != nil {
        log.Fatal(err)
    }
    if res.VciDeferredParseResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `serviceID`                                                                              | `string`                                                                                 | :heavy_check_mark:                                                                       | A service ID.                                                                            |
| `body`                                                                                   | [components.VciDeferredParseRequest](../../models/components/vcideferredparserequest.md) | :heavy_check_mark:                                                                       | N/A                                                                                      |
| `opts`                                                                                   | [][operations.Option](../../models/operations/option.md)                                 | :heavy_minus_sign:                                                                       | The options for this request.                                                            |

### Response

**[*operations.VciDeferredParseAPIResponse](../../models/operations/vcideferredparseapiresponse.md), error**

### Errors

| Error Type            | Status Code           | Content Type          |
| --------------------- | --------------------- | --------------------- |
| apierrors.ResultError | 400, 401, 403         | application/json      |
| apierrors.ResultError | 500                   | application/json      |
| apierrors.APIError    | 4XX, 5XX              | \*/\*                 |

## DeferredIssue

Issue a deferred verifiable credential

### Example Usage

<!-- UsageSnippet language="go" operationID="vci_deferred_issue_api" method="post" path="/api/{serviceId}/vci/deferred/issue" -->
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

    res, err := s.VerifiableCredentials.DeferredIssue(ctx, "<id>", components.VciDeferredIssueRequest{})
    if err != nil {
        log.Fatal(err)
    }
    if res.VciDeferredIssueResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `serviceID`                                                                              | `string`                                                                                 | :heavy_check_mark:                                                                       | A service ID.                                                                            |
| `body`                                                                                   | [components.VciDeferredIssueRequest](../../models/components/vcideferredissuerequest.md) | :heavy_check_mark:                                                                       | N/A                                                                                      |
| `opts`                                                                                   | [][operations.Option](../../models/operations/option.md)                                 | :heavy_minus_sign:                                                                       | The options for this request.                                                            |

### Response

**[*operations.VciDeferredIssueAPIResponse](../../models/operations/vcideferredissueapiresponse.md), error**

### Errors

| Error Type            | Status Code           | Content Type          |
| --------------------- | --------------------- | --------------------- |
| apierrors.ResultError | 400, 401, 403         | application/json      |
| apierrors.ResultError | 500                   | application/json      |
| apierrors.APIError    | 4XX, 5XX              | \*/\*                 |