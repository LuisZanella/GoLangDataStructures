package main

import (
	"fmt"
)

type BST struct {
	Value int

	Left  *BST
	Right *BST
}

func main() {
	tree := &BST{
		Value: 10,
		Left: &BST{
			Value: 5,
			Left: &BST{
				Value: 2,
				Left: &BST{
					Value: 1,
					Left:  nil,
					Right: nil,
				},
				Right: nil,
			},
			Right: &BST{
				Value: 5,
				Left:  nil,
				Right: nil,
			},
		},
		Right: &BST{
			Value: 15,
			Left: &BST{
				Value: 13,
				Left:  nil,
				Right: &BST{
					Value: 14,
					Left:  nil,
					Right: nil,
				},
			},
			Right: &BST{
				Value: 22,
				Left:  nil,
				Right: nil,
			},
		},
	}
	result := tree.FindClosestValue(12)
	fmt.Printf("%v", result)
}

func (tree *BST) FindClosestValue(target int) int {
	closest := tree.Value
	for tree != nil {
		if Abs(closest-target) > Abs(target-tree.Value) {
			closest = tree.Value
		}
		if target < tree.Value {
			tree = tree.Left
		} else {
			if target > tree.Value {
				tree = tree.Right
			} else {
				break
			}
		}
	}
	return closest
}

func Abs(number int) int {
	if number < 0 {
		return -number
	}
	return number
}
