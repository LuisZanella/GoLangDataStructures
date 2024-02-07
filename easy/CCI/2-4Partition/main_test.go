package main

import "testing"

func TestPartition(t *testing.T) {
	partition(&LinkedList{2,
		&LinkedList{5,
			&LinkedList{10,
				&LinkedList{1, &LinkedList{7, nil}}}}}, 5)
}
