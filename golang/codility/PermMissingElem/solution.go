package PermMissingElem

func Solution(A []int) int {
	answer := 1
	if len(A) == 0 {
		return answer
	}
	numberMap := map[int]bool{}
	maxNum := A[0]
	for _, number := range A {
		numberMap[number] = true
		if maxNum < number {
			maxNum = number
		}
	}

	for answer <= maxNum {
		if _, exist := numberMap[answer]; !exist {
			break
		}
		answer++
	}
	return answer
}
