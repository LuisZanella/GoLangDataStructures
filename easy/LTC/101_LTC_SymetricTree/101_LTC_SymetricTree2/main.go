package main

import "container/list"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type LinkedList struct {
	Node *TreeNode
	Next *LinkedList
}

func (l *LinkedList) append(node *TreeNode) {
	newList := LinkedList{Node: node}
	if l.Node == nil {
		l.Node = node
		return
	}
	newList.Next = l
	*l = newList
}
func (l *LinkedList) remove() *TreeNode {
	if l == nil {
		return nil
	}
	helper := l.Node
	l = l.Next
	return helper
}
func isSymmetric(root *TreeNode) bool {
	list.New()
	newList := &LinkedList{}
	newList.append(root.Left)
	newList.append(root.Right)
	for newList != nil {
		l := newList.remove()
		r := newList.remove()
		if l == nil && r == nil {
			continue
		}
		if l == nil || r == nil || r.Val != l.Val {
			return false
		}
		newList.append(l.Left)
		newList.append(r.Right)
		newList.append(l.Right)
		newList.append(r.Left)
	}
	return true
}
