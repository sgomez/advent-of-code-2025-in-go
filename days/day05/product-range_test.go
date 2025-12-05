package day05

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewProductRange(t *testing.T) {
	productRange := ProductRangeFromString("11-22")

	assert.Equal(t, 11, productRange.from)
	assert.Equal(t, 22, productRange.to)
}

func TestInvalidIds(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		id       int
		expected bool
	}{
		{"Check 11 in 3-5", "3-5", 1, false},
		{"Check 11 in 3-5", "3-5", 5, true},
		{"Check 11 in 3-5", "3-5", 8, false},
		{"Check 22 in 11-22", "11-22", 22, true},
		{"Check 10 not in 11-22", "11-22", 10, false},
		{"Check 23 not in 11-22", "11-22", 23, false},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			productRange := ProductRangeFromString(test.input)

			got := productRange.isFresh(test.id)

			assert.Equal(t, test.expected, got)
		})
	}
}

func TestToArray(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected []int
	}{
		{"Single value", "5-5", []int{5}},
		{"Range size 3", "1-3", []int{1, 2, 3}},
		{"Large range", "11-14", []int{11, 12, 13, 14}},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			productRange := ProductRangeFromString(test.input)

			assert.Equal(t, test.expected, productRange.ToArray())
		})
	}
}
