package PlusOne

func reverse(a []int) []int {
	idx := 0
	for idx < len(a)/2 {
		temp := a[idx]
		a[idx] = a[len(a)-idx-1]
		a[len(a)-idx-1] = temp
		idx++
	}
	return a
}

func plusOne(digits []int) []int {
	answer := []int{}
	additional := true

	idx := len(digits) - 1
	for idx >= 0 {
		n := digits[idx]
		if additional {
			n++
		}
		additional = false
		if n >= 10 {
			n -= 10
			additional = true
		}
		answer = append(answer, n)
		idx--
	}
	if additional {
		answer = append(answer, 1)
	}
	return reverse(answer)
}
