package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

var rulesMap = make(map[int][]int)
var pageNumbersMatrix = make([][]int, 204)

func handleRules(scanner *bufio.Scanner) {
	for scanner.Scan() {
		val := scanner.Text()
		rulesArray := strings.Split(val, "|")

		var rulesIntArray []int

		for _, rule := range rulesArray {
			ruleInt, err := strconv.Atoi(rule)
			check(err)
			rulesIntArray = append(rulesIntArray, ruleInt)
		}

		rulesMap[rulesIntArray[0]] = append(rulesMap[rulesIntArray[0]], rulesIntArray[1])
	}
}

func handlePageNumbers(scanner *bufio.Scanner) {
	i := 0
	for scanner.Scan() {
		val := scanner.Text()
		pageNumbersArray := strings.Split(val, ",")

		var pageNumbersIntArray []int

		for _, pageNumberString := range pageNumbersArray {
			pageNumberInt, err := strconv.Atoi(pageNumberString)
			check(err)
			pageNumbersIntArray = append(pageNumbersIntArray, pageNumberInt)
		}

		pageNumbersMatrix[i] = append(pageNumbersMatrix[i], pageNumbersIntArray...)
		i++
	}
}

func readFromFile(filename string, typeOfFile string) {
	file, err := os.Open(filename)
	check(err)

	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	switch typeOfFile {
	case "rules":
		handleRules(scanner)
	case "pageNumbers":
		handlePageNumbers(scanner)
	}
}

func main() {
	start := time.Now()
	readFromFile("./rules.txt", "rules")
	readFromFile("./page_numbers.txt", "pageNumbers")

	result := first()
	fmt.Println(result)
	resultSecond := second()
	fmt.Println(resultSecond)
	fmt.Println(time.Since(start))

}
