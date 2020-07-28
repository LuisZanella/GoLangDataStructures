package main

import (
	f "fmt"

	array "./src/Class"
)

func main() {
	f.Println("Stack 1")
	stack := array.New()
	stack.Push(1)
	stack.Push(3)
	stack.Pop()
	f.Println(stack)
	f.Println("Stack 2")
	stack2 := array.NewWithValue(5)
	f.Println(stack2)
	stack2.Push(1)
	stack2.Push(3)
	stack2.Push("hi")
	f.Println(stack2)
	stack2.Delete(0)
	f.Println(stack2)
	f.Println(stack2.Pop())
	f.Println(stack2.Get(0))
	stack2.Push(8)
	f.Println(stack2)
}
