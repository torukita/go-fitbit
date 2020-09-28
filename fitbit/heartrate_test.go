package fitbit

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestHeartRateParamParse(t *testing.T) {
	ctx := context.Background()
	t.Run("ok cases", func(t *testing.T) {
		cases := []struct {
			in *HeartRateParam
		}{
			{in: &HeartRateParam{Date: "2006-01-02", DetailLevel: "1sec"}},
			{in: &HeartRateParam{Date: "2006-01-02", DetailLevel: "1min", StartTime: "01:01", EndTime: "23:58"}},
		}

		for _, c := range cases {
			req, err := GetHeartRateIntradayTimeSeriesRequest(ctx, c.in)
			if assert.NoError(t, err, "in:%+v", c.in) {
				assert.NotNil(t, req)
				t.Log(req.URL.String())
			}
		}
	})

	t.Run("ng cases", func(t *testing.T) {
		cases := []struct {
			in *HeartRateParam
		}{
			{in: &HeartRateParam{}},
			{in: &HeartRateParam{Date: "2020-01-01"}},
			// not leaped year
			{in: &HeartRateParam{Date: "2019-02-29", DetailLevel: "1min"}},
			{in: &HeartRateParam{Date: "2019-02-40", DetailLevel: "1min"}}, // impossible date
			{in: &HeartRateParam{Date: "2020-01-01", DetailLevel: "1mins"}},
			{in: &HeartRateParam{Date: "2020-01-01", DetailLevel: "1min", StartTime: "03:04"}},
		}

		for _, c := range cases {
			_, err := GetHeartRateIntradayTimeSeriesRequest(ctx, c.in)
			assert.Error(t, err, "in:%+v", c.in)
		}
	})
}

func TestGetHeartRateIntradayTimeSeries(t *testing.T) {
	c := setup_client()
	t.Run("GetHeartRateIntradayTimeSeries", func(t *testing.T) {
		now := time.Now()
		m, _, err := c.GetHeartRateIntradayTimeSeries(&HeartRateParam{
			Date:        now.Format("2006-01-02"),
			DetailLevel: "1sec",
		})
		if assert.NoError(t, err) {
			assert.NotEmpty(t, m)
		}
	})
}
