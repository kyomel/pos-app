package create

import (
	"errors"

	"github.com/kyomel/pos-app/internal/constant"
	"github.com/kyomel/pos-app/internal/infra/response"
)

var (
	errNameRoleEmailPasswordProfileIsEmpty = errors.New("name, role, email, password, profile is empty")
	errEmailAlreadyExist                   = errors.New("email already exist")
	errRoleIsNotSupported                  = errors.New("role is not supported, should be one of cashier or warehouse")
	errPasswordInvalidLength               = errors.New("password must be at least 6 characters")
	errEmailNotFound                       = errors.New("email not found")
	errInternalServerError                 = errors.New("internal server error")
	errEmailIsNotValid                     = errors.New("email is not valid")
)

var (
	errCreateMap = map[string]response.Response{
		errNameRoleEmailPasswordProfileIsEmpty.Error(): response.NewErrorBadRequest("bad request", errNameRoleEmailPasswordProfileIsEmpty, response.WithStatusCode(
			generateStatusCode(constant.ServiceEmployeeCreate, "01"),
		)),
		errRoleIsNotSupported.Error(): response.NewErrorBadRequest("bad request", errRoleIsNotSupported, response.WithStatusCode(
			generateStatusCode(constant.ServiceEmployeeCreate, "02"),
		)),
		errPasswordInvalidLength.Error(): response.NewErrorBadRequest("bad request", errPasswordInvalidLength, response.WithStatusCode(
			generateStatusCode(constant.ServiceEmployeeCreate, "03"),
		)),
		errEmailIsNotValid.Error(): response.NewErrorBadRequest("bad request", errEmailIsNotValid, response.WithStatusCode(
			generateStatusCode(constant.ServiceEmployeeCreate, "04"),
		)),
		errEmailAlreadyExist.Error(): response.NewErrorConflict("data already exist", errEmailAlreadyExist, response.WithStatusCode(
			generateStatusCode(constant.ServiceEmployeeCreate, "01"),
		)),
	}
	errListMap = map[string]response.Response{}
)

func generateStatusCode(serviceCode string, statusCode string) string {
	return constant.ModuleEmployeeCode + serviceCode + statusCode
}

func getResponse(err error, errorMap map[string]response.Response) response.Response {
	if r, ok := errorMap[err.Error()]; ok {
		return r
	}
	return response.NewErrorGeneral("unknown error", errInternalServerError, response.WithStatusCode(
		generateStatusCode("00", "99"),
	))
}
