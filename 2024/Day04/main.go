package main

import (
	"AdventOfCode/utils"
	"fmt"
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

	fmt.Printf("The answer of Star 1 (%s) is: %d\n", fileType, question1(lines))
	fmt.Printf("The answer of Star 2 (%s) is: %d\n", fileType, question2(lines))
}

func question1(lines []string) int {
	searchMap := constructMap(lines)
	searchChars := []string{"X", "M", "A", "S"}
	masks := [][]int{{0, 1}, {1, 0}, {-1, 0}, {0, -1}, {-1, -1}, {1, 1}, {-1, 1}, {1, -1}}
	sum := 0

	for i, row := range searchMap {
		for j := range row {
			sum += matchStringFrom(i, j, masks, searchChars, searchMap)
		}
	}

	return sum
}

func question2(lines []string) int {
	searchMap := constructMap(lines)

	sum := 0

	for i, row := range searchMap {
		for j := range row {
			if searchMap[i][j] == "A" {
				sum += matchDiagonalXmas(i, j, searchMap)
			}
		}
	}

	return sum
}

func constructMap(lines []string) [][]string {
	searchMap := make([][]string, len(lines))

	for i, line := range lines {
		searchMap[i] = strings.Split(line, "")
	}

	return searchMap
}

func matchStringFrom(x int, y int, masks [][]int, str []string, searchMap [][]string) int {
	if searchMap[x][y] != str[0] {
		return 0
	}

	if len(str) == 1 {
		return 1
	}

	sum := 0
	for _, mask := range masks {
		nx, ny := x+mask[0], y+mask[1]
		if nx < 0 || ny < 0 || nx >= len(searchMap) || ny >= len(searchMap[0]) {
			continue
		}
		sum += matchStringFrom(nx, ny, [][]int{mask}, str[1:], searchMap)
	}
	return sum
}

func matchDiagonalXmas(x int, y int, searchMap [][]string) int {
	directions := [][][2]int{
		{{-1, -1}, {1, 1}},
		{{-1, 1}, {1, -1}},
	}

	for _, diagonal := range directions {
		combo := ""
		for _, offset := range diagonal {
			nx, ny := x+offset[0], y+offset[1]

			if nx < 0 || ny < 0 || nx >= len(searchMap) || ny >= len(searchMap[0]) {
				return 0
			}
			combo += searchMap[nx][ny]
		}
		if combo != "MS" && combo != "SM" {
			return 0
		}
	}
	return 1
}
