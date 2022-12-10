package Sqrt

import (
	"math"
)

func mySqrt(x int) int {
	if x < 2 {
		return x
	}
	maxN := x
	minN := 0
	var n int
	for true {
		n = (maxN + minN) / 2
		a := int(math.Pow(float64(n), 2))
		b := int(math.Pow(float64(n+1), 2))
		if a <= (x) && b > (x) {
			break
		}
		if a > (x) {
			maxN = n
		} else {

			minN = n
		}
	}
	return n
}
