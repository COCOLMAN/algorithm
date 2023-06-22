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

	start := 0
	end := len(arr) - 2
	for start <= end {

		target := (start + end) / 2
		if arr[target+1]-arr[target] != diff {
			return arr[target] + diff
		}
		if arr[target] != arr[0]+target*diff {
			end = target - 1
		} else {
			start = target + 1
		}
	}

	return arr[len(arr)-1] + diff
}
