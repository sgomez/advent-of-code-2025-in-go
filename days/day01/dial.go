package day01

import "math"

type Dial struct {
	value        int
	zeroHits     int
	zeroCrossing int
}

func NewDial() *Dial {
	return &Dial{value: 50, zeroHits: 0, zeroCrossing: 0}
}

func (d *Dial) Value() int {
	return d.value
}

func (d *Dial) ZeroHits() int {
	return d.zeroHits
}

func (d *Dial) ZeroCrossings() int {
	return d.zeroCrossing
}

func (d *Dial) Move(order Order) {
	d.rotate(order.Steps())
}

func (d *Dial) rotate(delta int) {
	if delta == 0 {
		return
	}

	d.zeroCrossing += countZeroCrossings(d.value, delta)
	d.value = flooredMod(d.value+delta, 100)

	if d.value == 0 {
		d.zeroHits++
	}
}

func countZeroCrossings(start, delta int) int {
	if delta < 0 {
		start = (100 - start) % 100
	}

	steps := absInt(delta)
	return (start + steps) / 100
}

func flooredMod(a int, b int) int {
	return ((a % b) + b) % b
}

func absInt(a int) int {
	return int(math.Abs(float64(a)))
}
