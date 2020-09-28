package fitbit

import (
	"errors"
	"io"
	"io/ioutil"
	"net/http"
	"time"

	"golang.org/x/oauth2"
)

var (
	DefaultHttpClient = &http.Client{Timeout: 30 * time.Second}
)

type Client struct {
	httpClient *http.Client
}

func New(access_token string) *Client {
	token := Token{AccessToken: access_token}
	return NewClient(token.Client())
}

func NewClient(httpClient *http.Client) *Client {
	if httpClient == nil {
		httpClient = DefaultHttpClient // without authorization
	}
	return &Client{
		httpClient: httpClient,
	}
}

func (c *Client) Token() (Token, error) {
	v, ok := c.httpClient.Transport.(*oauth2.Transport)
	if !ok {
		return Token{}, errors.New("should be aauth2.Transport to use Token()")
	}
	t, err := v.Source.Token()
	if err != nil {
		return Token{}, err
	}
	return fromToken(t), nil
}

// Do returns raw response without processing Body responsed.
// The Caller must handle Body.Close() function
func (c *Client) Do(r *http.Request) (*Response, error) {
	resp, err := c.httpClient.Do(r)
	if err != nil {
		return nil, err
	}
	// Body.Close() depends on caller
	return newResponse(resp), nil
}

// to be replaced with current Do function
func (c *Client) do_request(r *http.Request, v interface{}) (*Response, error) {
	resp, err := c.httpClient.Do(r)
	if err != nil {
		return nil, err
	}
	defer func() { // in the case of that body is not processed or not closed
		io.Copy(ioutil.Discard, resp.Body)
		resp.Body.Close()
	}()
	response := newResponse(resp)
	if v == nil {
		return nil, errors.New("v interface must not be null")
	}
	switch val := v.(type) {
	case io.Writer:
		if err := response.Output(val); err != nil {
			return response, err
		}
	default:
		if err := response.Decode(v); err != nil {
			return response, err
		}
	}
	return response, nil
}
