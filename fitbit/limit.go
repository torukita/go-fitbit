package fitbit

import (
	"net/http"
	"strconv"
)

const (
	headerRateLimitLimit     = "Fitbit-Rate-Limit-Limit"
	headerRateLimitRemaining = "Fitbit-Rate-Limit-Remaining"
	headerRateLimitReset     = "Fitbit-Rate-Limit-Reset"
)

type RateLimit struct {
	Limit     int
	Remaining int
	Reset     int
}

func newRateLimit(r *http.Response) (rate RateLimit) {
	if limit := r.Header.Get(headerRateLimitLimit); limit != "" {
		rate.Limit, _ = strconv.Atoi(limit)
	}
	if remain := r.Header.Get(headerRateLimitRemaining); remain != "" {
		rate.Remaining, _ = strconv.Atoi(remain)
	}
	if reset := r.Header.Get(headerRateLimitReset); reset != "" {
		rate.Reset, _ = strconv.Atoi(reset)
	}
	return
}
