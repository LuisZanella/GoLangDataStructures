package main

type LinkedList struct {
	data int
	next *LinkedList
}

//func partition(list *LinkedList, target int) *LinkedList {
//	var head *LinkedList
//	tail := list
//	for list != nil {
//		next := list.next
//		if list.data < target {
//			list.next = head
//			head = list
//		} else {
//			tail.next = list
//			tail = list
//		}
//		list = next
//	}
//	tail.next = nil
//	return head
//}

func partition(list *LinkedList, target int) *LinkedList {
	if list == nil {
		return list
	}
	var head = &LinkedList{}
	var tail = &LinkedList{}
	htail := head
	ttail := tail
	for list != nil {
		if list.data < target {
			htail.next = list
			htail = htail.next
		} else {
			ttail.next = list
			ttail = ttail.next
		}
		list = list.next
	}
	htail.next = tail.next
	return head.next
}
