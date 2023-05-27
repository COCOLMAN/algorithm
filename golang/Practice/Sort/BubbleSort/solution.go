package BubbleSort

func bubbleSort(numbers []int) {
	for i := len(numbers) - 1; i >= 1; i-- {
		for j := 0; j < i; j++ {
			if numbers[j] > numbers[j+1] {
				tmp := numbers[j+1]
				numbers[j+1] = numbers[j]
				numbers[j] = tmp
			}
		}
	}
}
