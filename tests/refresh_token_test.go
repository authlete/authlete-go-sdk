package tests

import (
	"context"
	"fmt"
	"net/url"
	"testing"

	authlete "github.com/authlete/authlete-go-sdk"
	"github.com/authlete/authlete-go-sdk/models/components"
	"github.com/stretchr/testify/require"
)

// doAuthCodeFlow runs an auth-code flow end-to-end and returns the token
// response so the caller can assert on the refresh_token / access_token.
func doAuthCodeFlow(
	t *testing.T,
	ctx context.Context,
	sdk *authlete.Authlete,
	svcID, clientID, clientSecret string,
) *components.TokenResponse {
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
	require.Equal(t, components.AuthorizationResponseActionInteraction, *auth.Action,
		"expected INTERACTION, got %s", *auth.Action)
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
	})
	require.NoError(t, err, "token request failed")
	require.NotNil(t, tokenResp.TokenResponse)
	tok := tokenResp.TokenResponse
	require.NotNil(t, tok.Action)
	require.Equal(t, components.TokenResponseActionOk, *tok.Action,
		"expected OK at token endpoint, got %s: %s", *tok.Action, ptrStr(tok.ResultMessage))
	return tok
}

// =============================================================================
// Service with both AUTHORIZATION_CODE and REFRESH_TOKEN grant types enabled.
// =============================================================================
func TestRefreshTokenFlow(t *testing.T) {
	ctx := context.Background()
	mgmtClient := createSDKClient(mgmtToken)
	sdkClient := createSDKClient(serviceToken)

	_, err := mgmtClient.Service.Update(ctx, serviceID, &components.ServiceInput{
		SupportedGrantTypes: []components.GrantType{
			components.GrantTypeAuthorizationCode,
			components.GrantTypeRefreshToken,
		},
		AccessTokenDuration:  ptr(tokenDurationSeconds),
		RefreshTokenDuration: ptr(tokenDurationSeconds),
	})
	require.NoError(t, err, "failed to update service")

	clientID, clientSecret := createTestClient(t, mgmtClient, serviceID)
	t.Cleanup(func() {
		_, _ = mgmtClient.Client.Delete(context.Background(), serviceID, clientID)
	})

	t.Run("RefreshTokenIssued", func(t *testing.T) {
		tok := doAuthCodeFlow(t, ctx, sdkClient, serviceID, clientID, clientSecret)
		require.NotNil(t, tok.RefreshToken,
			"refresh token must be issued when the refresh_token grant is supported")
	})

	t.Run("RefreshTokenFlow", func(t *testing.T) {
		tok := doAuthCodeFlow(t, ctx, sdkClient, serviceID, clientID, clientSecret)
		require.NotNil(t, tok.RefreshToken)

		refreshResp, err := sdkClient.Token.Process(ctx, serviceID, components.TokenRequest{
			Parameters:   fmt.Sprintf("grant_type=refresh_token&refresh_token=%s", *tok.RefreshToken),
			ClientID:     ptr(clientID),
			ClientSecret: ptr(clientSecret),
		})
		require.NoError(t, err, "refresh token exchange failed")
		require.NotNil(t, refreshResp.TokenResponse)
		refreshed := refreshResp.TokenResponse
		require.NotNil(t, refreshed.Action)
		require.Equal(t, components.TokenResponseActionOk, *refreshed.Action,
			"expected OK for refresh exchange, got %s: %s", *refreshed.Action, ptrStr(refreshed.ResultMessage))
		require.NotNil(t, refreshed.AccessToken, "new access token must be issued on refresh")

		assertTokenValid(t, sdkClient, serviceID, *refreshed.AccessToken)
	})

	t.Run("RefreshTokenRevocation", func(t *testing.T) {
		tok := doAuthCodeFlow(t, ctx, sdkClient, serviceID, clientID, clientSecret)
		require.NotNil(t, tok.RefreshToken)

		revokeResp, err := sdkClient.Revocation.Process(ctx, serviceID, components.RevocationRequest{
			Parameters:   fmt.Sprintf("token=%s", *tok.RefreshToken),
			ClientID:     ptr(clientID),
			ClientSecret: ptr(clientSecret),
		})
		require.NoError(t, err, "revocation request failed")
		require.NotNil(t, revokeResp.RevocationResponse)
		require.NotNil(t, revokeResp.RevocationResponse.Action)
		require.Equal(t, components.RevocationResponseActionOk, *revokeResp.RevocationResponse.Action,
			"expected OK for refresh-token revocation")
	})
}

// =============================================================================
// Service with only AUTHORIZATION_CODE grant type (REFRESH_TOKEN not supported).
// =============================================================================
func TestRefreshTokenNotSupported(t *testing.T) {
	ctx := context.Background()
	mgmtClient := createSDKClient(mgmtToken)
	sdkClient := createSDKClient(serviceToken)

	_, err := mgmtClient.Service.Update(ctx, serviceID, &components.ServiceInput{
		SupportedGrantTypes: []components.GrantType{
			components.GrantTypeAuthorizationCode,
		},
		AccessTokenDuration:  ptr(tokenDurationSeconds),
		RefreshTokenDuration: ptr(tokenDurationSeconds),
	})
	require.NoError(t, err, "failed to restrict supported grant types")

	t.Cleanup(func() {
		_, _ = mgmtClient.Service.Update(context.Background(), serviceID, &components.ServiceInput{
			SupportedGrantTypes: []components.GrantType{
				components.GrantTypeAuthorizationCode,
				components.GrantTypeRefreshToken,
			},
			AccessTokenDuration:  ptr(tokenDurationSeconds),
			RefreshTokenDuration: ptr(tokenDurationSeconds),
		})
	})

	clientID, clientSecret := createTestClient(t, mgmtClient, serviceID)
	t.Cleanup(func() {
		_, _ = mgmtClient.Client.Delete(context.Background(), serviceID, clientID)
	})

	t.Run("RefreshTokenNotIssued", func(t *testing.T) {
		tok := doAuthCodeFlow(t, ctx, sdkClient, serviceID, clientID, clientSecret)
		require.Nil(t, tok.RefreshToken,
			"refresh token must not be issued when the refresh_token grant is not supported")
	})

	t.Run("RefreshTokenRejected", func(t *testing.T) {
		tokenResp, err := sdkClient.Token.Process(ctx, serviceID, components.TokenRequest{
			Parameters:   "grant_type=refresh_token&refresh_token=dummy_token",
			ClientID:     ptr(clientID),
			ClientSecret: ptr(clientSecret),
		})
		require.NoError(t, err, "token request failed")
		require.NotNil(t, tokenResp.TokenResponse)
		require.NotNil(t, tokenResp.TokenResponse.Action)
		require.NotEqual(t, components.TokenResponseActionOk, *tokenResp.TokenResponse.Action,
			"refresh_token grant must be rejected when not supported by the service")
	})
}
