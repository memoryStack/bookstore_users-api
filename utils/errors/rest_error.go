package errors

import "net/http"

// RestError ...
type RestError struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
	Error   string `json:"error"`
}

// BadRequestError ... this func will create error response for bad requests
func BadRequestError(errMsg string) *RestError {
	return &RestError{
		Message: errMsg,
		Status:  http.StatusBadRequest,
		Error:   "bad request",
	}
}

/*
	we defined a common error messages structure to be in sync with all the error happening in any API.
	if each API has their different different error responses structure then it's going to be a log of overload.
*/
