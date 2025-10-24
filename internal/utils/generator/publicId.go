package generator

import "github.com/oklog/ulid/v2"

func GeneratePublicId() string {
	id := ulid.Make()
	return id.String()
}
