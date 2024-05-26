package main

import "fmt"

type BinaryTree struct {
	Value int
	Left  *BinaryTree
	Right *BinaryTree
}

func main() {
	tree := &BinaryTree{
		Value: 10,
		Left: &BinaryTree{
			Value: 5,
			Left: &BinaryTree{
				Value: 2,
				Left: &BinaryTree{
					Value: 1,
					Left:  nil,
					Right: nil,
				},
				Right: nil,
			},
			Right: &BinaryTree{
				Value: 5,
				Left:  nil,
				Right: nil,
			},
		},
		Right: &BinaryTree{
			Value: 15,
			Left: &BinaryTree{
				Value: 13,
				Left:  nil,
				Right: &BinaryTree{
					Value: 14,
					Left:  nil,
					Right: nil,
				},
			},
			Right: &BinaryTree{
				Value: 22,
				Left:  nil,
				Right: nil,
			},
		},
	}
	depths := calculateDepth(tree, 0)
	fmt.Printf("%v", depths)
}

func calculateDepth(node *BinaryTree, depth int) int {
	if node == nil {
		return 0
	}
	return depth + calculateDepth(node.Left, depth+1) + calculateDepth(node.Right, depth+1)
}
