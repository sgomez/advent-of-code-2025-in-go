package day02

import (
	"embed"
	"fmt"
	"strings"

	"github.com/sgomez/advent-of-code-2025-in-go/utils"
)

const message = "Day 02. First problem: %d. Second problem: %d"

//go:embed *.txt
var inputFS embed.FS

func Run() string {
	lines, _ := utils.ReadLines(inputFS, "day02.txt")

	parts := strings.Split(lines[0], ",")
	total1 := 0
	total2 := 0

	for _, part := range parts {
		productRange := ProductRangeFromString(part)
		doubleSum, repeatSum := productRange.SumInvalidIDs()
		total1 += doubleSum
		total2 += repeatSum
	}

	return fmt.Sprintf(message, total1, total2)
}
