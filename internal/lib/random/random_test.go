package random

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestNewRandomString(t *testing.T) {
	tests := []struct {
		name string
		size int
	}{
		{
			name: "size = 1",
			size: 1,
		},
		{
			name: "size = 6",
			size: 6,
		},
		{
			name: "size = 12",
			size: 12,
		},
		{
			name: "size = 24",
			size: 24,
		},
		{
			name: "size = 36",
			size: 36,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			str1 := NewRandomString(tt.size)
			time.Sleep(1)
			str2 := NewRandomString(tt.size)

			assert.Len(t, str1, tt.size)
			assert.Len(t, str2, tt.size)

			assert.NotEqual(t, str1, str2)
		})
	}
}
