package fitbit

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestSleepLogsParamParse(t *testing.T) {
	ctx := context.Background()
	t.Run("ok cases", func(t *testing.T) {
		cases := []struct {
			in *SleepLogsParam
		}{
			{in: &SleepLogsParam{Date: "2006-01-02"}},
			{in: &SleepLogsParam{Date: "2006-01-02", StartDate: "2006-01-02"}},
			{in: &SleepLogsParam{Date: "2006-01-02", StartDate: "2006-01-02", EndDate: "2006-01-03"}},
			{in: &SleepLogsParam{StartDate: "2006-01-02", EndDate: "2006-01-03"}},
		}

		for _, c := range cases {
			req, err := GetSleepLogsRequest(ctx, c.in)
			assert.NoError(t, err, "in:%+v", c.in)
			assert.NotNil(t, req)
			t.Log(req.URL.String())
		}
	})

	t.Run("ng cases", func(t *testing.T) {
		cases := []struct {
			in *SleepLogsParam
		}{
			{in: &SleepLogsParam{}},
			// impossible date
			{in: &SleepLogsParam{Date: "2020-01-40"}},
			// lack of EndDate
			{in: &SleepLogsParam{StartDate: "2020-01-01"}},
		}

		for _, c := range cases {
			_, err := GetSleepLogsRequest(ctx, c.in)
			assert.Error(t, err, "in:%+v", c.in)
			t.Log(err)
		}
	})
}
func TestGetSleepLogs(t *testing.T) {
	c := setup_client()
	t.Run("GetSleepLogs", func(t *testing.T) {
		now := time.Now()
		m, _, err := c.GetSleepLogs(&SleepLogsParam{Date: now.Format("2006-01-02")})
		if assert.NoError(t, err) {
			assert.NotEmpty(t, m)
		}
	})
}
