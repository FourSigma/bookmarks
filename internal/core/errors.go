package core

import "fmt"

func NewErrorBadRequest(e error) ErrorBadRequest {
	return ErrorBadRequest{
		err: e,
	}
}

type ErrorBadRequest struct {
	err error
}

func (e ErrorBadRequest) Error() string {
	return fmt.Sprintf("ErrorBadRequest: %s", e.err.Error())
}
