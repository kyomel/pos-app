package response

import (
	"encoding/json"
	"net/http"
)

func newError(httpStatusCode int, msg string, err error, opts ...OptResponse) Response {
	var resp = Response{
		Success:    false,
		HttpStatus: httpStatusCode,
		Message:    msg,
		Error:      err.Error(),
	}

	for _, opt := range opts {
		opt(&resp)
	}

	return resp
}

func NewErrorGeneral(msg string, err error, opts ...OptResponse) Response {
	return newError(http.StatusInternalServerError, msg, err, opts...)
}

func NewErrorBadRequest(msg string, err error, opts ...OptResponse) Response {
	return newError(http.StatusBadRequest, msg, err, opts...)
}

func NewErrorUnauthorized(msg string, err error, opts ...OptResponse) Response {
	return newError(http.StatusUnauthorized, msg, err, opts...)
}

func NewErrorUnprocessableEntity(msg string, err error, opts ...OptResponse) Response {
	return newError(http.StatusUnprocessableEntity, msg, err, opts...)
}

func (r Response) JSON(rw http.ResponseWriter) {
	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(r.HttpStatus)
	json.NewEncoder(rw).Encode(r)
}
