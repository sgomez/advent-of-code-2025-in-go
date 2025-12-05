package day05

import "sort"

type ProductRangeCollection struct {
	ranges []ProductRange
}

func NewFromLines(lines []string) ProductRangeCollection {
	productRanges := make([]ProductRange, 0, len(lines))
	for _, line := range lines {
		if line == "" {
			continue
		}

		productRanges = append(productRanges, ProductRangeFromString(line))
	}

	return ProductRangeCollection{ranges: productRanges}
}

func (p ProductRangeCollection) IsFresh(id int) bool {
	for _, productRange := range p.ranges {
		if productRange.isFresh(id) {
			return true
		}
	}

	return false
}

func (p ProductRangeCollection) CountFreshProducts(ids []int) int {
	count := 0
	for _, id := range ids {
		if p.IsFresh(id) {
			count++
		}
	}

	return count
}

func (p ProductRangeCollection) UniqueProductIDs() []int {
	unique := make([]int, 0)
	seen := make(map[int]struct{})

	for _, productRange := range p.ranges {
		for id := productRange.from; id <= productRange.to; id++ {
			if _, exists := seen[id]; exists {
				continue
			}

			seen[id] = struct{}{}
			unique = append(unique, id)
		}
	}

	return unique
}

func (p ProductRangeCollection) UniqueProductCount() int {
	if len(p.ranges) == 0 {
		return 0
	}

	ranges := make([]ProductRange, len(p.ranges))
	copy(ranges, p.ranges)

	sort.Slice(ranges, func(i, j int) bool {
		return ranges[i].from < ranges[j].from
	})

	total := 0
	current := ranges[0]

	for i := 1; i < len(ranges); i++ {
		next := ranges[i]
		if next.from <= current.to {
			if next.to > current.to {
				current.to = next.to
			}
			continue
		}

		total += current.to - current.from + 1
		current = next
	}

	total += current.to - current.from + 1

	return total
}
