package main

import (
	"AdventOfCode/utils"
	"fmt"
	"strconv"
	"strings"
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

	fmt.Printf("The answer for Star 1 (%s) is: %d\n", fileType, question1(lines))
	fmt.Printf("The answer for Star 2 (%s) is: %d\n", fileType, question2(lines))
}

func question1(lines []string) int {
	board := utils.ConstructBoard(lines)
	return solve(board, nil, false)
}

func question2(lines []string) int {
	board := utils.ConstructBoard(lines)
	return solve(board, make([]string, 0), true)
}

func solve(board [][]string, paths []string, distinctPath bool) int {
	sum := 0
	for i := range board {
		for j := range board[i] {
			if board[i][j] == "0" {
				visited := make(map[string]bool)
				sum += hillExplore(i, j, board, visited, paths, distinctPath)
			}
		}
	}
	return sum
}

func hillExplore(i, j int, m [][]string, visited map[string]bool, paths []string, distinctPath bool) int {
	if distinctPath {
		paths = append(paths, fmt.Sprintf("%d,%d", i, j))
	}

	if m[i][j] == "9" {
		var pathKey string
		if distinctPath {
			pathKey = strings.Join(paths, "|")
		} else {
			pathKey = fmt.Sprintf("%d,%d", i, j)
		}

		if !visited[pathKey] {
			visited[pathKey] = true
			return 1
		}
		return 0
	}

	// for test cases
	if m[i][j] == "." {
		return 0
	}

	count := 0
	directions := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	currentHeight, _ := strconv.Atoi(m[i][j])

	for _, dir := range directions {
		newI := i + dir[0]
		newJ := j + dir[1]

		if newI < 0 || newI >= len(m) || newJ < 0 || newJ >= len(m[0]) {
			continue
		}

		nextHeight, _ := strconv.Atoi(m[newI][newJ])
		if currentHeight+1 == nextHeight {
			count += hillExplore(newI, newJ, m, visited, paths, distinctPath)
		}
	}

	return count
}
