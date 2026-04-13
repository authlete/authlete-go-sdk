package tests

import (
	"context"
	"fmt"
	"os"
	"testing"
	"time"

	authlete "github.com/authlete/authlete-go-sdk"
	"github.com/authlete/authlete-go-sdk/models/components"
	"github.com/stretchr/testify/require"
)

var (
	apiBaseURL   = requireEnv("API_BASE_URL")
	serviceID    = requireEnv("SERVICE_ID")
	serviceToken = requireEnv("SERVICE_TOKEN")
	mgmtToken    = envOrDefault("ORG_TOKEN", serviceToken)
)

const (
	redirectURI          = "https://client.example.com/callback"
	testState            = "testState"
	subject              = "testuser"
	tokenDurationSeconds = int64(600)
)

func requireEnv(key string) string {
	v := os.Getenv(key)
	if v == "" {
		panic(fmt.Sprintf("required environment variable %s is not set", key))
	}
	return v
}

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

	resp, err := sdk.Client.Create(ctx, svcID, &clientInput)
	require.NoError(t, err, "failed to create test client")
	require.NotNil(t, resp.Client, "client response must not be nil")
	require.NotNil(t, resp.Client.ClientID, "client ID must not be nil")
	require.NotNil(t, resp.Client.ClientSecret, "client secret must not be nil")

	return fmt.Sprintf("%d", *resp.Client.ClientID), *resp.Client.ClientSecret
}

func assertTokenValid(t *testing.T, sdk *authlete.Authlete, svcID string, accessToken string) {
	t.Helper()

	ctx := context.Background()
	resp, err := sdk.Introspection.Process(ctx, svcID, components.IntrospectionRequest{
		Token: accessToken,
	})
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

func ptrStr(s *string) string {
	if s == nil {
		return "<nil>"
	}
	return *s
}
