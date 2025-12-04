package day04

import (
	"embed"
	"fmt"

	"github.com/sgomez/advent-of-code-2025-in-go/utils"
)

const message = "Day 03. Question 1: %d || Question 2: %d"

//go:embed *.txt
var inputFS embed.FS

func Run(filename string) string {
	lines, _ := utils.ReadLines(inputFS, filename)

	q1 := 0
	q2 := 0

	m := NewMapFromLines(lines)
	q1 = m.MarkValidPositions()
	m = NewMapFromLines(lines)
	q2 = m.RemoveRollsUntilStable()

	return fmt.Sprintf(message, q1, q2)
}
