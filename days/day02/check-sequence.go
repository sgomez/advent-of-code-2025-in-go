package day02

import (
	"slices"
	"strconv"
)

func CheckSequenceDoubles(n int) bool {
	if n < 0 {
		return false
	}

	s := strconv.Itoa(n)

	half := len(s) / 2
	return s[:half] == s[half:]
}

func CheckSequence(n int) bool {
	if n < 0 {
		return false
	}

	s := strconv.Itoa(n)
	l := len(s)

	for size := 1; size <= l/2; size++ {
		chunks, valid := chunk(s, size)
		if !valid {
			continue
		}

		equal := slices.Min(chunks) == slices.Max(chunks)
		if equal == true {
			return true
		}
	}

	return false
}

func chunk(s string, size int) ([]string, bool) {
	if size <= 0 || len(s)%size != 0 {
		return nil, false
	}
	var out []string
	for i := 0; i < len(s); i += size {
		out = append(out, s[i:i+size])
	}
	return out, true
}
