package MissingNumberInArithmeticProgression

func getDiff(a [2]int, b [2]int) int {
	if a[1]-a[0] > b[1]-b[0] {
		if a[1]-a[0] < 0 {
			return a[1] - a[0]
		}
		return b[1] - b[0]
	}
	if a[1]-a[0] < 0 {
		return b[1] - b[0]
	}
	return a[1] - a[0]
}

func missingNumber(arr []int) int {
	diff := getDiff([2]int{arr[0], arr[1]}, [2]int{arr[len(arr)-2], arr[len(arr)-1]})
	for idx, num := range arr {
		if idx != len(arr)-1 {
			if (arr[idx+1] - num) != diff {
				return num + diff
			}
		}
	}
	return arr[len(arr)-1] + diff
}
