package fitbit

import (
	"context"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func setup_client() *Client {
	return New(os.Getenv("TEST_TOKEN"))
}

func TestClientToken(t *testing.T) {
	t.Run("success case", func(t *testing.T) {
		token := &Token{
			AccessToken: os.Getenv("TEST_TOKEN"),
		}
		c := NewClient(token.Client())
		got, err := c.Token()
		if assert.NoError(t, err) {
			assert.Equal(t, *token, got)
		}
	})
	t.Run("fail case", func(t *testing.T) {
		c := NewClient(nil) // not oauth2 transport
		_, err := c.Token()
		assert.Error(t, err)
	})
}

func TestClientDo(t *testing.T) {
	ctx := context.Background()
	t.Run("success case", func(t *testing.T) {
		token := &Token{
			AccessToken: os.Getenv("TEST_TOKEN"),
		}
		c := NewClient(token.Client())
		req, err := GetTokenStateRequest(ctx, token.AccessToken)
		assert.NoError(t, err)
		resp, err := c.Do(req)
		assert.NoError(t, err)
		defer resp.Body.Close()
		var m TokenState
		err = resp.Decode(&m)
		if assert.NoError(t, err) {
			assert.Equal(t, true, m.Active)
		}
	})
	t.Run("fail case", func(t *testing.T) {
		token := &Token{
			AccessToken: "",
		}
		c := NewClient(token.Client()) // client with no token
		req, err := GetTokenStateRequest(ctx, token.AccessToken)
		if assert.NoError(t, err) {
			_, err = c.Do(req)
			assert.Error(t, err)
		}
	})
}
