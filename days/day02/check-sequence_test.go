package day02

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCheckSequenceDoubles(t *testing.T) {
	tests := []struct {
		name     string
		input    int
		expected bool
	}{
		{"Check 11", 11, true},
		{"Check 111", 111, false},
		{"Check 1111", 1111, true},
		{"Check 1110", 1110, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := CheckSequenceDoubles(tt.input)

			assert.Equal(t, tt.expected, got)
		})
	}
}

func TestCheckSequence(t *testing.T) {
	tests := []struct {
		name     string
		input    int
		expected bool
	}{
		{"Check 11", 11, true},
		{"Check 111", 111, true},
		{"Check 824824824", 824824824, true},
		{"Check 1110", 1110, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := CheckSequence(tt.input)

			assert.Equal(t, tt.expected, got)
		})
	}
}
