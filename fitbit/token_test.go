package fitbit

import (
	"context"
	"encoding/json"
	"io"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func decodeResponse(body io.Reader, v interface{}) error {
	err := json.NewDecoder(body).Decode(&v)
	if err != nil && err != io.EOF { // in the case of empty body
		return err
	}
	return nil
}

func TestClient(t *testing.T) {
	ctx := context.Background()
	t.Run("success case", func(t *testing.T) {
		token := &Token{
			AccessToken: os.Getenv("TEST_TOKEN"),
		}
		c := token.Client()
		req, err := GetTokenStateRequest(ctx, token.AccessToken)
		assert.NoError(t, err)
		resp, err := c.Do(req)
		assert.NoError(t, err)
		defer resp.Body.Close()

		var m TokenState
		err = decodeResponse(resp.Body, &m)
		assert.NoError(t, err)
		assert.Equal(t, true, m.Active)
	})
	t.Run("fail case", func(t *testing.T) {
		token := &Token{
			AccessToken: "",
		}
		c := NewClient(token.Client()) // client with no token
		req, err := GetTokenStateRequest(ctx, token.AccessToken)
		assert.NoError(t, err)
		_, err = c.Do(req)
		assert.Error(t, err)
	})
}

func TestClientWithID(t *testing.T) {
	ctx := context.Background()
	t.Run("success case", func(t *testing.T) {
		token := &Token{
			AccessToken: os.Getenv("TEST_TOKEN"),
		}
		id := os.Getenv("TEST_CLIENT_ID")
		secret := os.Getenv("TEST_CLIENT_SECRET")
		c := token.ClientWithID(id, secret)
		req, err := GetTokenStateRequest(ctx, token.AccessToken)
		assert.NoError(t, err)
		resp, err := c.Do(req)
		assert.NoError(t, err)
		var m TokenState
		defer resp.Body.Close()
		err = decodeResponse(resp.Body, &m)
		assert.NoError(t, err)
		assert.Equal(t, true, m.Active)
		assert.Equal(t, id, m.ClientID)
	})
}
