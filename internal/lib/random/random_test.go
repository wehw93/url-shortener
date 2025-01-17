package random

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewRandomString(t *testing.T) {
	tests := []struct {
		name string
		size int
	}{{
		name: "1",
		size: 1,
	},
		{
			name: "2",
			size: 2,
		},
		{
			name: "3",
			size: 3,
		},
		{
			name: "4",
			size: 4,
		},
		{
			name: "5",
			size: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			str1 := NewRandomString(tt.size)
			str2 := NewRandomString(tt.size)
			assert.Len(t, str1, tt.size)
			assert.Len(t, str2, tt.size)
			assert.NotEqual(t, str1, str2)
		})
	}
}
