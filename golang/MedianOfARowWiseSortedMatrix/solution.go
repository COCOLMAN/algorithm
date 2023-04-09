package MedianOfARowWiseSortedMatrix

func matrixMedian(grid [][]int) int {
	totalSize := len(grid) * len(grid[0])
	wantIdx := totalSize / 2
	current := 0
	var currentNum int
	rowCurrentPositionList := make([]int, len(grid))
	for true {
		var rowIdx, currentPosition int

		var minNumAndRow []int

		for rowIdx, currentPosition = range rowCurrentPositionList {

			row := grid[rowIdx]

			if len(row) > currentPosition {
				if len(minNumAndRow) == 0 {
					minNumAndRow = append(minNumAndRow, row[currentPosition])
					minNumAndRow = append(minNumAndRow, rowIdx)
				} else if minNumAndRow[0] > row[currentPosition] {
					minNumAndRow[0] = row[currentPosition]
					minNumAndRow[1] = rowIdx
				}
			}
		}

		rowCurrentPositionList[minNumAndRow[1]]++
		current++

		if current-1 == wantIdx {
			//fmt.Println(rowCurrentPositionList, rowIdx)
			return minNumAndRow[0]
		}
	}
	return currentNum
}
