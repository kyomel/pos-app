package create

type Auth struct{}

func (a Auth) IsExists() bool {
	return a != Auth{}
}

type Employee struct{}
