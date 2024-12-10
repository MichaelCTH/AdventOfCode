package main

import (
	"AdventOfCode/utils"
	"fmt"
	"sort"
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

	fmt.Printf("The answer of Star 1 (%s) is: %d\n", fileType, question1(lines))
	fmt.Printf("The answer of Star 2 (%s) is: %d\n", fileType, question2(lines))
}

func question1(lines []string) int {
	leftValues := make([]int, len(lines))
	rightValues := make([]int, len(lines))

	for i, line := range lines {
		values := strings.Split(line, "   ")
		left, _ := strconv.Atoi(values[0])
		right, _ := strconv.Atoi(values[1])
		leftValues[i] = left
		rightValues[i] = right
	}

	sort.Ints(leftValues)
	sort.Ints(rightValues)

	var totalDif int
	for i := range leftValues {
		totalDif += utils.Abs(leftValues[i] - rightValues[i])
	}

	return totalDif
}

func question2(lines []string) int {
	leftValues := make([]int, len(lines))
	rightMap := make(map[int]int)

	for i, line := range lines {
		values := strings.Split(line, "   ")
		left, _ := strconv.Atoi(values[0])
		right, _ := strconv.Atoi(values[1])

		leftValues[i] = left
		rightMap[right]++
	}

	var totalDif int
	for _, num := range leftValues {
		if times, ok := rightMap[num]; ok {
			totalDif += num * times
		}
	}

	return totalDif
}
