package PerformStringShifts

func stringShift(s string, shift [][]int) string {
	shiftSize := 0
	for _, move := range shift {
		if move[0] == 0 {
			shiftSize -= move[1]
		} else {
			shiftSize += move[1]
		}
	}

	shiftSize %= len(s)

	if shiftSize > 0 {
		return s[len(s)-shiftSize:] + s[:len(s)-shiftSize]
	}
	shiftSize = -shiftSize

	return s[shiftSize:] + s[:shiftSize]
}
