package day05

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

func (p ProductRange) isFresh(id int) bool {
	return p.from <= id && id <= p.to
}

func (p ProductRange) ToArray() []int {
	size := p.to - p.from + 1
	if size <= 0 {
		return []int{}
	}

	values := make([]int, 0, size)
	for value := p.from; value <= p.to; value++ {
		values = append(values, value)
	}

	return values
}
