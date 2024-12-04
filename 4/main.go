package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

const x = 88
const m = 77
const a = 65
const s = 83

func check(err error) {
	if err != nil {
		panic(err)
	}
}

var dataMatrix = make([][]byte, 140)

func readFromFile() {
	file, err := os.Open("./data.txt")
	check(err)

	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	i := 0
	for scanner.Scan() {
		val := scanner.Bytes()
		dataMatrix[i] = make([]byte, len(val))
		dataMatrix[i] = append(dataMatrix[i], val...)
		i++
	}
}

func lookHorizontal(row int, col int) bool {
	switch dataMatrix[row][col] {
	case x:
		if len(dataMatrix[row]) > col+1 && dataMatrix[row][col+1] == m {
			if len(dataMatrix[row]) > col+2 && dataMatrix[row][col+2] == a {
				if len(dataMatrix[row]) > col+3 && dataMatrix[row][col+3] == s {
					return true
				}
			}
		}
	case s:
		if len(dataMatrix[row]) > col+1 && dataMatrix[row][col+1] == a {
			if len(dataMatrix[row]) > col+2 && dataMatrix[row][col+2] == m {
				if len(dataMatrix[row]) > col+3 && dataMatrix[row][col+3] == x {
					return true
				}
			}
		}
	}
	return false
}

func lookVertical(row int, col int) bool {
	switch dataMatrix[row][col] {
	case x:
		if len(dataMatrix) > row+1 && dataMatrix[row+1][col] == m {
			if len(dataMatrix) > row+2 && dataMatrix[row+2][col] == a {
				if len(dataMatrix) > row+3 && dataMatrix[row+3][col] == s {
					return true
				}
			}
		}
	case s:
		if len(dataMatrix) > row+1 && dataMatrix[row+1][col] == a {
			if len(dataMatrix) > row+2 && dataMatrix[row+2][col] == m {
				if len(dataMatrix) > row+3 && dataMatrix[row+3][col] == x {
					return true
				}
			}
		}
	}
	return false
}

func lookDiagonal(row int, col int) bool {
	switch dataMatrix[row][col] {
	case x:
		if len(dataMatrix) > row+1 && len(dataMatrix[row]) > col+1 && dataMatrix[row+1][col+1] == m {
			if len(dataMatrix) > row+2 && len(dataMatrix[row]) > col+2 && dataMatrix[row+2][col+2] == a {
				if len(dataMatrix) > row+3 && len(dataMatrix[row]) > col+3 && dataMatrix[row+3][col+3] == s {
					return true
				}
			}
		}
	case s:
		if len(dataMatrix) > row+1 && len(dataMatrix[row]) > col+1 && dataMatrix[row+1][col+1] == a {
			if len(dataMatrix) > row+2 && len(dataMatrix[row]) > col+2 && dataMatrix[row+2][col+2] == m {
				if len(dataMatrix) > row+3 && len(dataMatrix[row]) > col+3 && dataMatrix[row+3][col+3] == x {
					return true
				}
			}
		}
	}
	return false
}

func lookDiagonalLeft(row int, col int) bool {
	switch dataMatrix[row][col] {
	case x:
		if len(dataMatrix) > row+1 && 0 <= col-1 && dataMatrix[row+1][col-1] == m {
			if len(dataMatrix) > row+2 && 0 <= col-2 && dataMatrix[row+2][col-2] == a {
				if len(dataMatrix) > row+3 && 0 <= col-3 && dataMatrix[row+3][col-3] == s {
					return true
				}
			}
		}
	case s:
		if len(dataMatrix) > row+1 && 0 <= col-1 && dataMatrix[row+1][col-1] == a {
			if len(dataMatrix) > row+2 && 0 <= col-2 && dataMatrix[row+2][col-2] == m {
				if len(dataMatrix) > row+3 && 0 <= col-3 && dataMatrix[row+3][col-3] == x {
					return true
				}
			}
		}
	}
	return false
}

func main() {
	start := time.Now()
	readFromFile()
	resultFirst := first()
	fmt.Println(resultFirst)
	resultSecond := second()
	fmt.Println(resultSecond)
	timeElapsed := time.Since(start)
	fmt.Println("execution time:", timeElapsed)
}
