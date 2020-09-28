package fitbit

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

// Response is a wrapper for http.Response received from fitbit api
type Response struct {
	*http.Response
	rateLimit RateLimit
}

func newResponse(r *http.Response) *Response {
	return &Response{
		Response:  r,
		rateLimit: newRateLimit(r),
	}
}

func (r *Response) Decode(v interface{}) error {
	if v == nil {
		return errors.New("TODO: define error message")
	}
	body := r.Response.Body
	defer r.Response.Body.Close()
	if r.Response.StatusCode != http.StatusOK {
		var e ErrorResponse
		err := json.NewDecoder(body).Decode(&e)
		if err != nil {
			return fmt.Errorf("failed to decode error response: %w", err)
		}
		return &e
	}
	err := json.NewDecoder(body).Decode(v)
	if err != nil && err != io.EOF { // in the case of empty body
		return fmt.Errorf("failed to decode body response: %w", err)
	}
	return nil
}

// TODO refactor
func (r *Response) OutputWithIndent(w io.Writer) error {
	body := r.Response.Body
	defer r.Response.Body.Close()
	b, err := ioutil.ReadAll(body)
	if err != nil {
		return err
	}
	var out bytes.Buffer
	if err := json.Indent(&out, b, "", "  "); err != nil {
		return err
	}
	_, err = out.WriteTo(w)
	return err
}

func (r *Response) Output(w io.Writer) error {
	body := r.Response.Body
	defer r.Response.Body.Close()
	_, err := io.Copy(w, body)
	return err
}

func (r *Response) RateLimit() RateLimit {
	return r.rateLimit
}
