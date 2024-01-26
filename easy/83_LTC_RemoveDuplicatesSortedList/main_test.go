package main

import (
	"fmt"
	"testing"
)

func TestDeleteDuplicates(t *testing.T) {
	var list ListNode
	t.Run("nil", func(t *testing.T) {
		got := deleteDuplicates(nil)
		printAssert(t, got, nil)
	})
	t.Run("1,1,2", func(t *testing.T) {
		list = ListNode{1, &ListNode{1, &ListNode{2, nil}}}
		got := deleteDuplicates(&list)
		want := &ListNode{1, &ListNode{2, nil}}
		printAssert(t, got, want)
	})
	t.Run("1,2", func(t *testing.T) {
		list = ListNode{1, &ListNode{2, nil}}
		got := deleteDuplicates(&list)
		want := &ListNode{1, &ListNode{2, nil}}
		printAssert(t, got, want)
	})
	t.Run("1,1,2,3,3", func(t *testing.T) {
		list = ListNode{1, &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{3, nil}}}}}
		got := deleteDuplicates(&list)
		want := &ListNode{1, &ListNode{2, &ListNode{3, nil}}}
		printAssert(t, got, want)
	})

}

func printAssert(t *testing.T, got *ListNode, want *ListNode) {
	if fmt.Sprintf("%T", got) != fmt.Sprintf("%T", want) {
		t.Errorf("Got: %v Wanted: %v", got, want)
	}
}
