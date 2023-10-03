package choruspro

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func setup() (client *Client, mux *http.ServeMux, teardown func()) {
	mux = http.NewServeMux()

	// OAuth token
	mux.HandleFunc("/api/oauth/token", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, `{"access_token":"token","token_type":"Bearer","expires_in":3600}`)
	})

	server := httptest.NewServer(mux)

	client = NewClient()
	url, _ := client.BaseUrl.Parse(server.URL + "/")
	client.BaseUrl = url
	client.AuthUrl = url

	return client, mux, server.Close
}

// Test function under s.client.DoRequest failure.
// Method f should be a regular call that would normally succeed, but
// should return an error when NewRequest or s.client.DoRequest fails.
func testDoRequestFailure(t *testing.T, methodName string, client *Client, f func() error) {
	t.Helper()
	if methodName == "" {
		t.Error("testNewRequestAndDoFailure : method name should be provided")
	}

	client.BaseUrl.Path = ""
	err := f()

	if err == nil {
		t.Errorf("client.BaseURL.Path='' %v err = nil, want error", methodName)
	}
}

func testMethod(t *testing.T, r *http.Request, want string) {
	t.Helper()
	if got := r.Method; got != want {
		t.Errorf("Request method : %v, got %v", got, want)
	}
}

func assertNilError(t *testing.T, err error) {
	t.Helper()
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
}
