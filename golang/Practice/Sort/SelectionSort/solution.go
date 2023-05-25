package SelectionSort

func selectSort(numbers []int) {
	for currentIdx, currentNum := range numbers {
		var smallestNumIdx int
		var smallestNum int = currentNum
		for targetIdx, targetNum := range numbers[currentIdx:] {
			if targetNum < smallestNum {
				smallestNumIdx = targetIdx
				smallestNum = targetNum
			}
		}
		numbers[currentIdx] = smallestNum
		numbers[smallestNumIdx] = currentNum
	}
}
