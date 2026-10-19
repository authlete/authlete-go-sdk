package tests

import (
	"context"
	"fmt"
	"os"
	"testing"
	"time"

	authlete "github.com/authlete/authlete-go-sdk"
	"github.com/authlete/authlete-go-sdk/models/components"
	"github.com/authlete/authlete-go-sdk/models/operations"
	"github.com/authlete/authlete-go-sdk/optionalnullable"
	"github.com/stretchr/testify/require"
)

var (
	apiBaseURL   = os.Getenv("API_BASE_URL")
	serviceID    = os.Getenv("SERVICE_ID")
	serviceToken = os.Getenv("SERVICE_TOKEN")
	mgmtToken    = envOrDefault("ORG_TOKEN", serviceToken)
)

const (
	redirectURI          = "https://client.example.com/callback"
	testState            = "testState"
	subject              = "testuser"
	tokenDurationSeconds = int64(600)
)

func envOrDefault(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}

func createSDKClient(token string) *authlete.Authlete {
	return authlete.New(
		authlete.WithSecurity(token),
		authlete.WithServerURL(apiBaseURL),
		authlete.WithServiceID(serviceID),
	)
}

func createTestClient(t *testing.T, sdk *authlete.Authlete, svcID string) (clientID string, clientSecret string) {
	t.Helper()

	ctx := context.Background()
	clientInput := components.ClientInput{
		ClientName:    ptr(fmt.Sprintf("go-sdk-test-client-%d", time.Now().UnixMilli())),
		ClientType:    ptr(components.ClientTypeConfidential),
		GrantTypes:    []components.GrantType{components.GrantTypeAuthorizationCode, components.GrantTypeRefreshToken},
		ResponseTypes: []components.ResponseType{components.ResponseTypeCode},
		RedirectUris:  []string{redirectURI},
	}

	resp, err := sdk.Client.Create(ctx, operations.ClientCreateAPIRequest{ServiceID: &svcID, Body: &clientInput})
	require.NoError(t, err, "failed to create test client")
	require.NotNil(t, resp.Client, "client response must not be nil")
	require.NotNil(t, resp.Client.ClientID, "client ID must not be nil")
	require.NotNil(t, resp.Client.ClientSecret, "client secret must not be nil")

	return fmt.Sprintf("%d", *resp.Client.ClientID), *resp.Client.ClientSecret
}

func assertTokenValid(t *testing.T, sdk *authlete.Authlete, svcID string, accessToken string) {
	t.Helper()

	ctx := context.Background()
	resp, err := sdk.Introspection.Process(ctx, operations.AuthIntrospectionAPIRequest{ServiceID: &svcID, Body: components.IntrospectionRequest{
		Token: accessToken,
	}})
	require.NoError(t, err, "introspection request failed")
	require.NotNil(t, resp.IntrospectionResponse, "introspection response must not be nil")
	require.NotNil(t, resp.IntrospectionResponse.Action, "introspection action must not be nil")
	require.Equal(t, components.IntrospectionResponseActionOk, *resp.IntrospectionResponse.Action,
		"expected introspection action OK, got %s: %s",
		*resp.IntrospectionResponse.Action,
		ptrStr(resp.IntrospectionResponse.ResultMessage))
}

func ptr[T any](v T) *T {
	return &v
}

// setValue wraps v for an OptionalNullable field, so it is sent as v.
func setValue[T any](v T) optionalnullable.OptionalNullable[T] {
	return optionalnullable.From(&v)
}

// setUnlessNil wraps p for an OptionalNullable field, leaving the field unset
// (omitted) when p is nil rather than sending an explicit null.
func setUnlessNil[T any](p *T) optionalnullable.OptionalNullable[T] {
	if p == nil {
		return nil
	}
	return optionalnullable.From(p)
}

func ptrStr(s *string) string {
	if s == nil {
		return "<nil>"
	}
	return *s
}
