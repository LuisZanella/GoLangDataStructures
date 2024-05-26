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
	results := BranchSums(tree)
	fmt.Printf("%v", results)
}
func BranchSums(tree *BinaryTree) []int {
	var sums []int
	preOrderTransversal(tree, 0, &sums)
	return sums
}

func preOrderTransversal(tree *BinaryTree, sum int, elements *[]int) {
	if tree == nil {
		return
	}
	sum += tree.Value
	if tree.Left == nil && tree.Right == nil {
		*elements = append(*elements, sum)
		return
	}
	preOrderTransversal(tree.Left, sum, elements)
	preOrderTransversal(tree.Right, sum, elements)
}
