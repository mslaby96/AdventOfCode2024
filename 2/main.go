package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strconv"
	"time"
	"unicode"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func readFromFile() map[int][]int {
	f, err := os.Open("./data.txt")
	check(err)

	defer f.Close()

	reports := make(map[int][]int)

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanBytes)
	i := 0
	var arr []int
	var tempString string
	for scanner.Scan() {
		val := scanner.Bytes()
		valRune := bytes.Runes(val)

		if unicode.IsSpace(valRune[0]) {
			number, err := strconv.Atoi(string(tempString))
			check(err)
			arr = append(arr, number)
			tempString = ""
		}

		if string(val) == "\n" {
			reports[i] = arr
			arr = nil
			i++
		}

		if !unicode.IsSpace(valRune[0]) {
			tempString = tempString + string(val)
		}

	}
	return reports
}

func main() {
	start := time.Now()
	safeReportsFirst := first()
	safeReportsSecond := second()
	fmt.Println("first exercise safe reports:", safeReportsFirst)
	fmt.Println("second excercise safe reports:", safeReportsSecond)
	timeElapsed := time.Since(start)
	fmt.Println("execution time:", timeElapsed)
}
