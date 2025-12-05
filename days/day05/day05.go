package day05

import (
	"embed"
	"fmt"
	"strconv"

	"github.com/sgomez/advent-of-code-2025-in-go/utils"
)

const message = "Day 03. Question 1: %d || Question 2: %d"

//go:embed *.txt
var inputFS embed.FS

func Run(filename string) string {

	lines, _ := utils.ReadLines(inputFS, filename)

	var productRangeLines []string
	var ids []int
	parsingRanges := true

	for _, line := range lines {
		if line == "" {
			parsingRanges = false
			continue
		}

		if parsingRanges {
			productRangeLines = append(productRangeLines, line)
			continue
		}

		id, err := strconv.Atoi(line)
		if err != nil {
			continue
		}

		ids = append(ids, id)
	}

	productRangeCollection := NewFromLines(productRangeLines)

	q1 := productRangeCollection.CountFreshProducts(ids)

	q2 := productRangeCollection.UniqueProductCount()

	return fmt.Sprintf(message, q1, q2)

}
