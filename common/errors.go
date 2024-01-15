package common

import (
	"errors"
	"fmt"
)

type ErrorFromGoroutine struct {
	Err          error
	FailedAutoId string
}

type TooManyRequestsError struct {
	Err        error
	StatusCode int
}

func NewTooManyRequestsError(err error, status int) error {
	return &TooManyRequestsError{err, status}
}

func (tmre *TooManyRequestsError) Error() string {
	return tmre.Err.Error()
}

func IsTooManyRequestsError(err error) bool {
	var tmre *TooManyRequestsError
	return errors.As(err, &tmre)
}

type RequestError struct {
	Err        error
	StatusCode int
	Status     string
	Body       string
}

func NewRequestError(err error, statusCode int, status, body string) error {
	return &RequestError{err, statusCode, status, body}
}

func (re *RequestError) Error() string {
	return fmt.Sprintf("err: %s, status code: %d, status: %s, body: %s", re.Err.Error(), re.StatusCode, re.Status, re.Body)
}

func isRequestError(err error) bool {
	var re *RequestError
	return errors.As(err, &re)
}
