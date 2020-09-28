package fitbit

import (
	"context"
	"fmt"
	"net/http"
	"strings"
)

type TokenState struct {
	Active    bool   `json:"active"`
	Scope     string `json:"scope"`
	ClientID  string `json:"client_id"`
	UserID    string `json:"user_id"`
	TokenType string `json:"token_type"`
	Exp       int64  `json:"exp"`
	Iat       int64  `json:"iat"`
}

func (api *Client) GetTokenState(access_token string) (*TokenState, *Response, error) {
	return api.GetTokenStateContext(context.Background(), access_token)
}

func (api *Client) GetTokenStateContext(ctx context.Context, access_token string) (*TokenState, *Response, error) {
	var m TokenState
	req, err := GetTokenStateRequest(ctx, access_token)
	if err != nil {
		return nil, nil, err
	}
	resp, err := api.do_request(req, &m)
	return &m, resp, err
}

func GetTokenState(c *Client, access_token string) (*TokenState, error) {
	return GetTokenStateContext(context.Background(), c, access_token)
}

func GetTokenStateContext(ctx context.Context, c *Client, access_token string) (*TokenState, error) {
	m, _, err := c.GetTokenStateContext(ctx, access_token)
	return m, err
}

func GetTokenStateRequest(ctx context.Context, access_token string) (*http.Request, error) {
	url := "https://api.fitbit.com/1.1/oauth2/introspect"
	body := fmt.Sprintf("token=%s", access_token)
	req, err := http.NewRequestWithContext(ctx, "POST", url, strings.NewReader(body))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	return req, err
}
