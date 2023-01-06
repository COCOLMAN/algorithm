package FrogJmp

func Solution(X int, Y int, D int) int {
	d := Y - X
	jmp := d / D
	if float64(d)/float64(D) != float64(d/D) {
		jmp++
	}
	return jmp
}
