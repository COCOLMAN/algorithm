package PassingCars

func Solution(A []int) int {
	passingCount := 0
	eastCount := 0
	for _, a := range A {
		if a == 0 {
			eastCount++
		} else {
			passingCount += eastCount
			if passingCount > 1000000000 {
				return -1
			}
		}
	}
	return passingCount
}
