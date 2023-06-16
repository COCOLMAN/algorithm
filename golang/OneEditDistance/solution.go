package OneEditDistance

import (
	"math"
)

func isOneEditDistance(s string, t string) bool {
	i := 0
	for i < len(s) && i < len(t) {
		if s[i] != t[i] {
			return s[i+1:] == t[i+1:] || s[i:] == t[i+1:] || s[i+1:] == t[i:]
		}
		i++
	}

	if math.Abs(float64(len(s)-len(t))) == 1 {
		return true
	}
	return false
}
