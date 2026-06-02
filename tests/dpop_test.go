package tests

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"net/url"
	"testing"
	"time"

	authlete "github.com/authlete/authlete-go-sdk"
	"github.com/authlete/authlete-go-sdk/models/components"
	"github.com/stretchr/testify/require"
)

// obtainDpopBoundToken runs auth → issue → token (with DPoP at the token
// endpoint) and returns the access token. Auth + issue do not need DPoP.
func obtainDpopBoundToken(
	t *testing.T,
	ctx context.Context,
	sdk *authlete.Authlete,
	svcID, clientID, clientSecret string,
	key *ecdsa.PrivateKey,
) string {
	t.Helper()
	encodedRedirect := url.QueryEscape(redirectURI)

	authResp, err := sdk.Authorization.ProcessRequest(ctx, svcID, components.AuthorizationRequest{
		Parameters: fmt.Sprintf(
			"response_type=code&client_id=%s&redirect_uri=%s&state=%s",
			clientID, encodedRedirect, testState,
		),
	})
	require.NoError(t, err, "authorization request failed")
	require.NotNil(t, authResp.AuthorizationResponse)
	auth := authResp.AuthorizationResponse
	require.NotNil(t, auth.Action)
	require.Equal(t, components.AuthorizationResponseActionInteraction, *auth.Action)
	require.NotNil(t, auth.Ticket)

	issueResp, err := sdk.Authorization.Issue(ctx, svcID, components.AuthorizationIssueRequest{
		Ticket:  *auth.Ticket,
		Subject: subject,
	})
	require.NoError(t, err, "authorization issue failed")
	require.NotNil(t, issueResp.AuthorizationIssueResponse)
	issue := issueResp.AuthorizationIssueResponse
	require.NotNil(t, issue.Action)
	require.Equal(t, components.AuthorizationIssueResponseActionLocation, *issue.Action)
	require.NotNil(t, issue.AuthorizationCode)

	tokenResp, err := sdk.Token.Process(ctx, svcID, components.TokenRequest{
		Parameters: fmt.Sprintf(
			"grant_type=authorization_code&code=%s&redirect_uri=%s",
			*issue.AuthorizationCode, encodedRedirect,
		),
		ClientID:     ptr(clientID),
		ClientSecret: ptr(clientSecret),
		Dpop:         ptr(dpopProof(t, key, "POST", dpopTokenEndpoint)),
		Htm:          ptr("POST"),
		Htu:          ptr(dpopTokenEndpoint),
	})
	require.NoError(t, err, "token request failed")
	require.NotNil(t, tokenResp.TokenResponse)
	tok := tokenResp.TokenResponse
	require.NotNil(t, tok.Action)
	require.Equal(t, components.TokenResponseActionOk, *tok.Action,
		"expected OK at DPoP token endpoint, got %s: %s", *tok.Action, ptrStr(tok.ResultMessage))
	require.NotNil(t, tok.AccessToken)
	return *tok.AccessToken
}

// =============================================================================
// Standard service — DPoP is optional.
// =============================================================================
func TestDpopFlow(t *testing.T) {
	ctx := context.Background()
	mgmtClient := createSDKClient(mgmtToken)
	sdkClient := createSDKClient(serviceToken)

	_, err := mgmtClient.Service.Update(ctx, serviceID, &components.ServiceInput{
		TokenEndpoint:       ptr(dpopTokenEndpoint),
		AccessTokenDuration: ptr(tokenDurationSeconds),
	})
	require.NoError(t, err, "failed to update service")

	clientID, clientSecret := createTestClient(t, mgmtClient, serviceID)
	t.Cleanup(func() {
		_, _ = mgmtClient.Client.Delete(context.Background(), serviceID, clientID)
	})

	t.Run("DpopBasicFlow", func(t *testing.T) {
		key := generateECKey(t)
		accessToken := obtainDpopBoundToken(t, ctx, sdkClient, serviceID, clientID, clientSecret, key)
		require.NotEmpty(t, accessToken)
	})

	t.Run("DpopIntrospectionValid", func(t *testing.T) {
		key := generateECKey(t)
		accessToken := obtainDpopBoundToken(t, ctx, sdkClient, serviceID, clientID, clientSecret, key)

		introResp, err := sdkClient.Introspection.Process(ctx, serviceID, components.IntrospectionRequest{
			Token: accessToken,
			Dpop:  ptr(dpopProof(t, key, "GET", dpopResourceURL, dpopProofOpts{accessToken: accessToken})),
			Htm:   ptr("GET"),
			Htu:   ptr(dpopResourceURL),
		})
		require.NoError(t, err, "introspection request failed")
		require.NotNil(t, introResp.IntrospectionResponse)
		require.NotNil(t, introResp.IntrospectionResponse.Action)
		require.Equal(t, components.IntrospectionResponseActionOk, *introResp.IntrospectionResponse.Action,
			"expected OK introspection with valid DPoP, got %s: %s",
			*introResp.IntrospectionResponse.Action,
			ptrStr(introResp.IntrospectionResponse.ResultMessage))
	})

	t.Run("DpopIntrospectionWithoutProofRejected", func(t *testing.T) {
		key := generateECKey(t)
		accessToken := obtainDpopBoundToken(t, ctx, sdkClient, serviceID, clientID, clientSecret, key)

		introResp, err := sdkClient.Introspection.Process(ctx, serviceID, components.IntrospectionRequest{
			Token: accessToken,
		})
		require.NoError(t, err, "introspection request failed")
		require.NotNil(t, introResp.IntrospectionResponse)
		require.NotNil(t, introResp.IntrospectionResponse.Action)
		require.NotEqual(t, components.IntrospectionResponseActionOk,
			*introResp.IntrospectionResponse.Action,
			"introspection without DPoP proof must not return OK for a DPoP-bound token")
	})
}

// =============================================================================
// Client with dpopRequired: true.
// =============================================================================
func TestDpopRequired(t *testing.T) {
	ctx := context.Background()
	mgmtClient := createSDKClient(mgmtToken)
	sdkClient := createSDKClient(serviceToken)

	_, err := mgmtClient.Service.Update(ctx, serviceID, &components.ServiceInput{
		TokenEndpoint:       ptr(dpopTokenEndpoint),
		AccessTokenDuration: ptr(tokenDurationSeconds),
	})
	require.NoError(t, err, "failed to update service")

	// Create a confidential client with dpopRequired: true.
	clientResp, err := mgmtClient.Client.Create(ctx, serviceID, &components.ClientInput{
		ClientName:    ptr(fmt.Sprintf("go-sdk-test-dpop-required-%d", time.Now().UnixMilli())),
		ClientType:    ptr(components.ClientTypeConfidential),
		GrantTypes:    []components.GrantType{components.GrantTypeAuthorizationCode},
		ResponseTypes: []components.ResponseType{components.ResponseTypeCode},
		RedirectUris:  []string{redirectURI},
		DpopRequired:  ptr(true),
	})
	require.NoError(t, err, "failed to create dpopRequired client")
	require.NotNil(t, clientResp.Client)
	require.NotNil(t, clientResp.Client.ClientID)
	require.NotNil(t, clientResp.Client.ClientSecret)
	clientID := fmt.Sprintf("%d", *clientResp.Client.ClientID)
	clientSecret := *clientResp.Client.ClientSecret
	t.Cleanup(func() {
		_, _ = mgmtClient.Client.Delete(context.Background(), serviceID, clientID)
	})

	t.Run("TokenWithoutDpopRejected", func(t *testing.T) {
		encodedRedirect := url.QueryEscape(redirectURI)

		authResp, err := sdkClient.Authorization.ProcessRequest(ctx, serviceID, components.AuthorizationRequest{
			Parameters: fmt.Sprintf(
				"response_type=code&client_id=%s&redirect_uri=%s&state=%s",
				clientID, encodedRedirect, testState,
			),
		})
		require.NoError(t, err, "authorization request failed")
		require.NotNil(t, authResp.AuthorizationResponse)
		auth := authResp.AuthorizationResponse
		require.NotNil(t, auth.Action)
		require.Equal(t, components.AuthorizationResponseActionInteraction, *auth.Action)
		require.NotNil(t, auth.Ticket)

		issueResp, err := sdkClient.Authorization.Issue(ctx, serviceID, components.AuthorizationIssueRequest{
			Ticket:  *auth.Ticket,
			Subject: subject,
		})
		require.NoError(t, err, "authorization issue failed")
		require.NotNil(t, issueResp.AuthorizationIssueResponse)
		issue := issueResp.AuthorizationIssueResponse
		require.NotNil(t, issue.Action)
		require.Equal(t, components.AuthorizationIssueResponseActionLocation, *issue.Action)
		require.NotNil(t, issue.AuthorizationCode)

		// Token request WITHOUT DPoP — must be rejected.
		tokenResp, err := sdkClient.Token.Process(ctx, serviceID, components.TokenRequest{
			Parameters: fmt.Sprintf(
				"grant_type=authorization_code&code=%s&redirect_uri=%s",
				*issue.AuthorizationCode, encodedRedirect,
			),
			ClientID:     ptr(clientID),
			ClientSecret: ptr(clientSecret),
		})
		require.NoError(t, err, "token request failed")
		require.NotNil(t, tokenResp.TokenResponse)
		require.NotNil(t, tokenResp.TokenResponse.Action)
		require.NotEqual(t, components.TokenResponseActionOk, *tokenResp.TokenResponse.Action,
			"token request without DPoP must be rejected when dpopRequired=true")
	})

	t.Run("DpopFlowSucceedsWhenRequired", func(t *testing.T) {
		key := generateECKey(t)
		accessToken := obtainDpopBoundToken(t, ctx, sdkClient, serviceID, clientID, clientSecret, key)
		require.NotEmpty(t, accessToken)
	})
}
