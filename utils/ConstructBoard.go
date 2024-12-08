package utils

import "strings"

func ConstructBoard(lines []string) [][]string {
	board := make([][]string, len(lines))
	for idx, line := range lines {
		board[idx] = strings.Split(line, "")
	}
	return board
}
