package utils

import (
	"strconv"
)

func StringsToIntegers(inputs []string) ([]int, error) {
	integers := make([]int, 0, len(inputs))
	for _, line := range inputs {
		n, err := strconv.Atoi(line)
		if err != nil {
			return nil, err
		}
		integers = append(integers, n)
	}
	return integers, nil
}
