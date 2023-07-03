package ArmstrongNumber

import (
	"math"
	"strconv"
)

func isArmstrong(n int) bool {
	strNs := strconv.Itoa(n)
	k := len(strNs)
	result := 0
	for _, strN := range strNs {
		val, _ := strconv.Atoi(string(strN))
		result += int(math.Pow(float64(val), float64(k)))
	}
	return result == n
}
