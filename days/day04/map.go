package day04

import "errors"

const (
	emptyCell  byte = '.'
	rollCell   byte = '@'
	markedCell byte = 'x'
)

type Map struct {
	cells     []byte
	dimension int
}

func NewMapFromLines(lines []string) *Map {
	dimension := len(lines)
	m := &Map{cells: make([]byte, dimension*dimension), dimension: dimension}

	for i, line := range lines {
		for j := range line {
			_ = m.Set(i, j, line[j])
		}
	}

	return m
}

func (s *Map) Set(x int, y int, value byte) error {
	if x < 0 || x >= s.dimension || y < 0 || y >= s.dimension {
		return errors.New("out of range")
	}

	s.cells[x*s.dimension+y] = value

	return nil
}

func (s *Map) Get(x int, y int) (byte, error) {
	if x < 0 || x >= s.dimension || y < 0 || y >= s.dimension {
		return emptyCell, errors.New("out of range")
	}

	return s.cells[x*s.dimension+y], nil
}

func (s *Map) countRollsAt(x int, y int) int {
	total := 0

	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			if i == 0 && j == 0 {
				continue
			}

			row := x + i
			col := y + j

			cell, err := s.Get(row, col)

			if err == nil && cell != emptyCell {
				total++
			}
		}
	}

	return total
}

func (s *Map) MarkValidPositions() int {
	total := 0

	for i := 0; i < s.dimension; i++ {
		for j := 0; j < s.dimension; j++ {
			cell, _ := s.Get(i, j)
			if cell != rollCell {
				continue
			}

			if s.countRollsAt(i, j) < 4 {
				total++
				_ = s.Set(i, j, markedCell)
			}
		}
	}

	return total
}

func (s *Map) clearMarked() {
	for i, cell := range s.cells {
		if cell == markedCell {
			s.cells[i] = emptyCell
		}
	}
}

func (s *Map) RemoveRollsUntilStable() int {
	total := 0

	for {
		marked := s.MarkValidPositions()
		if marked == 0 {
			break
		}
		total += marked
		s.clearMarked()
	}

	return total

}
