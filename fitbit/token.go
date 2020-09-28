package fitbit

import (
	"context"
	"net/http"
	"time"

	"golang.org/x/oauth2"
)

type Token struct {
	AccessToken  string    `json:"access_token"`
	RefreshToken string    `json:"refresh_token,omitempty"`
	Expiry       time.Time `json:"expiry,omitempty"`
}

func toToken(token *Token) oauth2.Token {
	return oauth2.Token{
		AccessToken:  token.AccessToken,
		RefreshToken: token.RefreshToken,
		Expiry:       token.Expiry,
		TokenType:    "Bearer",
	}
}

func fromToken(token *oauth2.Token) Token {
	return Token{
		AccessToken:  token.AccessToken,
		RefreshToken: token.RefreshToken,
		Expiry:       token.Expiry,
	}
}

func (t *Token) ClientWithID(id, secret string) *http.Client {
	httpClient := DefaultHttpClient
	ctx := context.WithValue(context.Background(), oauth2.HTTPClient, httpClient)
	cfg := newConfig()
	cfg.ClientID = id
	cfg.ClientSecret = secret
	tk := toToken(t)
	return cfg.Client(ctx, &tk)
}

func (t *Token) Client() *http.Client {
	return t.ClientWithID("", "")
}
