package fitbit

import (
	"context"
	"fmt"
	"net/http"
)

type HeartRateData struct {
	ActivitiesHeart         []ActivitiesHeart       `json:"activities-heart"`
	ActivitiesHeartIntraday ActivitiesHeartIntraday `json:"activities-heart-intraday"`
}

type ActivitiesHeart struct {
	DateTime string `json:"dateTime"`
	Value    struct {
		CustomHeartRateZones []interface{} `json:"customHeartRateZones"`
		HeartRateZones       []struct {
			CaloriesOut float64 `json:"caloriesOut"`
			Max         int     `json:"max"`
			Min         int     `json:"min"`
			Minutes     int     `json:"minutes"`
			Name        string  `json:"name"`
		} `json:"heartRateZones"`
		RestingHeartRate int `json:"restingHeartRate"`
	} `json:"value"`
}

type ActivitiesHeartIntraday struct {
	Dataset []struct {
		Time  string `json:"time"`
		Value int    `json:"value"`
	} `json:"dataset"`
	DatasetInterval int    `json:"datasetInterval"`
	DatasetType     string `json:"datasetType"`
}

func (api *Client) GetHeartRateIntradayTimeSeries(param *HeartRateParam) (*HeartRateData, *Response, error) {
	return api.GetHeartRateIntradayTimeSeriesContext(context.Background(), param)
}

func (api *Client) GetHeartRateIntradayTimeSeriesContext(ctx context.Context, param *HeartRateParam) (*HeartRateData, *Response, error) {
	var m HeartRateData
	req, err := GetHeartRateIntradayTimeSeriesRequest(ctx, param)
	if err != nil {
		return nil, nil, err
	}
	resp, err := api.do_request(req, &m)
	return &m, resp, err
}

func GetHeartRateIntradayTimeSeries(c *Client, param *HeartRateParam) (*HeartRateData, error) {
	return GetHeartRateIntradayTimeSeriesContext(context.Background(), c, param)
}

func GetHeartRateIntradayTimeSeriesContext(ctx context.Context, c *Client, param *HeartRateParam) (*HeartRateData, error) {
	m, _, err := c.GetHeartRateIntradayTimeSeriesContext(ctx, param)
	return m, err
}

// Get Heart Rate Intraday Time Series
// Personal App Type is requeired
/*
GET https://api.fitbit.com/1/user/-/activities/heart/date/[date]/[end-date]/[detail-level].json
GET https://api.fitbit.com/1/user/-/activities/heart/date/[date]/[end-date]/[detail-level]/time/[start-time]/[end-time].json
GET https://api.fitbit.com/1/user/-/activities/heart/date/[date]/1d/[detail-level].json`
GET https://api.fitbit.com/1/user/-/activities/heart/date/[date]/1d/[detail-level]/time/[start-time]/[end-time].json
*/
type HeartRateParam struct {
	Date        string `validate:"len=0|datetime=2006-01-02"`
	EndDate     string `validate:"len=0|datetime=2006-01-02"`
	DetailLevel string `validate:"oneof=1sec 1min"`
	StartTime   string `validate:"len=0|datetime=15:04"`
	EndTime     string `validate:"len=0|datetime=15:04"`
}

func (c *HeartRateParam) parse() error {
	if err := validate.Struct(*c); err != nil {
		return err
	}
	return nil
}

func GetHeartRateIntradayTimeSeriesRequest(ctx context.Context, param *HeartRateParam) (*http.Request, error) {
	if err := param.parse(); err != nil {
		return nil, err
	}
	var url string
	endDate := param.EndDate
	if endDate == "" {
		endDate = "1d"
	}
	if param.StartTime == "" && param.EndTime == "" {
		url = fmt.Sprintf("https://api.fitbit.com/1/user/-/activities/heart/date/%s/%s/%s.json", param.Date, endDate, param.DetailLevel)
	}
	if param.StartTime != "" && param.EndTime != "" {
		url = fmt.Sprintf("https://api.fitbit.com/1/user/-/activities/heart/date/%s/%s/%s/time/%s/%s.json",
			param.Date, endDate, param.DetailLevel, param.StartTime, param.EndTime)
	}
	if url == "" {
		return nil, ErrRequestUnknown
	}
	return http.NewRequestWithContext(ctx, "GET", url, nil)
}
