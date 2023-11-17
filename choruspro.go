package choruspro

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"strings"
	"time"

	"golang.org/x/oauth2"
)

const (
	defaultBaseURL = "https://sandbox-api.piste.gouv.fr/"
	defaultAuthURL = "https://sandbox-oauth.piste.gouv.fr/"
)

// ClientConfig contains the configuration for the client
type ClientConfig struct {
	// Piste URL
	BaseUrl string

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
	BaseUrl *url.URL

	// Piste OAuth URL
	AuthUrl *url.URL

	// Piste client ID
	clientId string

	// Piste client secret
	clientSecret string

	// OAuth token used for authentication
	token *oauth2.Token

	// Chorus Pro technical account credentials (login:password base64 encoded)
	login string

	// Shared between services
	common service

	// Services
	Transverses *TransversesService
	Factures    *FacturesService
}

type service struct {
	client *Client
}

func NewClient() *Client {
	c := &Client{client: http.DefaultClient}
	c.initialize()
	return c
}

func (c *Client) WithConfig(config *ClientConfig) (*Client, error) {
	var err error

	c.BaseUrl, err = url.Parse(config.BaseUrl)
	if err != nil {
		return nil, err
	}

	c.AuthUrl, err = url.Parse(config.AuthUrl)
	if err != nil {
		return nil, err
	}

	c.clientId = config.ClientId
	c.clientSecret = config.ClientSecret
	c.login = config.Login

	return c, err
}

func (c *Client) initialize() {
	c.common.client = c
	c.BaseUrl, _ = url.Parse(defaultBaseURL)
	c.AuthUrl, _ = url.Parse(defaultAuthURL)
	c.Transverses = (*TransversesService)(&c.common)
	c.Factures = (*FacturesService)(&c.common)
}

func (c *Client) newRequest(ctx context.Context, method, url string, body interface{}) (*http.Request, error) {
	if !strings.HasSuffix(c.BaseUrl.Path, "/") {
		return nil, fmt.Errorf("BaseURL must have a trailing slash, but %q does not", c.BaseUrl)
	}

	u, err := c.BaseUrl.Parse(url)
	if err != nil {
		return nil, err
	}

	data, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, method, u.String(), bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", "application/json;charset=utf-8")
	req.Header.Add("cpro-account", c.login)

	return req, nil
}

func (c *Client) doRequest(ctx context.Context, req *http.Request, obj interface{}) error {
	// Check if token is valid, if not, get a new one
	if !c.token.Valid() {
		token, err := getOAuthToken(c.clientId, c.clientSecret, c.AuthUrl)
		if err != nil {
			return err
		}

		// Update token
		c.token = token

		req.Header.Add("Authorization", fmt.Sprintf("Bearer %v", c.token.AccessToken))
	}

	res, err := c.client.Do(req)
	if err != nil {
		return err
	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		if res.StatusCode == http.StatusBadRequest {
			return parseError(res)
		}

		return fmt.Errorf("choruspro: %v", res.Status)
	}

	data, err := io.ReadAll(res.Body)
	// Print
	log.Println(string(data))
	if err != nil {
		return nil
	}

	return json.Unmarshal(data, obj)
}

func getOAuthToken(clientId, clientSecret string, authUrl *url.URL) (*oauth2.Token, error) {
	u, _ := authUrl.Parse("api/oauth/token")
	c := http.DefaultClient

	data := url.Values{}
	data.Set("client_id", clientId)
	data.Set("client_secret", clientSecret)
	data.Set("grant_type", "client_credentials")

	encodedData := data.Encode()

	req, err := http.NewRequest(http.MethodPost, u.String(), strings.NewReader(encodedData))
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	var token struct {
		AccessToken string `json:"access_token"`
		TokenType   string `json:"token_type,omitempty"`
		ExpiresIn   int16  `json:"expires_in,omitempty"`
	}

	res, err := c.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("choruspro: %v", res.Status)
	}

	err = json.NewDecoder(res.Body).Decode(&token)
	if err != nil {
		return nil, err
	}

	tok := &oauth2.Token{
		AccessToken: token.AccessToken,
		TokenType:   token.TokenType,
		Expiry:      time.Now().Add(time.Duration(token.ExpiresIn) * time.Second),
	}

	return tok, nil
}
