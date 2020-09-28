package fitbit

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetDevices(t *testing.T) {
	c := setup_client()
	id := ""
	t.Run("GetDevices", func(t *testing.T) {
		m, _, err := c.GetDevices()
		if assert.NoError(t, err) {
			assert.NotEmpty(t, m)
			id = (*m)[0].ID
		}
	})
	t.Run("GetAlarms", func(t *testing.T) {
		m, _, err := c.GetAlarms(&GetAlarmsParam{TrackerID: id})
		if assert.NoError(t, err) {
			assert.NotEmpty(t, m)
		}
	})

}
