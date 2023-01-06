package main

import "fmt"

type LinkedList struct {
	Val  int
	Next *LinkedList
}

func PreOrderTraverse(node *LinkedList) {
	if node == nil {
		return
	}
	fmt.Printf("%d -> ", node.Val)
	PreOrderTraverse(node.Next)
}

func PostOrderTraverse(node *LinkedList) {
	if node == nil {
		return
	}
	PostOrderTraverse(node.Next)
	fmt.Printf("%d -> ", node.Val)

}

func main() {
	list := &LinkedList{
		Val: 1,
		Next: &LinkedList{
			Val: 2,
			Next: &LinkedList{
				Val: 3,
				Next: &LinkedList{
					Val: 4,
					Next: &LinkedList{
						Val: 5,
					},
				},
			},
		},
	}

	PreOrderTraverse(list)
	fmt.Println()
	PostOrderTraverse(list)
}
