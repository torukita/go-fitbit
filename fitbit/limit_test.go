package fitbit

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRateLimit(t *testing.T) {
	type in struct {
		limit     string
		remaining string
		reset     string
	}
	cases := []struct {
		in  in
		out RateLimit
	}{
		{
			in:  in{limit: "0", remaining: "0", reset: "0"},
			out: RateLimit{},
		},
		{
			in:  in{limit: "150", remaining: "149", reset: "3500"},
			out: RateLimit{150, 149, 3500},
		},
		{
			in:  in{remaining: "149", reset: "3500"},
			out: RateLimit{0, 149, 3500},
		},
	}

	for _, c := range cases {
		h := &http.Response{
			Header: make(http.Header, 0),
		}
		h.Header.Set(headerRateLimitLimit, c.in.limit)
		h.Header.Set(headerRateLimitRemaining, c.in.remaining)
		h.Header.Set(headerRateLimitReset, c.in.reset)
		got := newRateLimit(h)
		assert.Equal(t, c.out, got)
	}
}
