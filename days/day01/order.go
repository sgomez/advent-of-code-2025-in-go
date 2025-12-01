package day01

import (
	"regexp"
	"strconv"
)

var orderRe = regexp.MustCompile(`^([RL])(\d+)$`)

type Order struct {
	direction string
	positions int
}

func NewOrder(order string) Order {
	parts := orderRe.FindStringSubmatch(order)
	direction := parts[1]
	positions, _ := strconv.Atoi(parts[2])

	return Order{direction: direction, positions: positions}
}

func (o Order) Steps() int {
	if o.direction == "L" {
		return -o.positions
	} else {
		return o.positions
	}
}
