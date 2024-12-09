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
	input := strings.Split(lines[0], "")
	spaces := readSpaces(input)
	moveSpaces(spaces)
	return checksum(spaces)
}

func question2(lines []string) int {
	input := strings.Split(lines[0], "")
	spaces := readSpaces(input)
	moveSpacesWholeFile(spaces)
	return checksum(spaces)
}

func readSpaces(input []string) []string {
	spaces := []string{}
	fileNum := 0

	for idx := 0; idx < len(input); idx++ {
		numI, _ := strconv.Atoi(input[idx])
		if idx%2 == 0 {
			strFileNum := strconv.Itoa(fileNum)
			for i := 0; i < numI; i++ {
				spaces = append(spaces, strFileNum)
			}
			fileNum++
		} else {
			for i := 0; i < numI; i++ {
				spaces = append(spaces, ".")
			}
		}
	}

	return spaces
}

func moveSpaces(input []string) {
	start := 0
	end := len(input) - 1

	for start < end {
		for start < len(input) && input[start] != "." {
			start++
		}

		for end >= 0 && input[end] == "." {
			end--
		}

		if start < end {
			input[start], input[end] = input[end], "."
		}
	}
}

func checksum(input []string) int {
	result := 0

	for idx, val := range input {
		if val == "." {
			continue
		} else {
			num, _ := strconv.Atoi(val)
			result += idx * num
		}
	}

	return result
}

func moveSpacesWholeFile(arr []string) {
	end := len(arr) - 1
	seen := make(map[string]bool)

	for end > 0 {
		for end >= 0 && arr[end] == "." {
			end--
		}

		if end < 0 {
			break
		}

		current := arr[end]
		id := current
		lenRun := 0
		for (end-lenRun) >= 0 && arr[end-lenRun] != "." && arr[end-lenRun] == id {
			lenRun++
		}

		if !seen[id] {
			head := findAvailableSpaesMatchLen(arr, end, lenRun)
			if head > -1 {
				for ptr := 0; ptr < lenRun; ptr++ {
					arr[head+ptr] = current
					arr[end-ptr] = "."
				}
			}
			seen[id] = true
		}

		end -= lenRun
	}
}

func findAvailableSpaesMatchLen(arr []string, end, lenRun int) int {
	head := 0
	for head < end-lenRun {
		for head < len(arr) && arr[head] != "." {
			head++
		}

		if head >= len(arr) {
			break
		}

		emptyLen := 0
		for (head+emptyLen) < len(arr) && arr[head+emptyLen] == "." {
			emptyLen++
		}

		if head > end-lenRun {
			break
		}

		if emptyLen >= lenRun {
			return head
		}

		head += emptyLen
	}

	return -1
}
