package tests

import (
	"fmt"
	"os"
	"testing"
)

// TestMain skips the integration tests when credentials are not configured, so
// that SDK generation (which compiles and runs this package) does not fail.
// CI sets REQUIRE_INTEGRATION_ENV so missing credentials fail instead of skip.
func TestMain(m *testing.M) {
	if apiBaseURL == "" || serviceID == "" || serviceToken == "" {
		msg := "API_BASE_URL, SERVICE_ID and SERVICE_TOKEN must be set to run integration tests"
		if os.Getenv("REQUIRE_INTEGRATION_ENV") != "" {
			fmt.Fprintln(os.Stderr, msg)
			os.Exit(1)
		}
		fmt.Println("skipping:", msg)
		os.Exit(0)
	}
	os.Exit(m.Run())
}
