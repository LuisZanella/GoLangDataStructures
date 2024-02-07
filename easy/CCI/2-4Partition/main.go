package main

type LinkedList struct {
	data int
	next *LinkedList
}

func partition(list *LinkedList, target int) *LinkedList {
	var head *LinkedList
	tail := list
	for list != nil {
		next := list.next
		if list.data < target {
			list.next = head
			head = list
		} else {
			tail.next = list
			tail = list
		}
		list = next
	}
	tail.next = nil
	return head
}
