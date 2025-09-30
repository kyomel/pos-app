package login

import (
	"errors"

	"github.com/kyomel/pos-app/internal/constant"
	"github.com/kyomel/pos-app/internal/infra/response"
)

var (
	errEmailOrPasswordEmpty  = errors.New("email or password is empty")
	errPasswordInvalidLength = errors.New("password length minimum is 6")

	errEmailOrPasswordIsNotMatch = errors.New("email or password is not match")
	errAccountIsNotActive        = errors.New("account is not active")

	errInternalServerError = errors.New("internal server error")
)

func generateStatusCode(statusCode string) string {
	return constant.ModuleAuthCode + constant.ServiceAuthLoginCode + statusCode
}

var (
	errMap = map[string]response.Response{
		errEmailOrPasswordEmpty.Error(): response.NewErrorBadRequest("bad request", errEmailOrPasswordEmpty, response.WithStatusCode(
			generateStatusCode("01"),
		)),
		errPasswordInvalidLength.Error(): response.NewErrorBadRequest("password length minimum is 6", errPasswordInvalidLength, response.WithStatusCode(
			generateStatusCode("02"),
		)),
		errEmailOrPasswordIsNotMatch.Error(): response.NewErrorUnauthorized("unauthorized", errEmailOrPasswordIsNotMatch, response.WithStatusCode(
			generateStatusCode("01"),
		)),
		errAccountIsNotActive.Error(): response.NewErrorUnprocessableEntity("unprocessable entity", errAccountIsNotActive, response.WithStatusCode(
			generateStatusCode("01"),
		)),
	}
)

func getResponse(err error) response.Response {
	if r, ok := errMap[err.Error()]; ok {
		return r
	}
	return response.NewErrorGeneral("uknown error", errInternalServerError, response.WithStatusCode(
		generateStatusCode("99"),
	))
}
