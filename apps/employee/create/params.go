package create

import (
	"github.com/kyomel/pos-app/internal/constant"
	"github.com/kyomel/pos-app/internal/utils/encryption"
	"github.com/kyomel/pos-app/internal/utils/generator"
	"github.com/kyomel/pos-app/internal/utils/validation"
)

type CreateEmployeeRequest struct {
	Name     string `json:"name"`
	Role     string `json:"role"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Profile  string `json:"profile"`
}

func (r CreateEmployeeRequest) Validate() error {
	if r.Name == "" || r.Role == "" || r.Email == "" || r.Password == "" || r.Profile == "" {
		return errNameRoleEmailPasswordProfileIsEmpty
	}

	if len(r.Password) < 6 {
		return errPasswordInvalidLength
	}

	if !constant.IsRoleCanBeCreated(r.Role) {
		return errRoleIsNotSupported
	}

	if !validation.IsValidEmail(r.Email) {
		return errEmailIsNotValid
	}

	return nil
}

func (r CreateEmployeeRequest) ToAuthModel() (auth Auth) {
	hashed, _ := encryption.GenerateFromPassword(r.Password)
	auth = Auth{
		PublicId: generator.GeneratePublicId(),
		Email:    r.Email,
		Password: hashed,
		Role:     r.Role,
		IsActive: true,
	}

	return
}

func (r CreateEmployeeRequest) ToEmployeeModel(authId string) (employee Employee) {
	employee = Employee{
		PublicId: generator.GeneratePublicId(),
		Name:     r.Name,
		Profile:  r.Profile,
		AuthId:   authId,
	}
	return
}
