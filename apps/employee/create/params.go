package create

import "github.com/kyomel/pos-app/internal/constant"

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

	return nil
}
