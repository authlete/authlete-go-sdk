package tests

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"
	"testing"

	authlete "github.com/authlete/authlete-go-sdk"
	"github.com/authlete/authlete-go-sdk/models/components"
	"github.com/stretchr/testify/require"
)

// =============================================================================
// Extra Properties — attached at authorization issue or at the token endpoint.
// Hidden properties must not appear in response_content but must appear in
// introspection.
// Ref: https://www.authlete.com/developers/definitive_guide/extra_properties/
// =============================================================================

var (
	visibleProp = components.Property{
		Key:   ptr("tenant_id"),
		Value: ptr("acme-corp"),
	}
	hiddenProp = components.Property{
		Key:    ptr("internal_user_tier"),
		Value:  ptr("premium"),
		Hidden: ptr(true),
	}
)

func obtainTicket(
	t *testing.T,
	ctx context.Context,
	sdk *authlete.Authlete,
	svcID, clientID string,
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
	require.NotNil(t, authResp.AuthorizationResponse.Action)
	require.Equal(t, components.AuthorizationResponseActionInteraction,
		*authResp.AuthorizationResponse.Action)
	require.NotNil(t, authResp.AuthorizationResponse.Ticket)
	return *authResp.AuthorizationResponse.Ticket
}

// findProperty returns the first property whose Key matches `key`, or nil.
func findProperty(props []components.Property, key string) *components.Property {
	for i := range props {
		if props[i].Key != nil && *props[i].Key == key {
			return &props[i]
		}
	}
	return nil
}

func TestExtraProperties(t *testing.T) {
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

	t.Run("PropertiesAtAuthorizationIssue", func(t *testing.T) {
		encodedRedirect := url.QueryEscape(redirectURI)

		// Authorization issue — attach properties here.
		issueResp, err := sdkClient.Authorization.Issue(ctx, serviceID, components.AuthorizationIssueRequest{
			Ticket:     obtainTicket(t, ctx, sdkClient, serviceID, clientID),
			Subject:    subject,
			Properties: []components.Property{visibleProp, hiddenProp},
		})
		require.NoError(t, err, "authorization issue failed")
		require.NotNil(t, issueResp.AuthorizationIssueResponse)
		issue := issueResp.AuthorizationIssueResponse
		require.NotNil(t, issue.Action)
		require.Equal(t, components.AuthorizationIssueResponseActionLocation, *issue.Action)
		require.NotNil(t, issue.AuthorizationCode)

		// Token request — no properties here.
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
		tok := tokenResp.TokenResponse
		require.NotNil(t, tok.Action)
		require.Equal(t, components.TokenResponseActionOk, *tok.Action,
			"expected OK, got %s: %s", *tok.Action, ptrStr(tok.ResultMessage))
		require.NotNil(t, tok.AccessToken)
		require.NotNil(t, tok.ResponseContent)

		// SDK deserializes properties array with correct key/value/hidden fields.
		visible := findProperty(tok.Properties, *visibleProp.Key)
		hidden := findProperty(tok.Properties, *hiddenProp.Key)
		require.NotNil(t, visible, "visible property must be in token response")
		require.NotNil(t, hidden, "hidden property must be in token response")
		require.NotNil(t, visible.Value)
		require.Equal(t, *visibleProp.Value, *visible.Value)
		require.NotNil(t, hidden.Value)
		require.Equal(t, *hiddenProp.Value, *hidden.Value)
		require.NotNil(t, hidden.Hidden)
		require.True(t, *hidden.Hidden, "hidden flag must be true")

		// Only visible property appears in response_content.
		var responseJSON map[string]any
		require.NoError(t, json.Unmarshal([]byte(*tok.ResponseContent), &responseJSON))
		require.Equal(t, *visibleProp.Value, responseJSON[*visibleProp.Key])
		_, hasHidden := responseJSON[*hiddenProp.Key]
		require.False(t, hasHidden, "hidden property must not appear in response_content")

		// Both accessible via introspection.
		introResp, err := sdkClient.Introspection.Process(ctx, serviceID, components.IntrospectionRequest{
			Token: *tok.AccessToken,
		})
		require.NoError(t, err, "introspection request failed")
		require.NotNil(t, introResp.IntrospectionResponse)
		introVisible := findProperty(introResp.IntrospectionResponse.Properties, *visibleProp.Key)
		introHidden := findProperty(introResp.IntrospectionResponse.Properties, *hiddenProp.Key)
		require.NotNil(t, introVisible, "visible property must be exposed via introspection")
		require.NotNil(t, introHidden, "hidden property must be exposed via introspection")
	})

	t.Run("PropertiesAtTokenEndpoint", func(t *testing.T) {
		encodedRedirect := url.QueryEscape(redirectURI)

		// Authorization issue — no properties here.
		issueResp, err := sdkClient.Authorization.Issue(ctx, serviceID, components.AuthorizationIssueRequest{
			Ticket:  obtainTicket(t, ctx, sdkClient, serviceID, clientID),
			Subject: subject,
		})
		require.NoError(t, err, "authorization issue failed")
		require.NotNil(t, issueResp.AuthorizationIssueResponse)
		issue := issueResp.AuthorizationIssueResponse
		require.NotNil(t, issue.Action)
		require.Equal(t, components.AuthorizationIssueResponseActionLocation, *issue.Action)
		require.NotNil(t, issue.AuthorizationCode)

		// Token request — attach properties here.
		tokenResp, err := sdkClient.Token.Process(ctx, serviceID, components.TokenRequest{
			Parameters: fmt.Sprintf(
				"grant_type=authorization_code&code=%s&redirect_uri=%s",
				*issue.AuthorizationCode, encodedRedirect,
			),
			ClientID:     ptr(clientID),
			ClientSecret: ptr(clientSecret),
			Properties:   []components.Property{visibleProp},
		})
		require.NoError(t, err, "token request failed")
		require.NotNil(t, tokenResp.TokenResponse)
		tok := tokenResp.TokenResponse
		require.NotNil(t, tok.Action)
		require.Equal(t, components.TokenResponseActionOk, *tok.Action,
			"expected OK, got %s: %s", *tok.Action, ptrStr(tok.ResultMessage))
		require.NotNil(t, tok.ResponseContent)

		var responseJSON map[string]any
		require.NoError(t, json.Unmarshal([]byte(*tok.ResponseContent), &responseJSON))
		require.Equal(t, *visibleProp.Value, responseJSON[*visibleProp.Key])
	})
}
