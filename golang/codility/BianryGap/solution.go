package BianryGap

func toBin(n int) string {
	var binary string

	for n >= 2 {
		var bit string
		if float64(n/2) == float64(n)/2 {
			bit = "0"
		} else {
			bit = "1"
		}
		binary = bit + binary
		n /= 2
	}
	if n == 1 {
		binary = "1" + binary
	}
	return binary
}

func Solution(N int) int {
	binN := toBin(N)
	longestGap := 0
	gap := 0
	for _, bit := range binN {
		if string(bit) == "1" {
			if longestGap < gap {
				longestGap = gap
			}
			gap = 0
		} else {
			gap++
		}
	}
	return longestGap
}
