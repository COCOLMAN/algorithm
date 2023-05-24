package PartitioningIntoMinimumNumberOfDeciBinaryNumbers

func minPartitions(n string) int {
	maxInt := 1
	for _, digit := range n {
		if maxInt < (int(digit) - 48) {
			maxInt = int(digit) - 48
		}
	}
	return maxInt
}
