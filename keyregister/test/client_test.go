package keyregister

import (
	"fmt"
	"math/rand"
	"os"
	"testing"
)

func getDefaultOrg() string {
	defaultorg := os.Getenv("keyregister_DEFAULT_ORG")

	if defaultorg == "" {
		return "keyregister"
	}

	return defaultorg
}

func newTestClient(t *testing.T) *Client {
	authtoken := os.Getenv("keyregister_AUTH_TOKEN")
	endpoint := os.Getenv("keyregister_ENDPOINT")

	if authtoken == "" {
		t.Fatal("Need an authtoken to continue. Please set keyregister_AUTH_TOKEN")
	}

	if endpoint == "" {
		endpoint = "http://keyregister.io/api/0/"
	}

	client, clienterr := NewClient(authtoken, &endpoint, nil)
	if clienterr != nil {
		t.Fatal(clienterr)
	}

	return client

}

func generateIdentifier(prefix string) string {
	return fmt.Sprintf("%s %d", prefix, rand.Int())
}

func TestClientBadEndpoint(t *testing.T) {
	t.Parallel()
	badendpoint := ""

	authtoken := os.Getenv("keyregister_AUTH_TOKEN")

	_, berr := NewClient(authtoken, &badendpoint, nil)
	if berr == nil {
		t.Error("Should have gotten an error for an empty endpoint")
	}
}

func TestClientKnownGoodEndpoint(t *testing.T) {
	bclient, berr := NewClient("testauthclient", nil, nil)
	if berr != nil {
		t.Error(berr)
	}
	if bclient.Endpoint != "https://keyregister.io/api/0" {
		t.Errorf("Endpoint is not https://keyregister.io/api/0 got %s", bclient.Endpoint)
	}
}

func TestNewRequestWillNotAddExtraTrailingSlashToEndpoint(t *testing.T) {
	endpoint := "some-endpoint/"
	bclient, berr := NewClient("testauthclient", nil, nil)
	if berr != nil {
		t.Error(berr)
	}
	req, err := bclient.newRequest("get", endpoint, nil)
	if req == nil || err != nil {
		t.Errorf("can't generate request: %v", err)
	}

	if req.URL.String() != "https://keyregister.io/api/0/some-endpoint/" {
		t.Errorf("Endpoint is not https://keyregister.io/api/0/some-endpoint/ got %s", req.URL.String())
	}
}
