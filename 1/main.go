package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func readFromFile(listOne *[]int, listTwo *[]int) {
	f, err := os.Open("./data.txt")
	check(err)

	for {
		l1 := make([]byte, 5)
		spaces := make([]byte, 3)
		l2 := make([]byte, 5)
		nl := make([]byte, 1)

		_, err := f.Read(l1)
		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Println(err)
			return
		}

		_, err = f.Read(spaces)
		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Println(err)
			return
		}

		_, err = f.Read(l2)
		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Println(err)
			return
		}
		_, err = f.Read(nl)
		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Println(err)
			return
		}
		l1Int, err := strconv.Atoi(string(l1))
		check(err)

		l2Int, err := strconv.Atoi(string(l2))
		check(err)

		*listOne = append(*listOne, l1Int)
		*listTwo = append(*listTwo, l2Int)
	}
}

func main() {
	firstExercise()
	secondExercise()
}
