package fitbit

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestWeightLogsParamParse(t *testing.T) {
	ctx := context.Background()
	t.Run("ok cases", func(t *testing.T) {
		cases := []struct {
			in *WeightLogsParam
		}{
			{in: &WeightLogsParam{Date: "2006-01-02"}},
			{in: &WeightLogsParam{Date: "2006-01-02", Period: "1d"}},
			{in: &WeightLogsParam{Date: "2019-01-02", Period: "1m"}},
			{in: &WeightLogsParam{Date: "2019-02-01", Period: "7d", BaseDate: "2019-03-01"}},
			{in: &WeightLogsParam{Date: "", Period: "7d", BaseDate: "2019-03-01"}},
		}

		for _, c := range cases {
			req, err := GetWeightLogsRequest(ctx, c.in)
			if assert.NoError(t, err, "in:%+v", c.in) {
				assert.NotNil(t, req)
				t.Log(req.URL.String())
			}
		}
	})

	t.Run("ng cases", func(t *testing.T) {
		cases := []struct {
			in *WeightLogsParam
		}{
			{in: &WeightLogsParam{}},
			{in: &WeightLogsParam{Date: "2020-01-40", Period: "7d"}}, // impossible date
			{in: &WeightLogsParam{Date: "2020-01-01", BaseDate: "03:04"}},
			{in: &WeightLogsParam{Date: "2020-01"}},
		}

		for _, c := range cases {
			_, err := GetWeightLogsRequest(ctx, c.in)
			if assert.Error(t, err, "in:%+v", c.in) {
				t.Log(err)
			}
		}
	})
}

func TestBodyFatLogsParamParse(t *testing.T) {
	ctx := context.Background()
	t.Run("ok cases", func(t *testing.T) {
		cases := []struct {
			in *BodyFatLogsParam
		}{
			{in: &BodyFatLogsParam{Date: "2006-01-02"}},
			{in: &BodyFatLogsParam{Date: "2006-01-02", Period: "1d"}},
			{in: &BodyFatLogsParam{Date: "2019-01-02", Period: "1m"}},
			{in: &BodyFatLogsParam{Date: "2019-02-01", Period: "7d", BaseDate: "2019-03-01"}},
			{in: &BodyFatLogsParam{Date: "", Period: "7d", BaseDate: "2019-03-01"}},
		}

		for _, c := range cases {
			req, err := GetBodyFatLogsRequest(ctx, c.in)
			if assert.NoError(t, err, "in:%+v", c.in) {
				assert.NotNil(t, req)
				t.Log(req.URL.String())
			}
		}
	})

	t.Run("ng cases", func(t *testing.T) {
		cases := []struct {
			in *BodyFatLogsParam
		}{
			{in: &BodyFatLogsParam{}},
			// impossible date
			{in: &BodyFatLogsParam{Date: "2020-01-40", Period: "7d"}},
			{in: &BodyFatLogsParam{Date: "2020-01-01", BaseDate: "03:04"}},
		}

		for _, c := range cases {
			_, err := GetBodyFatLogsRequest(ctx, c.in)
			if assert.Error(t, err, "in:%+v", c.in) {
				t.Log(err)
			}
		}
	})
}

func TestGetWeightLogs(t *testing.T) {
	c := setup_client()
	t.Run("GetWeightLogs", func(t *testing.T) {
		now := time.Now()
		m, resp, err := c.GetWeightLogs(&WeightLogsParam{Date: now.Format("2006-01-02")})
		if assert.NoError(t, err) {
			assert.NotEmpty(t, m)
			assert.NotNil(t, resp)
		}
	})
}

func TestGetBodyFatLogs(t *testing.T) {
	c := setup_client()
	t.Run("GetBodyFatLogs", func(t *testing.T) {
		now := time.Now()
		m, resp, err := c.GetBodyFatLogs(&BodyFatLogsParam{Date: now.Format("2006-01-02")})
		if assert.NoError(t, err) {
			assert.NotEmpty(t, m)
			assert.NotNil(t, resp)
		}
	})
}
