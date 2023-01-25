package OddOccurrencesInArray

func Solution(A []int) int {
	odds := map[int]bool{}

	for _, n := range A {
		if _, ok := odds[n]; ok {
			delete(odds, n)
		} else {
			odds[n] = true
		}
	}
	for key := range odds {
		return key
	}
	return 0
}
