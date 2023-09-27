package choruspro

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func setup() (client *Client, mux *http.ServeMux, teardown func()) {
	mux = http.NewServeMux()

	// OAuth token mock
	mux.HandleFunc("/api/oauth/token", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, `{"access_token":"token","token_type":"Bearer","expires_in":3600}`)
	})

	server := httptest.NewServer(mux)

	client = NewClient()
	client.BaseUrl = server.URL
	client.AuthUrl = server.URL + "/api/oauth/token"

	return client, mux, server.Close
}

func testMethod(t *testing.T, r *http.Request, want string) {
	if got := r.Method; got != want {
		t.Errorf("Request method : %v, got %v", got, want)
	}
}
