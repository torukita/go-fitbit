package fitbit

import "fmt"

type ErrorResponse struct {
	Errors []struct {
		ErrorType string `json:"errorType"`
		Message   string `json:"message"`
	} `json:"errors"`
	Success bool `json:"success"`
}

func (e *ErrorResponse) Error() string {
	return fmt.Sprintf("errorType:%s message:%s", e.Errors[0].ErrorType, e.Errors[0].Message)
}
