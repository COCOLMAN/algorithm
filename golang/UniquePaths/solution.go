package UniquePaths

func uniquePaths(m int, n int) int {
	path := make([][]int, m)
	for i := 0; i < m; i++ {
		row := make([]int, n)
		for i := range row {
			row[i] = 1
		}
		path[i] = row
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			path[i][j] = path[i-1][j] + path[i][j-1]
		}
	}
	return path[m-1][n-1]
}
