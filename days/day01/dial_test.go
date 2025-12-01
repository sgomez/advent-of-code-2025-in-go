package day01

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewDial(t *testing.T) {
	dial := NewDial()

	expected := 50
	assert.Equal(t, expected, dial.value)
}

func TestTurnRight(t *testing.T) {
	tests := []struct {
		name     string
		move     string
		expected int
	}{
		{"turn right 49", "R49", 99},
		{"turn right 50", "R50", 0},
		{"turn right 51", "R51", 1},
		{"turn right 150", "R150", 0},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			dial := NewDial()

			dial.Move(NewOrder(test.move))
			got := dial.Value()

			assert.Equal(t, test.expected, got)
		})
	}
}

func TestTurnLeft(t *testing.T) {
	tests := []struct {
		name     string
		move     string
		expected int
	}{
		{"turn left 49", "L49", 1},
		{"turn left 50", "L50", 0},
		{"turn left 51", "L51", 99},
		{"turn left 150", "L150", 0},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			dial := NewDial()

			dial.Move(NewOrder(test.move))
			got := dial.Value()

			assert.Equal(t, test.expected, got)
		})
	}
}

func TestCountHitZero(t *testing.T) {
	tests := []struct {
		name     string
		movement string
		expected int
	}{
		{"turn left 49", "L49", 1},
		{"turn right 1", "R1", 2},
		{"turn right 100", "R100", 2},
		{"turn left 1", "L1", 1},
	}
	dial := NewDial()

	for _, test := range tests {
		order := NewOrder(test.movement)
		t.Run(test.name, func(t *testing.T) {
			dial.Move(order)
			got := dial.Value()

			assert.Equal(t, test.expected, got)
		})
	}
}

func TestCountPassZero(t *testing.T) {
	tests := []struct {
		name     string
		movement string
		expected int
	}{
		{"turn right 49", "R49", 0},
		{"turn right 50", "R50", 1},
		{"turn left 49", "L49", 0},
		{"turn left 50", "L50", 1},
		{"turn right 149", "R149", 1},
		{"turn right 150", "R150", 2},
		{"turn left 149", "L149", 1},
		{"turn left 150", "L150", 2},
		{"turn right 1000", "R1000", 10},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			dial := NewDial()

			order := NewOrder(test.movement)
			dial.Move(order)
			got := dial.ZeroCrossings()

			assert.Equal(t, test.expected, got)
		})
	}
}
