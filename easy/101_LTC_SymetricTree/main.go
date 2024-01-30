package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Node struct {
	data *TreeNode
	next *Node
}
type Queue struct {
	front *Node
	last  *Node
}

// *
//
//	Recursive way
//
// /**
// * Definition for a binary tree node.
// * type TreeNode struct {
// *     Val int
// *     Left *TreeNode
// *     Right *TreeNode
// * }
// */
//
//	func isSymmetric(root *TreeNode) bool {
//	   if root == nil {
//	       return true
//	   }
//	   return isMirror(root.Left, root.Right)
//	}
//
//	func isMirror(a, b *TreeNode) bool {
//	   // If both nodes are nil, they are mirrors of each other
//	   if a == nil && b == nil {
//	       return true
//	   }
//	   // If one of the nodes is nil, or their values differ, they aren't mirrors
//	   if a == nil || b == nil || a.Val != b.Val {
//	       return false
//	   }
//	   // Check if the left subtree of 'a' is a mirror of the right subtree of 'b' and vice versa
//	   return isMirror(a.Left, b.Right) && isMirror(a.Right, b.Left)
//	}
//
// */
func (q *Queue) add(node *TreeNode) {
	newNode := &Node{data: node}
	if q.last == nil {
		q.front = newNode
		q.last = newNode
	} else {
		q.last.next = newNode
		q.last = newNode
	}
}
func (q *Queue) remove() *TreeNode {
	if q.front == nil {
		return nil
	}
	helper := q.front.data
	q.front = q.front.next
	if q.front == nil {
		q.last = nil
	}
	return helper
}
func (q *Queue) isEmpty() bool {
	return q.front == nil
}

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	queue := &Queue{}
	queue.add(root.Left)
	queue.add(root.Right)
	for !queue.isEmpty() {
		l := queue.remove()
		r := queue.remove()
		if l == nil && r == nil {
			continue
		}
		if l == nil || r == nil || l.Val != r.Val {
			return false
		}

		queue.add(l.Left)
		queue.add(r.Right)
		queue.add(l.Right)
		queue.add(r.Left)
	}
	return true
}
