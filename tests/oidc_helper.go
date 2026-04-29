package tests

import (
	"context"
	"crypto/rand"
	"crypto/rsa"
	"encoding/base64"
	"encoding/binary"
	"encoding/json"
	"math/big"
	"strings"
	"testing"

	authlete "github.com/authlete/authlete-go-sdk"
	"github.com/authlete/authlete-go-sdk/models/components"
	"github.com/stretchr/testify/require"
)

const oidcIssuer = "https://as.example.com"

// oidcJWKS returns a JSON-encoded JWKS containing one RSA-2048 private key
// suitable for signing id_tokens (RS256). The kid is also returned so the
// caller can wire it up via Service.idTokenSignatureKeyId.
func generateRSAJWKS(t *testing.T) (jwks string, kid string) {
	t.Helper()
	key, err := rsa.GenerateKey(rand.Reader, 2048)
	require.NoError(t, err, "rsa key generation failed")
	key.Precompute()
	require.GreaterOrEqual(t, len(key.Primes), 2, "rsa key must have two primes")

	kid = randomKID(t)

	jwk := map[string]string{
		"kty": "RSA",
		"kid": kid,
		"use": "sig",
		"alg": "RS256",
		"n":   b64BigInt(key.N),
		"e":   b64Int(key.E),
		"d":   b64BigInt(key.D),
		"p":   b64BigInt(key.Primes[0]),
		"q":   b64BigInt(key.Primes[1]),
		"dp":  b64BigInt(key.Precomputed.Dp),
		"dq":  b64BigInt(key.Precomputed.Dq),
		"qi":  b64BigInt(key.Precomputed.Qinv),
	}
	out, err := json.Marshal(map[string]any{"keys": []map[string]string{jwk}})
	require.NoError(t, err)
	return string(out), kid
}

// setupOIDCService updates the service with the OIDC settings needed for
// id_token issuance: issuer, RSA JWKS, the kid to use for id_token signing,
// the openid scope, and (optionally) a token endpoint for DPoP-style flows.
func setupOIDCService(
	t *testing.T,
	ctx context.Context,
	mgmt *authlete.Authlete,
	svcID string,
	tokenEndpoint *string,
) {
	t.Helper()
	jwks, kid := generateRSAJWKS(t)

	openidName := "openid"
	defaultEntry := false
	_, err := mgmt.Service.Update(ctx, svcID, &components.ServiceInput{
		Issuer:                ptr(oidcIssuer),
		Jwks:                  ptr(jwks),
		IDTokenSignatureKeyID: ptr(kid),
		TokenEndpoint:         tokenEndpoint,
		AccessTokenDuration:   ptr(tokenDurationSeconds),
		SupportedScopes: []components.Scope{
			{Name: &openidName, DefaultEntry: &defaultEntry},
		},
	})
	require.NoError(t, err, "OIDC service setup failed")
}

// decodeJWTPayload returns the JSON-decoded payload (claims) of `jwt`. It does
// NOT verify the signature — tests run against a known-good signer.
func decodeJWTPayload(t *testing.T, jwt string) map[string]any {
	t.Helper()
	parts := strings.Split(jwt, ".")
	require.GreaterOrEqual(t, len(parts), 2, "id_token must have at least header.payload")

	payload, err := base64.RawURLEncoding.DecodeString(parts[1])
	require.NoError(t, err, "id_token payload base64 decode failed")

	var claims map[string]any
	require.NoError(t, json.Unmarshal(payload, &claims), "id_token payload JSON decode failed")
	return claims
}

func assertOIDCClaims(
	t *testing.T,
	claims map[string]any,
	expectedSub, expectedNonce, expectedClientID string,
) {
	t.Helper()
	require.Equal(t, expectedSub, claims["sub"], "id_token sub must equal expected subject")
	require.Equal(t, expectedNonce, claims["nonce"], "id_token nonce must match the request nonce")
	require.NotEmpty(t, claims["iss"], "id_token must include an iss claim")

	switch aud := claims["aud"].(type) {
	case string:
		require.Equal(t, expectedClientID, aud, "id_token aud must equal client_id")
	case []any:
		var found bool
		for _, a := range aud {
			if s, ok := a.(string); ok && s == expectedClientID {
				found = true
				break
			}
		}
		require.True(t, found, "id_token aud must include client_id %s, got %v", expectedClientID, aud)
	default:
		t.Fatalf("id_token aud has unexpected type %T", aud)
	}
}

func b64BigInt(n *big.Int) string {
	return base64.RawURLEncoding.EncodeToString(n.Bytes())
}

func b64Int(n int) string {
	// RFC 7518 §6.3.1.2: the integer must be encoded big-endian with no
	// leading zero bytes. Strip leading zeros from a fixed 4-byte buffer.
	var buf [4]byte
	binary.BigEndian.PutUint32(buf[:], uint32(n))
	i := 0
	for i < 3 && buf[i] == 0 {
		i++
	}
	return base64.RawURLEncoding.EncodeToString(buf[i:])
}

func randomKID(t *testing.T) string {
	t.Helper()
	buf := make([]byte, 16)
	_, err := rand.Read(buf)
	require.NoError(t, err)
	return base64.RawURLEncoding.EncodeToString(buf)
}
