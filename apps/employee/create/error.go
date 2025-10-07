package create

import "errors"

var (
	errNameRoleEmailPasswordProfileIsEmpty = errors.New("name, role, email, password, profile is empty")
	errEmailAlreadyExist                   = errors.New("email already exist")
	errRoleIsNotSupported                  = errors.New("role is not supported, should be one of cashier or warehouse")
	errPasswordInvalidLength               = errors.New("password must be at least 6 characters")
	errEmailNotFound                       = errors.New("email not found")
)
