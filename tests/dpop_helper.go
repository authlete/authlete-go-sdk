package tests

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

const (
	dpopTokenEndpoint = "https://as.example.com/token"
	dpopResourceURL   = "https://rs.example.com/api/resource"
	dpopUserinfoURL   = "https://as.example.com/userinfo"
)

// generateECKey returns a fresh P-256 ECDSA key for signing DPoP proofs.
func generateECKey(t *testing.T) *ecdsa.PrivateKey {
	t.Helper()
	key, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	require.NoError(t, err, "failed to generate EC key")
	return key
}

// dpopProofOpts is an optional bag of DPoP proof claims.
type dpopProofOpts struct {
	accessToken string
	nonce       string
}

// dpopProof builds and signs a DPoP proof JWT (RFC 9449) as `header.payload.sig`.
//
// header: {typ: "dpop+jwt", alg: "ES256", jwk: {kty:EC, crv:P-256, x, y}}
// payload: {jti, htm, htu, iat, ath?, nonce?}
// signature: ES256 (R||S) over base64url(header) + "." + base64url(payload).
func dpopProof(t *testing.T, key *ecdsa.PrivateKey, htm, htu string, opts ...dpopProofOpts) string {
	t.Helper()

	x := base64URL(leftPad(key.PublicKey.X.Bytes(), 32))
	y := base64URL(leftPad(key.PublicKey.Y.Bytes(), 32))

	header := map[string]any{
		"typ": "dpop+jwt",
		"alg": "ES256",
		"jwk": map[string]any{"kty": "EC", "crv": "P-256", "x": x, "y": y},
	}
	payload := map[string]any{
		"jti": randomJTI(t),
		"htm": htm,
		"htu": htu,
		"iat": time.Now().Unix(),
	}
	if len(opts) > 0 {
		o := opts[0]
		if o.accessToken != "" {
			ath := sha256.Sum256([]byte(o.accessToken))
			payload["ath"] = base64URL(ath[:])
		}
		if o.nonce != "" {
			payload["nonce"] = o.nonce
		}
	}

	headerB, err := json.Marshal(header)
	require.NoError(t, err)
	payloadB, err := json.Marshal(payload)
	require.NoError(t, err)
	signingInput := base64URL(headerB) + "." + base64URL(payloadB)

	digest := sha256.Sum256([]byte(signingInput))
	r, s, err := ecdsa.Sign(rand.Reader, key, digest[:])
	require.NoError(t, err, "ecdsa sign failed")
	sig := append(leftPad(r.Bytes(), 32), leftPad(s.Bytes(), 32)...)

	return fmt.Sprintf("%s.%s", signingInput, base64URL(sig))
}

func randomJTI(t *testing.T) string {
	t.Helper()
	buf := make([]byte, 16)
	_, err := rand.Read(buf)
	require.NoError(t, err)
	return base64URL(buf)
}

func base64URL(b []byte) string {
	return base64.RawURLEncoding.EncodeToString(b)
}

// leftPad pads `b` with leading zero bytes to length `size`. ECDSA R/S values
// can be shorter than the curve byte length when leading bytes are zero, but
// DPoP requires fixed-width R||S.
func leftPad(b []byte, size int) []byte {
	if len(b) >= size {
		return b
	}
	out := make([]byte, size)
	copy(out[size-len(b):], b)
	return out
}
