package Pow_x_n

func pow(x float64, n int) float64 {
	if n == 0 {
		return 1
	}
	if n == 1 {
		return x
	}
	if n == 2 {
		return x * x
	}
	if n%2 == 0 {
		v := pow(x, n/2)
		return v * v
	}
	v := pow(x, (n-1)/2)
	return v * v * x
}

func myPow(x float64, n int) float64 {
	var neg bool
	if n < 0 {
		neg = true
		n = -n
	}
	v := pow(x, n)
	if neg {
		return 1 / v
	}
	return v
}
