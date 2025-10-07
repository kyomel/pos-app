package create

import "net/http"

type handler struct {
	svc service
}

func newHandler(svc service) handler {
	return handler{svc: svc}
}

func (h handler) createEmployee(rw http.ResponseWriter, r *http.Request) {

}
