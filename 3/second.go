package main

import (
	"regexp"
	"strconv"
)

const stopInstruction = "don't()"
const releaseInstruction = "do()"

func calculate(instruction string) int {
	re := regexp.MustCompile("[0-9]+")
	foundNumbers := re.FindAllString(instruction, -1)
	number1, err := strconv.Atoi(foundNumbers[0])
	check(err)
	number2, err := strconv.Atoi(foundNumbers[1])
	check(err)
	return number1 * number2
}

func second() int {
	data := readFromFile()
	r, _ := regexp.Compile(`don't\(\)|do\(\)|mul\([0-9]+,[0-9]+\)`)

	foundInstructions := (r.FindAllString(data, -1))

	isStop := false
	sum := 0

	for _, instruction := range foundInstructions {
		if isStop {
			if instruction == releaseInstruction {
				isStop = false
			}
		} else if instruction == stopInstruction {
			isStop = true
		} else if instruction != stopInstruction && instruction != releaseInstruction {
			result := calculate(instruction)
			sum += result
		}
	}
	return sum
}
