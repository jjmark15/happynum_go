package cli

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInterpretArg(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected int
	}{
		{"scientific notation", "1e6", 1000000},
		{"plain integer", "1000", 1000},
		{"invalid string", "abc", 1},
		{"empty string", "", 1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expected, interpretArg(tt.input))
		})
	}
}

func TestNewCommand(t *testing.T) {
	cmd := newCommand()

	assert.Equal(t, "happynum", cmd.Name)
	assert.Equal(t, "Distinct Happy Number Range Counter", cmd.Usage)
	assert.NotEmpty(t, cmd.Flags)
	assert.Len(t, cmd.Authors, 1)
}

func TestNewCommandRunsWithRangeFlag(t *testing.T) {
	cmd := newCommand()

	err := cmd.Run(context.Background(), []string{"test", "--range", "100"})
	assert.NoError(t, err)
}

func TestNewCommandRunsWithSingleFlag(t *testing.T) {
	cmd := newCommand()

	err := cmd.Run(context.Background(), []string{"test", "--range", "100", "--single"})
	assert.NoError(t, err)
}
