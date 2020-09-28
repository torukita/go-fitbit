package fitbit

import (
	"context"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/fitbit"
)

type AuthConfig struct {
	ClientID     string
	ClientSecret string
	RedirectURL  string
}

func newConfig() *oauth2.Config {
	return &oauth2.Config{
		Endpoint: fitbit.Endpoint,
		Scopes:   []string{"activity", "location", "social", "heartrate", "settings", "sleep", "weight", "profile", "nutrition"},
	}
}

// GenerateToken returns the Token generated by authorization code grant flow.
func GenerateToken(c *AuthConfig, code string) (Token, error) {
	cfg := newConfig()
	cfg.ClientID = c.ClientID
	cfg.ClientSecret = c.ClientSecret
	cfg.RedirectURL = c.RedirectURL
	ctx := context.WithValue(context.Background(), oauth2.HTTPClient, DefaultHttpClient)
	token, err := cfg.Exchange(ctx, code)
	if err != nil {
		return Token{}, err
	}
	return fromToken(token), nil
}

// AuthCodeURL returns the URL to fitibit auth page with authorization code grant flow.
// expires_in parameter is default 8hours.
func AuthCodeURL(c *AuthConfig) string {
	cfg := newConfig()
	cfg.ClientID = c.ClientID
	cfg.ClientSecret = c.ClientSecret
	cfg.RedirectURL = c.RedirectURL
	return cfg.AuthCodeURL("", oauth2.SetAuthURLParam("prompt", "consent"))
}

// ImplicitURL returns the URL to fitbit auth page with implicit grant flow.
// Expires parameter can be one of 86400(1day) 604800(1week) 2592000(30days) 31536000(1year)
func ImplicitURL(c *AuthConfig, expires string) string {
	switch expires {
	case "86400", "604800", "2592000", "31536000":
	default:
		expires = "86400"
	}
	cfg := newConfig()
	cfg.ClientID = c.ClientID
	cfg.ClientSecret = c.ClientSecret
	cfg.RedirectURL = c.RedirectURL
	// response_type is just token instead of code.
	return cfg.AuthCodeURL("",
		oauth2.SetAuthURLParam("response_type", "token"),
		oauth2.SetAuthURLParam("prompt", "consent"),
		oauth2.SetAuthURLParam("expires_in", expires),
	)
}
