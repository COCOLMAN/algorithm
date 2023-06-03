package ConfusingNumber

func confusingNumber(n int) bool {
	confusingNumbers := map[int]bool{
		0: true,
		1: true,
		6: true,
		8: true,
		9: true,
	}

	isAllOne := true
	mo := 1
	base := 10
	for n != 0 {
		diff := n % base
		target := diff / mo
		if _, exist := confusingNumbers[target]; !exist {
			return false
		} else {
			if target != 1 {
				isAllOne = false
			}
		}

		n -= diff
		base *= 10
		mo *= 10
	}
	return !isAllOne
}
