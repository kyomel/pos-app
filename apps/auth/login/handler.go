package login

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/kyomel/pos-app/internal/config"
	"github.com/kyomel/pos-app/internal/constant"
	"github.com/kyomel/pos-app/internal/infra/response"
)

type handler struct {
	svc service
}

func newHandler(svc service) handler {
	return handler{svc: svc}
}

func (h handler) Login(rw http.ResponseWriter, r *http.Request) {
	var req = LoginRequest{}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		resp := response.NewErrorBadRequest("bad request", err, response.WithStatusCode(
			fmt.Sprintf("%v%v%v%v", http.StatusBadRequest, constant.ModuleAuthCode, constant.ServiceAuthLoginCode, "00"),
		))

		resp.JSON(rw)
		return
	}

	if err := req.Validate(); err != nil {
		resp := getResponse(err)
		resp.JSON(rw)
		return
	}

	token, role, err := h.svc.login(r.Context(), req)
	if err != nil {
		resp := getResponse(err)
		resp.JSON(rw)
		return
	}

	cfg := config.GetConfig()

	resp := response.NewSuccessOk("login success", response.WithPayload(
		map[string]interface{}{
			"token":      token,
			"role":       role,
			"token_type": cfg.App.TokenType,
		},
	),
		response.WithStatusCode(generateStatusCode("00")),
	)

	resp.JSON(rw)
}
