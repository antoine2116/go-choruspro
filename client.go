package choruspro

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

// ClientConfig contains the configuration for the client
type ClientConfig struct {
	// Piste URL
	Url string

	// Piste OAuth URL
	AuthUrl string

	// Piste client ID
	ClientId string

	// Piste client secret
	ClientSecret string

	// Chorus Pro technical credentials (login:password base64 encoded)
	Login string
}

// Client is used for HTTP requests to the Chorus Pro API
// and it routes these requests through the Piste API
type Client struct {
	// HTTP client used to communicate with the API.
	client *http.Client

	// Piste URL
	url string

	// Piste OAuth URL
	authUrl string

	// Piste client ID
	clientId string

	// Piste client secret
	clientSecret string

	// Chorus Pro technical credentials (login:password base64 encoded)
	login string
}

type OAuthParams struct {
}

type OAuthToken struct {
	// AccessToken is the token that authorizes and authenticates the requests
	AccessToken string `json:"access_token"`

	// TokenType is the type of token.
	TokenType string `json:"token_type,omitempty"`

	// ExpiresIn is the optional expiration time of the access token.
	ExpiresIn int16 `json:"expires_in,omitempty"`

	//  Scope specifies optional requested permissions.
	Scope string `json:"scope,omitempty"`
}

func NewClient(config *ClientConfig) *Client {
	return &Client{
		client:       http.DefaultClient,
		url:          config.Url,
		authUrl:      config.AuthUrl,
		clientId:     config.ClientId,
		clientSecret: config.ClientSecret,
		login:        config.Login,
	}
}

func (c *Client) doRequest(ctx context.Context, req *http.Request, obj interface{}) error {
	res, err := c.client.Do(req)
	if err != nil {
		return err
	}

	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("error : status code %v", res.Status)
	}

	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return nil
	}

	return json.Unmarshal(data, obj)
}

func (c *Client) newRequest(ctx context.Context, method, url string, body interface{}) (*http.Request, error) {
	// TODO : cache management
	token, err := getOAuthToken(c.authUrl, c.clientId, c.clientSecret)
	if err != nil {
		return nil, err
	}

	data, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, method, c.url+url, bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}

	req.Header.Add("Authorization", fmt.Sprintf("Bearer %v", token.AccessToken))
	req.Header.Add("Content-Type", "application/json;charset=utf-8")
	req.Header.Add("cpro-account", c.login)

	return req, nil
}

func getOAuthToken(authUrl, clientId, clientSecret string) (*OAuthToken, error) {
	c := http.DefaultClient

	data := url.Values{}
	data.Set("client_id", clientId)
	data.Set("client_secret", clientSecret)
	data.Set("grant_type", "client_credentials")

	encodedData := data.Encode()

	req, err := http.NewRequest(http.MethodPost, authUrl+"/api/oauth/token", strings.NewReader(encodedData))
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	res, err := c.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	token := OAuthToken{}

	err = json.NewDecoder(res.Body).Decode(&token)
	if err != nil {
		return nil, err
	}

	return &token, nil
}
