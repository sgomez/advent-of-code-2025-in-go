package day02

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewProductRange(t *testing.T) {
	productRange := ProductRangeFromString("11-22")

	assert.Equal(t, 11, productRange.from)
	assert.Equal(t, 22, productRange.to)
}

func TestInvalidDoubleIds(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected []int
	}{
		{"Check 11-22", "11-22", []int{11, 22}},
		{"Check 95-115", "95-115", []int{99}},
		{"Check 1698522-1698528", "1698522-1698528", []int{}},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			productRange := ProductRangeFromString(test.input)

			got := productRange.invalidDoubleIds()

			assert.Equal(t, test.expected, got)
		})
	}
}

func TestInvalidIds(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected []int
	}{
		{"Check 11-22", "11-22", []int{11, 22}},
		{"Check 95-115", "95-115", []int{99, 111}},
		{"Check 998-1012", "998-1012", []int{999, 1010}},
		{"Check 1188511880-1188511890", "1188511880-1188511890", []int{1188511885}},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			productRange := ProductRangeFromString(test.input)

			got := productRange.invalidIds()

			assert.Equal(t, test.expected, got)
		})
	}
}
