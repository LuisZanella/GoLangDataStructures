package main

type node struct {
	value string
}

type linkedList struct {
	value node
	next  *linkedList
}
