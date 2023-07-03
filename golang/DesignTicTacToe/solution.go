package DesignTicTacToe

type playerStone struct {
	horizontal map[int]int
	vertical   map[int]int
	diagonal   map[int]int
	size       int
}

func (p *playerStone) place(x, y int) bool {
	p.horizontal[x]++
	if p.horizontal[x] == p.size {
		return true
	}
	p.vertical[y]++
	if p.vertical[y] == p.size {
		return true
	}
	if x == y {
		p.diagonal[0]++
		if p.diagonal[0] == p.size {
			return true
		}
		if x == (p.size / 2) {
			p.diagonal[1]++
			if p.diagonal[1] == p.size {
				return true
			}
		}
	} else if x == p.size-1-y {
		p.diagonal[1]++
		if p.diagonal[1] == p.size {
			return true
		}
	}
	return false
}

type TicTacToe struct {
	size int
	p1   playerStone
	p2   playerStone
}

func Constructor(n int) TicTacToe {
	return TicTacToe{
		size: n,
		p1: playerStone{
			horizontal: map[int]int{},
			vertical:   map[int]int{},
			diagonal:   map[int]int{},
			size:       n,
		},
		p2: playerStone{
			horizontal: map[int]int{},
			vertical:   map[int]int{},
			diagonal:   map[int]int{},
			size:       n,
		},
	}
}

func (this *TicTacToe) Move(row int, col int, player int) int {
	if player == 1 {
		if this.p1.place(row, col) {
			return player
		}
	}
	if player == 2 {
		if this.p2.place(row, col) {
			return player
		}
	}
	return 0
}
