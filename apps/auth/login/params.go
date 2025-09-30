package login

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (r LoginRequest) Validate() error {
	if r.Email == "" || r.Password == "" {
		return errEmailOrPasswordEmpty
	}

	if len(r.Password) < 6 {
		return errPasswordInvalidLength
	}

	return nil
}
