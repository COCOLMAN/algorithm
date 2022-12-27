package SearchA2DMatrix

func searchMatrix(matrix [][]int, target int) bool {
	rowIdxStart, rowIdxEnd := 0, len(matrix)-1
	var rowIdx int
	for rowIdxStart <= rowIdxEnd {
		rowIdxMid := (rowIdxStart + rowIdxEnd) / 2
		if matrix[rowIdxMid][0] <= target && target <= matrix[rowIdxMid][len(matrix[rowIdxMid])-1] {
			rowIdx = rowIdxMid
			break
		} else {
			if target < matrix[rowIdxMid][0] {
				rowIdxEnd = rowIdxMid - 1
			} else if target > matrix[rowIdxMid][len(matrix[rowIdxMid])-1] {
				rowIdxStart = rowIdxMid + 1
			}
		}
	}
	var colIdx int
	colIdxStart, colIdxEnd := 0, len(matrix[rowIdx])-2
	for colIdxStart <= colIdxEnd {
		colIdxMid := (colIdxStart + colIdxEnd) / 2
		if matrix[rowIdx][colIdxMid] <= target && target < matrix[rowIdx][colIdxMid+1] {
			colIdx = colIdxMid
			break
		}
		if target == matrix[rowIdx][colIdxMid+1] {
			colIdx = colIdxMid + 1
			break
		}
		if target < matrix[rowIdx][colIdxMid] {
			colIdxEnd = colIdxMid - 1
		} else if target > matrix[rowIdx][colIdxMid+1] {
			colIdxStart = colIdxMid + 1
		}

	}

	return matrix[rowIdx][colIdx] == target
}
