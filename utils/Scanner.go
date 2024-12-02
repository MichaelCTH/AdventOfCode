package utils

import (
	"bufio"
	"os"
)

func ReadLinesFromFile(filename string) ([]string, error) {
	var lines []string

	// Open the file
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// Read the file line by line using a scanner
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text()) // Append each line to the slice
	}

	// Check for errors during the reading process
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return lines, nil
}
