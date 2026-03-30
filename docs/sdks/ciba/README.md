# Ciba

## Overview

### Available Operations

* [ProcessAuthentication](#processauthentication) - Process Backchannel Authentication Request
* [Issue](#issue) - Issue Backchannel Authentication Response
* [Fail](#fail) - Fail Backchannel Authentication Request
* [Complete](#complete) - Complete Backchannel Authentication

## ProcessAuthentication

This API parses request parameters of a [backchannel authentication request](https://openid.net/specs/openid-client-initiated-backchannel-authentication-core-1_0.html#auth_request)
and returns necessary data for the authorization server implementation to process the backchannel
authentication request further.


### Example Usage

<!-- UsageSnippet language="go" operationID="backchannel_authentication_api" method="post" path="/api/{serviceId}/backchannel/authentication" -->
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

    res, err := s.Ciba.ProcessAuthentication(ctx, "<id>", components.BackchannelAuthenticationRequest{
        Parameters: "login_hint=john&scope=openid&client_notification_token=my-client-notification-token&user_code=my-user-code",
        ClientID: authlete.Pointer("26862190133482"),
        ClientSecret: authlete.Pointer("8J9pAEX6IQw7lYtYGsc_s9N4jlEz_DfkoCHIswJjFjfgKZX-nC4EvKtaHXcP9mHBfS7IU4jytjSZZpaK9UJ77A"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.BackchannelAuthenticationResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                  | Type                                                                                                       | Required                                                                                                   | Description                                                                                                |
| ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                      | :heavy_check_mark:                                                                                         | The context to use for the request.                                                                        |
| `serviceID`                                                                                                | `string`                                                                                                   | :heavy_check_mark:                                                                                         | A service ID.                                                                                              |
| `body`                                                                                                     | [components.BackchannelAuthenticationRequest](../../models/components/backchannelauthenticationrequest.md) | :heavy_check_mark:                                                                                         | N/A                                                                                                        |
| `opts`                                                                                                     | [][operations.Option](../../models/operations/option.md)                                                   | :heavy_minus_sign:                                                                                         | The options for this request.                                                                              |

### Response

**[*operations.BackchannelAuthenticationAPIResponse](../../models/operations/backchannelauthenticationapiresponse.md), error**

### Errors

| Error Type            | Status Code           | Content Type          |
| --------------------- | --------------------- | --------------------- |
| apierrors.ResultError | 400, 401, 403         | application/json      |
| apierrors.ResultError | 500                   | application/json      |
| apierrors.APIError    | 4XX, 5XX              | \*/\*                 |

## Issue

This API prepares JSON that contains an `auth_req_id`. The JSON should be used as the response body
of the response which is returned to the client from the [backchannel authentication endpoint](https://openid.net/specs/openid-client-initiated-backchannel-authentication-core-1_0.html#auth_backchannel_endpoint)


### Example Usage

<!-- UsageSnippet language="go" operationID="backchannel_authentication_issue_api" method="post" path="/api/{serviceId}/backchannel/authentication/issue" -->
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

    res, err := s.Ciba.Issue(ctx, "<id>", components.BackchannelAuthenticationIssueRequest{
        Ticket: "NFIHGx_btVrWmtAD093D-87JxvT4DAtuijEkLVHbS4Q",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.BackchannelAuthenticationIssueResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                            | Type                                                                                                                 | Required                                                                                                             | Description                                                                                                          |
| -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                                | :heavy_check_mark:                                                                                                   | The context to use for the request.                                                                                  |
| `serviceID`                                                                                                          | `string`                                                                                                             | :heavy_check_mark:                                                                                                   | A service ID.                                                                                                        |
| `body`                                                                                                               | [components.BackchannelAuthenticationIssueRequest](../../models/components/backchannelauthenticationissuerequest.md) | :heavy_check_mark:                                                                                                   | N/A                                                                                                                  |
| `opts`                                                                                                               | [][operations.Option](../../models/operations/option.md)                                                             | :heavy_minus_sign:                                                                                                   | The options for this request.                                                                                        |

### Response

**[*operations.BackchannelAuthenticationIssueAPIResponse](../../models/operations/backchannelauthenticationissueapiresponse.md), error**

### Errors

| Error Type            | Status Code           | Content Type          |
| --------------------- | --------------------- | --------------------- |
| apierrors.ResultError | 400, 401, 403         | application/json      |
| apierrors.ResultError | 500                   | application/json      |
| apierrors.APIError    | 4XX, 5XX              | \*/\*                 |

## Fail

The API prepares JSON that contains an error. The JSON should be used as the response body of the
response which is returned to the client from the [backchannel authentication endpoint](https://openid.net/specs/openid-client-initiated-backchannel-authentication-core-1_0.html#auth_backchannel_endpoint).


### Example Usage

<!-- UsageSnippet language="go" operationID="backchannel_authentication_fail_api" method="post" path="/api/{serviceId}/backchannel/authentication/fail" -->
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

    res, err := s.Ciba.Fail(ctx, "<id>", components.BackchannelAuthenticationFailRequest{
        Ticket: "<value>",
        Reason: components.BackchannelAuthenticationFailRequestReasonMissingUserCode,
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.BackchannelAuthenticationFailResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                          | Type                                                                                                               | Required                                                                                                           | Description                                                                                                        |
| ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                                              | :heavy_check_mark:                                                                                                 | The context to use for the request.                                                                                |
| `serviceID`                                                                                                        | `string`                                                                                                           | :heavy_check_mark:                                                                                                 | A service ID.                                                                                                      |
| `body`                                                                                                             | [components.BackchannelAuthenticationFailRequest](../../models/components/backchannelauthenticationfailrequest.md) | :heavy_check_mark:                                                                                                 | N/A                                                                                                                |
| `opts`                                                                                                             | [][operations.Option](../../models/operations/option.md)                                                           | :heavy_minus_sign:                                                                                                 | The options for this request.                                                                                      |

### Response

**[*operations.BackchannelAuthenticationFailAPIResponse](../../models/operations/backchannelauthenticationfailapiresponse.md), error**

### Errors

| Error Type            | Status Code           | Content Type          |
| --------------------- | --------------------- | --------------------- |
| apierrors.ResultError | 400, 401, 403         | application/json      |
| apierrors.ResultError | 500                   | application/json      |
| apierrors.APIError    | 4XX, 5XX              | \*/\*                 |

## Complete

This API returns information about what action the authorization server should take after it receives
the result of end-user's decision about whether the end-user has approved or rejected a client application's
request on the authentication device.


### Example Usage

<!-- UsageSnippet language="go" operationID="backchannel_authentication_complete_api" method="post" path="/api/{serviceId}/backchannel/authentication/complete" -->
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

    res, err := s.Ciba.Complete(ctx, "<id>", components.BackchannelAuthenticationCompleteRequest{
        Ticket: "NFIHGx_btVrWmtAD093D-87JxvT4DAtuijEkLVHbS4Q",
        Result: components.BackchannelAuthenticationCompleteRequestResultAuthorized,
        Subject: "john",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.BackchannelAuthenticationCompleteResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                  | Type                                                                                                                       | Required                                                                                                                   | Description                                                                                                                |
| -------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                                      | :heavy_check_mark:                                                                                                         | The context to use for the request.                                                                                        |
| `serviceID`                                                                                                                | `string`                                                                                                                   | :heavy_check_mark:                                                                                                         | A service ID.                                                                                                              |
| `body`                                                                                                                     | [components.BackchannelAuthenticationCompleteRequest](../../models/components/backchannelauthenticationcompleterequest.md) | :heavy_check_mark:                                                                                                         | N/A                                                                                                                        |
| `opts`                                                                                                                     | [][operations.Option](../../models/operations/option.md)                                                                   | :heavy_minus_sign:                                                                                                         | The options for this request.                                                                                              |

### Response

**[*operations.BackchannelAuthenticationCompleteAPIResponse](../../models/operations/backchannelauthenticationcompleteapiresponse.md), error**

### Errors

| Error Type            | Status Code           | Content Type          |
| --------------------- | --------------------- | --------------------- |
| apierrors.ResultError | 400, 401, 403         | application/json      |
| apierrors.ResultError | 500                   | application/json      |
| apierrors.APIError    | 4XX, 5XX              | \*/\*                 |