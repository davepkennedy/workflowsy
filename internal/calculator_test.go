package internal_test

import (
	"testing"

	"workflowsy/internal"

	"github.com/stretchr/testify/assert"
)

func TestCalculator(t *testing.T) {
	tests := []struct {
		input string
		expected float64
	}{
		{"3 4 +", 7},
		{"10 2 -", 8},
		{"5 6 *", 30},
		{"8 4 /", 2},
		{"3 4 + 2 *", 14},
	}

	for _, tt := range tests {
		t.Run(tt.input, func (t *testing.T) {
			calculator := internal.NewCalculator()
			result, err := calculator.Process(tt.input)
			assert.NoError(t, err)
			assert.Equal(t, tt.expected, result)
		})
	}
}