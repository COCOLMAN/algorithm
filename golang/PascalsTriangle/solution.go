package PascalsTriangle

func generate(numRows int) [][]int {
	if numRows == 1 {
		return [][]int{{1}}
	}
	if numRows == 2 {
		return [][]int{{1}, {1, 1}}
	}
	beforeTriangle := generate(numRows - 1)
	beforeRow := beforeTriangle[len(beforeTriangle)-1]
	var lastRow []int

	lastRow = append(lastRow, 1)
	for i := 1; i < numRows-1; i++ {
		item := beforeRow[i-1] + beforeRow[i]
		lastRow = append(lastRow, item)
	}
	lastRow = append(lastRow, 1)
	beforeTriangle = append(beforeTriangle, lastRow)
	return beforeTriangle
}
