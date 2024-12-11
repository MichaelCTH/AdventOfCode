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
	stones := strings.Split(lines[0], " ")
	return solve(stones, 25)
}

func question2(lines []string) int {
	stones := strings.Split(lines[0], " ")
	return solve(stones, 75)
}

func solve(stones []string, times int) int {
	count := 0
	for i := range stones {
		count += blink(stones[i], times, make(map[string]int, 0))
	}

	return count
}

func blink(stone string, times int, memo map[string]int) int {
	key := stone + "|" + strconv.Itoa(times)
	if val, ok := memo[key]; ok {
		return val
	}

	if times == 0 {
		memo[key] = 1
		return 1
	}

	num, _ := strconv.Atoi(stone)
	if num == 0 {
		memo[key] = blink("1", times-1, memo)
	} else if len(stone)%2 == 0 {
		left := stone[:len(stone)/2]
		right := stone[len(stone)/2:]
		cleanRight, _ := strconv.Atoi(right)
		memo[key] = blink(left, times-1, memo) + blink(strconv.Itoa(cleanRight), times-1, memo)
	} else {
		memo[key] = blink(strconv.Itoa(num*2024), times-1, memo)
	}

	return memo[key]
}
