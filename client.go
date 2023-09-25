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

	// OAuth token used for authentication
	token *oauth2.Token

	// Chorus Pro technical account credentials (login:password base64 encoded)
	login string

	// Shared between services
	common service

	// Services
	Transverses *TransversesService
}

type service struct {
	client *Client
}

func NewClient(config *ClientConfig) *Client {
	c := &Client{
		client:       http.DefaultClient,
		url:          config.Url,
		authUrl:      config.AuthUrl,
		clientId:     config.ClientId,
		clientSecret: config.ClientSecret,
		login:        config.Login,
	}

	c.initialize()
	return c
}

func (c *Client) initialize() {
	c.common.client = c
	c.Transverses = (*TransversesService)(&c.common)
}

func (c *Client) newRequest(ctx context.Context, method, url string, body interface{}) (*http.Request, error) {
	// Check if token is valid, if not, get a new one
	if !c.token.Valid() {
		token, err := getOAuthToken(c.authUrl, c.clientId, c.clientSecret)
		if err != nil {
			return nil, err
		}

		// Update token
		c.token = token
	}

	data, err := json.Marshal(body)

	// Temp
	log.Printf("Request body : \n%s", string(data))

	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, method, c.url+url, bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}

	req.Header.Add("Authorization", fmt.Sprintf("Bearer %v", c.token.AccessToken))
	req.Header.Add("Content-Type", "application/json;charset=utf-8")
	req.Header.Add("cpro-account", c.login)

	return req, nil
}

func (c *Client) doRequest(ctx context.Context, req *http.Request, obj interface{}) error {
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
	log.Printf("Response body : \n%s", string(data))
	if err != nil {
		return nil
	}

	return json.Unmarshal(data, obj)
}

func getOAuthToken(authUrl, clientId, clientSecret string) (*oauth2.Token, error) {
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
