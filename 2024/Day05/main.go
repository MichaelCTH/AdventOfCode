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
	return ProcessSequences(lines, true)
}

func question2(lines []string) int {
	return ProcessSequences(lines, false)
}

func ProcessSequences(lines []string, checkEquality bool) int {
	orders, sequences := parseLines(lines)
	orderMap := buildOrderMap(orders)

	sum := 0
	for _, seq := range sequences {
		originalSeq := append([]string(nil), seq...)

		sort.Slice(seq, func(i, j int) bool {
			return compareOrder(seq[i], seq[j], orderMap)
		})

		isEqual := strings.Join(originalSeq, ",") == strings.Join(seq, ",")
		if (checkEquality && isEqual) || (!checkEquality && !isEqual) {
			sum += getMiddleValue(seq)
		}
	}

	return sum
}

func parseLines(lines []string) ([]string, [][]string) {
	var orders []string
	var sequences [][]string
	isOrderSection := true

	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			isOrderSection = false
			continue
		}

		if isOrderSection {
			orders = append(orders, line)
		} else {
			seq := strings.Split(line, ",")
			sequences = append(sequences, seq)
		}
	}

	return orders, sequences
}

func buildOrderMap(orders []string) map[string]map[string]bool {
	orderMap := make(map[string]map[string]bool)
	for _, o := range orders {
		parts := strings.Split(o, "|")
		a, b := strings.TrimSpace(parts[0]), strings.TrimSpace(parts[1])
		if orderMap[a] == nil {
			orderMap[a] = make(map[string]bool)
		}
		orderMap[a][b] = true
	}
	return orderMap
}

func compareOrder(a, b string, orderMap map[string]map[string]bool) bool {
	if orderMap[a][b] {
		return true
	}
	if orderMap[b][a] {
		return false
	}

	return a < b
}

func getMiddleValue(seq []string) int {
	mid := len(seq) / 2
	val, _ := strconv.Atoi(seq[mid])
	return val
}
