package day05

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewFromLines(t *testing.T) {
	lines := []string{
		"1-3",
		"5-7",
	}

	collection := NewFromLines(lines)

	assert.Equal(t, 2, len(collection.ranges))
	assert.Equal(t, ProductRange{from: 1, to: 3}, collection.ranges[0])
	assert.Equal(t, ProductRange{from: 5, to: 7}, collection.ranges[1])
}

func TestCollectionIsFresh(t *testing.T) {
	collection := NewFromLines([]string{"1-3", "5-7"})

	assert.True(t, collection.IsFresh(2))
	assert.True(t, collection.IsFresh(6))
	assert.False(t, collection.IsFresh(4))
	assert.False(t, collection.IsFresh(8))
}

func TestCountFreshProducts(t *testing.T) {
	collection := NewFromLines([]string{"1-3", "5-7"})

	assert.Equal(t, 3, collection.CountFreshProducts([]int{1, 2, 4, 6, 8}))
	assert.Equal(t, 0, collection.CountFreshProducts([]int{0, 4, 8}))
}

func TestUniqueProductIDs(t *testing.T) {
	collection := NewFromLines([]string{"1-3", "2-4", "6-6"})

	assert.Equal(t, []int{1, 2, 3, 4, 6}, collection.UniqueProductIDs())
}

func TestUniqueProductCount(t *testing.T) {
	tests := []struct {
		name     string
		input    []string
		expected int
	}{
		{"Single range", []string{"1-3"}, 3},
		{"Overlapping ranges", []string{"1-3", "2-4"}, 4},
		{"Disjoint ranges", []string{"1-2", "4-5"}, 4},
		{"Contained range", []string{"1-10", "3-5"}, 10},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			collection := NewFromLines(test.input)

			assert.Equal(t, test.expected, collection.UniqueProductCount())
		})
	}
}
