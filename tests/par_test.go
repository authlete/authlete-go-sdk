package tests

import (
	"context"
	"fmt"
	"net/url"
	"testing"

	"github.com/authlete/authlete-go-sdk/models/components"
	"github.com/stretchr/testify/require"
)

// =============================================================================
// Standard service — PAR is optional.
// =============================================================================
func TestParFlow(t *testing.T) {
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

	t.Run("ParBasicFlow", func(t *testing.T) {
		encodedRedirect := url.QueryEscape(redirectURI)

		// Step 1: Push authorization parameters
		parParams := fmt.Sprintf("response_type=code&client_id=%s&redirect_uri=%s&state=%s",
			clientID, encodedRedirect, testState)

		parResp, err := sdkClient.PushedAuthorization.Create(ctx, serviceID, components.PushedAuthorizationRequest{
			Parameters:   parParams,
			ClientID:     ptr(clientID),
			ClientSecret: ptr(clientSecret),
		})
		require.NoError(t, err, "PAR create failed")
		require.NotNil(t, parResp.PushedAuthorizationResponse, "PAR response must not be nil")

		par := parResp.PushedAuthorizationResponse
		require.NotNil(t, par.Action, "PAR action must not be nil")
		require.Equal(t, components.PushedAuthorizationResponseActionCreated, *par.Action,
			"expected CREATED, got %s: %s", *par.Action, ptrStr(par.ResultMessage))
		require.NotNil(t, par.RequestURI, "request_uri must be present after PAR")

		// Step 2: Authorization request using request_uri
		authParams := fmt.Sprintf("client_id=%s&request_uri=%s",
			clientID, url.QueryEscape(*par.RequestURI))

		authResp, err := sdkClient.Authorization.ProcessRequest(ctx, serviceID, components.AuthorizationRequest{
			Parameters: authParams,
		})
		require.NoError(t, err, "authorization request failed")
		require.NotNil(t, authResp.AuthorizationResponse)
		auth := authResp.AuthorizationResponse
		require.NotNil(t, auth.Action)
		require.Equal(t, components.AuthorizationResponseActionInteraction, *auth.Action,
			"expected INTERACTION, got %s", *auth.Action)
		require.NotNil(t, auth.Ticket)

		// Step 3: Issue authorization code
		issueResp, err := sdkClient.Authorization.Issue(ctx, serviceID, components.AuthorizationIssueRequest{
			Ticket:  *auth.Ticket,
			Subject: subject,
		})
		require.NoError(t, err, "authorization issue failed")
		require.NotNil(t, issueResp.AuthorizationIssueResponse)
		issue := issueResp.AuthorizationIssueResponse
		require.NotNil(t, issue.Action)
		require.Equal(t, components.AuthorizationIssueResponseActionLocation, *issue.Action,
			"expected LOCATION, got %s", *issue.Action)
		require.NotNil(t, issue.AuthorizationCode)

		// Step 4: Token exchange
		tokenParams := fmt.Sprintf("grant_type=authorization_code&code=%s&redirect_uri=%s",
			*issue.AuthorizationCode, encodedRedirect)

		tokenResp, err := sdkClient.Token.Process(ctx, serviceID, components.TokenRequest{
			Parameters:   tokenParams,
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

		assertTokenValid(t, sdkClient, serviceID, *tok.AccessToken)
	})

	t.Run("ParMissingClientSecretRejected", func(t *testing.T) {
		encodedRedirect := url.QueryEscape(redirectURI)

		parParams := fmt.Sprintf("response_type=code&client_id=%s&redirect_uri=%s&state=%s",
			clientID, encodedRedirect, testState)

		parResp, err := sdkClient.PushedAuthorization.Create(ctx, serviceID, components.PushedAuthorizationRequest{
			Parameters: parParams,
			ClientID:   ptr(clientID),
			// ClientSecret intentionally omitted.
		})
		require.NoError(t, err, "PAR create failed")
		require.NotNil(t, parResp.PushedAuthorizationResponse)
		require.NotNil(t, parResp.PushedAuthorizationResponse.Action)
		require.NotEqual(t, components.PushedAuthorizationResponseActionCreated,
			*parResp.PushedAuthorizationResponse.Action,
			"PAR request without client_secret must not succeed for a confidential client")
	})
}

// =============================================================================
// Service with parRequired: true.
// =============================================================================
func TestParRequired(t *testing.T) {
	ctx := context.Background()
	mgmtClient := createSDKClient(mgmtToken)
	sdkClient := createSDKClient(serviceToken)

	_, err := mgmtClient.Service.Update(ctx, serviceID, &components.ServiceInput{
		ParRequired:         ptr(true),
		AccessTokenDuration: ptr(tokenDurationSeconds),
	})
	require.NoError(t, err, "failed to enable parRequired")

	t.Cleanup(func() {
		_, _ = mgmtClient.Service.Update(context.Background(), serviceID, &components.ServiceInput{
			ParRequired:         ptr(false),
			AccessTokenDuration: ptr(tokenDurationSeconds),
		})
	})

	clientID, clientSecret := createTestClient(t, mgmtClient, serviceID)
	t.Cleanup(func() {
		_, _ = mgmtClient.Client.Delete(context.Background(), serviceID, clientID)
	})

	t.Run("DirectAuthRequestRejected", func(t *testing.T) {
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
			"direct auth request must be rejected when parRequired=true")
	})

	t.Run("ParFlowSucceedsWhenRequired", func(t *testing.T) {
		encodedRedirect := url.QueryEscape(redirectURI)

		parParams := fmt.Sprintf("response_type=code&client_id=%s&redirect_uri=%s&state=%s",
			clientID, encodedRedirect, testState)

		parResp, err := sdkClient.PushedAuthorization.Create(ctx, serviceID, components.PushedAuthorizationRequest{
			Parameters:   parParams,
			ClientID:     ptr(clientID),
			ClientSecret: ptr(clientSecret),
		})
		require.NoError(t, err, "PAR create failed")
		require.NotNil(t, parResp.PushedAuthorizationResponse)
		par := parResp.PushedAuthorizationResponse
		require.NotNil(t, par.Action)
		require.Equal(t, components.PushedAuthorizationResponseActionCreated, *par.Action,
			"expected CREATED, got %s: %s", *par.Action, ptrStr(par.ResultMessage))
		require.NotNil(t, par.RequestURI)

		authResp, err := sdkClient.Authorization.ProcessRequest(ctx, serviceID, components.AuthorizationRequest{
			Parameters: fmt.Sprintf("client_id=%s&request_uri=%s",
				clientID, url.QueryEscape(*par.RequestURI)),
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

		tokenParams := fmt.Sprintf("grant_type=authorization_code&code=%s&redirect_uri=%s",
			*issue.AuthorizationCode, encodedRedirect)

		tokenResp, err := sdkClient.Token.Process(ctx, serviceID, components.TokenRequest{
			Parameters:   tokenParams,
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

		assertTokenValid(t, sdkClient, serviceID, *tok.AccessToken)
	})
}
