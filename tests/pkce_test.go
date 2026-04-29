package tests

import (
	"context"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"net/url"
	"testing"

	authlete "github.com/authlete/authlete-go-sdk"
	"github.com/authlete/authlete-go-sdk/models/components"
	"github.com/stretchr/testify/require"
)

// generateCodeVerifier returns a 43-character RFC 7636 code verifier built
// from 32 random bytes encoded as base64url (no padding).
func generateCodeVerifier(t *testing.T) string {
	t.Helper()
	buf := make([]byte, 32)
	_, err := rand.Read(buf)
	require.NoError(t, err)
	return base64.RawURLEncoding.EncodeToString(buf)
}

func s256CodeChallenge(verifier string) string {
	sum := sha256.Sum256([]byte(verifier))
	return base64.RawURLEncoding.EncodeToString(sum[:])
}

// runPkceFlow runs the auth-code flow with a code_challenge in the auth
// request and the given code_verifier in the token request, returning the
// token response so callers can assert on it.
func runPkceFlow(
	t *testing.T,
	ctx context.Context,
	sdk *authlete.Authlete,
	svcID, clientID, clientSecret string,
	codeVerifier, codeChallenge, codeChallengeMethod string,
) *components.TokenResponse {
	t.Helper()
	encodedRedirect := url.QueryEscape(redirectURI)

	authResp, err := sdk.Authorization.ProcessRequest(ctx, svcID, components.AuthorizationRequest{
		Parameters: fmt.Sprintf(
			"response_type=code&client_id=%s&redirect_uri=%s&state=%s&code_challenge=%s&code_challenge_method=%s",
			clientID, encodedRedirect, testState, codeChallenge, codeChallengeMethod,
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
			"grant_type=authorization_code&code=%s&redirect_uri=%s&code_verifier=%s",
			*issue.AuthorizationCode, encodedRedirect, codeVerifier,
		),
		ClientID:     ptr(clientID),
		ClientSecret: ptr(clientSecret),
	})
	require.NoError(t, err, "token request failed")
	require.NotNil(t, tokenResp.TokenResponse)
	return tokenResp.TokenResponse
}

// =============================================================================
// Standard service — PKCE is optional.
// =============================================================================
func TestPkceFlow(t *testing.T) {
	ctx := context.Background()
	mgmtClient := createSDKClient(mgmtToken)
	sdkClient := createSDKClient(serviceToken)

	_, err := mgmtClient.Service.Update(ctx, serviceID, &components.ServiceInput{
		AccessTokenDuration: ptr(tokenDurationSeconds),
	})
	require.NoError(t, err, "failed to update service")

	clientID, clientSecret := createTestClient(t, mgmtClient, serviceID)
	t.Cleanup(func() {
		_, _ = mgmtClient.Client.Delete(context.Background(), serviceID, clientID)
	})

	t.Run("PkceS256Flow", func(t *testing.T) {
		verifier := generateCodeVerifier(t)
		challenge := s256CodeChallenge(verifier)

		tok := runPkceFlow(t, ctx, sdkClient, serviceID, clientID, clientSecret,
			verifier, challenge, "S256")

		require.NotNil(t, tok.Action)
		require.Equal(t, components.TokenResponseActionOk, *tok.Action,
			"expected OK, got %s: %s", *tok.Action, ptrStr(tok.ResultMessage))
		require.NotNil(t, tok.AccessToken)
		assertTokenValid(t, sdkClient, serviceID, *tok.AccessToken)
	})

	t.Run("PkcePlainFlow", func(t *testing.T) {
		verifier := generateCodeVerifier(t)

		tok := runPkceFlow(t, ctx, sdkClient, serviceID, clientID, clientSecret,
			verifier, verifier, "plain")

		require.NotNil(t, tok.Action)
		require.Equal(t, components.TokenResponseActionOk, *tok.Action,
			"expected OK, got %s: %s", *tok.Action, ptrStr(tok.ResultMessage))
		require.NotNil(t, tok.AccessToken)
		assertTokenValid(t, sdkClient, serviceID, *tok.AccessToken)
	})

	t.Run("WrongCodeVerifierRejected", func(t *testing.T) {
		verifier := generateCodeVerifier(t)
		challenge := s256CodeChallenge(verifier)
		wrongVerifier := generateCodeVerifier(t)

		tok := runPkceFlow(t, ctx, sdkClient, serviceID, clientID, clientSecret,
			wrongVerifier, challenge, "S256")

		require.NotNil(t, tok.Action)
		require.NotEqual(t, components.TokenResponseActionOk, *tok.Action,
			"token request with wrong code_verifier must not succeed")
	})
}

// =============================================================================
// Service with pkceRequired: true.
// =============================================================================
func TestPkceRequired(t *testing.T) {
	ctx := context.Background()
	mgmtClient := createSDKClient(mgmtToken)
	sdkClient := createSDKClient(serviceToken)

	_, err := mgmtClient.Service.Update(ctx, serviceID, &components.ServiceInput{
		PkceRequired:        ptr(true),
		AccessTokenDuration: ptr(tokenDurationSeconds),
	})
	require.NoError(t, err, "failed to enable pkceRequired")

	t.Cleanup(func() {
		_, _ = mgmtClient.Service.Update(context.Background(), serviceID, &components.ServiceInput{
			PkceRequired:        ptr(false),
			AccessTokenDuration: ptr(tokenDurationSeconds),
		})
	})

	clientID, clientSecret := createTestClient(t, mgmtClient, serviceID)
	t.Cleanup(func() {
		_, _ = mgmtClient.Client.Delete(context.Background(), serviceID, clientID)
	})

	t.Run("MissingCodeChallengeRejected", func(t *testing.T) {
		encodedRedirect := url.QueryEscape(redirectURI)
		params := fmt.Sprintf("response_type=code&client_id=%s&redirect_uri=%s&state=%s",
			clientID, encodedRedirect, testState)

		authResp, err := sdkClient.Authorization.ProcessRequest(ctx, serviceID, components.AuthorizationRequest{
			Parameters: params,
		})
		require.NoError(t, err, "authorization request failed")
		require.NotNil(t, authResp.AuthorizationResponse)
		require.NotNil(t, authResp.AuthorizationResponse.Action)
		require.NotEqual(t, components.AuthorizationResponseActionInteraction,
			*authResp.AuthorizationResponse.Action,
			"auth request without code_challenge must be rejected when pkceRequired=true")
	})

	t.Run("PkceS256FlowSucceeds", func(t *testing.T) {
		verifier := generateCodeVerifier(t)
		challenge := s256CodeChallenge(verifier)

		tok := runPkceFlow(t, ctx, sdkClient, serviceID, clientID, clientSecret,
			verifier, challenge, "S256")

		require.NotNil(t, tok.Action)
		require.Equal(t, components.TokenResponseActionOk, *tok.Action,
			"expected OK, got %s: %s", *tok.Action, ptrStr(tok.ResultMessage))
		require.NotNil(t, tok.AccessToken)
		assertTokenValid(t, sdkClient, serviceID, *tok.AccessToken)
	})
}

// =============================================================================
// Service with pkceS256Required: true.
// =============================================================================
func TestPkceS256Required(t *testing.T) {
	ctx := context.Background()
	mgmtClient := createSDKClient(mgmtToken)
	sdkClient := createSDKClient(serviceToken)

	_, err := mgmtClient.Service.Update(ctx, serviceID, &components.ServiceInput{
		PkceS256Required:    ptr(true),
		AccessTokenDuration: ptr(tokenDurationSeconds),
	})
	require.NoError(t, err, "failed to enable pkceS256Required")

	t.Cleanup(func() {
		_, _ = mgmtClient.Service.Update(context.Background(), serviceID, &components.ServiceInput{
			PkceS256Required:    ptr(false),
			AccessTokenDuration: ptr(tokenDurationSeconds),
		})
	})

	clientID, clientSecret := createTestClient(t, mgmtClient, serviceID)
	t.Cleanup(func() {
		_, _ = mgmtClient.Client.Delete(context.Background(), serviceID, clientID)
	})

	t.Run("PlainMethodRejected", func(t *testing.T) {
		verifier := generateCodeVerifier(t)
		encodedRedirect := url.QueryEscape(redirectURI)

		params := fmt.Sprintf(
			"response_type=code&client_id=%s&redirect_uri=%s&state=%s&code_challenge=%s&code_challenge_method=plain",
			clientID, encodedRedirect, testState, verifier,
		)

		authResp, err := sdkClient.Authorization.ProcessRequest(ctx, serviceID, components.AuthorizationRequest{
			Parameters: params,
		})
		require.NoError(t, err, "authorization request failed")
		require.NotNil(t, authResp.AuthorizationResponse)
		require.NotNil(t, authResp.AuthorizationResponse.Action)
		require.NotEqual(t, components.AuthorizationResponseActionInteraction,
			*authResp.AuthorizationResponse.Action,
			"plain code_challenge_method must be rejected when pkceS256Required=true")
	})

	t.Run("S256FlowSucceeds", func(t *testing.T) {
		verifier := generateCodeVerifier(t)
		challenge := s256CodeChallenge(verifier)

		tok := runPkceFlow(t, ctx, sdkClient, serviceID, clientID, clientSecret,
			verifier, challenge, "S256")

		require.NotNil(t, tok.Action)
		require.Equal(t, components.TokenResponseActionOk, *tok.Action,
			"expected OK, got %s: %s", *tok.Action, ptrStr(tok.ResultMessage))
		require.NotNil(t, tok.AccessToken)
		assertTokenValid(t, sdkClient, serviceID, *tok.AccessToken)
	})
}
