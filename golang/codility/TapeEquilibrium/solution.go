package TapeEquilibrium

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func Solution(A []int) int {
	sum := 0
	for _, n := range A {
		sum += n
	}

	var leftSum, rightSum, minDiff int

	leftSum = A[0]
	rightSum = sum - A[0]
	minDiff = abs(leftSum - rightSum)
	for _, n := range A[1 : len(A)-1] {
		leftSum += n
		rightSum -= n
		diff := abs(leftSum - rightSum)
		if minDiff > diff {
			minDiff = diff
		}
	}
	return minDiff
}
