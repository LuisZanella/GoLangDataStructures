package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	cleanedList := head
	for cleanedList != nil && cleanedList.Next != nil {
		if cleanedList.Val != cleanedList.Next.Val {
			cleanedList = cleanedList.Next
		} else {
			cleanedList.Next = cleanedList.Next.Next
		}
	}
	return head
}
