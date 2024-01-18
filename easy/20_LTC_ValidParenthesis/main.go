package main

import (
	"errors"
	"fmt"
)

type Stack []byte

func (st *Stack) push(v byte) {
	*st = append(*st, v)
}

func (st *Stack) isValidPopItem(v byte) bool {
	inverseValue := map[byte]byte{'}': '{', ']': '[', ')': '('}
	return inverseValue[v] == st.peek()
}
func (st *Stack) pop(v byte) byte {
	if st.isEmpty() {
		errors.New("Trying to remove from an empty stack")
	}
	if st.isValidPopItem(v) {
		top := (*st)[len(*st)-1]
		*st = (*st)[:len(*st)-1]
		return top
	}
	return 0
}
func (st *Stack) isEmpty() bool {
	return len(*st) == 0
}
func (st *Stack) peek() byte {
	if st.isEmpty() {
		return 0
	}
	top := (*st)[len(*st)-1]
	return top
}

func isOpenSign(character byte) bool {
	return character == '{' || character == '(' || character == '['
}

func areParenthesisValid(word string) bool {
	var wordSize = len(word)
	var openParenthesis Stack
	for i := 0; i < wordSize; i++ {
		if isOpenSign(word[i]) {
			openParenthesis.push(word[i])
		} else {
			poppedSign := openParenthesis.pop(word[i])
			if poppedSign == 0 {
				return false
			}
		}
	}
	return openParenthesis.isEmpty()
}

/*
*

	func isValid(s string) bool {
	    mp := map[byte]byte{
	        '{': '}',
	        '[': ']',
	        '(': ')',
	    }
	    st := make([]byte, 0)
	    for i := 0; i < len(s); i++ {
	        if val, ok := mp[s[i]]; ok {
	            st = append(st, val)
	            continue
	        }
	        if len(st) == 0 {
	            return false
	        }
	        ch := st[len(st) - 1]
	        if ch != s[i] {
	            return false
	        }
	        st = st[:len(st) - 1]
	    }
	    if len(st) == 0 {
	        return true
	    }
	    return false
	}
*/
func main() {
	fmt.Println(areParenthesisValid("()"))
	fmt.Println(areParenthesisValid("()[]{}"))
	fmt.Println(areParenthesisValid("(]"))
	fmt.Println(areParenthesisValid(")))]"))
}
