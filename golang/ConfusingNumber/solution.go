package ConfusingNumber

func confusingNumber(n int) bool {
	confusingNumbers := map[int]int{
		0: 0,
		1: 1,
		6: 9,
		8: 8,
		9: 6,
	}

	mo := 1
	base := 10
	numbers := []int{}
	originN := n
	for n != 0 {
		diff := n % base
		target := diff / mo

		if val, exist := confusingNumbers[target]; !exist {
			return false
		} else {
			numbers = append(numbers, val)
		}
		n -= diff
		base *= 10
		mo *= 10
	}

	rotateNumber := 0
	i := len(numbers) - 1
	base = 1
	for i >= 0 {
		rotateNumber += (numbers[i] * base)
		base *= 10
		i--
	}
	return rotateNumber != originN
}
