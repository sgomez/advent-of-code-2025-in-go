package day01

import (
	"embed"
	"fmt"

	"github.com/sgomez/advent-of-code-2025-in-go/utils"
)

const message = "Count clicks %d pass %d"

//go:embed *.txt
var inputFS embed.FS

func Run() string {
	dial := NewDial()

	lines, _ := utils.ReadLines(inputFS, "input01.txt")

	for line := range lines {
		order := NewOrder(lines[line])
		dial.Move(order)
	}

	return fmt.Sprintf(message, dial.ZeroHits(), dial.ZeroCrossings())
}
