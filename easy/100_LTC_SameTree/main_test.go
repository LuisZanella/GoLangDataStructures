package main

import "testing"

func TestSameTree(t *testing.T) {
	var root1 *TreeNode
	var root2 *TreeNode
	t.Run("same", func(t *testing.T) {
		root1 = &TreeNode{1, &TreeNode{Val: 2}, &TreeNode{Val: 3}}
		root2 = &TreeNode{1, &TreeNode{Val: 2}, &TreeNode{Val: 3}}
		want := true
		got := isSameTree(root1, root2)
		assert(t, got, want)
	})
	t.Run("empty", func(t *testing.T) {
		root1 = &TreeNode{}
		root2 = &TreeNode{}
		want := true
		got := isSameTree(root1, root2)
		assert(t, got, want)
	})
	t.Run("different", func(t *testing.T) {
		root1 = &TreeNode{1, &TreeNode{Val: 2}, &TreeNode{Val: 3}}
		root2 = &TreeNode{1, &TreeNode{Val: 3}, &TreeNode{Val: 3}}
		want := false
		got := isSameTree(root1, root2)
		assert(t, got, want)

	})
	t.Run("one empty", func(t *testing.T) {
		root1 = &TreeNode{1, &TreeNode{Val: 2}, &TreeNode{Val: 3}}
		root2 = &TreeNode{}
		want := false
		got := isSameTree(root1, root2)
		assert(t, got, want)

	})
}

func assert(t *testing.T, got, want bool) {
	t.Helper()
	if got != want {
		t.Errorf("Got: %t Want: %t", got, want)
	}
}
