package main

func lookDiagonalSecond(row int, col int) bool {
	switch dataMatrix[row][col] {
	case m:
		if len(dataMatrix) > row+1 && len(dataMatrix[row]) > col+1 && dataMatrix[row+1][col+1] == a {
			if len(dataMatrix) > row+2 && len(dataMatrix[row]) > col+2 && dataMatrix[row+2][col+2] == s {
				return true
			}
		}
	case s:
		if len(dataMatrix) > row+1 && len(dataMatrix[row]) > col+1 && dataMatrix[row+1][col+1] == a {
			if len(dataMatrix) > row+2 && len(dataMatrix[row]) > col+2 && dataMatrix[row+2][col+2] == m {
				return true
			}
		}
	}
	return false
}

func lookDiagonalLeftSecond(row int, col int) bool {
	switch dataMatrix[row][col] {
	case m:
		if len(dataMatrix) > row+1 && 0 <= col-1 && dataMatrix[row+1][col-1] == a {
			if len(dataMatrix) > row+2 && 0 <= col-2 && dataMatrix[row+2][col-2] == s {
				return true
			}
		}
	case s:
		if len(dataMatrix) > row+1 && 0 <= col-1 && dataMatrix[row+1][col-1] == a {
			if len(dataMatrix) > row+2 && 0 <= col-2 && dataMatrix[row+2][col-2] == m {
				return true
			}
		}
	}
	return false
}

func second() int {
	foundXmasInShapeOfX := 0

	for row, dataArray := range dataMatrix {
		for col := range dataArray {
			result := lookDiagonalSecond(row, col)
			if result && len(dataArray) > col+2 {
				secondResult := lookDiagonalLeftSecond(row, col+2)
				if secondResult {
					foundXmasInShapeOfX++
				}
			}
		}
	}

	return foundXmasInShapeOfX
}
