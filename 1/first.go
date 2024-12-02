package main

import (
	"fmt"
	"slices"
)

func firstExercise() {
	listOne := []int{}
	listTwo := []int{}

	readFromFile(&listOne, &listTwo)

	slices.Sort(listOne)
	slices.Sort(listTwo)

	sum := 0

	for index, element := range listOne {
		distance := element - listTwo[index]
		var distanceAbs = distance
		if distanceAbs < 0 {
			distanceAbs = -distance
		}
		sum += distanceAbs
	}

	fmt.Printf("result: %d", sum)

}
