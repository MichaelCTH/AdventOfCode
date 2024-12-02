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

func isSafeReport(nums []int) bool {
	if nums[0] == nums[1] {
		return false
	}

	increasing := nums[0] < nums[1]
	isSafe := true

	for i := 0; i < len(nums)-1; i++ {
		diff := nums[i+1] - nums[i]
		if (increasing && (diff <= 0 || diff > 3)) || (!increasing && (diff >= 0 || diff < -3)) {
			isSafe = false
			break
		}
	}

	return isSafe

}

func question1(lines []string) int {
	var numOfSafe int

	for _, line := range lines {
		values := strings.Split(line, " ")
		nums, err := utils.StringsToIntegers(values)
		if err != nil {
			fmt.Printf("Error converting string to integers: %v\n", err)
			continue
		}

		if isSafeReport(nums) {
			numOfSafe++
		}
	}

	return numOfSafe
}

func question2(lines []string) int {
	var numOfSafe int

	for _, line := range lines {
		values := strings.Split(line, " ")
		nums, err := utils.StringsToIntegers(values)
		if err != nil {
			fmt.Printf("Error converting string to integers: %v\n", err)
			continue
		}

		if isSafeReport(nums) {
			numOfSafe++
			continue
		}

		for i := 0; i < len(nums); i++ {
			newSlice := make([]int, 0, len(nums)-1)
			newSlice = append(newSlice, nums[:i]...)
			newSlice = append(newSlice, nums[i+1:]...)

			if isSafeReport(newSlice) {
				numOfSafe++
				break
			}
		}
	}

	return numOfSafe
}
