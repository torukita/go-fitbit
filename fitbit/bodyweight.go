package fitbit

import (
	"context"
	"errors"
	"fmt"
	"net/http"
)

var (
	ErrRequestUnknown = errors.New("unknown error") // TODO put into other file
)

type WeightLogs struct {
	Weight []Weight `json:"weight"`
}

type Weight struct {
	Bmi  float64 `json:"bmi"`
	Date string  `json:"date"`
	// Fat    float64 `json:"fat"`  deplicated ?
	LogID  int64   `json:"logId"`
	Source string  `json:"source"`
	Time   string  `json:"time"`
	Weight float64 `json:"weight"`
}

func (api *Client) GetWeightLogs(param *WeightLogsParam) (*WeightLogs, *Response, error) {
	return api.GetWeightLogsContext(context.Background(), param)
}

func (api *Client) GetWeightLogsContext(ctx context.Context, param *WeightLogsParam) (*WeightLogs, *Response, error) {
	var m WeightLogs
	req, err := GetWeightLogsRequest(ctx, param)
	if err != nil {
		return nil, nil, err
	}
	resp, err := api.do_request(req, &m)
	return &m, resp, err
}

func GetWeightLogs(c *Client, param *WeightLogsParam) (*WeightLogs, error) {
	return GetWeightLogsContext(context.Background(), c, param)
}

func GetWeightLogsContext(ctx context.Context, c *Client, param *WeightLogsParam) (*WeightLogs, error) {
	m, _, err := c.GetWeightLogsContext(ctx, param)
	return m, err
}

/*
GET https://api.fitbit.com/1/user/[user-id]/body/log/weight/date/[date].json
GET https://api.fitbit.com/1/user/[user-id]/body/log/weight/date/[base-date]/[period].json
GET https://api.fitbit.com/1/user/[user-id]/body/log/weight/date/[base-date]/[end-date].json
*/
type WeightLogsParam struct {
	Date     string `validate:"len=0|datetime=2006-01-02"`
	Period   string `validate:"len=0|oneof=1d 7d 30d 1w 1m"`
	BaseDate string `validate:"len=0|datetime=2006-01-02"`
	EndDate  string `validate:"len=0|datetime=2006-01-02"`
}

func (c *WeightLogsParam) parse() error {
	if err := validate.Struct(*c); err != nil {
		return err
	}
	return nil
}

func GetWeightLogsRequest(ctx context.Context, param *WeightLogsParam) (*http.Request, error) {
	if err := param.parse(); err != nil {
		return nil, err
	}
	var url string
	if param.EndDate != "" && param.BaseDate != "" {
		url = fmt.Sprintf("https://api.fitbit.com/1/user/-/body/log/weight/date/%s/%s.json", param.BaseDate, param.EndDate)
	}
	if param.Period != "" && param.BaseDate != "" {
		url = fmt.Sprintf("https://api.fitbit.com/1/user/-/body/log/weight/date/%s/%s.json", param.BaseDate, param.Period)
	}
	if param.Date != "" {
		url = fmt.Sprintf("https://api.fitbit.com/1/user/-/body/log/weight/date/%s.json", param.Date)
	}
	if url == "" {
		return nil, ErrRequestUnknown
	}
	return http.NewRequestWithContext(ctx, "GET", url, nil)
}

type BodyFatLogs struct {
	Fat []BodyFat `json:"fat"`
}

type BodyFat struct {
	Date   string  `json:"date"`
	Fat    float64 `json:"fat"`
	LogID  int64   `json:"logId"`
	Source string  `json:"source"`
	Time   string  `json:"time"`
}

func (api *Client) GetBodyFatLogs(param *BodyFatLogsParam) (*BodyFatLogs, *Response, error) {
	return api.GetBodyFatLogsContext(context.Background(), param)
}

func (api *Client) GetBodyFatLogsContext(ctx context.Context, param *BodyFatLogsParam) (*BodyFatLogs, *Response, error) {
	var m BodyFatLogs
	req, err := GetBodyFatLogsRequest(ctx, param)
	if err != nil {
		return nil, nil, err
	}
	resp, err := api.do_request(req, &m)
	return &m, resp, err
}

func GetBodyFatLogs(c *Client, param *BodyFatLogsParam) (*BodyFatLogs, error) {
	return GetBodyFatLogsContext(context.Background(), c, param)
}

func GetBodyFatLogsContext(ctx context.Context, c *Client, param *BodyFatLogsParam) (*BodyFatLogs, error) {
	m, _, err := c.GetBodyFatLogsContext(ctx, param)
	return m, err
}

/*
GET https://api.fitbit.com/1/user/[user-id]/body/log/fat/date/[date].json
GET https://api.fitbit.com/1/user/[user-id]/body/log/fat/date/[date]/[period].json
GET https://api.fitbit.com/1/user/[user-id]/body/log/fat/date/[base-date]/[end-date].json
*/
type BodyFatLogsParam struct {
	Date     string `validate:"len=0|datetime=2006-01-02"`
	Period   string `validate:"len=0|oneof=1d 7d 30d 1w 1m"`
	BaseDate string `validate:"len=0|datetime=2006-01-02"`
	EndDate  string `validate:"len=0|datetime=2006-01-02"`
}

func (c *BodyFatLogsParam) parse() error {
	if err := validate.Struct(*c); err != nil {
		return err
	}
	return nil
}

func GetBodyFatLogsRequest(ctx context.Context, param *BodyFatLogsParam) (*http.Request, error) {
	if err := param.parse(); err != nil {
		return nil, err
	}
	var url string
	if param.EndDate != "" && param.BaseDate != "" {
		url = fmt.Sprintf("https://api.fitbit.com/1/user/-/body/log/fat/date/%s/%s.json", param.BaseDate, param.EndDate)
	}
	if param.Period != "" && param.BaseDate != "" {
		url = fmt.Sprintf("https://api.fitbit.com/1/user/-/body/log/fat/date/%s/%s.json", param.BaseDate, param.Period)
	}
	if param.Date != "" {
		url = fmt.Sprintf("https://api.fitbit.com/1/user/-/body/log/fat/date/%s.json", param.Date)
	}
	if url == "" {
		return nil, ErrRequestUnknown
	}
	return http.NewRequestWithContext(ctx, "GET", url, nil)
}
