package main

import (
	"AdventOfCode/utils"
	"fmt"
	"strings"
)

var DirectionSymbols = []string{"^", ">", "v", "<"}

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

	fmt.Printf("The answer for Star 1 (%s) is: %d\n", fileType, question1(lines))
	fmt.Printf("The answer for Star 2 (%s) is: %d\n", fileType, question2(lines))
}

func question1(lines []string) int {
	board := parseBoard(lines)
	guardX, guardY := locateGuard(board)
	direction := board[guardX][guardY]

	for step := 0; step < len(board)*len(board[0]); step++ {
		nextX, nextY := moveForward(guardX, guardY, direction)

		if isOutOfBounds(nextX, nextY, board) {
			break
		}

		switch cell := board[nextX][nextY]; cell {
		case ".":
			board[nextX][nextY] = direction
			guardX, guardY = nextX, nextY
		case "#":
			direction = getNextDirection(direction)
		default:
			if isDirectionSymbol(cell) {
				if cell == direction {
					break
				}
				board[nextX][nextY] = direction
				guardX, guardY = nextX, nextY
			}
		}
	}

	return countGuardMovements(board)
}

func question2(lines []string) int {
	board := parseBoard(lines)
	totalValidPositions := 0
	maxSteps := len(board) * len(board[0])

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if board[i][j] == "." {
				modifiedBoard := utils.DeepCopyBoard(board)
				modifiedBoard[i][j] = "#"

				guardX, guardY := locateGuard(modifiedBoard)
				direction := modifiedBoard[guardX][guardY]

				nextX, nextY := -1, -1
				for step := 0; step < maxSteps; step++ {
					nextX, nextY = moveForward(guardX, guardY, direction)

					if isOutOfBounds(nextX, nextY, modifiedBoard) {
						break
					}

					switch cell := modifiedBoard[nextX][nextY]; cell {
					case ".":
						modifiedBoard[nextX][nextY] = direction
						guardX, guardY = nextX, nextY
					case "#":
						direction = getNextDirection(direction)
					default:
						if isDirectionSymbol(cell) {
							if cell == direction {
								break
							}
							modifiedBoard[nextX][nextY] = direction
							guardX, guardY = nextX, nextY
						}
					}
				}

				if isValidPosition(nextX, nextY, modifiedBoard) && modifiedBoard[nextX][nextY] != "#" {
					totalValidPositions++
				}
			}
		}
	}

	return totalValidPositions
}

func parseBoard(lines []string) [][]string {
	board := make([][]string, len(lines))
	for idx, line := range lines {
		board[idx] = strings.Split(line, "")
	}
	return board
}

func locateGuard(board [][]string) (int, int) {
	for i, row := range board {
		for j, cell := range row {
			if isDirectionSymbol(cell) {
				return i, j
			}
		}
	}
	return -1, -1
}

func getNextDirection(currentDir string) string {
	switch currentDir {
	case "^":
		return ">"
	case ">":
		return "v"
	case "v":
		return "<"
	case "<":
		return "^"
	default:
		return ""
	}
}

func isDirectionSymbol(symbol string) bool {
	for _, dir := range DirectionSymbols {
		if dir == symbol {
			return true
		}
	}
	return false
}

func moveForward(x, y int, direction string) (int, int) {
	switch direction {
	case "^":
		return x - 1, y
	case ">":
		return x, y + 1
	case "v":
		return x + 1, y
	case "<":
		return x, y - 1
	default:
		return -1, -1
	}
}

func countGuardMovements(board [][]string) int {
	count := 0
	for _, row := range board {
		for _, cell := range row {
			if isDirectionSymbol(cell) {
				count++
			}
		}
	}
	return count
}

func isOutOfBounds(x, y int, board [][]string) bool {
	return x < 0 || y < 0 || x >= len(board) || y >= len(board[0])
}

func isValidPosition(x, y int, board [][]string) bool {
	return x >= 0 && y >= 0 && x < len(board) && y < len(board[0])
}
