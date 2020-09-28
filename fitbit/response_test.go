package fitbit

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHandleOrgResponse(t *testing.T) {
	cases := []struct {
		in string
	}{
		{in: `{"k1": "v1"}`},
		{in: `{"k2": "v2"}`},
		{in: ``},
	}
	var v interface{}
	for _, c := range cases {
		reader := strings.NewReader(c.in)
		err := json.NewDecoder(reader).Decode(&v)
		if err != nil && err == io.EOF {
			err = nil // ignore EOF errors caused by empty response body
		}
		assert.NoError(t, err)
	}
}

func TestResponseDecode(t *testing.T) {
	c := setup_client()
	t.Run("success case", func(t *testing.T) {
		req, _ := GetDevicesRequest(context.Background())
		resp, err := c.Do(req)
		if assert.NoError(t, err) {
			var m Devices
			err = resp.Decode(&m)
			assert.NoError(t, err)
		}
	})

	t.Run("fail case", func(t *testing.T) {
		req, _ := http.NewRequest("GET", "https://api.fitbit.com/1/user/-", nil) // error and expects error response type
		resp, err := c.Do(req)
		if assert.NoError(t, err) {
			var m Devices
			err = resp.Decode(&m)
			if assert.Error(t, err) {
				assert.IsType(t, &ErrorResponse{}, err)
			}
		}
	})
}
