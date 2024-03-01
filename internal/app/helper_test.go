package app

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestToYaml(t *testing.T) {
	tests := []struct {
		name     string
		input    interface{}
		expected string
	}{
		{
			name:     "Test struct to YAML",
			input:    struct{ Name string }{Name: "John"},
			expected: "name: John\n",
		},
		{
			name:     "Test map to YAML",
			input:    map[string]string{"key": "value"},
			expected: "key: value\n",
		},
		{
			name:     "Test empty input",
			input:    nil,
			expected: "null\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ToYaml(tt.input)
			assert.Equal(t, tt.expected, result)
		})
	}
}
