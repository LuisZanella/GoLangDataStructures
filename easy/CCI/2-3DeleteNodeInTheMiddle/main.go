package main

import "errors"

type LinkedList struct {
	data int
	next *LinkedList
}

func DeleteNode(l *LinkedList) (*LinkedList, error) {

	if l == nil || l.next == nil {
		return nil, errors.New("Not a proper 'middle' node")
	}
	l.data = l.next.data
	l.next = l.next.next
	return l, nil
}
