# authlete

Developer-friendly & type-safe Go SDK specifically catered to leverage *authlete* API.

[![Built by Speakeasy](https://img.shields.io/badge/Built_by-SPEAKEASY-374151?style=for-the-badge&labelColor=f3f4f6)](https://www.speakeasy.com/?utm_source=authlete&utm_campaign=go)
[![License: MIT](https://img.shields.io/badge/LICENSE_//_MIT-3b5bdb?style=for-the-badge&labelColor=eff6ff)](https://opensource.org/licenses/MIT)


<br /><br />
> [!IMPORTANT]
> This SDK is not yet ready for production use. To complete setup please follow the steps outlined in your [workspace](https://app.speakeasy.com/org/authlete/sdk-workspace). Delete this section before > publishing to a package manager.

<!-- Start Summary [summary] -->
## Summary

Authlete API: Welcome to the **Authlete API documentation**. Authlete is an **API-first service** where every aspect of the 
platform is configurable via API. This documentation will help you authenticate and integrate with Authlete to 
build powerful OAuth 2.0 and OpenID Connect servers.

At a high level, the Authlete API is grouped into two categories:

- **Management APIs**: Enable you to manage services and clients.
- **Runtime APIs**: Allow you to build your own Authorization Servers or Verifiable Credential (VC) issuers.

## 🌐 API Servers

Authlete is a global service with clusters available in multiple regions across the world:

- 🇺🇸 **US**: `https://us.authlete.com`
- 🇯🇵 **Japan**: `https://jp.authlete.com`
- 🇪🇺 **Europe**: `https://eu.authlete.com`
- 🇧🇷 **Brazil**: `https://br.authlete.com`

Our customers can host their data in the region that best meets their requirements.

## 🔑 Authentication

All API endpoints are secured using **Bearer token authentication**. You must include an access token in every request:

```
Authorization: Bearer YOUR_ACCESS_TOKEN
```

### Getting Your Access Token

Authlete supports two types of access tokens:

**Service Access Token** - Scoped to a single service (authorization server instance)

1. Log in to [Authlete Console](https://console.authlete.com)
2. Navigate to your service → **Settings** → **Access Tokens**
3. Click **Create Token** and select permissions (e.g., `service.read`, `client.write`)
4. Copy the generated token

**Organization Token** - Scoped to your entire organization

1. Log in to [Authlete Console](https://console.authlete.com)
2. Navigate to **Organization Settings** → **Access Tokens**
3. Click **Create Token** and select org-level permissions
4. Copy the generated token

> ⚠️ **Important Note**: Tokens inherit the permissions of the account that creates them. Service tokens can only 
> access their specific service, while organization tokens can access all services within your org.

### Token Security Best Practices

- **Never commit tokens to version control** - Store in environment variables or secure secret managers
- **Rotate regularly** - Generate new tokens periodically and revoke old ones
- **Scope appropriately** - Request only the permissions your application needs
- **Revoke unused tokens** - Delete tokens you're no longer using from the console

### Quick Test

Verify your token works with a simple API call:

```bash
curl -X GET https://us.authlete.com/api/service/get/list \
  -H "Authorization: Bearer YOUR_ACCESS_TOKEN"
```

## 🎓 Tutorials

If you're new to Authlete or want to see sample implementations, these resources will help you get started:

- [Getting Started with Authlete](https://www.authlete.com/developers/getting_started/)
- [From Sign-Up to the First API Request](https://www.authlete.com/developers/tutorial/signup/)

## 🛠 Contact Us

If you have any questions or need assistance, our team is here to help:

- [Contact Page](https://www.authlete.com/contact/)
<!-- End Summary [summary] -->

<!-- Start Table of Contents [toc] -->
## Table of Contents
<!-- $toc-max-depth=2 -->
* [authlete](#authlete)
  * [🌐 API Servers](#api-servers)
  * [🔑 Authentication](#authentication)
  * [🎓 Tutorials](#tutorials)
  * [🛠 Contact Us](#contact-us)
  * [SDK Installation](#sdk-installation)
  * [SDK Example Usage](#sdk-example-usage)
  * [Authentication](#authentication-1)
  * [Available Resources and Operations](#available-resources-and-operations)
  * [Retries](#retries)
  * [Error Handling](#error-handling)
  * [Server Selection](#server-selection)
  * [Custom HTTP Client](#custom-http-client)
* [Development](#development)
  * [Maturity](#maturity)
  * [Contributions](#contributions)

<!-- End Table of Contents [toc] -->

<!-- Start SDK Installation [installation] -->
## SDK Installation

To add the SDK as a dependency to your project:
```bash
go get github.com/authlete/authlete-go-sdk
```
<!-- End SDK Installation [installation] -->

<!-- Start SDK Example Usage [usage] -->
## SDK Example Usage

### Example

```go
package main

import (
	"context"
	authlete "github.com/authlete/authlete-go-sdk"
	"log"
	"os"
)

func main() {
	ctx := context.Background()

	s := authlete.New(
		authlete.WithSecurity(os.Getenv("AUTHLETE_BEARER")),
	)

	res, err := s.Service.Get(ctx, "<id>")
	if err != nil {
		log.Fatal(err)
	}
	if res.Service != nil {
		// handle response
	}
}

```
<!-- End SDK Example Usage [usage] -->

<!-- Start Authentication [security] -->
## Authentication

### Per-Client Security Schemes

This SDK supports the following security scheme globally:

| Name     | Type | Scheme      | Environment Variable |
| -------- | ---- | ----------- | -------------------- |
| `Bearer` | http | HTTP Bearer | `AUTHLETE_BEARER`    |

You can configure it using the `WithSecurity` option when initializing the SDK client instance. For example:
```go
package main

import (
	"context"
	authlete "github.com/authlete/authlete-go-sdk"
	"log"
	"os"
)

func main() {
	ctx := context.Background()

	s := authlete.New(
		authlete.WithSecurity(os.Getenv("AUTHLETE_BEARER")),
	)

	res, err := s.Service.Get(ctx, "<id>")
	if err != nil {
		log.Fatal(err)
	}
	if res.Service != nil {
		// handle response
	}
}

```
<!-- End Authentication [security] -->

<!-- Start Available Resources and Operations [operations] -->
## Available Resources and Operations

<details open>
<summary>Available methods</summary>

### [Authorization](docs/sdks/authorization/README.md)

* [ProcessRequest](docs/sdks/authorization/README.md#processrequest) - Process Authorization Request
* [Fail](docs/sdks/authorization/README.md#fail) - Fail Authorization Request
* [Issue](docs/sdks/authorization/README.md#issue) - Issue Authorization Response

#### [Authorization.Management](docs/sdks/authorizationmanagement/README.md)

* [GetTicketInfo](docs/sdks/authorizationmanagement/README.md#getticketinfo) - Get Ticket Information
* [UpdateTicket](docs/sdks/authorizationmanagement/README.md#updateticket) - Update Ticket Information

### [Ciba](docs/sdks/ciba/README.md)

* [ProcessAuthentication](docs/sdks/ciba/README.md#processauthentication) - Process Backchannel Authentication Request
* [Issue](docs/sdks/ciba/README.md#issue) - Issue Backchannel Authentication Response
* [Fail](docs/sdks/ciba/README.md#fail) - Fail Backchannel Authentication Request
* [Complete](docs/sdks/ciba/README.md#complete) - Complete Backchannel Authentication

### [Client](docs/sdks/client/README.md)

* [Get](docs/sdks/client/README.md#get) - Get Client
* [List](docs/sdks/client/README.md#list) - List Clients
* [Create](docs/sdks/client/README.md#create) - Create Client
* [Update](docs/sdks/client/README.md#update) - Update Client
* [UpdateForm](docs/sdks/client/README.md#updateform) - Update Client
* [Delete](docs/sdks/client/README.md#delete) - Delete Client ⚡

#### [Client.Management](docs/sdks/clientmanagement1/README.md)

* [UpdateLockFlag](docs/sdks/clientmanagement1/README.md#updatelockflag) - Update Client Lock
* [RefreshSecret](docs/sdks/clientmanagement1/README.md#refreshsecret) - Rotate Client Secret
* [UpdateSecret](docs/sdks/clientmanagement1/README.md#updatesecret) - Update Client Secret
* [ListAuthorizations](docs/sdks/clientmanagement1/README.md#listauthorizations) - Get Authorized Applications (by Subject)
* [UpdateAuthorizations](docs/sdks/clientmanagement1/README.md#updateauthorizations) - Update Client Tokens
* [DeleteAuthorizations](docs/sdks/clientmanagement1/README.md#deleteauthorizations) - Delete Client Tokens (by Subject)
* [GetGrantedScopes](docs/sdks/clientmanagement1/README.md#getgrantedscopes) - Get Granted Scopes (by Subject)
* [DeleteGrantedScopes](docs/sdks/clientmanagement1/README.md#deletegrantedscopes) - Delete Granted Scopes (by Subject)
* [GetRequestableScopes](docs/sdks/clientmanagement1/README.md#getrequestablescopes) - Get Requestable Scopes
* [UpdateRequestableScopes](docs/sdks/clientmanagement1/README.md#updaterequestablescopes) - Update Requestable Scopes
* [DeleteRequestableScopes](docs/sdks/clientmanagement1/README.md#deleterequestablescopes) - Delete Requestable Scopes

### [ClientAuthorizations](docs/sdks/clientauthorizations/README.md)

* [List](docs/sdks/clientauthorizations/README.md#list) - Get Authorized Applications

### [ClientManagement](docs/sdks/clientmanagement2/README.md)

* [ListAuthorizedApplications](docs/sdks/clientmanagement2/README.md#listauthorizedapplications) - Get Authorized Applications
* [RevokeClientTokens](docs/sdks/clientmanagement2/README.md#revokeclienttokens) - Delete Client Tokens
* [RevokeGrantedScopes](docs/sdks/clientmanagement2/README.md#revokegrantedscopes) - Delete Granted Scopes

### [Clients.GrantedScopes](docs/sdks/grantedscopes/README.md)

* [Get](docs/sdks/grantedscopes/README.md#get) - Get Granted Scopes
* [Create](docs/sdks/grantedscopes/README.md#create) - Get Granted Scopes

### [Clients.RequestableScopes](docs/sdks/requestablescopes/README.md)

* [Update](docs/sdks/requestablescopes/README.md#update) - Update Requestable Scopes

### [Clients.Tokens](docs/sdks/tokens/README.md)

* [Delete](docs/sdks/tokens/README.md#delete) - Delete Client Tokens

### [DeviceFlow](docs/sdks/deviceflow/README.md)

* [Authorization](docs/sdks/deviceflow/README.md#authorization) - Process Device Authorization Request
* [Verification](docs/sdks/deviceflow/README.md#verification) - Process Device Verification Request
* [Complete](docs/sdks/deviceflow/README.md#complete) - Complete Device Authorization

### [DynamicClientRegistration](docs/sdks/dynamicclientregistration/README.md)

* [Register](docs/sdks/dynamicclientregistration/README.md#register) - Register Client
* [Get](docs/sdks/dynamicclientregistration/README.md#get) - Get Client
* [Update](docs/sdks/dynamicclientregistration/README.md#update) - Update Client
* [Delete](docs/sdks/dynamicclientregistration/README.md#delete) - Delete Client

### [Federation](docs/sdks/federation/README.md)

* [Configuration](docs/sdks/federation/README.md#configuration) - Process Entity Configuration Request
* [Registration](docs/sdks/federation/README.md#registration) - Process Federation Registration Request

### [GrantManagement](docs/sdks/grantmanagement/README.md)

* [ProcessRequest](docs/sdks/grantmanagement/README.md#processrequest) - Process Grant Management Request

### [HardwareSecurityKeys](docs/sdks/hardwaresecuritykeys/README.md)

* [Create](docs/sdks/hardwaresecuritykeys/README.md#create) - Create Security Key
* [Delete](docs/sdks/hardwaresecuritykeys/README.md#delete) - Delete Security Key
* [Get](docs/sdks/hardwaresecuritykeys/README.md#get) - Get Security Key
* [List](docs/sdks/hardwaresecuritykeys/README.md#list) - List Security Keys

### [Introspection](docs/sdks/introspection/README.md)

* [Process](docs/sdks/introspection/README.md#process) - Process Introspection Request
* [StandardProcess](docs/sdks/introspection/README.md#standardprocess) - Process OAuth 2.0 Introspection Request

### [Jose](docs/sdks/jose/README.md)

* [Verify](docs/sdks/jose/README.md#verify) - Verify JOSE

### [Jwks](docs/sdks/jwks/README.md)

* [Get](docs/sdks/jwks/README.md#get) - Get JWK Set

### [Lifecycle](docs/sdks/lifecycle/README.md)

* [Healthcheck](docs/sdks/lifecycle/README.md#healthcheck) - Health Check

### [NativeSso](docs/sdks/nativesso/README.md)

* [Process](docs/sdks/nativesso/README.md#process) - Native SSO Processing
* [Logout](docs/sdks/nativesso/README.md#logout) - Native SSO Logout Processing

### [PushedAuthorization](docs/sdks/pushedauthorization/README.md)

* [Create](docs/sdks/pushedauthorization/README.md#create) - Process Pushed Authorization Request

### [Revocation](docs/sdks/revocation/README.md)

* [Process](docs/sdks/revocation/README.md#process) - Process Revocation Request

### [Service](docs/sdks/service/README.md)

* [Get](docs/sdks/service/README.md#get) - Get Service
* [List](docs/sdks/service/README.md#list) - List Services
* [Update](docs/sdks/service/README.md#update) - Update Service
* [Delete](docs/sdks/service/README.md#delete) - Delete Service ⚡
* [GetConfiguration](docs/sdks/service/README.md#getconfiguration) - Get Service Configuration

### [Token](docs/sdks/token/README.md)

* [Process](docs/sdks/token/README.md#process) - Process Token Request
* [Fail](docs/sdks/token/README.md#fail) - Fail Token Request
* [Issue](docs/sdks/token/README.md#issue) - Issue Token Response

#### [Token.Management](docs/sdks/tokenmanagement/README.md)

* [ReissueIDToken](docs/sdks/tokenmanagement/README.md#reissueidtoken) - Reissue ID Token
* [List](docs/sdks/tokenmanagement/README.md#list) - List Issued Tokens
* [Create](docs/sdks/tokenmanagement/README.md#create) - Create Access Token
* [Update](docs/sdks/tokenmanagement/README.md#update) - Update Access Token
* [Delete](docs/sdks/tokenmanagement/README.md#delete) - Delete Access Token
* [Revoke](docs/sdks/tokenmanagement/README.md#revoke) - Revoke Access Token

### [Userinfo](docs/sdks/userinfo/README.md)

* [Process](docs/sdks/userinfo/README.md#process) - Process UserInfo Request
* [Issue](docs/sdks/userinfo/README.md#issue) - Issue UserInfo Response

### [VerifiableCredentials](docs/sdks/verifiablecredentials/README.md)

* [GetMetadata](docs/sdks/verifiablecredentials/README.md#getmetadata) - Get Verifiable Credential Issuer Metadata
* [GetJwtIssuer](docs/sdks/verifiablecredentials/README.md#getjwtissuer) - Get JWT Issuer Information
* [GetJwks](docs/sdks/verifiablecredentials/README.md#getjwks) - Get JSON Web Key Set
* [CreateOffer](docs/sdks/verifiablecredentials/README.md#createoffer) - Create Credential Offer
* [GetOfferInfo](docs/sdks/verifiablecredentials/README.md#getofferinfo) - Get Credential Offer Information
* [Parse](docs/sdks/verifiablecredentials/README.md#parse) - Parse Single Credential
* [Issue](docs/sdks/verifiablecredentials/README.md#issue) - Issue Single Credential
* [BatchParse](docs/sdks/verifiablecredentials/README.md#batchparse) - Parse Batch Credentials
* [BatchIssue](docs/sdks/verifiablecredentials/README.md#batchissue) - Issue Batch Credentials
* [DeferredParse](docs/sdks/verifiablecredentials/README.md#deferredparse) - Parse Deferred Credential
* [DeferredIssue](docs/sdks/verifiablecredentials/README.md#deferredissue) - Issue Deferred Credential

</details>
<!-- End Available Resources and Operations [operations] -->

<!-- Start Retries [retries] -->
## Retries

Some of the endpoints in this SDK support retries. If you use the SDK without any configuration, it will fall back to the default retry strategy provided by the API. However, the default retry strategy can be overridden on a per-operation basis, or across the entire SDK.

To change the default retry strategy for a single API call, simply provide a `retry.Config` object to the call by using the `WithRetries` option:
```go
package main

import (
	"context"
	authlete "github.com/authlete/authlete-go-sdk"
	"github.com/authlete/authlete-go-sdk/retry"
	"log"
	"models/operations"
	"os"
)

func main() {
	ctx := context.Background()

	s := authlete.New(
		authlete.WithSecurity(os.Getenv("AUTHLETE_BEARER")),
	)

	res, err := s.Service.Get(ctx, "<id>", operations.WithRetries(
		retry.Config{
			Strategy: "backoff",
			Backoff: &retry.BackoffStrategy{
				InitialInterval: 1,
				MaxInterval:     50,
				Exponent:        1.1,
				MaxElapsedTime:  100,
			},
			RetryConnectionErrors: false,
		}))
	if err != nil {
		log.Fatal(err)
	}
	if res.Service != nil {
		// handle response
	}
}

```

If you'd like to override the default retry strategy for all operations that support retries, you can use the `WithRetryConfig` option at SDK initialization:
```go
package main

import (
	"context"
	authlete "github.com/authlete/authlete-go-sdk"
	"github.com/authlete/authlete-go-sdk/retry"
	"log"
	"os"
)

func main() {
	ctx := context.Background()

	s := authlete.New(
		authlete.WithRetryConfig(
			retry.Config{
				Strategy: "backoff",
				Backoff: &retry.BackoffStrategy{
					InitialInterval: 1,
					MaxInterval:     50,
					Exponent:        1.1,
					MaxElapsedTime:  100,
				},
				RetryConnectionErrors: false,
			}),
		authlete.WithSecurity(os.Getenv("AUTHLETE_BEARER")),
	)

	res, err := s.Service.Get(ctx, "<id>")
	if err != nil {
		log.Fatal(err)
	}
	if res.Service != nil {
		// handle response
	}
}

```
<!-- End Retries [retries] -->

<!-- Start Error Handling [errors] -->
## Error Handling

Handling errors in this SDK should largely match your expectations. All operations return a response object or an error, they will never return both.

By Default, an API error will return `apierrors.APIError`. When custom error responses are specified for an operation, the SDK may also return their associated error. You can refer to respective *Errors* tables in SDK docs for more details on possible error types for each operation.

For example, the `Get` function may return the following errors:

| Error Type            | Status Code   | Content Type     |
| --------------------- | ------------- | ---------------- |
| apierrors.ResultError | 400, 401, 403 | application/json |
| apierrors.ResultError | 500           | application/json |
| apierrors.APIError    | 4XX, 5XX      | \*/\*            |

### Example

```go
package main

import (
	"context"
	"errors"
	authlete "github.com/authlete/authlete-go-sdk"
	"github.com/authlete/authlete-go-sdk/models/apierrors"
	"log"
	"os"
)

func main() {
	ctx := context.Background()

	s := authlete.New(
		authlete.WithSecurity(os.Getenv("AUTHLETE_BEARER")),
	)

	res, err := s.Service.Get(ctx, "<id>")
	if err != nil {

		var e *apierrors.ResultError
		if errors.As(err, &e) {
			// handle error
			log.Fatal(e.Error())
		}

		var e *apierrors.ResultError
		if errors.As(err, &e) {
			// handle error
			log.Fatal(e.Error())
		}

		var e *apierrors.APIError
		if errors.As(err, &e) {
			// handle error
			log.Fatal(e.Error())
		}
	}
}

```
<!-- End Error Handling [errors] -->

<!-- Start Server Selection [server] -->
## Server Selection

### Select Server by Index

You can override the default server globally using the `WithServerIndex(serverIndex int)` option when initializing the SDK client instance. The selected server will then be used as the default on the operations that use it. This table lists the indexes associated with the available servers:

| #   | Server                    | Description         |
| --- | ------------------------- | ------------------- |
| 0   | `https://us.authlete.com` | 🇺🇸 US Cluster     |
| 1   | `https://jp.authlete.com` | 🇯🇵 Japan Cluster  |
| 2   | `https://eu.authlete.com` | 🇪🇺 Europe Cluster |
| 3   | `https://br.authlete.com` | 🇧🇷 Brazil Cluster |

#### Example

```go
package main

import (
	"context"
	authlete "github.com/authlete/authlete-go-sdk"
	"log"
	"os"
)

func main() {
	ctx := context.Background()

	s := authlete.New(
		authlete.WithServerIndex(0),
		authlete.WithSecurity(os.Getenv("AUTHLETE_BEARER")),
	)

	res, err := s.Service.Get(ctx, "<id>")
	if err != nil {
		log.Fatal(err)
	}
	if res.Service != nil {
		// handle response
	}
}

```

### Override Server URL Per-Client

The default server can also be overridden globally using the `WithServerURL(serverURL string)` option when initializing the SDK client instance. For example:
```go
package main

import (
	"context"
	authlete "github.com/authlete/authlete-go-sdk"
	"log"
	"os"
)

func main() {
	ctx := context.Background()

	s := authlete.New(
		authlete.WithServerURL("https://br.authlete.com"),
		authlete.WithSecurity(os.Getenv("AUTHLETE_BEARER")),
	)

	res, err := s.Service.Get(ctx, "<id>")
	if err != nil {
		log.Fatal(err)
	}
	if res.Service != nil {
		// handle response
	}
}

```
<!-- End Server Selection [server] -->

<!-- Start Custom HTTP Client [http-client] -->
## Custom HTTP Client

The Go SDK makes API calls that wrap an internal HTTP client. The requirements for the HTTP client are very simple. It must match this interface:

```go
type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}
```

The built-in `net/http` client satisfies this interface and a default client based on the built-in is provided by default. To replace this default with a client of your own, you can implement this interface yourself or provide your own client configured as desired. Here's a simple example, which adds a client with a 30 second timeout.

```go
import (
	"net/http"
	"time"

	"github.com/authlete/authlete-go-sdk"
)

var (
	httpClient = &http.Client{Timeout: 30 * time.Second}
	sdkClient  = authlete.New(authlete.WithClient(httpClient))
)
```

This can be a convenient way to configure timeouts, cookies, proxies, custom headers, and other low-level configuration.
<!-- End Custom HTTP Client [http-client] -->

<!-- Placeholder for Future Speakeasy SDK Sections -->

# Development

## Maturity

This SDK is in beta, and there may be breaking changes between versions without a major version update. Therefore, we recommend pinning usage
to a specific package version. This way, you can install the same version each time without breaking changes unless you are intentionally
looking for the latest version.

## Contributions

While we value open-source contributions to this SDK, this library is generated programmatically. Any manual changes added to internal files will be overwritten on the next generation. 
We look forward to hearing your feedback. Feel free to open a PR or an issue with a proof of concept and we'll do our best to include it in a future release. 

### SDK Created by [Speakeasy](https://www.speakeasy.com/?utm_source=authlete&utm_campaign=go)
