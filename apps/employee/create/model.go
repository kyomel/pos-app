package create

import "time"

type Auth struct {
	PublicId string
	Email    string
	Password string
	Role     string
	IsActive bool
}

func (a Auth) IsExists() bool {
	return a != Auth{}
}

type Employee struct {
	id        int
	PublicId  string
	Name      string
	Profile   string
	AuthId    string
	CreatedAt time.Time
	UpdatedAt time.Time
}
