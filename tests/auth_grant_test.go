package tests

import (
	"context"
	"fmt"
	"net/url"
	"testing"

	"github.com/authlete/authlete-go-sdk/models/components"
	"github.com/stretchr/testify/require"
)

func TestAuthorizationCodeFlow(t *testing.T) {
	ctx := context.Background()
	mgmtClient := createSDKClient(mgmtToken)
	sdkClient := createSDKClient(serviceToken)

	// Configure service with a known access token duration.
	_, err := mgmtClient.Service.Update(ctx, serviceID, &components.ServiceInput{
		AccessTokenDuration: ptr(tokenDurationSeconds),
	})
	require.NoError(t, err, "failed to update service")

	// Create a test client and schedule cleanup.
	clientID, clientSecret := createTestClient(t, mgmtClient, serviceID)
	t.Cleanup(func() {
		_, _ = mgmtClient.Client.Delete(context.Background(), serviceID, clientID)
	})

	// --- Step 1: Authorization Request ---
	params := fmt.Sprintf("response_type=code&client_id=%s&redirect_uri=%s&state=%s",
		clientID, url.QueryEscape(redirectURI), testState)

	authResp, err := sdkClient.Authorization.ProcessRequest(ctx, serviceID, components.AuthorizationRequest{
		Parameters: params,
	})
	require.NoError(t, err, "authorization request failed")
	require.NotNil(t, authResp.AuthorizationResponse, "authorization response must not be nil")

	authResult := authResp.AuthorizationResponse
	require.NotNil(t, authResult.Action, "authorization action must not be nil")
	require.Equal(t, components.AuthorizationResponseActionInteraction, *authResult.Action,
		"expected INTERACTION, got %s", *authResult.Action)
	require.NotNil(t, authResult.Ticket, "authorization ticket must not be nil")

	// --- Step 2: Authorization Issue (simulate user consent) ---
	issueResp, err := sdkClient.Authorization.Issue(ctx, serviceID, components.AuthorizationIssueRequest{
		Ticket:  *authResult.Ticket,
		Subject: subject,
	})
	require.NoError(t, err, "authorization issue failed")
	require.NotNil(t, issueResp.AuthorizationIssueResponse, "authorization issue response must not be nil")

	issueResult := issueResp.AuthorizationIssueResponse
	require.NotNil(t, issueResult.Action, "issue action must not be nil")
	require.Equal(t, components.AuthorizationIssueResponseActionLocation, *issueResult.Action,
		"expected LOCATION, got %s", *issueResult.Action)
	require.NotNil(t, issueResult.AuthorizationCode, "authorization code must not be nil")
	require.NotNil(t, issueResult.ResponseContent, "response content must not be nil")
	require.Contains(t, *issueResult.ResponseContent, "code=")
	require.Contains(t, *issueResult.ResponseContent, fmt.Sprintf("state=%s", testState))

	// --- Step 3: Token Request ---
	tokenParams := fmt.Sprintf("grant_type=authorization_code&code=%s&redirect_uri=%s",
		*issueResult.AuthorizationCode, url.QueryEscape(redirectURI))

	tokenResp, err := sdkClient.Token.Process(ctx, serviceID, components.TokenRequest{
		Parameters:   tokenParams,
		ClientID:     ptr(clientID),
		ClientSecret: ptr(clientSecret),
	})
	require.NoError(t, err, "token request failed")
	require.NotNil(t, tokenResp.TokenResponse, "token response must not be nil")

	tokenResult := tokenResp.TokenResponse
	require.NotNil(t, tokenResult.Action, "token action must not be nil")
	require.Equal(t, components.TokenResponseActionOk, *tokenResult.Action,
		"expected OK, got %s", *tokenResult.Action)
	require.NotNil(t, tokenResult.AccessToken, "access token must not be nil")

	// --- Step 4: Introspection ---
	assertTokenValid(t, sdkClient, serviceID, *tokenResult.AccessToken)

	// --- Step 5: Revocation ---
	revocationParams := fmt.Sprintf("token=%s", *tokenResult.AccessToken)

	revokeResp, err := sdkClient.Revocation.Process(ctx, serviceID, components.RevocationRequest{
		Parameters:   revocationParams,
		ClientID:     ptr(clientID),
		ClientSecret: ptr(clientSecret),
	})
	require.NoError(t, err, "revocation request failed")
	require.NotNil(t, revokeResp.RevocationResponse, "revocation response must not be nil")
	require.NotNil(t, revokeResp.RevocationResponse.Action, "revocation action must not be nil")
	require.Equal(t, components.RevocationResponseActionOk, *revokeResp.RevocationResponse.Action,
		"expected OK, got %s", *revokeResp.RevocationResponse.Action)
}
