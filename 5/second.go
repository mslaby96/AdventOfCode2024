package main

import (
	"slices"
)

func checkRulesSecond(currentPage int, row int, col int) (int, bool) {
	for j := col; j >= 0; j-- {
		if slices.Contains(rulesMap[currentPage], pageNumbersMatrix[row][j]) {
			return j, false
		}
	}
	return -1, true
}

func fixUpdate(incorrectUpdate []int, indexOfIncorrectUpdate int, indexOfCurrentRule int) []int {
	incorrectUpdate[indexOfIncorrectUpdate], incorrectUpdate[indexOfCurrentRule] = incorrectUpdate[indexOfCurrentRule], incorrectUpdate[indexOfIncorrectUpdate]
	return incorrectUpdate
}

var isFixed = false

func findIncorrectUpdates(row int) []int {
	isError := false
	for j := range pageNumbersMatrix[row] {
		index, ok := checkRulesSecond(pageNumbersMatrix[row][j], row, j)
		if !ok {
			isError = true
			fixUpdate(pageNumbersMatrix[row], index, j)
			isFixed = true
		}
	}
	if isError {
		return findIncorrectUpdates(row)
	}
	if isFixed {
		isFixed = false
		return pageNumbersMatrix[row]
	}
	return nil
}

var incorrectRules = make([][]int, 0)

func second() int {

	sum := 0
	for i := range pageNumbersMatrix {
		fixedUpdate := findIncorrectUpdates(i)
		if fixedUpdate != nil {
			incorrectRules = append(incorrectRules, fixedUpdate)
		}
	}

	for _, rules := range incorrectRules {
		middlePoint := (len(rules) / 2)
		sum += rules[middlePoint]
	}

	return sum
}
