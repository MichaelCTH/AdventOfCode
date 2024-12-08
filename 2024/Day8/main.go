package main

import (
	"AdventOfCode/utils"
	"fmt"
)

type IntPair struct {
	X int
	Y int
}

const (
	occupied = "#"
	empty    = "."
)

func main() {
	processFile("./example.txt", "example")
	processFile("./input.txt", "input")
}

func processFile(filePath, fileType string) {
	lines, err := utils.ReadLinesFromFile(filePath)
	if err != nil {
		fmt.Printf("Error reading %s file: %v\n", fileType, err)
		return
	}

	fmt.Printf("The answer of Star 1 (%s) is: %d\n", fileType, question1(lines))
	fmt.Printf("The answer of Star 2 (%s) is: %d\n", fileType, question2(lines))
}

func question1(lines []string) int {
	board := utils.ConstructBoard(lines)
	signalMap := aggregateSignals(board)
	count := 0

	for _, pairs := range signalMap {
		for i := 0; i < len(pairs); i++ {
			for j := i + 1; j < len(pairs); j++ {
				drawLineOrDiagonal(pairs[i], pairs[j], board, true)
				drawLineOrDiagonal(pairs[j], pairs[i], board, true)
			}
		}
	}

	for _, line := range board {
		for _, loc := range line {
			if loc == occupied {
				count++
			}
		}
	}

	return count
}

func question2(lines []string) int {
	board := utils.ConstructBoard(lines)
	signalMap := aggregateSignals(board)
	count := 0

	for _, pairs := range signalMap {
		if len(pairs) == 1 {
			continue
		}
		for i := 0; i < len(pairs); i++ {
			for j := i + 1; j < len(pairs); j++ {
				drawLineOrDiagonal(pairs[i], pairs[j], board, false)
				drawLineOrDiagonal(pairs[j], pairs[i], board, false)
			}
		}
	}

	for _, line := range board {
		for _, loc := range line {
			if loc != empty {
				count++
			}
		}
	}

	return count
}

func aggregateSignals(board [][]string) map[string][]IntPair {
	signalMap := make(map[string][]IntPair)

	for i := range board {
		for j := range board[i] {
			if board[i][j] == empty {
				continue
			}

			signalMap[board[i][j]] = append(signalMap[board[i][j]], IntPair{X: i, Y: j})
		}
	}

	return signalMap
}

func drawLineOrDiagonal(a, b IntPair, board [][]string, isDiagonal bool) bool {
	disX := b.X - a.X
	disY := b.Y - a.Y

	newX := a.X - disX
	newY := a.Y - disY

	drawn := false

	for newX >= 0 && newY >= 0 && newX < len(board) && newY < len(board[0]) {
		board[newX][newY] = occupied
		drawn = true

		if isDiagonal {
			break
		}

		newX -= disX
		newY -= disY
	}

	return drawn
}
