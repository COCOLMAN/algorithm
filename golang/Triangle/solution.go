package Triangle

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func minimumTotal(triangle [][]int) int {
	height := len(triangle) - 1
	for i := 1; i < len(triangle); i++ {
		floor := height - i
		row := triangle[floor]
		for j := 0; j < len(row); j++ {
			left := triangle[floor+1][j]
			right := triangle[floor+1][j+1]
			row[j] = row[j] + min(left, right)
		}
	}
	return triangle[0][0]
}
