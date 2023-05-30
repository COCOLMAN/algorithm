package UniquePath2

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	for i := 0; i < len(obstacleGrid); i++ {
		for j := 0; j < len(obstacleGrid[0]); j++ {
			if obstacleGrid[i][j] == 1 {
				obstacleGrid[i][j] = -1
			}
		}
	}

	for i := 0; i < len(obstacleGrid); i++ {
		if obstacleGrid[i][0] == -1 {
			break
		}
		obstacleGrid[i][0] = 1
	}
	for i := 0; i < len(obstacleGrid[0]); i++ {
		if obstacleGrid[0][i] == -1 {
			break
		}
		obstacleGrid[0][i] = 1
	}

	for i := 1; i < len(obstacleGrid); i++ {
		for j := 1; j < len(obstacleGrid[0]); j++ {
			if obstacleGrid[i][j] == -1 {
				continue
			}
			var fromLeft, fromTop int
			if obstacleGrid[i-1][j] != -1 {
				fromLeft = obstacleGrid[i-1][j]
			}
			if obstacleGrid[i][j-1] != -1 {
				fromTop = obstacleGrid[i][j-1]
			}
			obstacleGrid[i][j] = fromLeft + fromTop
		}
	}

	if obstacleGrid[len(obstacleGrid)-1][len(obstacleGrid[0])-1] == -1 {
		return 0
	}
	return obstacleGrid[len(obstacleGrid)-1][len(obstacleGrid[0])-1]
}
