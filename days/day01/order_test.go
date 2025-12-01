package day01

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewRightOrder(t *testing.T) {
	order := NewOrder("R10")

	assert.Equal(t, order.Steps(), 10)
}

func TestNewLeftOrder(t *testing.T) {
	order := NewOrder("L5")

	assert.Equal(t, order.Steps(), -5)
}
