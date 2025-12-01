package utils

import (
	"bufio"
	"io/fs"
)

func ReadLines(source fs.FS, filename string) ([]string, error) {
	file, err := source.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return lines, nil
}
