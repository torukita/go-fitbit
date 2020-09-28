package fitbit

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetProfile(t *testing.T) {
	c := setup_client()
	t.Run("GetProfile", func(t *testing.T) {
		m, _, err := c.GetProfile()
		if assert.NoError(t, err) {
			assert.NotEmpty(t, m)
		}
	})
}
