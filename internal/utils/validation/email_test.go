package validation

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestValidateEmail(t *testing.T) {
	type test struct {
		title    string
		email    string
		expected bool
	}

	tests := []test{
		{
			title:    "valid email",
			email:    "john.doe@example.com",
			expected: true,
		},
		{
			title:    "invalid email",
			email:    "johndoeexample",
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.title, func(t *testing.T) {
			got := IsValidEmail(tt.email)
			require.Equal(t, tt.expected, got)
		})
	}
}
