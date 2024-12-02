package main

func remove(slice []int, s int) []int {
	newSlice := make([]int, len(slice))
	copy(newSlice, slice)
	return append(newSlice[:s], newSlice[s+1:]...)
}

func isArrayAscending(arr []int) bool {
	sum := 0
	for index, element := range arr {
		if index+1 < len(arr) {
			if element-arr[index+1] < 0 {
				sum++
			} else if element-arr[index+1] > 0 {
				sum--
			}
		}
	}
	return sum > 0

}

func checkIfReportsAreSafe(arr []int) int {
	ascending := isArrayAscending(arr)
	for index, element := range arr {
		if index+1 < len(arr) {
			if ascending {
				if !(arr[index+1]-element >= 1 && arr[index+1]-element <= 3) {
					return index
				}
			}
			if !ascending {
				if !(element-arr[index+1] >= 1 && element-arr[index+1] <= 3) {
					return index
				}
			}
		}
	}
	return -1
}

func second() int {
	reports := readFromFile()

	safeReports := 0
	i := 0
	for _, arr := range reports {

		isFail := false

		indexFromArray := checkIfReportsAreSafe(arr)
		if indexFromArray != -1 {
			firstPossibleArray := remove(arr, indexFromArray)
			secondPossbileArray := remove(arr, indexFromArray+1)
			secondTryFirstPossbileArray := checkIfReportsAreSafe(firstPossibleArray)
			secondTrySecondPossbileArray := checkIfReportsAreSafe(secondPossbileArray)
			if secondTryFirstPossbileArray != -1 && secondTrySecondPossbileArray != -1 {
				isFail = true
			}
		}

		if !isFail {
			safeReports++
		}
		i++
	}
	return safeReports
}
