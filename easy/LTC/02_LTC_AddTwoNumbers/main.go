package main

import "fmt"

//**
//You are given two non-empty linked lists representing two non-negative integers. The digits are stored in reverse order,
//and each of their nodes contains a single digit. Add the two numbers and return the sum as a linked list.
//You may assume the two numbers do not contain any leading zero, except the number 0 itself.
//*/

type Node struct {
	value    int
	nextNode *Node
}

func main() {
	l1 := &Node{2, &Node{4, &Node{3, nil}}}
	l2 := &Node{5, &Node{6, &Node{4, nil}}}

	dummy := &Node{}
	current := dummy
	carry := 0
	sum := 0
	quotient := 0
	for l1 != nil || l2 != nil || carry > 0 {
		sum = carry
		if l1 != nil {
			sum += l1.value
			l1 = l1.nextNode
		}
		if l2 != nil {
			sum += l2.value
			l2 = l2.nextNode
		}
		carry = sum / 10
		quotient = sum % 10

		current.nextNode = &Node{value: quotient}
		current = current.nextNode
	}
	printList(dummy)
}
func printList(head *Node) {
	for head != nil {
		fmt.Print(head.value)
		if head.nextNode != nil {
			fmt.Print(" -> ")
		}
		head = head.nextNode
	}
	fmt.Println()
}
