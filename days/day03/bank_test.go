package day03

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewBankFromString(t *testing.T) {
	bank := NewBankFromString("1234")

	expected := []int{1, 2, 3, 4}
	assert.Equal(t, expected, bank.batteries)
}

func TestMaxJoltage(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected int
	}{
		{"Check 987654321111111", "987654321111111", 98},
		{"Check 811111111111119", "811111111111119", 89},
		{"Check 234234234234278", "234234234234278", 78},
		{"Check 818181911112111", "818181911112111", 92},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			bank := NewBankFromString(test.input)

			got := bank.LargestJoltage(2)

			assert.Equal(t, test.expected, got)
		})
	}
}

func TestLargestJoltage(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected int
	}{
		{"Check 987654321111111", "987654321111111", 987654321111},
		{"Check 811111111111119", "811111111111119", 811111111119},
		{"Check 234234234234278", "234234234234278", 434234234278},
		{"Check 818181911112111", "818181911112111", 888911112111},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			bank := NewBankFromString(test.input)

			got := bank.LargestJoltage(12)

			assert.Equal(t, test.expected, got)
		})
	}
}
