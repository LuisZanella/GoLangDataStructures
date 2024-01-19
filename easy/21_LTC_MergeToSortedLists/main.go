package main

import "fmt"

/*
*
You are given the heads of two sorted linked lists list1 and list2.

Merge the two lists into one sorted list. The list should be made by splicing together the nodes of the first two lists.

Return the head of the merged linked list.

	Example 1:


	Input: list1 = [1,2,4], list2 = [1,3,4]
	Output: [1,1,2,3,4,4]
	Example 2:

	Input: list1 = [], list2 = []
	Output: []
	Example 3:

	Input: list1 = [], list2 = [0]
	Output: [0]

Constraints:

The number of nodes in both lists is in the range [0, 50].
-100 <= Node.val <= 100
Both list1 and list2 are sorted in non-decreasing order.
*/
type ListNode struct {
	Val  int
	Next *ListNode
}

func addValue(list *ListNode, value int) *ListNode {
	newNode := &ListNode{Val: value}
	if list == nil {
		return newNode
	}
	helper := list.Next
	list.Next.Val = value
	list.Next.Next = helper
	return list
}

func mergeSortedLinkedList(firstList *ListNode, secondList *ListNode) *ListNode {
	helper := firstList
	for helper != nil || secondList != nil {
		if secondList != nil && (helper == nil || secondList.Val > helper.Val) {
			firstList = addValue(firstList, secondList.Val)
			secondList = secondList.Next
		} else {
			helper = helper.Next
		}
	}
	return firstList
}
func main() {
	l1_0 := &ListNode{Val: 10, Next: nil}
	l1_1 := &ListNode{Val: 2, Next: l1_0}
	l1_2 := &ListNode{Val: 3, Next: l1_1}
	l1_3 := &ListNode{Val: 1, Next: l1_2}
	l1_4 := &ListNode{Val: 1, Next: l1_3}
	l1_5 := &ListNode{Val: 1, Next: l1_4}

	l2_0 := &ListNode{Val: 10, Next: nil}
	l2_1 := &ListNode{Val: 2, Next: l2_0}
	l2_2 := &ListNode{Val: 1, Next: l2_1}
	l2_3 := &ListNode{Val: 1, Next: l2_2}
	l2_4 := &ListNode{Val: 1, Next: l2_3}
	l2_5 := &ListNode{Val: 1, Next: l2_4}

	fmt.Println(mergeSortedLinkedList(l1_5, l2_5))
	l1_5 = nil
	l2_5 = nil
	fmt.Println(mergeSortedLinkedList(l1_5, l2_5))
	l2_5 = &ListNode{Val: 0}
	fmt.Println(mergeSortedLinkedList(l1_5, l2_5))
}
