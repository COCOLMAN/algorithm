package main

import (
	"fmt"
)

type BinaryTreeNode struct {
	Val   int
	Right *BinaryTreeNode
	Left  *BinaryTreeNode
}

func PreOrderTraverse(node *BinaryTreeNode) {
	if node == nil {
		return
	}
	fmt.Printf("%d -> ", node.Val)
	PreOrderTraverse(node.Left)
	PreOrderTraverse(node.Right)
}

func InOrderTraverse(node *BinaryTreeNode) {
	if node == nil {
		return
	}
	InOrderTraverse(node.Left)
	fmt.Printf("%d -> ", node.Val)
	InOrderTraverse(node.Right)
}

func PostOrderTraverse(node *BinaryTreeNode) {
	if node == nil {
		return
	}
	PostOrderTraverse(node.Left)
	PostOrderTraverse(node.Right)
	fmt.Printf("%d -> ", node.Val)
}

func main() {
	node := &BinaryTreeNode{
		Val: 0,
		Left: &BinaryTreeNode{
			Val: 1,
			Left: &BinaryTreeNode{
				Val: 3,
				Left: &BinaryTreeNode{
					Val: 7,
				},
				Right: &BinaryTreeNode{
					Val: 8,
				},
			},
			Right: &BinaryTreeNode{
				Val: 4,
				Left: &BinaryTreeNode{
					Val: 9,
				},
				Right: &BinaryTreeNode{
					Val: 10,
				},
			},
		},
		Right: &BinaryTreeNode{
			Val: 2,
			Left: &BinaryTreeNode{
				Val: 5,
				Left: &BinaryTreeNode{
					Val: 11,
				},
			},
			Right: &BinaryTreeNode{
				Val: 6,
			},
		},
	}

	PreOrderTraverse(node)
	fmt.Println()
	InOrderTraverse(node)
	fmt.Println()
	PostOrderTraverse(node)
}
