package CyclicRotation

func Solution(A []int, K int) []int {
	getNewIdx := func(idx int) int {
		newIdx := K + idx
		for newIdx >= len(A) {
			newIdx -= len(A)
		}
		return newIdx
	}

	newA := make([]int, len(A))
	for idx, n := range A {
		newIdx := getNewIdx(idx)
		newA[newIdx] = n
	}
	return newA
}
