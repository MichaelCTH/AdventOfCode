package main

import (
	"AdventOfCode/utils"
	"fmt"
	"regexp"
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

	fmt.Printf("The answer of Star 1 (%s) is: %d\n", fileType, question1(strings.Join(lines, "")))
	fmt.Printf("The answer of Star 2 (%s) is: %d\n", fileType, question2(strings.Join(lines, "")))
}

func question1(line string) int {
	re := regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)`)
	matches := re.FindAllStringSubmatch(line, -1)
	sum := 0

	for _, match := range matches {
		num1, err1 := strconv.Atoi(match[1])
		num2, err2 := strconv.Atoi(match[2])
		if err1 == nil && err2 == nil {
			sum += num1 * num2
		}
	}

	return sum
}

func question2(line string) int {
	re := regexp.MustCompile(`(mul\((\d{1,3}),(\d{1,3})\)|do\(\)|don't\(\))`)
	sum := 0

	matches := re.FindAllStringSubmatch(line, -1)
	enabled := true
	for _, match := range matches {
		switch match[0] {
		case "do()":
			enabled = true
		case "don't()":
			enabled = false
		default:
			if enabled {
				num1, err1 := strconv.Atoi(match[2])
				num2, err2 := strconv.Atoi(match[3])
				if err1 == nil && err2 == nil {
					sum += num1 * num2
				}
			}
		}
	}

	return sum
}
