package day02

import (
	"regexp"
	"strconv"
)

type ProductRange struct {
	from int
	to   int
}

var rangeRE = regexp.MustCompile(`^(\d+)-(\d+)$`)

func ProductRangeFromString(input string) ProductRange {
	parts := rangeRE.FindStringSubmatch(input)
	from, _ := strconv.Atoi(parts[1])
	to, _ := strconv.Atoi(parts[2])

	return ProductRange{from: from, to: to}
}

func (r ProductRange) SumInvalidIDs() (doubleSum, repeatSum int) {
	for i := r.from; i <= r.to; i++ {
		if CheckSequenceDoubles(i) {
			doubleSum += i
		}
		if CheckSequence(i) {
			repeatSum += i
		}
	}
	return doubleSum, repeatSum
}

func (r ProductRange) invalidDoubleIds() []int {
	invalidIds := []int{}
	for i := r.from; i <= r.to; i++ {
		if CheckSequenceDoubles(i) {
			invalidIds = append(invalidIds, i)
		}
	}

	return invalidIds
}

func (r ProductRange) invalidIds() []int {
	invalidIds := []int{}
	for i := r.from; i <= r.to; i++ {
		if CheckSequence(i) {
			invalidIds = append(invalidIds, i)
		}
	}

	return invalidIds
}
