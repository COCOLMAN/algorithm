package SpiralMatrix

func right(x, y int) (int, int) {
	return x, y + 1
}

func left(x, y int) (int, int) {
	return x, y - 1
}

func up(x, y int) (int, int) {
	return x - 1, y
}

func down(x, y int) (int, int) {
	return x + 1, y
}

func spiralOrder(matrix [][]int) []int {
	totalLen := len(matrix) * len(matrix[0])
	result := []int{}
	var x, y int
	ban := 9999

	operators := []func(int, int) (int, int){right, down, left, up}
	fIndex := 0
	nextPoint := operators[fIndex]
	for true {
		num := matrix[x][y]
		result = append(result, num)
		matrix[x][y] = ban
		if len(result) >= totalLen {
			break
		}

	getNewPoint:
		newX, newY := nextPoint(x, y)
		if newX < 0 || newX >= len(matrix) || newY < 0 || newY >= len(matrix[0]) || matrix[newX][newY] == ban {

			if fIndex == 3 {
				fIndex = 0
				nextPoint = right
			} else {
				fIndex++
				nextPoint = operators[fIndex]
			}
			goto getNewPoint
		}
		x, y = newX, newY
	}
	return result
}
