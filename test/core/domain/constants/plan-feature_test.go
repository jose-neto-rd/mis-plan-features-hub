package constants_test

import (
	"mis-plan-features-hub/internal/core/domain/constants"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGeneratePermissions(t *testing.T) {
	tests := []struct {
		name     string
		all      []string
		exclude  []string
		expected []string
	}{
		{
			name:     "empty lists",
			all:      []string{},
			exclude:  []string{},
			expected: []string{},
		},
		{
			name:     "no permissions to exclude",
			all:      []string{"read", "write", "execute"},
			exclude:  []string{},
			expected: []string{"read", "write", "execute"},
		},
		{
			name:     "exclude some permissions",
			all:      []string{"read", "write", "execute", "delete"},
			exclude:  []string{"write", "delete"},
			expected: []string{"read", "execute"},
		},
		{
			name:     "exclude all permissions",
			all:      []string{"read", "write", "execute"},
			exclude:  []string{"read", "write", "execute"},
			expected: []string{},
		},
		{
			name:     "exclude non-existent permissions",
			all:      []string{"read", "write"},
			exclude:  []string{"execute", "delete"},
			expected: []string{"read", "write"},
		},
		{
			name:     "duplicate permissions in all",
			all:      []string{"read", "read", "write"},
			exclude:  []string{"read"},
			expected: []string{"write"},
		},
		{
			name:     "duplicate permissions in exclude",
			all:      []string{"read", "write", "execute"},
			exclude:  []string{"write", "write"},
			expected: []string{"read", "execute"},
		},
		{
			name:     "case sensitive permissions",
			all:      []string{"Read", "WRITE", "execute"},
			exclude:  []string{"read", "Write"},
			expected: []string{"Read", "WRITE", "execute"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := constants.GeneratePermissions(tt.all, tt.exclude)
			assert.Equal(t, tt.expected, result, "they should be equal")
		})
	}
}
