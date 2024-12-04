package main

func first() int {

	foundXmas := 0

	for row, dataArray := range dataMatrix {
		for col := range dataArray {
			result := lookHorizontal(row, col)
			if result {
				foundXmas++
			}
			resultVertical := lookVertical(row, col)
			if resultVertical {
				foundXmas++
			}

			resultDiagonal := lookDiagonal(row, col)
			if resultDiagonal {
				foundXmas++
			}

			resultDiagonalLeft := lookDiagonalLeft(row, col)
			if resultDiagonalLeft {
				foundXmas++
			}
		}
	}
	return foundXmas
}
