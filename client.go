package bing

import (
	"context"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"golang.org/x/oauth2"
)

// Session contains the client information needed for a current session.
type Session struct {
	AccountId      string
	CustomerId     string
	DeveloperToken string
	ClientID       string
	ClientSecret   string
	Username       string
	Password       string
	HTTPClient     HttpClient
	TokenSource    oauth2.TokenSource
}

// AuthConfig contains the fields needed to start an authentication process.
type AuthConfig struct {
	AccountId      string
	CustomerID     string
	DeveloperToken string
	Oauth2Config   *oauth2.Config
	Oauth2Token    *oauth2.Token
	TokenSource    oauth2.TokenSource
}

type HttpClient interface {
	Do(*http.Request) (*http.Response, error)
}

type OAuthTokens struct {
	AccessToken  string `json:"access_token"`
	ExpiresIn    int    `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
}

// NewAuth starts a new bing session.
func NewAuth(config AuthConfig) *Session {
	ctx := context.TODO()
	return &Session{
		AccountId:      config.AccountId,
		CustomerId:     config.CustomerID,
		DeveloperToken: config.DeveloperToken,
		ClientID:       config.Oauth2Config.ClientID,
		ClientSecret:   config.Oauth2Config.ClientSecret,
		HTTPClient:     config.Oauth2Config.Client(ctx, config.Oauth2Token),
		TokenSource: config.Oauth2Config.TokenSource(context.TODO(), &oauth2.Token{
			AccessToken:  config.Oauth2Token.AccessToken,
			RefreshToken: config.Oauth2Token.RefreshToken,
		}),
	}
}

// refreshToken is a process to renew the token once it has expired.
func (s *Session) refreshToken() error {
	tokenUrl := "https://login.microsoftonline.com/common/oauth2/v2.0/token"
	contentType := "application/x-www-form-urlencoded"
	token, _ := s.TokenSource.Token()

	// Set body values
	data := url.Values{}
	data.Set("client_id", s.ClientID)
	data.Set("scope", "https://ads.microsoft.com/msads.manage offline_access")
	data.Set("refresh_token", token.RefreshToken)
	data.Set("client_secret", s.ClientSecret)
	data.Set("grant_type", "refresh_token")
	encodedData := data.Encode()
	body := strings.NewReader(encodedData)

	// Http request
	resp, err := http.Post(tokenUrl, contentType, body)
	if err != nil {
		return errors.New("Error refreshtoken http Post request: " + err.Error())
	}
	defer resp.Body.Close()

	// Unmarshal
	bytes, _ := ioutil.ReadAll(resp.Body)
	tokens := &OAuthTokens{}
	err = json.Unmarshal(bytes, tokens)
	if err != nil {
		return errors.New("Error Unmarshaling refreshtoken response: " + err.Error())
	}

	// Overwriting tokens on the session
	config := oauth2.Config{}
	ts := config.TokenSource(context.TODO(), &oauth2.Token{
		AccessToken:  tokens.AccessToken,
		RefreshToken: tokens.RefreshToken,
	})
	s.TokenSource = ts
	return nil
}
