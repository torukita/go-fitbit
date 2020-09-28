package fitbit

import (
	"context"
	"fmt"
	"net/http"
)

type SleepLogs struct {
	Sleep   []Sleep `json:"sleep"`
	Summary Summary `json:"summary"`
}

type Sleep struct {
	DateOfSleep string `json:"dateOfSleep"`
	Duration    int    `json:"duration"`
	Efficiency  int    `json:"efficiency"`
	EndTime     string `json:"endTime"`
	InfoCode    int    `json:"infoCode"`
	IsMainSleep bool   `json:"isMainSleep"`
	Levels      struct {
		Data []struct {
			DateTime string `json:"dateTime"`
			Level    string `json:"level"`
			Seconds  int    `json:"seconds"`
		} `json:"data"`
		ShortData []struct {
			DateTime string `json:"dateTime"`
			Level    string `json:"level"`
			Seconds  int    `json:"seconds"`
		} `json:"shortData"`
		Summary struct {
			Deep struct {
				Count               int `json:"count"`
				Minutes             int `json:"minutes"`
				ThirtyDayAvgMinutes int `json:"thirtyDayAvgMinutes"`
			} `json:"deep"`
			Light struct {
				Count               int `json:"count"`
				Minutes             int `json:"minutes"`
				ThirtyDayAvgMinutes int `json:"thirtyDayAvgMinutes"`
			} `json:"light"`
			Rem struct {
				Count               int `json:"count"`
				Minutes             int `json:"minutes"`
				ThirtyDayAvgMinutes int `json:"thirtyDayAvgMinutes"`
			} `json:"rem"`
			Wake struct {
				Count               int `json:"count"`
				Minutes             int `json:"minutes"`
				ThirtyDayAvgMinutes int `json:"thirtyDayAvgMinutes"`
			} `json:"wake"`
		} `json:"summary"`
	} `json:"levels"`
	LogID               int64  `json:"logId"`
	MinutesAfterWakeup  int    `json:"minutesAfterWakeup"`
	MinutesAsleep       int    `json:"minutesAsleep"`
	MinutesAwake        int    `json:"minutesAwake"`
	MinutesToFallAsleep int    `json:"minutesToFallAsleep"`
	StartTime           string `json:"startTime"`
	TimeInBed           int    `json:"timeInBed"`
	Type                string `json:"type"`
}

type Summary struct {
	Stages struct {
		Deep  int `json:"deep"`
		Light int `json:"light"`
		Rem   int `json:"rem"`
		Wake  int `json:"wake"`
	} `json:"stages"`
	TotalMinutesAsleep int `json:"totalMinutesAsleep"`
	TotalSleepRecords  int `json:"totalSleepRecords"`
	TotalTimeInBed     int `json:"totalTimeInBed"`
}

func (api *Client) GetSleepLogs(param *SleepLogsParam) (*SleepLogs, *Response, error) {
	return api.GetSleepLogsContext(context.Background(), param)
}

func (api *Client) GetSleepLogsContext(ctx context.Context, param *SleepLogsParam) (*SleepLogs, *Response, error) {
	var m SleepLogs
	req, err := GetSleepLogsRequest(ctx, param)
	if err != nil {
		return nil, nil, err
	}
	resp, err := api.do_request(req, &m)
	return &m, resp, err
}

func GetSleepLogs(c *Client, param *SleepLogsParam) (*SleepLogs, error) {
	return GetSleepLogsContext(context.Background(), c, param)
}

func GetSleepLogsContext(ctx context.Context, c *Client, param *SleepLogsParam) (*SleepLogs, error) {
	m, _, err := c.GetSleepLogsContext(ctx, param)
	return m, err
}

// GET https://api.fitbit.com/1.2/user/[user-id]/sleep/date/[date].json
// GET https://api.fitbit.com/1.2/user/[user-id]/sleep/date/[startDate]/[endDate].json
type SleepLogsParam struct {
	Date      string `validate:"len=0|datetime=2006-01-02"`
	StartDate string `validate:"len=0|datetime=2006-01-02"`
	EndDate   string `validate:"len=0|datetime=2006-01-02"`
}

func (c *SleepLogsParam) parse() error {
	if err := validate.Struct(*c); err != nil {
		return err
	}
	return nil
}

func GetSleepLogsRequest(ctx context.Context, param *SleepLogsParam) (*http.Request, error) {
	if err := param.parse(); err != nil {
		return nil, err
	}
	var url string
	if param.StartDate != "" && param.EndDate != "" {
		url = fmt.Sprintf("https://api.fitbit.com/1.2/user/-/sleep/date/%s/%s.json", param.StartDate, param.EndDate)
	}
	if param.Date != "" {
		url = fmt.Sprintf("https://api.fitbit.com/1.2/user/-/sleep/date/%s.json", param.Date)
	}
	if url == "" {
		return nil, ErrRequestUnknown
	}
	return http.NewRequestWithContext(ctx, "GET", url, nil)
}

// TODO: GetSleepLogsLIst
