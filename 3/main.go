package main

import (
	"fmt"
	"os"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func readFromFile() string {
	f, err := os.ReadFile("./data.txt")
	check(err)

	return string(f)

}

func main() {
	firstAnswer := first()
	secondAnswer := second()
	fmt.Println(firstAnswer)
	fmt.Println(secondAnswer)
}
