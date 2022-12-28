package SumOfSquareNumbers

import (
	"math"
)

func judgeSquareSum(c int) bool {
	n := 0
	for n <= (int(math.Sqrt(float64(c))))+2 {
		n2 := int(math.Pow(float64(n), 2))
		start, end := n, int(math.Sqrt(float64(c-n2)))+2
		for start <= end {
			mid := (start + end) / 2
			mid2 := int(math.Pow(float64(mid), 2))
			if n2+mid2 == c {
				return true
			}
			if n2+mid2 < c {
				start = mid + 1
			} else {
				end = mid - 1
			}
		}
		n++
	}

	return false
}
