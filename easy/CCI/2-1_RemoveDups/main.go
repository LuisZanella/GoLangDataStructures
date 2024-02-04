package main

type node struct {
	value int
}

type linkedList struct {
	node *node
	next *linkedList
}

func removeDups(root *linkedList) *linkedList {
	var m = map[int]bool{}
	var p *linkedList
	helper := root
	for helper != nil && helper.node != nil {

		if m[helper.node.value] {
			p.next = p.next.next
			helper = helper.next
			continue
		}
		m[helper.node.value] = true
		p = helper
		helper = helper.next
	}
	return root
}
