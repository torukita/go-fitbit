package fitbit

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAuthCodeURL(t *testing.T) {
	cfg := &AuthConfig{
		ClientID: os.Getenv("TEST_CLIENT_ID"),
	}
	want := fmt.Sprintf("https://www.fitbit.com/oauth2/authorize?client_id=%s&prompt=consent&response_type=code&scope=activity+location+social+heartrate+settings+sleep+weight+profile+nutrition", cfg.ClientID)
	got := AuthCodeURL(cfg)
	assert.Equal(t, got, want)
}

func TestImplicitURL(t *testing.T) {
	cfg := &AuthConfig{
		ClientID: os.Getenv("TEST_CLIENT_ID"),
	}
	cases := []string{"86400", "86400", "604800", "2592000", "31536000"}
	for _, c := range cases {
		expires := c
		want := fmt.Sprintf("https://www.fitbit.com/oauth2/authorize?client_id=%s&expires_in=%s&prompt=consent&response_type=token&scope=activity+location+social+heartrate+settings+sleep+weight+profile+nutrition", cfg.ClientID, expires)
		got := ImplicitURL(cfg, expires)
		assert.Equal(t, got, want)
	}
	cases = []string{"864000", "3600"}
	for _, c := range cases { // default is 86400
		expires := c
		want := fmt.Sprintf("https://www.fitbit.com/oauth2/authorize?client_id=%s&expires_in=86400&prompt=consent&response_type=token&scope=activity+location+social+heartrate+settings+sleep+weight+profile+nutrition", cfg.ClientID)
		got := ImplicitURL(cfg, expires)
		assert.Equal(t, got, want)
	}
}

func TestGenerateToken(t *testing.T) {
	cfg := &AuthConfig{
		ClientID:     os.Getenv("TEST_CLIENT_ID"),
		ClientSecret: os.Getenv("TEST_CLIENT_SECRET"),
	}
	_, err := GenerateToken(cfg, "xxxx")
	if assert.Error(t, err) { // TODO: check if status code is 401 Unauthorized
		t.Log(err)
	}
}
