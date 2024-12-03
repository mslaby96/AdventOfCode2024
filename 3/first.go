package main

import (
	"regexp"
	"strconv"
)

func first() int {

	data := readFromFile()
	r, _ := regexp.Compile(`mul\([0-9]+,[0-9]+\)`)

	foundMulInstructions := (r.FindAllString(data, -1))

	sum := 0
	for _, instruction := range foundMulInstructions {
		re := regexp.MustCompile("[0-9]+")
		foundNumbers := re.FindAllString(instruction, -1)
		number1, err := strconv.Atoi(foundNumbers[0])
		check(err)
		number2, err := strconv.Atoi(foundNumbers[1])
		check(err)
		sum += number1 * number2
	}
	return sum
}
