package day03

import "slices"

type Bank struct {
	batteries []int
}

func NewBankFromString(s string) *Bank {
	numbers := make([]int, len(s))
	for i, ch := range s {
		numbers[i] = int(ch - '0')
	}
	return &Bank{numbers}
}

func (b Bank) LargestJoltage(size int) int {
	var selectedBatteries []int
	bank := b.batteries
	for len(selectedBatteries) < size {
		pos := findMaxJoltageValidIndex(bank, size-len(selectedBatteries))
		battery := bank[pos]
		selectedBatteries = append(selectedBatteries, battery)
		bank = bank[pos+1:]
	}

	return vectorToNumber(selectedBatteries)
}

func findMaxJoltageValidIndex(batteries []int, remaining int) int {
	for i := 9; i >= 0; i-- {
		pos := slices.Index(batteries, i)
		if pos == -1 {
			continue
		}
		if pos+remaining <= len(batteries) {
			return pos
		}
	}

	panic("no solution found")
}

func vectorToNumber(v []int) int {
	n := 0
	for _, d := range v {
		n = n*10 + d
	}

	return n
}
