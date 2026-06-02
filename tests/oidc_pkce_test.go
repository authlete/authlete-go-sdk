package tests

import (
	"context"
	"fmt"
	"net/url"
	"testing"

	"github.com/authlete/authlete-go-sdk/models/components"
	"github.com/stretchr/testify/require"
)

func TestOIDCPkceFlow(t *testing.T) {
	ctx := context.Background()
	mgmtClient := createSDKClient(mgmtToken)
	sdkClient := createSDKClient(serviceToken)

	setupOIDCService(t, ctx, mgmtClient, serviceID, nil)

	clientID, clientSecret := createTestClient(t, mgmtClient, serviceID)
	t.Cleanup(func() {
		_, _ = mgmtClient.Client.Delete(context.Background(), serviceID, clientID)
	})

	t.Run("OidcPkceS256Flow", func(t *testing.T) {
		nonce := randomNonce(t)
		verifier := generateCodeVerifier(t)
		challenge := s256CodeChallenge(verifier)
		encodedRedirect := url.QueryEscape(redirectURI)

		// Authorization request with openid scope, nonce, and PKCE.
		authResp, err := sdkClient.Authorization.ProcessRequest(ctx, serviceID, components.AuthorizationRequest{
			Parameters: fmt.Sprintf(
				"response_type=code&client_id=%s&redirect_uri=%s&scope=openid&nonce=%s&state=%s&code_challenge=%s&code_challenge_method=S256",
				clientID, encodedRedirect, nonce, testState, challenge,
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

		// Token request with code_verifier.
		tokenResp, err := sdkClient.Token.Process(ctx, serviceID, components.TokenRequest{
			Parameters: fmt.Sprintf(
				"grant_type=authorization_code&code=%s&redirect_uri=%s&code_verifier=%s",
				*issue.AuthorizationCode, encodedRedirect, verifier,
			),
			ClientID:     ptr(clientID),
			ClientSecret: ptr(clientSecret),
		})
		require.NoError(t, err, "token request failed")
		require.NotNil(t, tokenResp.TokenResponse)
		tok := tokenResp.TokenResponse
		require.NotNil(t, tok.Action)
		require.Equal(t, components.TokenResponseActionOk, *tok.Action,
			"expected OK, got %s: %s", *tok.Action, ptrStr(tok.ResultMessage))
		require.NotNil(t, tok.AccessToken)
		require.NotNil(t, tok.IDToken, "id_token must be issued for an openid PKCE flow")

		claims := decodeJWTPayload(t, *tok.IDToken)
		assertOIDCClaims(t, claims, subject, nonce, clientID)
	})
}
