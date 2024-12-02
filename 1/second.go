package main

import "fmt"

func secondExercise() {
	listOne := []int{}
	listTwo := []int{}

	readFromFile(&listOne, &listTwo)

	freq := make(map[int]int)

	for _, element := range listOne {
		for _, elementTwo := range listTwo {
			if element == elementTwo {
				freq[element]++
			}
		}
	}

	sum := 0

	for k, numberOfOccurances := range freq {
		sum += k * numberOfOccurances
	}

	fmt.Printf("sum of occurances: %d \n", sum)

}
