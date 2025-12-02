package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/sgomez/advent-of-code-2025-in-go/days/day01"
	"github.com/sgomez/advent-of-code-2025-in-go/days/day02"
	"github.com/sgomez/advent-of-code-2025-in-go/days/day03"
	"github.com/sgomez/advent-of-code-2025-in-go/days/day04"
	"github.com/sgomez/advent-of-code-2025-in-go/days/day05"
	"github.com/sgomez/advent-of-code-2025-in-go/days/day06"
	"github.com/sgomez/advent-of-code-2025-in-go/days/day07"
	"github.com/sgomez/advent-of-code-2025-in-go/days/day08"
	"github.com/sgomez/advent-of-code-2025-in-go/days/day09"
	"github.com/sgomez/advent-of-code-2025-in-go/days/day10"
	"github.com/sgomez/advent-of-code-2025-in-go/days/day11"
	"github.com/sgomez/advent-of-code-2025-in-go/days/day12"
)

var runners = map[string]func(string) string{
	"01": day01.Run,
	"02": day02.Run,
	"03": day03.Run,
	"04": day04.Run,
	"05": day05.Run,
	"06": day06.Run,
	"07": day07.Run,
	"08": day08.Run,
	"09": day09.Run,
	"10": day10.Run,
	"11": day11.Run,
	"12": day12.Run,
}

func main() {
	sample := flag.Bool("sample", false, "use sample input file")
	flag.Parse()

	if flag.NArg() != 1 {
		fmt.Println("Usage: go run . [--sample] <day (01-12)>")
		os.Exit(1)
	}

	day := flag.Arg(0)
	run, ok := runners[day]
	if !ok {
		fmt.Printf("Unknown day %q. Expected 01-12.\n", day)
		os.Exit(1)
	}

	filename := fmt.Sprintf("day%s.txt", day)
	if *sample {
		filename = fmt.Sprintf("day%s-sample.txt", day)
	}

	fmt.Println(run(filename))
}
