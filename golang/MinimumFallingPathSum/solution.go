package MinimumFallingPathSum

func min(numbers []int) int {
	minN := numbers[0]
	for _, n := range numbers[1:] {
		if minN > n {
			minN = n
		}
	}
	return minN
}

func minFallingPathSum(matrix [][]int) int {
	for i := 0; i < len(matrix)-1; i++ {
		height := len(matrix) - 2 - i

		for j := 0; j < len(matrix[height]); j++ {
			var bottoms []int
			bottoms = append(bottoms, matrix[height+1][j])
			if j > 0 {
				bottoms = append(bottoms, matrix[height+1][j-1])
			}
			if j < len(matrix[height])-1 {
				bottoms = append(bottoms, matrix[height+1][j+1])
			}
			matrix[height][j] += min(bottoms)
		}
	}
	return min(matrix[0])
}
