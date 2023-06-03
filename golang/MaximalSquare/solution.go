package MaximalSquare

func maximalSquare(matrix [][]byte) int {
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			cell := matrix[i][j]
			if cell != 0 {
				getRightDownSize(i, j, matrix)
				getLeftTopSize(i, j, matrix)
			}
		}
	}
	return 0
}

func getLeftTopSize(i int, j int, matrix [][]byte) {

}

func getRightDownSize(i int, j int, matrix [][]byte) {

}
