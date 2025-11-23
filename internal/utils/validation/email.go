package validation

import "strings"

func IsValidEmail(email string) bool {
	return len(strings.Split(email, "@")) == 2
}
