package day04

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewMap(t *testing.T) {
	lines := []string{
		"12",
		"34",
	}
	m := NewMapFromLines(lines)

	tests := []struct {
		name string
		x, y int
		want byte
	}{
		{"00", 0, 0, '1'},
		{"01", 0, 1, '2'},
		{"10", 1, 0, '3'},
		{"11", 1, 1, '4'},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got, _ := m.Get(tc.x, tc.y)
			assert.Equal(t, tc.want, got)
		})
	}
}

func TestMap_CountRollsAt(t *testing.T) {
	lines := []string{
		"@@@",
		"@.@",
		"@@@",
	}
	m := NewMapFromLines(lines)

	assert.Equal(t, 8, m.countRollsAt(1, 1))
	assert.Equal(t, 2, m.countRollsAt(0, 0))
	assert.Equal(t, 4, m.countRollsAt(1, 0))
}

func TestMap_CountValidPositions(t *testing.T) {
	lines := []string{
		"@@@",
		"@.@",
		"@@@",
	}
	m := NewMapFromLines(lines)

	got := m.MarkValidPositions()

	assert.Equal(t, 4, got)
}
