package main

import "errors"

type LinkedList struct {
	data int
	next *LinkedList
}

func ReturnKthFromLast(l *LinkedList, t int) (*LinkedList, error) {
	if t < 0 {
		return nil, errors.New("Out of the index")
	}
	p := l
	r := l
	for i := 0; i < t; i++ {
		if r == nil {
			return nil, errors.New("Out of the index")
		}
		r = r.next
	}
	for r != nil {
		p = p.next
		r = r.next
	}
	return p, nil
}
