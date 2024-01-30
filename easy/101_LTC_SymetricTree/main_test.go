package main

import "testing"

func TestIsSymetric(t *testing.T) {
	var tree *TreeNode
	var want bool
	t.Run("Not Symmetric", func(t *testing.T) {
		tree = &TreeNode{1, &TreeNode{Val: 2, Left: &TreeNode{Val: 3}}, &TreeNode{Val: 2, Left: &TreeNode{Val: 3, Left: &TreeNode{Val: 5}}}}
		want = false
		got := isSymmetric(tree)
		assert(t, got, want)
	})
	t.Run("Symmetric", func(t *testing.T) {
		tree = &TreeNode{1, &TreeNode{2, &TreeNode{Val: 3}, &TreeNode{Val: 4}}, &TreeNode{2, &TreeNode{Val: 4}, &TreeNode{Val: 3}}}
		want = true
		got := isSymmetric(tree)
		assert(t, got, want)
	})
	t.Run("Empty", func(t *testing.T) {
		tree = &TreeNode{}
		want = true
		got := isSymmetric(tree)
		assert(t, got, want)
	})
	t.Run("Not Center Symmetric", func(t *testing.T) {
		tree = &TreeNode{1, &TreeNode{Val: 2, Left: &TreeNode{Val: 3}}, &TreeNode{Val: 2, Left: &TreeNode{Val: 3}}}
		want = false
		got := isSymmetric(tree)
		assert(t, got, want)
	})
}

func assert(t *testing.T, got, want bool) {
	t.Helper()
	if got != want {
		t.Errorf("Got: %t Want: %t", got, want)
	}
}
