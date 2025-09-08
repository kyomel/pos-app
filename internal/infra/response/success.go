package response

import "net/http"

func NewSuccessCreated(msg string, opts ...OptResponse) Response {
	var resp = Response{
		HttpStatus: http.StatusCreated,
		Message:    msg,
		Success:    true,
	}

	for _, opt := range opts {
		opt(&resp)
	}

	return resp
}

func NewSuccessOk(msg string, opts ...OptResponse) Response {
	var resp = Response{
		HttpStatus: http.StatusOK,
		Message:    msg,
		Success:    true,
	}

	for _, opt := range opts {
		opt(&resp)
	}

	return resp
}
