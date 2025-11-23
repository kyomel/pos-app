package create

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/kyomel/pos-app/internal/constant"
	"github.com/kyomel/pos-app/internal/infra/response"
)

type handler struct {
	svc service
}

func newHandler(svc service) handler {
	return handler{svc: svc}
}

func (h handler) createEmployee(rw http.ResponseWriter, r *http.Request) {
	var req = CreateEmployeeRequest{}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		resp := response.NewErrorBadRequest("bad request", err, response.WithStatusCode(
			fmt.Sprintf("%v%v%v%v", http.StatusBadRequest, constant.ModuleEmployeeCode, constant.ServiceEmployeeCreate, "00"),
		))

		resp.JSON(rw)
		return
	}

	if err := req.Validate(); err != nil {
		resp := getResponse(err, errCreateMap)
		resp.JSON(rw)
		return
	}

	if err := h.svc.create(r.Context(), req); err != nil {
		resp := getResponse(err, errCreateMap)
		resp.JSON(rw)
		return
	}

	resp := response.NewSuccessCreated("employee created successfully", response.WithStatusCode(
		fmt.Sprintf("%v%v%v", constant.ModuleEmployeeCode, constant.ServiceEmployeeCreate, "00"),
	))

	resp.JSON(rw)

}
