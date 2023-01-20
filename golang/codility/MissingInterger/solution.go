package MissingInterger

func Solution(A []int) int {
	hash := map[int]bool{}

	for _, a := range A {
		if _, ok := hash[a]; !ok {
			hash[a] = true
		}
	}

	for n := 1; n < len(A)+1; n++ {
		exist := hash[n]
		if !exist {
			return n
		}
	}
	return len(A) + 1
}
