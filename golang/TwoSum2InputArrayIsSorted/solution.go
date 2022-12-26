package TwoSum2InputArrayIsSorted

func twoSum(numbers []int, target int) []int {
	for idx, num := range numbers {
		want := target - num

		start, end := idx+1, len(numbers)-1
		for start <= end {
			mid := (start + end) / 2
			if numbers[mid] == want {
				return []int{idx + 1, mid + 1}
			} else if numbers[mid] < want {
				start = mid + 1
			} else {
				end = mid - 1
			}
		}
	}
	return []int{}
}
