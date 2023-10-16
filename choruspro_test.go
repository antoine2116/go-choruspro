package choruspro

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"reflect"
	"testing"
	"time"

	"golang.org/x/oauth2"
)

const (
	baseAuthPath = "/api/oauth/token"
)

func setup() (client *Client, mux *http.ServeMux, teardown func()) {
	client, mux, teardown = bareSetup()

	// OAuth token
	mux.HandleFunc(baseAuthPath, func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, `{"access_token":"token","token_type":"Bearer","expires_in":3600}`)
	})

	return client, mux, teardown
}

func bareSetup() (client *Client, mux *http.ServeMux, teardown func()) {
	mux = http.NewServeMux()
	server := httptest.NewServer(mux)

	client = NewClient()

	url, _ := client.BaseUrl.Parse(server.URL + "/")
	client.BaseUrl = url
	client.AuthUrl = url

	return client, mux, server.Close
}

// Test function under s.client.NewRequest or s.client.DoRequest failure.
// Method f should be a regular call that would normally succeed, but
// should return an error when NewRequest or s.client.DoRequest fails.
func testNewRequestAndDoRequestFailure(t *testing.T, methodName string, client *Client, f func() error) {
	t.Helper()
	if methodName == "" {
		t.Error("testNewRequestAndDoFailure : method name should be provided")
	}

	client.BaseUrl.Path = ""

	err := f()
	if err == nil {
		t.Errorf("client.BaseURL.Path='' %v err = nil, want error", methodName)
	}

	client.BaseUrl.Path = "/v1/"

	err = f()
	if err == nil {
		t.Errorf("client.AuthURL.Path='' %v err = nil, want error", methodName)
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

func TestWithConfig(t *testing.T) {
	cfg := &ClientConfig{
		BaseUrl:      "https://api.com/v1/",
		AuthUrl:      "https://api.com/oauth/token/",
		ClientId:     "client_id",
		ClientSecret: "client_secret",
		Login:        "login",
	}

	c, err := NewClient().WithConfig(cfg)

	if err != nil {
		t.Errorf("Client.WithConfig returned err = %v, want nil", err)
	}

	if c.BaseUrl.String() != cfg.BaseUrl {
		t.Errorf("Client.BaseUrl = %v, want %v", c.BaseUrl, cfg.BaseUrl)
	}

	if c.AuthUrl.String() != cfg.AuthUrl {
		t.Errorf("Client.AuthUrl = %v, want %v", c.AuthUrl, cfg.AuthUrl)
	}

	if c.clientId != cfg.ClientId {
		t.Errorf("Client.ClientId = %v, want %v", c.clientId, cfg.ClientId)
	}

	if c.clientSecret != cfg.ClientSecret {
		t.Errorf("Client.ClientSecret = %v, want %v", c.clientSecret, cfg.ClientSecret)
	}

	if c.login != cfg.Login {
		t.Errorf("Client.Login = %v, want %v", c.login, cfg.Login)
	}
}

func TestWithConfig_badUrls(t *testing.T) {
	tests := []struct {
		name    string
		baseUrl string
		authUrl string
	}{
		{
			name:    "Bad base url",
			baseUrl: "bad\nbase\nurl",
			authUrl: "https://api.com/oauth/token/",
		},
		{
			name:    "Bad auth url",
			baseUrl: "https://api.com/v1/",
			authUrl: "bad\nauth\nurl",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := NewClient().WithConfig(&ClientConfig{
				BaseUrl: tt.baseUrl,
				AuthUrl: tt.authUrl,
			})
			if err == nil {
				t.Errorf("Client.WithConfig returned err = nil, want error")
			}
		})
	}
}

func TestNewRequest(t *testing.T) {
	c := NewClient()
	ctx := context.Background()

	inURL, outURL := "/foo", defaultBaseURL+"foo"
	inBody, outBody := &codeLangueOptions{CodeLangueFr}, `{"codeLangue":"FR"}`

	req, _ := c.newRequest(ctx, http.MethodPost, inURL, inBody)

	// test relative URL was expanded
	if got, want := req.URL.String(), outURL; got != want {
		t.Errorf("newRequest(%v) URL = %v, want %v", inURL, req.URL, outURL)
	}

	// test that body was JSON encoded
	body, _ := io.ReadAll(req.Body)
	if got, want := string(body), outBody; got != want {
		t.Errorf("NewRequest(%q) Body is %v, want %v", inBody, got, want)
	}
}

func TestNewRequest_badUrl(t *testing.T) {
	c := NewClient()
	ctx := context.Background()

	_, err := c.newRequest(ctx, http.MethodGet, ":", nil)
	if err == nil {
		t.Errorf("Client.NewRequest returned err = nil, want error")
	}
}

func TestNewRequest_badBody(t *testing.T) {
	c := NewClient()
	ctx := context.Background()

	_, err := c.newRequest(ctx, http.MethodGet, "/", make(chan int))
	if err == nil {
		t.Errorf("Client.NewRequest returned err = nil, want error")
	}
}

func TestNewRequest_badMethod(t *testing.T) {
	c := NewClient()
	ctx := context.Background()

	_, err := c.newRequest(ctx, "BAD METHOD", "/", nil)
	if err == nil {
		t.Errorf("Client.NewRequest returned err = nil, want error")
	}
}

func TestDoRequest(t *testing.T) {
	client, mux, terdown := setup()
	defer terdown()

	type testStruct struct {
		Login string
	}

	client.token = &oauth2.Token{
		AccessToken: "access_token",
		TokenType:   "Bearer",
		Expiry:      time.Now().Add(time.Hour),
	}

	mux.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprint(w, `{"login":"login"}`)
	})

	ctx := context.Background()
	req, _ := client.newRequest(ctx, http.MethodGet, "/test", nil)
	body := new(testStruct)

	err := client.doRequest(ctx, req, body)
	assertNilError(t, err)

	want := &testStruct{"login"}
	if !reflect.DeepEqual(body, want) {
		t.Errorf("Client.doRequest returned %+v, want %+v", body, want)
	}
}

func TestDoRequest_badAuthUrl(t *testing.T) {
	client, _, terdown := setup()
	defer terdown()

	client.AuthUrl, _ = url.Parse("/bad/url")

	ctx := context.Background()
	req, _ := client.newRequest(ctx, http.MethodGet, "/", nil)

	err := client.doRequest(ctx, req, nil)
	if err == nil {
		t.Errorf("Client.doRequest returned err = nil, want error")
	}
}

// T
func TestDo_badUrl(t *testing.T) {
	client, _, teardown := setup()
	defer teardown()

	ctx := context.Background()
	req, _ := client.newRequest(ctx, http.MethodGet, ".", nil)
	req.URL.Path = ":"

	err := client.doRequest(ctx, req, nil)
	if err == nil {
		t.Errorf("Client.doRequest returned err = nil, want error")
	}
}

// Test handling of an error caused by the internal http client's Do() function.
func TestDo_redirectLoop(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/api", http.StatusFound)
	})

	ctx := context.Background()
	req, _ := client.newRequest(ctx, http.MethodGet, ".", nil)

	err := client.doRequest(ctx, req, nil)
	if err == nil {
		t.Errorf("Client.doRequest returned err = nil, want error")
	}
}

func TestGetOAuthToken(t *testing.T) {
	c, _, teardown := setup()
	defer teardown()

	token, err := getOAuthToken("", "", c.AuthUrl)

	assertNilError(t, err)

	want := &oauth2.Token{
		AccessToken: "token",
		TokenType:   "Bearer",
		Expiry:      time.Now().Add(time.Hour), // unable to test this (time is not equal cause of nanoseconds)
	}

	if token.AccessToken != want.AccessToken {
		t.Errorf("getOAuthToken returned AccessToken = %+v, want %+v", token, want)
	}

	if token.TokenType != want.TokenType {
		t.Errorf("getOAuthToken returned TokenType = %+v, want %+v", token, want)
	}
}

func TestGetOAuthToken_httpError(t *testing.T) {
	client, mux, teardown := bareSetup()
	defer teardown()

	mux.HandleFunc(baseAuthPath, func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Bad Request", http.StatusBadRequest)
	})

	_, err := getOAuthToken("", "", client.AuthUrl)
	if err == nil {
		t.Errorf("getOAuthToken returned err = nil, want error")
	}
}

func TestGetOAuthToken_nilBody(t *testing.T) {
	client, mux, teardown := bareSetup()
	defer teardown()

	mux.HandleFunc(baseAuthPath, func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "")
	})

	_, err := getOAuthToken("", "", client.AuthUrl)
	if err == nil {
		t.Errorf("getOAuthToken returned err = nil, want error")
	}
}
