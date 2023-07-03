package DesignTicTacToe

import (
	"fmt"
	"testing"
)

func TestGame1(t *testing.T) {
	game := Constructor(3)
	fmt.Println(game.Move(0, 0, 1))
	fmt.Println(game.Move(0, 2, 2))
	fmt.Println(game.Move(2, 2, 1))
	fmt.Println(game.Move(1, 1, 2))
	fmt.Println(game.Move(2, 0, 1))
	fmt.Println(game.Move(1, 0, 2))
	fmt.Println(game.Move(2, 1, 1))
}

func TestGame2(t *testing.T) {
	game := Constructor(2)
	fmt.Println(game.Move(0, 0, 2))
	fmt.Println(game.Move(1, 1, 1))
	fmt.Println(game.Move(0, 1, 2))
}
