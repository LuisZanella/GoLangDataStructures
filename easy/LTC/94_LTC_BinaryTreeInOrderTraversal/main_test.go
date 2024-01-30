package main

import (
	"reflect"
	"testing"
)

func TestInorderTraversal(t *testing.T) {
	var root *TreeNode
	var want []int
	t.Run("2 layers", func(t *testing.T) {
		root = &TreeNode{2, &TreeNode{Val: 1}, &TreeNode{Val: 3}}
		want = []int{1, 2, 3}
		got := inorderTraversal(root)
		assert(t, got, want)
	})
	t.Run("2 layers", func(t *testing.T) {
		root = &TreeNode{1, nil, &TreeNode{Val: 2, Left: &TreeNode{Val: 3}}}
		want = []int{1, 3, 2}
		got := inorderTraversal(root)
		assert(t, got, want)
	})
	t.Run("1 layer", func(t *testing.T) {
		root = &TreeNode{Val: 1}
		want = []int{1}
		got := inorderTraversal(root)
		assert(t, got, want)
	})
	t.Run("nil", func(t *testing.T) {
		root = nil
		want = nil

		got := inorderTraversal(root)
		assert(t, got, want)
	})
	t.Run("0 layers", func(t *testing.T) {
		root = &TreeNode{}
		want = []int{0}

		got := inorderTraversal(root)
		assert(t, got, want)
	})
}

func TestInorderTraversal2(t *testing.T) {
	var root *TreeNode
	var want []int
	t.Run("2 layers", func(t *testing.T) {
		root = &TreeNode{2, &TreeNode{Val: 1}, &TreeNode{Val: 3}}
		want = []int{1, 2, 3}
		got := inorderTraversal2(root)
		assert(t, got, want)
	})
	t.Run("2 layers", func(t *testing.T) {
		root = &TreeNode{1, nil, &TreeNode{Val: 2, Left: &TreeNode{Val: 3}}}
		want = []int{1, 3, 2}
		got := inorderTraversal2(root)
		assert(t, got, want)
	})
	t.Run("1 layer", func(t *testing.T) {
		root = &TreeNode{Val: 1}
		want = []int{1}
		got := inorderTraversal2(root)
		assert(t, got, want)
	})
	t.Run("nil", func(t *testing.T) {
		root = nil
		want = []int{}

		got := inorderTraversal2(root)
		assert(t, got, want)
	})
	t.Run("0 layers", func(t *testing.T) {
		root = &TreeNode{}
		want = []int{0}

		got := inorderTraversal2(root)
		assert(t, got, want)
	})
}

func assert(t *testing.T, got, wanted []int) {
	t.Helper()
	if !reflect.DeepEqual(got, wanted) {
		t.Errorf("Got: %v Wanted: %v", got, wanted)
	}
}
