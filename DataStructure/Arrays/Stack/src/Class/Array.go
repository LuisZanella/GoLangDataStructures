package array

import (
	"errors"
)

type Array struct {
	length int
	data   []interface{}
}

func New() Array {
	return (Array{0, nil})
}
func NewWithValue(value interface{}) Array {
	array := Array{1, append(make([]interface{}, 0), value)}
	return array
}

func (stack *Array) Get(index int) interface{} {
	return stack.data[index]
}
func (stack *Array) Push(value interface{}) int {
	stack.data = append(stack.data, value)
	stack.length++
	return stack.length
}
func (stack *Array) Pop() (value interface{}, err error) {
	if stack.length <= 0 {
		return nil, errors.New("Stack empty, fill it")
	}
	stack.data = stack.data[:stack.length-1] // stack.data[:stack.length] = stack.data[0:stack.length]
	stack.length--
	return stack.length, err
}

func (stack *Array) Delete(index int) {
	ShiftItems(index, stack)
}
func ShiftItems(index int, stack *Array) {
	stack.data = append(stack.data[:index], stack.data[index+1:]...)
	// stack.data = append(stack.data[:index], stack.data[index+1:]...)
	// =
	// for i := index; i < stack.length-1; i++ {
	// 	stack.data[i] = stack.data[i+1]
	// }
	// stack.data = stack.data[:stack.length-1]
	stack.length--
}
