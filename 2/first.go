package main

func first() int {
	reports := readFromFile()

	safeReports := 0

	for _, arr := range reports {
		ascending := false
		if arr[0]-arr[1] < 0 {
			ascending = true
		}
		isFail := false
		for index, element := range arr {
			if index+1 < len(arr) {
				if ascending {
					if !(arr[index+1]-element >= 1 && arr[index+1]-element <= 3) {
						isFail = true
						break
					}
				}
				if !ascending {
					if !(element-arr[index+1] >= 1 && element-arr[index+1] <= 3) {
						isFail = true
						break
					}

				}
			}
		}
		if !isFail {
			safeReports++
		}
	}
	return safeReports
}
