package main

import (
	"slices"
)

func checkRules(currentPage int, row int, col int) bool {
	for j := col; j >= 0; j-- {
		if slices.Contains(rulesMap[currentPage], pageNumbersMatrix[row][j]) {
			return false
		}
	}
	return true
}

var correctRules = make([][]int, 0)

func first() int {
	sum := 0

	for i := range pageNumbersMatrix {
		isError := false
		for j := range pageNumbersMatrix[i] {
			result := checkRules(pageNumbersMatrix[i][j], i, j)
			if !result {
				isError = true
				break
			}
		}
		if !isError {
			correctRules = append(correctRules, pageNumbersMatrix[i])
		}
	}

	for _, rules := range correctRules {
		middlePoint := (len(rules) / 2)
		sum += rules[middlePoint]
	}
	return sum
}
