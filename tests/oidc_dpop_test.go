package tests

import (
	"context"
	"fmt"
	"net/url"
	"testing"

	"github.com/authlete/authlete-go-sdk/models/components"
	"github.com/stretchr/testify/require"
)

func TestOIDCDpopFlow(t *testing.T) {
	ctx := context.Background()
	mgmtClient := createSDKClient(mgmtToken)
	sdkClient := createSDKClient(serviceToken)

	setupOIDCService(t, ctx, mgmtClient, serviceID, ptr(dpopTokenEndpoint))

	clientID, clientSecret := createTestClient(t, mgmtClient, serviceID)
	t.Cleanup(func() {
		_, _ = mgmtClient.Client.Delete(context.Background(), serviceID, clientID)
	})

	t.Run("OidcDpopFlow", func(t *testing.T) {
		key := generateECKey(t)
		nonce := randomNonce(t)
		encodedRedirect := url.QueryEscape(redirectURI)

		// Step 1: Authorization request with openid scope and nonce.
		authResp, err := sdkClient.Authorization.ProcessRequest(ctx, serviceID, components.AuthorizationRequest{
			Parameters: fmt.Sprintf(
				"response_type=code&client_id=%s&redirect_uri=%s&scope=openid&nonce=%s&state=%s",
				clientID, encodedRedirect, nonce, testState,
			),
		})
		require.NoError(t, err, "authorization request failed")
		require.NotNil(t, authResp.AuthorizationResponse)
		auth := authResp.AuthorizationResponse
		require.NotNil(t, auth.Action)
		require.Equal(t, components.AuthorizationResponseActionInteraction, *auth.Action)
		require.NotNil(t, auth.Ticket)

		// Step 2: Issue authorization code.
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

		// Step 3: Token request with DPoP proof.
		tokenResp, err := sdkClient.Token.Process(ctx, serviceID, components.TokenRequest{
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
			"expected OK, got %s: %s", *tok.Action, ptrStr(tok.ResultMessage))
		require.NotNil(t, tok.AccessToken)
		require.NotNil(t, tok.IDToken, "id_token must be issued for an openid DPoP flow")

		claims := decodeJWTPayload(t, *tok.IDToken)
		assertOIDCClaims(t, claims, subject, nonce, clientID)
	})
}
