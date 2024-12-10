package main

import (
	"AdventOfCode/utils"
	"fmt"
	"math"
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
	sum := 0
	for _, line := range lines {
		parts := strings.Split(line, ":")
		target, err := strconv.Atoi(parts[0])
		if err != nil {
			fmt.Printf("Error converting target to int: %v\n", err)
			continue
		}
		numStr := strings.TrimSpace(parts[1])
		nums, err := utils.StringsToIntegers(strings.Split(numStr, " "))
		if err != nil {
			fmt.Printf("Error converting strings to integers: %v\n", err)
			continue
		}

		if canMatch(target, nums, []string{"+", "*"}) {
			sum += target
		}
	}

	return sum
}

func question2(lines []string) int {
	sum := 0
	for _, line := range lines {
		parts := strings.Split(line, ":")
		target, err := strconv.Atoi(parts[0])
		if err != nil {
			fmt.Printf("Error converting target to int: %v\n", err)
			continue
		}
		numStr := strings.TrimSpace(parts[1])
		nums, err := utils.StringsToIntegers(strings.Split(numStr, " "))
		if err != nil {
			fmt.Printf("Error converting strings to integers: %v\n", err)
			continue
		}

		if canMatch(target, nums, []string{"+", "*", "||"}) {
			sum += target
		}
	}

	return sum
}

func concat(a, b int) int {
	val, _ := strconv.Atoi(strconv.Itoa(a) + strconv.Itoa(b))
	return val
}

func evaluate(nums []int, ops []string) int {
	result := nums[0]
	opMap := map[string]func(int, int) int{
		"+":  func(a, b int) int { return a + b },
		"*":  func(a, b int) int { return a * b },
		"||": concat,
	}

	for i, op := range ops {
		if fn, ok := opMap[op]; ok {
			result = fn(result, nums[i+1])
		}
	}
	return result
}

func canMatch(target int, nums []int, operators []string) bool {
	if len(nums) == 1 {
		return nums[0] == target
	}
	opCount := len(operators)
	n := len(nums) - 1

	totalCombos := math.Pow(float64(opCount), float64(n))

	for opCombo := 0; opCombo < int(totalCombos); opCombo++ {
		ops := make([]string, n)
		curr := opCombo
		for i := 0; i < n; i++ {
			ops[i] = operators[curr%opCount]
			curr /= opCount
		}

		if evaluate(nums, ops) == target {
			return true
		}
	}

	return false
}
