package fitbit

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetTokenState(t *testing.T) {
	c := setup_client()
	t.Run("GetTokenState", func(t *testing.T) {
		m, _, err := c.GetTokenState(os.Getenv("TEST_TOKEN"))
		if assert.NoError(t, err) {
			assert.NotEmpty(t, m)
		}
	})
}
